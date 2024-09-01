package interpret

import (
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/yaml"
	"math"
	"slices"
	"strings"
)

type Interpreter interface {
	Interpret(*InterpretInput) *InterpretOutput
	InterpretExpr(*InterpretExprInput) *InterpretExprOutput
	InterpretEval(*InterpretExprInput) *InterpretExprOutput
	InterpretScalar(*InterpretExprInput) *InterpretExprOutput
	InterpretObj(*InterpretExprInput) *InterpretExprOutput
	InterpretArr(*InterpretExprInput) *InterpretExprOutput
	InterpretJson(*InterpretExprInput) *InterpretExprOutput
	InterpretRangeIter(*InterpretExprInput) *InterpretExprOutput
	InterpretGetElem(*InterpretExprInput) *InterpretExprOutput
	InterpretFunCall(*InterpretExprInput) *InterpretExprOutput
	InterpretCases(*InterpretExprInput) *InterpretExprOutput
	InterpretOpUnary(*InterpretExprInput) *InterpretExprOutput
	InterpretOpBinary(*InterpretExprInput) *InterpretExprOutput
	InterpretOpVariadic(*InterpretExprInput) *InterpretExprOutput
}

func NewInterpreter() Interpreter {
	return &interpreter{}
}

type interpreter struct{}

func (i *interpreter) Interpret(input *InterpretInput) *InterpretOutput {
	// Decode input
	v := yaml.NewDecoder().Decode(&yaml.DecodeInput{Yaml: input.Source})
	if v.IsError {
		return &InterpretOutput{
			Status:       InterpretOutput_DECODE_ERROR,
			ErrorMessage: v.ErrorMessage,
		}
	}

	// Validate input
	{
		v := NewValidator().Validate(&ValidateInput{Source: input.Source})
		if v.Status != ValidateOutput_OK {
			return &InterpretOutput{
				Status:       InterpretOutput_VALIDATE_ERROR,
				ErrorMessage: v.ErrorMessage,
			}
		}
	}

	// Interpret input
	e := i.InterpretExpr(&InterpretExprInput{Path: &Path{}, Defs: &FunDefList{}, Expr: v.Value})
	if e.Status != InterpretExprOutput_OK {
		return &InterpretOutput{
			Status:        InterpretOutput_EXPR_ERROR,
			ErrorMessage:  e.ErrorMessage,
			ExprStatus:    e.Status,
			ExprErrorPath: e.ErrorPath,
		}
	}
	return &InterpretOutput{Value: e.Value}
}

func (i *interpreter) InterpretExpr(input *InterpretExprInput) *InterpretExprOutput {
	v := input.Expr
	switch v.Type {
	case yaml.Type_TYPE_BOOL, yaml.Type_TYPE_NUM, yaml.Type_TYPE_STR:
		return i.InterpretScalar(input)
	case yaml.Type_TYPE_OBJ:
		switch {
		case hasKey(input.Expr, "eval"):
			return i.InterpretEval(input)
		case hasKey(input.Expr, "obj"):
			return i.InterpretObj(input)
		case hasKey(input.Expr, "arr"):
			return i.InterpretArr(input)
		case hasKey(input.Expr, "json"):
			return i.InterpretJson(input)
		case hasKey(input.Expr, "for"):
			return i.InterpretRangeIter(input)
		case hasKey(input.Expr, "get"):
			return i.InterpretGetElem(input)
		case hasKey(input.Expr, "ref"):
			return i.InterpretFunCall(input)
		case hasKey(input.Expr, "cases"):
			return i.InterpretCases(input)
		case hasKey(input.Expr, OpUnary_LEN.KeyName()),
			hasKey(input.Expr, OpUnary_NOT.KeyName()),
			hasKey(input.Expr, OpUnary_FLAT.KeyName()),
			hasKey(input.Expr, OpUnary_FLOOR.KeyName()),
			hasKey(input.Expr, OpUnary_CEIL.KeyName()),
			hasKey(input.Expr, OpUnary_ABORT.KeyName()):
			return i.InterpretOpUnary(input)
		case hasKey(input.Expr, OpBinary_SUB.KeyName()),
			hasKey(input.Expr, OpBinary_DIV.KeyName()),
			hasKey(input.Expr, OpBinary_MOD.KeyName()),
			hasKey(input.Expr, OpBinary_EQ.KeyName()),
			hasKey(input.Expr, OpBinary_NEQ.KeyName()),
			hasKey(input.Expr, OpBinary_LT.KeyName()),
			hasKey(input.Expr, OpBinary_LTE.KeyName()),
			hasKey(input.Expr, OpBinary_GT.KeyName()),
			hasKey(input.Expr, OpBinary_GTE.KeyName()):
			return i.InterpretOpBinary(input)
		case hasKey(input.Expr, OpVariadic_ADD.KeyName()),
			hasKey(input.Expr, OpVariadic_MUL.KeyName()),
			hasKey(input.Expr, OpVariadic_AND.KeyName()),
			hasKey(input.Expr, OpVariadic_OR.KeyName()),
			hasKey(input.Expr, OpVariadic_CAT.KeyName()),
			hasKey(input.Expr, OpVariadic_MIN.KeyName()),
			hasKey(input.Expr, OpVariadic_MAX.KeyName()),
			hasKey(input.Expr, OpVariadic_MERGE.KeyName()):
			return i.InterpretOpVariadic(input)
		}
	}
	return errorUnsupportedExpr(input.Path, v)
}

