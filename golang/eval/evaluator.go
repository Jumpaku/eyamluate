package eval

import (
	"fmt"
	"github.com/Jumpaku/eyamluate/golang/yaml"
	"math"
	"slices"
	"strings"
)

type Evaluator interface {
	Evaluate(*EvaluateInput) *EvaluateOutput
	EvaluateExpr(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateEval(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateScalar(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateObj(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateArr(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateJson(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateRangeIter(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateGetElem(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateFunCall(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateCases(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateOpUnary(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateOpBinary(*EvaluateExprInput) *EvaluateExprOutput
	EvaluateOpVariadic(*EvaluateExprInput) *EvaluateExprOutput
}

func NewEvaluator() Evaluator {
	return &evaluator{}
}

type evaluator struct{}

func (i *evaluator) Evaluate(input *EvaluateInput) *EvaluateOutput {
	// Decode input
	v := yaml.NewDecoder().Decode(&yaml.DecodeInput{Yaml: input.Source})
	if v.IsError {
		return &EvaluateOutput{
			Status:       EvaluateOutput_DECODE_ERROR,
			ErrorMessage: v.ErrorMessage,
		}
	}

	// Validate input
	{
		v := NewValidator().Validate(&ValidateInput{Source: input.Source})
		if v.Status != ValidateOutput_OK {
			return &EvaluateOutput{
				Status:       EvaluateOutput_VALIDATE_ERROR,
				ErrorMessage: v.ErrorMessage,
			}
		}
	}

	// Evaluate input
	e := i.EvaluateExpr(&EvaluateExprInput{Path: &Path{}, Defs: EmptyFunDefList(), Expr: v.Value})
	if e.Status != EvaluateExprOutput_OK {
		return &EvaluateOutput{
			Status:        EvaluateOutput_EXPR_ERROR,
			ErrorMessage:  e.ErrorMessage,
			ExprStatus:    e.Status,
			ExprErrorPath: e.ErrorPath,
		}
	}
	return &EvaluateOutput{Value: e.Value}
}

func (i *evaluator) EvaluateExpr(input *EvaluateExprInput) *EvaluateExprOutput {
	switch input.Expr.Type {
	case yaml.Type_TYPE_BOOL, yaml.Type_TYPE_NUM, yaml.Type_TYPE_STR:
		return i.EvaluateScalar(input)
	case yaml.Type_TYPE_OBJ:
		switch {
		case hasKey(input.Expr, "eval"):
			return i.EvaluateEval(input)
		case hasKey(input.Expr, "obj"):
			return i.EvaluateObj(input)
		case hasKey(input.Expr, "arr"):
			return i.EvaluateArr(input)
		case hasKey(input.Expr, "json"):
			return i.EvaluateJson(input)
		case hasKey(input.Expr, "for"):
			return i.EvaluateRangeIter(input)
		case hasKey(input.Expr, "get"):
			return i.EvaluateGetElem(input)
		case hasKey(input.Expr, "ref"):
			return i.EvaluateFunCall(input)
		case hasKey(input.Expr, "cases"):
			return i.EvaluateCases(input)
		case hasKey(input.Expr, OpUnary_LEN.KeyName()),
			hasKey(input.Expr, OpUnary_NOT.KeyName()),
			hasKey(input.Expr, OpUnary_FLAT.KeyName()),
			hasKey(input.Expr, OpUnary_FLOOR.KeyName()),
			hasKey(input.Expr, OpUnary_CEIL.KeyName()),
			hasKey(input.Expr, OpUnary_ABORT.KeyName()):
			return i.EvaluateOpUnary(input)
		case hasKey(input.Expr, OpBinary_SUB.KeyName()),
			hasKey(input.Expr, OpBinary_DIV.KeyName()),
			hasKey(input.Expr, OpBinary_EQ.KeyName()),
			hasKey(input.Expr, OpBinary_NEQ.KeyName()),
			hasKey(input.Expr, OpBinary_LT.KeyName()),
			hasKey(input.Expr, OpBinary_LTE.KeyName()),
			hasKey(input.Expr, OpBinary_GT.KeyName()),
			hasKey(input.Expr, OpBinary_GTE.KeyName()):
			return i.EvaluateOpBinary(input)
		case hasKey(input.Expr, OpVariadic_ADD.KeyName()),
			hasKey(input.Expr, OpVariadic_MUL.KeyName()),
			hasKey(input.Expr, OpVariadic_AND.KeyName()),
			hasKey(input.Expr, OpVariadic_OR.KeyName()),
			hasKey(input.Expr, OpVariadic_CAT.KeyName()),
			hasKey(input.Expr, OpVariadic_MIN.KeyName()),
			hasKey(input.Expr, OpVariadic_MAX.KeyName()),
			hasKey(input.Expr, OpVariadic_MERGE.KeyName()):
			return i.EvaluateOpVariadic(input)
		}
	}
	return errorUnsupportedExpr(input.Path, input.Expr)
}

func (i *evaluator) EvaluateEval(input *EvaluateExprInput) *EvaluateExprOutput {
	path := input.Path
	st := input.Defs
	if where, ok := input.Expr.Obj["where"]; ok {
		path := path.AppendKey("where")
		for pos, w := range where.Arr {
			def, value := w.Obj["def"], w.Obj["value"]
			funDef := &FunDef{Def: def.Str, Value: value, Path: path.AppendIndex(pos)}
			if with, ok := w.Obj["with"]; ok {
				for _, w := range with.Arr {
					funDef.With = append(funDef.With, w.Str)
				}
			}
			st = st.Register(funDef)
		}
	}
	return i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("eval"), Defs: st, Expr: input.Expr.Obj["eval"]})
}

func (i *evaluator) EvaluateScalar(input *EvaluateExprInput) *EvaluateExprOutput {
	return &EvaluateExprOutput{Value: input.Expr}
}

func (i *evaluator) EvaluateObj(input *EvaluateExprInput) *EvaluateExprOutput {
	obj := input.Expr.Obj["obj"]
	path := input.Path.AppendKey("obj")
	v := map[string]*yaml.Value{}
	for pos, val := range obj.Obj {
		expr := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey(pos), Defs: input.Defs, Expr: val})
		if expr.Status != EvaluateExprOutput_OK {
			return expr
		}
		v[pos] = expr.Value
	}
	return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: v}}
}

func (i *evaluator) EvaluateArr(input *EvaluateExprInput) *EvaluateExprOutput {
	arr := input.Expr.Obj["arr"]
	path := input.Path.AppendKey("arr")
	v := []*yaml.Value{}
	for pos, val := range arr.Arr {
		expr := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendIndex(pos), Defs: input.Defs, Expr: val})
		if expr.Status != EvaluateExprOutput_OK {
			return expr
		}
		v = append(v, expr.Value)
	}
	return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
}

func (i *evaluator) EvaluateJson(input *EvaluateExprInput) *EvaluateExprOutput {
	v := input.Expr.Obj["json"]
	return &EvaluateExprOutput{Value: v}
}

func (i *evaluator) EvaluateRangeIter(input *EvaluateExprInput) *EvaluateExprOutput {
	path := input.Path
	for_ := input.Expr.Obj["for"]
	forPos, forVal := for_.Arr[0].Str, for_.Arr[1].Str
	in := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("in"), Defs: input.Defs, Expr: input.Expr.Obj["in"]})
	if in.Status != EvaluateExprOutput_OK {
		return in
	}
	switch in.Value.Type {
	default:
		return errorUnexpectedType(path.AppendKey("in"), []yaml.Type{yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, in.Value.Type)
	case yaml.Type_TYPE_STR:
		v := []*yaml.Value{}
		for pos, val := range []rune(in.Value.Str) {
			st := input.Defs
			st = st.Register(&FunDef{
				Def:   forPos,
				Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(pos)},
				Path:  path.AppendKey("for").AppendIndex(0),
			})
			st = st.Register(&FunDef{
				Def:   forVal,
				Value: &yaml.Value{Type: yaml.Type_TYPE_STR, Str: string(val)},
				Path:  path.AppendKey("for").AppendIndex(1),
			})
			if if_, ok := input.Expr.Obj["if"]; ok {
				if_ := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("if"), Defs: st, Expr: if_})
				if if_.Status != EvaluateExprOutput_OK {
					return if_
				}
				if if_.Value.Type != yaml.Type_TYPE_BOOL {
					return errorUnexpectedType(path.AppendKey("if"), []yaml.Type{yaml.Type_TYPE_BOOL}, if_.Value.Type)
				}
				if !if_.Value.Bool {
					continue
				}
			}
			do := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("do"), Defs: st, Expr: input.Expr.Obj["do"]})
			if do.Status != EvaluateExprOutput_OK {
				return do
			}
			v = append(v, do.Value)
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
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
			if if_, ok := input.Expr.Obj["if"]; ok {
				if_ := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("if"), Defs: st, Expr: if_})
				if if_.Status != EvaluateExprOutput_OK {
					return if_
				}
				if if_.Value.Type != yaml.Type_TYPE_BOOL {
					return errorUnexpectedType(path.AppendKey("if"), []yaml.Type{yaml.Type_TYPE_BOOL}, if_.Value.Type)
				}
				if !if_.Value.Bool {
					continue
				}
			}
			do := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("do"), Defs: st, Expr: input.Expr.Obj["do"]})
			if do.Status != EvaluateExprOutput_OK {
				return do
			}
			v = append(v, do.Value)
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
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
			if if_, ok := input.Expr.Obj["if"]; ok {
				if_ := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("if"), Defs: st, Expr: if_})
				if if_.Status != EvaluateExprOutput_OK {
					return if_
				}
				if if_.Value.Type != yaml.Type_TYPE_BOOL {
					return errorUnexpectedType(path.AppendKey("if"), []yaml.Type{yaml.Type_TYPE_BOOL}, if_.Value.Type)
				}
				if !if_.Value.Bool {
					continue
				}
			}
			do := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("do"), Defs: st, Expr: input.Expr.Obj["do"]})
			if do.Status != EvaluateExprOutput_OK {
				return do
			}
			v[pos] = do.Value
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: v}}
	}
}