func (i *interpreter) InterpretEval(input *InterpretExprInput) *InterpretExprOutput {
	path := input.Path
	st := input.Defs
	if where, ok := input.Expr.Obj["where"]; ok {
		for pos, w := range where.Arr {
			path := path.AppendKey("where")
			def := w.Obj["def"]
			value := w.Obj["value"]
			funDef := &FunDef{Def: def.Str, Value: value, Path: path.AppendIndex(pos)}
			if with, ok := w.Obj["with"]; ok {
				for _, w := range with.Arr {
					funDef.With = append(funDef.With, w.Str)
				}
			}
			st = st.Register(funDef)
		}
	}
	return i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("eval"), Defs: st, Expr: input.Expr.Obj["eval"]})
}

func (i *interpreter) InterpretScalar(input *InterpretExprInput) *InterpretExprOutput {
	return &InterpretExprOutput{Value: input.Expr}
}

func (i *interpreter) InterpretObj(input *InterpretExprInput) *InterpretExprOutput {
	obj := input.Expr.Obj["obj"]
	path := input.Path.AppendKey("obj")
	v := map[string]*yaml.Value{}
	for pos, val := range obj.Obj {
		expr := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey(pos), Defs: input.Defs, Expr: val})
		if expr.Status != InterpretExprOutput_OK {
			return expr
		}
		v[pos] = expr.Value
	}
	return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: v}}
}

func (i *interpreter) InterpretArr(input *InterpretExprInput) *InterpretExprOutput {
	arr := input.Expr.Obj["arr"]
	path := input.Path.AppendKey("arr")
	v := []*yaml.Value{}
	for pos, val := range arr.Arr {
		expr := i.InterpretExpr(&InterpretExprInput{Path: path.AppendIndex(pos), Defs: input.Defs, Expr: val})
		if expr.Status != InterpretExprOutput_OK {
			return expr
		}
		v = append(v, expr.Value)
	}
	return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
}

func (i *interpreter) InterpretJson(input *InterpretExprInput) *InterpretExprOutput {
	v := input.Expr.Obj["json"]
	return &InterpretExprOutput{Value: v}
}