func (i *evaluator) EvaluateGetElem(input *EvaluateExprInput) *EvaluateExprOutput {
	path := input.Path
	get := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("get"), Defs: input.Defs, Expr: input.Expr.Obj["get"]})
	if get.Status != EvaluateExprOutput_OK {
		return get
	}
	from := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("from"), Defs: input.Defs, Expr: input.Expr.Obj["from"]})
	if from.Status != EvaluateExprOutput_OK {
		return from
	}

	switch from.Value.Type {
	default:
		return errorUnexpectedType(path.AppendKey("from"), []yaml.Type{yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, from.Value.Type)
	case yaml.Type_TYPE_STR:
		if get.Value.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(path.AppendKey("get"), []yaml.Type{yaml.Type_TYPE_NUM}, get.Value.Type)
		}
		if !get.Value.CanInt() {
			return errorArithmeticError(path.AppendKey("get"), fmt.Sprintf("index %v is not an integer", get.Value.Num))
		}
		pos := int(get.Value.Num)
		if pos < 0 || pos >= len([]rune(from.Value.Str)) {
			return errorIndexOutOfBounds(path.AppendKey("get"), 0, len(from.Value.Arr), pos)
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_STR, Str: string([]rune(from.Value.Str)[pos])}}
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
		return &EvaluateExprOutput{Value: from.Value.Arr[pos]}
	case yaml.Type_TYPE_OBJ:
		if get.Value.Type != yaml.Type_TYPE_STR {
			return errorUnexpectedType(path.AppendKey("get"), []yaml.Type{yaml.Type_TYPE_STR}, get.Value.Type)
		}
		pos := get.Value.Str
		if _, ok := from.Value.Obj[pos]; !ok {
			return errorKeyNotFound(path.AppendKey("get"), pos, from.Value.Keys())
		}
		return &EvaluateExprOutput{Value: from.Value.Obj[pos]}
	}
}

func (i *evaluator) EvaluateFunCall(input *EvaluateExprInput) *EvaluateExprOutput {
	path := input.Path
	funCall := input.Expr
	ref := funCall.Obj["ref"]
	funDef := input.Defs.Find(ref.Str)
	if funDef == nil {
		return errorReferenceNotFound(path.AppendKey("ref"), ref.Str)
	}
	st := funDef
	for _, argName := range funDef.Def.With {
		with, ok := funCall.Obj["with"]
		if !ok {
			return errorKeyNotFound(path, "with", funCall.Keys())
		}
		argVal, ok := with.Obj[argName]
		if !ok {
			return errorKeyNotFound(path.AppendKey("with"), argName, with.Keys())
		}
		arg := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("with").AppendKey(argName), Defs: input.Defs, Expr: argVal})
		if arg.Status != EvaluateExprOutput_OK {
			return arg
		}
		jsonExpr := &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: map[string]*yaml.Value{"json": arg.Value}}
		st = st.Register(&FunDef{Def: argName, Value: jsonExpr, Path: path.AppendKey("with").AppendKey(argName)})
	}
	return i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("ref"), Defs: st, Expr: funDef.Def.Value})
}

func (i *evaluator) EvaluateCases(input *EvaluateExprInput) *EvaluateExprOutput {
	path := input.Path
	cases := input.Expr.Obj["cases"]
	for pos, c := range cases.Arr {
		path := path.AppendKey("cases").AppendIndex(pos)
		switch {
		case hasKey(c, "when"):
			when := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("when"), Defs: input.Defs, Expr: c.Obj["when"]})
			if when.Status != EvaluateExprOutput_OK {
				return when
			}
			if when.Value.Type != yaml.Type_TYPE_BOOL {
				return errorUnexpectedType(path.AppendKey("when"), []yaml.Type{yaml.Type_TYPE_BOOL}, when.Value.Type)
			}
			if when.Value.Bool {
				then := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("then"), Defs: input.Defs, Expr: c.Obj["then"]})
				if then.Status != EvaluateExprOutput_OK {
					return then
				}
				return then
			}
		case hasKey(c, "otherwise"):
			otherwise := i.EvaluateExpr(&EvaluateExprInput{Path: path.AppendKey("otherwise"), Defs: input.Defs, Expr: c.Obj["otherwise"]})
			if otherwise.Status != EvaluateExprOutput_OK {
				return otherwise
			}
			return otherwise
		}
	}
	return errorCasesNotExhaustive(path)
}