func (i *interpreter) InterpretRangeIter(input *InterpretExprInput) *InterpretExprOutput {
	path := input.Path
	for_ := input.Expr.Obj["for"]
	forPos, forVal := for_.Arr[0].Str, for_.Arr[1].Str
	in := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("in"), Defs: input.Defs, Expr: input.Expr.Obj["in"]})
	if in.Status != InterpretExprOutput_OK {
		return in
	}
	switch in.Value.Type {
	default:
		return errorUnexpectedType(path.AppendKey("in"), []yaml.Type{yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, in.Value.Type)
	case yaml.Type_TYPE_ARR:
		v := []*yaml.Value{}
		for pos, val := range in.Value.Arr {
			st := input.Defs
			st = st.Register(&FunDef{
				Def:   forPos,
				Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(pos)},
				Path:  path.AppendKey("for").AppendIndex(0),
			})
			st = st.Register(&FunDef{
				Def:   forVal,
				Value: val,
				Path:  path.AppendKey("for").AppendIndex(1),
			})
			do := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("do"), Defs: st, Expr: input.Expr.Obj["do"]})
			if do.Status != InterpretExprOutput_OK {
				return do
			}
			v = append(v, do.Value)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
	case yaml.Type_TYPE_OBJ:
		v := map[string]*yaml.Value{}
		for pos, val := range in.Value.Obj {
			st := input.Defs
			st = st.Register(&FunDef{
				Def:   forPos,
				Value: &yaml.Value{Type: yaml.Type_TYPE_STR, Str: pos},
				Path:  path.AppendKey("for").AppendIndex(0),
			})
			st = st.Register(&FunDef{
				Def:   forVal,
				Value: val,
				Path:  path.AppendKey("for").AppendIndex(1),
			})
			do := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("do"), Defs: st, Expr: input.Expr.Obj["do"]})
			if do.Status != InterpretExprOutput_OK {
				return do
			}
			v[pos] = do.Value
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: v}}
	}
}

func (i *interpreter) InterpretGetElem(input *InterpretExprInput) *InterpretExprOutput {
	path := input.Path
	get := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("get"), Defs: input.Defs, Expr: input.Expr.Obj["get"]})
	if get.Status != InterpretExprOutput_OK {
		return get
	}
	from := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("from"), Defs: input.Defs, Expr: input.Expr.Obj["from"]})
	if from.Status != InterpretExprOutput_OK {
		return from
	}

	switch from.Value.Type {
	default:
		return errorUnexpectedType(path.AppendKey("from"), []yaml.Type{yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, from.Value.Type)
	case yaml.Type_TYPE_ARR:
		if get.Value.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(path.AppendKey("get"), []yaml.Type{yaml.Type_TYPE_NUM}, get.Value.Type)
		}
		if !get.Value.CanInt() {
			return errorArithmeticError(path.AppendKey("get"), fmt.Sprintf("index %v is not an integer", get.Value.Num))
		}
		pos := int(get.Value.Num)
		if pos < 0 || pos >= len(from.Value.Arr) {
			return errorIndexOutOfBounds(path.AppendKey("get"), 0, len(from.Value.Arr), pos)
		}
		return &InterpretExprOutput{Value: from.Value.Arr[pos]}
	case yaml.Type_TYPE_OBJ:
		if get.Value.Type != yaml.Type_TYPE_STR {
			return errorUnexpectedType(path.AppendKey("get"), []yaml.Type{yaml.Type_TYPE_STR}, get.Value.Type)
		}
		pos := get.Value.Str
		if _, ok := from.Value.Obj[pos]; ok {
			return errorKeyNotFound(path.AppendKey("get"), from.Value.Keys(), pos)
		}
		return &InterpretExprOutput{Value: from.Value.Obj[pos]}
	}
}

func (i *interpreter) InterpretFunCall(input *InterpretExprInput) *InterpretExprOutput {
	path := input.Path
	funCall := input.Expr
	ref := funCall.Obj["ref"]
	funDef := input.Defs.Find(ref.Str)
	if funDef == nil {
		return errorReferenceNotFound(path.AppendKey("ref"), ref.Str)
	}
	st := input.Defs
	for _, argName := range funDef.Def.With {
		with, ok := funCall.Obj["with"]
		if !ok {
			return errorKeyNotFound(path, funCall.Keys(), "with")
		}
		argVal, ok := with.Obj[argName]
		if !ok {
			return errorKeyNotFound(path.AppendKey("with"), with.Keys(), argName)
		}
		st = st.Register(&FunDef{Def: argName, Value: argVal, Path: path.AppendKey("with").AppendKey(argName)})
	}
	return i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("ref"), Defs: st, Expr: funDef.Def.Value})
}