func (i *evaluator) EvaluateOpUnary(input *EvaluateExprInput) *EvaluateExprOutput {
	var (
		operator string
		operand  *yaml.Value
	)
	for k, v := range input.Expr.Obj { // only one property exists
		operator = k
		o := i.EvaluateExpr(&EvaluateExprInput{Path: input.Path.AppendKey(k), Defs: input.Defs, Expr: v})
		if o.Status != EvaluateExprOutput_OK {
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
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, operand.Type)
		case yaml.Type_TYPE_STR:
			return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Str))}}
		case yaml.Type_TYPE_ARR:
			return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Arr))}}
		case yaml.Type_TYPE_OBJ:
			return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Obj))}}
		}
	case OpUnary_NOT.KeyName():
		if operand.Type != yaml.Type_TYPE_BOOL {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Type)
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: !operand.Bool}}
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
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: v}}
	case OpUnary_FLOOR.KeyName():
		if operand.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
		}
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: math.Floor(operand.Num)}
		if !isFiniteNumber(v) {
			return errorArithmeticError(input.Path.AppendKey(operator), fmt.Sprintf("floor(%v) is not a finite number", operand.Num))
		}
		return &EvaluateExprOutput{Value: v}
	case OpUnary_CEIL.KeyName():
		if operand.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
		}
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: math.Ceil(operand.Num)}
		if !isFiniteNumber(v) {
			return errorArithmeticError(input.Path.AppendKey(operator), fmt.Sprintf("ceil(%v) is not a finite number", operand.Num))
		}
		return &EvaluateExprOutput{Value: v}
	case OpUnary_ABORT.KeyName():
		if operand.Type != yaml.Type_TYPE_STR {
			return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_STR}, operand.Type)
		}
		return &EvaluateExprOutput{Status: EvaluateExprOutput_ABORTED, ErrorMessage: operand.Str}
	}
}

func (i *evaluator) EvaluateOpBinary(input *EvaluateExprInput) *EvaluateExprOutput {
	var (
		operator           string
		operandL, operandR *yaml.Value
	)
	for k, v := range input.Expr.Obj { // only one property exists
		operator = k
		ol := i.EvaluateExpr(&EvaluateExprInput{Path: input.Path.AppendKey(k).AppendIndex(0), Defs: input.Defs, Expr: v.Arr[0]})
		if ol.Status != EvaluateExprOutput_OK {
			return ol
		}
		or := i.EvaluateExpr(&EvaluateExprInput{Path: input.Path.AppendKey(k).AppendIndex(1), Defs: input.Defs, Expr: v.Arr[1]})
		if or.Status != EvaluateExprOutput_OK {
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
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: operandL.Num - operandR.Num}
		if !isFiniteNumber(v) {
			return errorArithmeticError(input.Path.AppendKey(operator), fmt.Sprintf("%v-%v is not a finite number", operandL.Num, operandR.Num))
		}
		return &EvaluateExprOutput{Value: v}
	case OpBinary_DIV.KeyName():
		if operandL.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(0), []yaml.Type{yaml.Type_TYPE_NUM}, operandL.Type)
		}
		if operandR.Type != yaml.Type_TYPE_NUM {
			return errorUnexpectedType(input.Path.AppendKey(operator).AppendIndex(1), []yaml.Type{yaml.Type_TYPE_NUM}, operandR.Type)
		}
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: operandL.Num / operandR.Num}
		if !isFiniteNumber(v) {
			return errorArithmeticError(input.Path.AppendKey(operator), fmt.Sprintf("%v/%v is not a finite number", operandL.Num, operandR.Num))
		}
		return &EvaluateExprOutput{Value: v}
	case OpBinary_EQ.KeyName():
		return equal(input.Path.AppendKey(operator), operandL, operandR)
	case OpBinary_NEQ.KeyName():
		eq := equal(input.Path.AppendKey(operator), operandL, operandR)
		if eq.Status != EvaluateExprOutput_OK {
			return eq
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: !eq.Value.Bool}}
	case OpBinary_LT.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != EvaluateExprOutput_OK {
			return cmp
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num < 0}}
	case OpBinary_LTE.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != EvaluateExprOutput_OK {
			return cmp
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num <= 0}}
	case OpBinary_GT.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != EvaluateExprOutput_OK {
			return cmp
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num > 0}}
	case OpBinary_GTE.KeyName():
		cmp := compare(input.Path.AppendKey(operator), operandL, operandR)
		if cmp.Status != EvaluateExprOutput_OK {
			return cmp
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp.Value.Num >= 0}}
	}
}

func (i *evaluator) EvaluateOpVariadic(input *EvaluateExprInput) *EvaluateExprOutput {
	var (
		operator string
		operands []*yaml.Value
	)
	for k, v := range input.Expr.Obj { // only one property exists
		operator = k
		for _, e := range v.Arr {
			o := i.EvaluateExpr(&EvaluateExprInput{Path: input.Path.AppendKey(k), Defs: input.Defs, Expr: e})
			if o.Status != EvaluateExprOutput_OK {
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
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: add}
		if !isFiniteNumber(v) {
			v := []string{}
			for _, operand := range operands {
				v = append(v, fmt.Sprintf("%v", operand.Num))
			}
			return errorArithmeticError(input.Path.AppendKey(operator), fmt.Sprintf("add(%v) is not a finite number", strings.Join(v, ",")))
		}
		return &EvaluateExprOutput{Value: v}
	case OpVariadic_MUL.KeyName():
		mul := 1.0
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			mul *= operand.Num
		}
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: mul}
		if !isFiniteNumber(v) {
			v := []string{}
			for _, operand := range operands {
				v = append(v, fmt.Sprintf("%v", operand.Num))
			}
			return errorArithmeticError(input.Path.AppendKey(operator), fmt.Sprintf("mul(%v) is not a finite number", strings.Join(v, ",")))
		}
		return &EvaluateExprOutput{Value: v}
	case OpVariadic_AND.KeyName():
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_BOOL {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Type)
			}
			if !operand.Bool {
				return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: false}}
			}
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: true}}
	case OpVariadic_OR.KeyName():
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_BOOL {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Type)
			}
			if operand.Bool {
				return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: true}}
			}
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: false}}
	case OpVariadic_CAT.KeyName():
		cat := ""
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_STR {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_STR}, operand.Type)
			}
			cat += operand.Str
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_STR, Str: cat}}
	case OpVariadic_MIN.KeyName():
		min_ := math.Inf(1)
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			min_ = math.Min(min_, operand.Num)
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: min_}}
	case OpVariadic_MAX.KeyName():
		max_ := math.Inf(-1)
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return errorUnexpectedType(input.Path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			max_ = math.Max(max_, operand.Num)
		}
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: max_}}
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
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: merge}}
	}
}