func (i *interpreter) InterpretCases(input *InterpretExprInput) *InterpretExprOutput {
	path := input.Path
	cases := input.Expr.Obj["cases"]
	for _, c := range cases.Arr {
		switch {
		case hasKey(c, "when"):
			when := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("when"), Defs: input.Defs, Expr: c.Obj["when"]})
			if when.Status != InterpretExprOutput_OK {
				return when
			}
			if when.Value.Type != yaml.Type_TYPE_BOOL {
				return errorUnexpectedType(path.AppendKey("when"), []yaml.Type{yaml.Type_TYPE_BOOL}, when.Value.Type)
			}
			if when.Value.Bool {
				then := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("then"), Defs: input.Defs, Expr: c.Obj["then"]})
				if then.Status != InterpretExprOutput_OK {
					return then
				}
				return then
			}
		case hasKey(c, "otherwise"):
			otherwise := i.InterpretExpr(&InterpretExprInput{Path: path.AppendKey("otherwise"), Defs: input.Defs, Expr: c.Obj["otherwise"]})
			if otherwise.Status != InterpretExprOutput_OK {
				return otherwise
			}
			return otherwise
		}
	}
	return errorCasesNotExhaustive(path)
}

func (i *interpreter) InterpretOpUnary(input *InterpretExprInput) *InterpretExprOutput {
	var (
		operator string
		operand  *yaml.Value
	)
	for k, v := range input.Expr.Obj { // only one property exists
		operator = k
		o := i.InterpretExpr(&InterpretExprInput{Path: input.Path.AppendKey(k), Defs: input.Defs, Expr: v})
		if o.Status != InterpretExprOutput_OK {
			return o
		}
		operand = o.Value
	}
	switch operator {
	default:
		return errorUnsupportedOperation(input.Path, operator)
	case OpUnary_LEN.KeyName():
		switch operand.Type {
		default:
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR}, operand.Type)
		case yaml.Type_TYPE_STR:
			return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Str))}}
		case yaml.Type_TYPE_ARR:
			return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Arr))}}
		case yaml.Type_TYPE_OBJ:
			return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Obj))}}
		}
	case OpUnary_NOT.KeyName():
		if operand.Type != yaml.Type_TYPE_BOOL {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Type)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: !operand.Bool}}
	case OpUnary_FLAT.KeyName():
		if operand.Type != yaml.Type_TYPE_ARR {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_ARR}, operand.Type)
		}
		v := []*yaml.Value{}
		for _, elem := range operand.Arr {
			if elem.Type != yaml.Type_TYPE_ARR {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_ARR}, elem.Type)
			}
			v = append(v, elem.Arr...)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
	case OpUnary_FLOOR.KeyName():
		if operand.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: math.Floor(operand.Num)}}
	case OpUnary_CEIL.KeyName():
		if operand.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: math.Ceil(operand.Num)}}
	}
}

func (i *interpreter) InterpretOpBinary(input *InterpretExprInput) *InterpretExprOutput {
	var (
		operator           string
		operandL, operandR *yaml.Value
	)
	for k, v := range input.Expr.Obj { // only one property exists
		operator = k
		ol := i.InterpretExpr(&InterpretExprInput{Path: input.Path.AppendKey(k).AppendIndex(0), Defs: input.Defs, Expr: v.Arr[0]})
		if ol.Status != InterpretExprOutput_OK {
			return ol
		}
		or := i.InterpretExpr(&InterpretExprInput{Path: input.Path.AppendKey(k).AppendIndex(1), Defs: input.Defs, Expr: v.Arr[1]})
		if or.Status != InterpretExprOutput_OK {
			return or
		}
		operandL, operandR = ol.Value, or.Value
	}
	switch operator {
	default:
		return errorUnsupportedOperation(input.Path, operator)
	case OpBinary_SUB.KeyName():
		if operandL.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(0), []yaml.Type{yaml.Type_TYPE_NUM}, operandL.Type)
		}
		if operandR.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(1), []yaml.Type{yaml.Type_TYPE_NUM}, operandR.Type)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: operandL.Num - operandR.Num}}
	case OpBinary_DIV.KeyName():
		if operandL.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(0), []yaml.Type{yaml.Type_TYPE_NUM}, operandL.Type)
		}
		if operandR.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(1), []yaml.Type{yaml.Type_TYPE_NUM}, operandR.Type)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: operandL.Num / operandR.Num}}
	case OpBinary_MOD.KeyName():
		if operandL.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(0), []yaml.Type{yaml.Type_TYPE_NUM}, operandL.Type)
		}
		if operandR.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(1), []yaml.Type{yaml.Type_TYPE_NUM}, operandR.Type)
		}
		if !operandL.CanInt() {
			return errorArithmeticError(input.Path.AppendKey(operator).AppendIndex(0), fmt.Sprintf("left operand %v is not an integer", operandL.Num))
		}
		if !operandR.CanInt() {
			return errorArithmeticError(input.Path.AppendKey(operator).AppendIndex(1), fmt.Sprintf("right operand %v is not an integer", operandR.Num))
		}
		if operandR.Num == 0 {
			return errorArithmeticError(input.Path.AppendKey(operator).AppendIndex(1), "division by zero")
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(int(operandL.Num) % int(operandR.Num))}}
	case OpBinary_EQ.KeyName():
		return equal(input.Path.AppendKey(operator), operandL, operandR)
	case OpBinary_NEQ.KeyName():
		eq := equal(input.Path.AppendKey(operator), operandL, operandR)
		if eq.Status != InterpretExprOutput_OK {
			return eq
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: !eq.Value.Bool}}
	case OpBinary_LT.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != InterpretExprOutput_OK {
			return cmp
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num < 0}}
	case OpBinary_LTE.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != InterpretExprOutput_OK {
			return cmp
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num <= 0}}
	case OpBinary_GT.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != InterpretExprOutput_OK {
			return cmp
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num > 0}}
	case OpBinary_GTE.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != InterpretExprOutput_OK {
			return cmp
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num >= 0}}
	}
}