func equal(path *Path, l, r *yaml.Value) *EvaluateExprOutput {
	falseValue := &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: false}}
	trueValue := &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: true}}
	switch {
	default:
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_NUM, yaml.Type_TYPE_BOOL, yaml.Type_TYPE_STR, yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, l.Type)
	case l.Type != r.Type:
		return falseValue
	case l.Type == yaml.Type_TYPE_NUM:
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: l.Num == r.Num}}
	case l.Type == yaml.Type_TYPE_BOOL:
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: l.Bool == r.Bool}}
	case l.Type == yaml.Type_TYPE_STR:
		return &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: l.Str == r.Str}}
	case l.Type == yaml.Type_TYPE_ARR:
		if len(l.Arr) != len(r.Arr) {
			return falseValue
		}
		for i, l := range l.Arr {
			r := r.Arr[i]
			eq := equal(path, l, r)
			if eq.Status != EvaluateExprOutput_OK {
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
			if eq.Status != EvaluateExprOutput_OK {
				return eq
			}
			if !eq.Value.Bool {
				return falseValue
			}
		}
		return trueValue
	}
}
func compare(path *Path, l, r *yaml.Value) *EvaluateExprOutput {
	ltValue := &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: -1}}
	gtValue := &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: 1}}
	eqValue := &EvaluateExprOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: 0}}
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
		n := len(l.Arr)
		if n > len(r.Arr) {
			n = len(r.Arr)
		}
		for i := 0; i < n; i++ {
			l, r := l.Arr[i], r.Arr[i]
			cmp := compare(path, l, r)
			if cmp.Status != EvaluateExprOutput_OK {
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
func isFiniteNumber(v *yaml.Value) bool {
	return v.Type == yaml.Type_TYPE_NUM && !math.IsInf(v.Num, 0) && !math.IsNaN(v.Num)
}

func errorUnsupportedExpr(path *Path, v *yaml.Value) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_UNSUPPORTED_EXPR,
		ErrorMessage: fmt.Sprintf("unsupported expr: got %v", v.Keys()),
		ErrorPath:    path,
	}
}
func errorUnexpectedType(path *Path, want []yaml.Type, got yaml.Type) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_UNEXPECTED_TYPE,
		ErrorMessage: fmt.Sprintf("unexpected type: want %v, got %v", want, got),
		ErrorPath:    path,
	}
}
func errorArithmeticError(path *Path, message string) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_ARITHMETIC_ERROR,
		ErrorMessage: fmt.Sprintf("arithmetic error: %v", message),
		ErrorPath:    path,
	}
}
func errorIndexOutOfBounds(path *Path, begin, end, index int) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_INDEX_OUT_OF_BOUNDS,
		ErrorMessage: fmt.Sprintf("index out of bounds: %v not in [%v, %v)", index, begin, end),
		ErrorPath:    path,
	}
}
func errorKeyNotFound(path *Path, want string, actual []string) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_KEY_NOT_FOUND,
		ErrorMessage: fmt.Sprintf("key not found: %q not in {%v}", want, strings.Join(actual, ",")),
		ErrorPath:    path,
	}
}
func errorReferenceNotFound(path *Path, ref string) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_REFERENCE_NOT_FOUND,
		ErrorMessage: fmt.Sprintf("reference not found: %q", ref),
		ErrorPath:    path,
	}
}
func errorCasesNotExhaustive(path *Path) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_CASES_NOT_EXHAUSTIVE,
		ErrorMessage: "cases not exhaustive",
		ErrorPath:    path,
	}
}
func errorUnsupportedOperation(path *Path, gotOp string) *EvaluateExprOutput {
	return &EvaluateExprOutput{
		Status:       EvaluateExprOutput_UNSUPPORTED_OPERATION,
		ErrorMessage: fmt.Sprintf("unsupported operation: %q", gotOp),
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