func (i *interpreter) InterpretOpVariadic(input *InterpretExprInput) *InterpretExprOutput {
	var (
		operator string
		operands []*yaml.Value
	)
	for k, v := range input.Expr.Obj { // only one property exists
		operator = k
		for _, e := range v.Arr {
			o := i.InterpretExpr(&InterpretExprInput{Path: input.Path.AppendKey(k), Defs: input.Defs, Expr: e})
			if o.Status != InterpretExprOutput_OK {
				return o
			}
			operands = append(operands, o.Value)
		}
	}
	switch operator {
	default:
		return errorUnsupportedOperation(input.Path, operator)
	case OpVariadic_ADD.KeyName():
		add := 0.0
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			add += operand.Num
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: add}}
	case OpVariadic_MUL.KeyName():
		mul := 1.0
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			mul *= operand.Num
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: mul}}
	case OpVariadic_AND.KeyName():
		and := true
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_BOOL {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Type)
			}
			and = and && operand.Bool
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: and}}
	case OpVariadic_OR.KeyName():
		or := false
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_BOOL {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Type)
			}
			or = or || operand.Bool
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: or}}
	case OpVariadic_CAT.KeyName():
		cat := ""
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_STR {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_STR}, operand.Type)
			}
			cat += operand.Str
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_STR, Str: cat}}
	case OpVariadic_MIN.KeyName():
		min_ := math.Inf(1)
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			min_ = math.Min(min_, operand.Num)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: min_}}
	case OpVariadic_MAX.KeyName():
		max_ := math.Inf(-1)
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			max_ = math.Max(max_, operand.Num)
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: max_}}
	case OpVariadic_MERGE.KeyName():
		merge := map[string]*yaml.Value{}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_OBJ {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_OBJ}, operand.Type)
			}
			for k, v := range operand.Obj {
				merge[k] = v
			}
		}
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: merge}}
	}
}

func equal(path *Path, l, r *yaml.Value) *InterpretExprOutput {
	falseValue := &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: false}}
	trueValue := &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: true}}
	switch {
	default:
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_NUM, yaml.Type_TYPE_BOOL, yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, l.Type)
	case l.Type != r.Type:
		return falseValue
	case l.Type == yaml.Type_TYPE_NUM:
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: l.Num == r.Num}}
	case l.Type == yaml.Type_TYPE_BOOL:
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: l.Bool == r.Bool}}
	case l.Type == yaml.Type_TYPE_STR:
		return &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: l.Str == r.Str}}
	case l.Type == yaml.Type_TYPE_ARR:
		if len(l.Arr) != len(r.Arr) {
			return falseValue
		}
		for i, l := range l.Arr {
			r := r.Arr[i]
			eq := equal(path, l, r)
			if eq.Status != InterpretExprOutput_OK {
				return eq
			}
			if !eq.Value.Bool {
				return falseValue
			}
		}
		return trueValue
	case l.Type == yaml.Type_TYPE_OBJ:
		lk, rk := l.Keys(), r.Keys()
		slices.Sort(lk)
		slices.Sort(rk)
		if !slices.Equal(lk, rk) {
			return falseValue
		}
		for k, l := range l.Obj {
			r := r.Obj[k]
			eq := equal(path, l, r)
			if eq.Status != InterpretExprOutput_OK {
				return eq
			}
			if !eq.Value.Bool {
				return falseValue
			}
		}
		return trueValue
	}
}
func compare(path *Path, l, r *yaml.Value) *InterpretExprOutput {
	ltValue := &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: -1}}
	gtValue := &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: 1}}
	eqValue := &InterpretExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: 0}}
	switch {
	default:
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_NUM, yaml.Type_TYPE_BOOL, yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR}, l.Type)
	case l.Type == yaml.Type_TYPE_NUM && r.Type == yaml.Type_TYPE_NUM:
		if l.Num < r.Num {
			return ltValue
		}
		if l.Num > r.Num {
			return gtValue
		}
		return eqValue
	case l.Type == yaml.Type_TYPE_BOOL && r.Type == yaml.Type_TYPE_BOOL:
		if !l.Bool && r.Bool {
			return ltValue
		}
		if l.Bool && !r.Bool {
			return gtValue
		}
		return eqValue
	case l.Type == yaml.Type_TYPE_STR && r.Type == yaml.Type_TYPE_STR:
		if l.Str < r.Str {
			return ltValue
		}
		if l.Str > r.Str {
			return gtValue
		}
		return eqValue
	case l.Type == yaml.Type_TYPE_ARR && r.Type == yaml.Type_TYPE_ARR:
		for i, l := range l.Arr {
			r := r.Arr[i]
			cmp := compare(path, l, r)
			if cmp.Status != InterpretExprOutput_OK {
				return cmp
			}
			if cmp.Value.Num != 0 {
				return cmp
			}
		}
		if len(l.Arr) < len(r.Arr) {
			return ltValue
		}
		if len(l.Arr) > len(r.Arr) {
			return gtValue
		}
		return eqValue
	}
}

func errorUnsupportedExpr(path *Path, v *yaml.Value) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_UNSUPPORTED_EXPR,
		ErrorMessage: fmt.Sprintf("unsupported expr: got %v", v.Keys()),
		ErrorPath:    path,
	}
}
func errorUnexpectedType(path *Path, want []yaml.Type, got yaml.Type) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_UNEXPECTED_TYPE,
		ErrorMessage: fmt.Sprintf("unexpected type: want %v, got %v", want, got),
		ErrorPath:    path,
	}
}
func errorArithmeticError(path *Path, message string) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_ARITHMETIC_ERROR,
		ErrorMessage: fmt.Sprintf("arithmetic error: %v", message),
		ErrorPath:    path,
	}
}
func errorIndexOutOfBounds(path *Path, begin, end, index int) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_INDEX_OUT_OF_BOUNDS,
		ErrorMessage: fmt.Sprintf("index out of bounds: %v not in [%v, %v)", index, begin, end),
		ErrorPath:    path,
	}
}
func errorKeyNotFound(path *Path, keys []string, key string) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_KEY_NOT_FOUND,
		ErrorMessage: fmt.Sprintf("key not found: %v not in {%v}", key, strings.Join(keys, ",")),
		ErrorPath:    path,
	}
}
func errorReferenceNotFound(path *Path, ref string) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_REFERENCE_NOT_FOUND,
		ErrorMessage: fmt.Sprintf("reference not found: %v", ref),
		ErrorPath:    path,
	}
}
func errorCasesNotExhaustive(path *Path) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_CASES_NOT_EXHAUSTIVE,
		ErrorMessage: fmt.Sprintf("cases not exhaustive: %v", path.Format()),
		ErrorPath:    path,
	}
}
func errorUnsupportedOperation(path *Path, op string) *InterpretExprOutput {
	return &InterpretExprOutput{
		Status:       InterpretExprOutput_UNSUPPORTED_OPERATION,
		ErrorMessage: fmt.Sprintf("unsupported operation: %v", op),
		ErrorPath:    path,
	}
}
func hasKey(v *yaml.Value, k string) bool {
	for _, key := range v.Keys() {
		if key == k {
			return true
		}
	}
	return false
}
