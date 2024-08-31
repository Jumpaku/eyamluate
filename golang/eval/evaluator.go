package eval

import (
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/ast"
	"github.com/Jumpaku/eyamlate/golang/yaml"
	"math"
	"slices"
	"strings"
)

type Evaluator interface {
	Evaluate(*EvaluateInput) *EvaluateOutput
	EvaluateExpr(*EvaluateExprInput) *EvaluateOutput
	EvaluateEval(*EvaluateEvalInput) *EvaluateOutput
	EvaluateScalar(*EvaluateScalarInput) *EvaluateOutput
	EvaluateObj(*EvaluateObjInput) *EvaluateOutput
	EvaluateArr(*EvaluateArrInput) *EvaluateOutput
	EvaluateJson(*EvaluateJsonInput) *EvaluateOutput
	EvaluateRangeIter(*EvaluateRangeIterInput) *EvaluateOutput
	EvaluateGetElem(*EvaluateGetElemInput) *EvaluateOutput
	EvaluateFunCall(*EvaluateFunCallInput) *EvaluateOutput
	EvaluateCases(*EvaluateCasesInput) *EvaluateOutput
	EvaluateOpUnary(*EvaluateOpUnaryInput) *EvaluateOutput
	EvaluateOpBinary(*EvaluateOpBinaryInput) *EvaluateOutput
	EvaluateOpVariadic(*EvaluateOpVariadicInput) *EvaluateOutput
}

type BasicEvaluator struct{}

var _ Evaluator = &BasicEvaluator{}

func (e *BasicEvaluator) Evaluate(input *EvaluateInput) *EvaluateOutput {
	return e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: input.Expr})
}

func (e *BasicEvaluator) EvaluateExpr(input *EvaluateExprInput) *EvaluateOutput {
	switch input.Expr.Kind {
	default:
		return errorUnsupportedExpr(input.Expr)
	case ast.Expr_SCALAR:
		return e.EvaluateScalar(&EvaluateScalarInput{Path: input.Expr.Path, Defs: input.Defs, Scalar: input.Expr.Scalar})
	case ast.Expr_OBJ:
		return e.EvaluateObj(&EvaluateObjInput{Path: input.Expr.Path, Defs: input.Defs, Obj: input.Expr.Obj})
	case ast.Expr_ARR:
		return e.EvaluateArr(&EvaluateArrInput{Path: input.Expr.Path, Defs: input.Defs, Arr: input.Expr.Arr})
	case ast.Expr_JSON:
		return e.EvaluateJson(&EvaluateJsonInput{Path: input.Expr.Path, Defs: input.Defs, Json: input.Expr.Json})
	case ast.Expr_RANGE_ITER:
		return e.EvaluateRangeIter(&EvaluateRangeIterInput{Path: input.Expr.Path, Defs: input.Defs, RangeIter: input.Expr.RangeIter})
	case ast.Expr_GET_ELEM:
		return e.EvaluateGetElem(&EvaluateGetElemInput{Path: input.Expr.Path, Defs: input.Defs, GetElem: input.Expr.GetElem})
	case ast.Expr_FUN_CALL:
		return e.EvaluateFunCall(&EvaluateFunCallInput{Path: input.Expr.Path, Defs: input.Defs, FunCall: input.Expr.FunCall})
	case ast.Expr_CASES:
		return e.EvaluateCases(&EvaluateCasesInput{Path: input.Expr.Path, Defs: input.Defs, Cases: input.Expr.Cases})
	case ast.Expr_OP_UNARY:
		return e.EvaluateOpUnary(&EvaluateOpUnaryInput{Path: input.Expr.Path, Defs: input.Defs, OpUnary: input.Expr.OpUnary})
	case ast.Expr_OP_BINARY:
		return e.EvaluateOpBinary(&EvaluateOpBinaryInput{Path: input.Expr.Path, Defs: input.Defs, OpBinary: input.Expr.OpBinary})
	case ast.Expr_OP_VARIADIC:
		return e.EvaluateOpVariadic(&EvaluateOpVariadicInput{Path: input.Expr.Path, Defs: input.Defs, OpVariadic: input.Expr.OpVariadic})
	}
}

func (e *BasicEvaluator) EvaluateEval(input *EvaluateEvalInput) *EvaluateOutput {
	st := input.Defs
	for _, def := range input.Eval.Where {
		st = st.Register(def)
	}
	return e.EvaluateExpr(&EvaluateExprInput{Defs: st, Expr: input.Eval.Eval})
}

func (e *BasicEvaluator) EvaluateScalar(input *EvaluateScalarInput) *EvaluateOutput {
	v := input.Scalar
	switch v.Val.Type {
	default:
		return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM, yaml.Type_BOOL, yaml.Type_STR}, v.Val.Type)
	case yaml.Type_NUM, yaml.Type_BOOL, yaml.Type_STR:
		return &EvaluateOutput{Value: v.Val}
	}
}

func (e *BasicEvaluator) EvaluateObj(input *EvaluateObjInput) *EvaluateOutput {
	v := input.Obj
	obj := &yaml.Value{Type: yaml.Type_OBJ, Obj: map[string]*yaml.Value{}}
	for k, val := range v.Obj {
		expr := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: val})
		if expr.Status != EvaluateOutput_OK {
			return expr
		}
		obj.Obj[k] = expr.Value
	}
	return &EvaluateOutput{Value: obj}
}

func (e *BasicEvaluator) EvaluateArr(input *EvaluateArrInput) *EvaluateOutput {
	v := input.Arr
	arr := &yaml.Value{Type: yaml.Type_ARR}
	for _, val := range v.Arr {
		expr := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: val})
		if expr.Status != EvaluateOutput_OK {
			return expr
		}
		arr.Arr = append(arr.Arr, expr.Value)
	}
	return &EvaluateOutput{Value: arr}
}

func (e *BasicEvaluator) EvaluateJson(input *EvaluateJsonInput) *EvaluateOutput {
	return &EvaluateOutput{Value: input.Json.Json}
}

func (e *BasicEvaluator) EvaluateRangeIter(input *EvaluateRangeIterInput) *EvaluateOutput {
	in := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: input.RangeIter.In})
	if in.Status != EvaluateOutput_OK {
		return in
	}
	switch in.Value.Type {
	default:
		return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_ARR, yaml.Type_OBJ}, in.Value.Type)
	case yaml.Type_ARR:
		arr := &yaml.Value{Type: yaml.Type_ARR}
		for i, elem := range in.Value.Arr {
			st := input.Defs
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForPos,
				Value: exprOfValue(&yaml.Value{Type: yaml.Type_NUM, Num: float64(i)}),
			})
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForVal,
				Value: exprOfValue(elem),
			})
			if_ := e.EvaluateExpr(&EvaluateExprInput{Defs: st, Expr: input.RangeIter.If})
			if if_.Status != EvaluateOutput_OK {
				return if_
			}
			if if_.Value.Bool {
				do := e.EvaluateExpr(&EvaluateExprInput{Defs: st, Expr: input.RangeIter.Do})
				if do.Status != EvaluateOutput_OK {
					return do
				}
				arr.Arr = append(arr.Arr, do.Value)
			}
		}
		return &EvaluateOutput{Value: arr}
	case yaml.Type_OBJ:
		obj := &yaml.Value{Type: yaml.Type_OBJ, Obj: map[string]*yaml.Value{}}
		for k, elem := range in.Value.Obj {
			st := input.Defs
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForPos,
				Value: exprOfValue(&yaml.Value{Type: yaml.Type_STR, Str: k}),
			})
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForVal,
				Value: exprOfValue(elem),
			})
			if_ := e.EvaluateExpr(&EvaluateExprInput{Defs: st, Expr: input.RangeIter.If})
			if if_.Status != EvaluateOutput_OK {
				return if_
			}
			if if_.Value.Bool {
				do := e.EvaluateExpr(&EvaluateExprInput{Defs: st, Expr: input.RangeIter.Do})
				if do.Status != EvaluateOutput_OK {
					return do
				}
				obj.Obj[k] = do.Value
			}
		}
		return &EvaluateOutput{Value: obj}
	}
}

func (e *BasicEvaluator) EvaluateGetElem(input *EvaluateGetElemInput) *EvaluateOutput {
	get := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: input.GetElem.Get})
	if get.Status != EvaluateOutput_OK {
		return get
	}
	from := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: input.GetElem.From})
	if from.Status != EvaluateOutput_OK {
		return from
	}

	switch from.Value.Type {
	default:
		return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_ARR, yaml.Type_OBJ}, from.Value.Type)
	case yaml.Type_ARR:
		if !get.Value.CanInt() {
			return errorArithmeticError(input.Path, "index for an array must be integer")
		}
		index := int(get.Value.Num)
		if index < 0 || index >= len(from.Value.Arr) {
			return errorIndexOutOfBounds(input.Path, 0, len(from.Value.Arr), index)
		}
		return &EvaluateOutput{Value: from.Value.Arr[index]}
	case yaml.Type_OBJ:
		if get.Value.Type != yaml.Type_STR {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_STR}, get.Value.Type)
		}
		key := get.Value.Str
		val, ok := from.Value.Obj[key]
		if !ok {
			return errorKeyNotFound(input.Path, from.Value.Keys(), key)
		}
		return &EvaluateOutput{Value: val}
	}
}

func (e *BasicEvaluator) EvaluateFunCall(input *EvaluateFunCallInput) *EvaluateOutput {
	ref := input.FunCall.Ref
	st := input.Defs.Find(ref)
	if st == nil {
		return errorReferenceNotFound(input.Path, ref)
	}
	for _, defArg := range st.Def.With {
		if _, ok := input.FunCall.With[defArg]; !ok {
			return errorKeyNotFound(input.Path, st.Def.With, defArg)
		}
	}
	for k, v := range input.FunCall.With {
		arg := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: v})
		if arg.Status != EvaluateOutput_OK {
			return arg
		}
		st = st.Register(&ast.FunDef{
			Def:   k,
			Value: exprOfValue(arg.Value),
		})
	}
	return e.EvaluateExpr(&EvaluateExprInput{Defs: st, Expr: st.Def.Value})
}

func (e *BasicEvaluator) EvaluateCases(input *EvaluateCasesInput) *EvaluateOutput {
	for _, branch := range input.Cases.Branches {
		if branch.IsOtherwise {
			return e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: branch.Otherwise})
		}
		when := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: branch.When})
		if when.Status != EvaluateOutput_OK {
			return when
		}
		if !when.Value.Bool {
			continue
		}
		return e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: branch.Then})
	}
	return errorCasesNotExhaustive(input.Path)
}

func (e *BasicEvaluator) EvaluateOpUnary(input *EvaluateOpUnaryInput) *EvaluateOutput {
	op := input.OpUnary
	operand := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: op.Operand})
	if operand.Status != EvaluateOutput_OK {
		return operand
	}
	switch op.Operator {
	default:
		return errorUnsupportedOperation(input.Path, op.Operator.String())
	case ast.OpUnary_NOT:
		if operand.Value.Type != yaml.Type_BOOL {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_BOOL}, operand.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: !operand.Value.Bool},
		}
	case ast.OpUnary_LEN:
		switch operand.Value.Type {
		default:
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_ARR, yaml.Type_OBJ}, operand.Value.Type)
		case yaml.Type_ARR:
			return &EvaluateOutput{
				Value: &yaml.Value{Type: yaml.Type_NUM, Num: float64(len(operand.Value.Arr))},
			}
		case yaml.Type_OBJ:
			return &EvaluateOutput{
				Value: &yaml.Value{Type: yaml.Type_NUM, Num: float64(len(operand.Value.Obj))},
			}
		}
	case ast.OpUnary_FLOOR:
		if operand.Value.Type != yaml.Type_NUM {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_NUM, Num: math.Floor(operand.Value.Num)},
		}
	case ast.OpUnary_CEIL:
		if operand.Value.Type != yaml.Type_NUM {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_NUM, Num: math.Ceil(operand.Value.Num)},
		}
	case ast.OpUnary_FLAT:
		if operand.Value.Type != yaml.Type_ARR {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_ARR}, operand.Value.Type)
		}
		flat := []*yaml.Value{}
		for _, elem := range operand.Value.Arr {
			if elem.Type != yaml.Type_ARR {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_ARR}, elem.Type)
			}
			flat = append(flat, elem.Arr...)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_ARR, Arr: flat},
		}
	case ast.OpUnary_ABORT:
		if operand.Value.Type != yaml.Type_STR {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_STR}, operand.Value.Type)
		}
		return &EvaluateOutput{
			Status:       EvaluateOutput_ABORT_ERROR,
			ErrorMessage: operand.Value.Str,
			ErrorPath:    input.Path,
		}
	}
}

func (e *BasicEvaluator) EvaluateOpBinary(input *EvaluateOpBinaryInput) *EvaluateOutput {
	op := input.OpBinary
	operandLeft := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: op.OperandLeft})
	if operandLeft.Status != EvaluateOutput_OK {
		return operandLeft
	}
	operandRight := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: op.OperandRight})
	if operandRight.Status != EvaluateOutput_OK {
		return operandRight
	}
	switch op.Operator {
	default:
		return errorUnsupportedOperation(input.Path, op.Operator.String())
	case ast.OpBinary_SUB:
		if operandLeft.Value.Type != yaml.Type_NUM {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operandLeft.Value.Type)
		}
		if operandRight.Value.Type != yaml.Type_NUM {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operandRight.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_NUM,
				Num:  operandLeft.Value.Num - operandRight.Value.Num,
			},
		}
	case ast.OpBinary_DIV:
		if operandLeft.Value.Type != yaml.Type_NUM {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operandLeft.Value.Type)
		}
		if operandRight.Value.Type != yaml.Type_NUM {
			return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operandRight.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_NUM,
				Num:  operandLeft.Value.Num / operandRight.Value.Num,
			},
		}
	case ast.OpBinary_MOD:
		if !operandLeft.Value.CanInt() {
			return errorArithmeticError(input.Path, "first operand must be integer")
		}
		if !operandRight.Value.CanInt() {
			return errorArithmeticError(input.Path, "second operand must be integer")
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_NUM,
				Num:  float64(int64(operandLeft.Value.Num) % int64(operandRight.Value.Num)),
			},
		}
	case ast.OpBinary_EQ:
		return equal(input.Path, operandLeft.Value, operandRight.Value)
	case ast.OpBinary_NEQ:
		eq := equal(input.Path, operandLeft.Value, operandRight.Value)
		if eq.Status != EvaluateOutput_OK {
			return eq
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: !eq.Value.Bool},
		}
	case ast.OpBinary_LT:
		cmp := compare(input.Path, operandLeft.Value, operandRight.Value)
		if cmp.Status != EvaluateOutput_OK {
			return cmp
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: cmp.Value.Num < 0},
		}
	case ast.OpBinary_LTE:
		cmp := compare(input.Path, operandLeft.Value, operandRight.Value)
		if cmp.Status != EvaluateOutput_OK {
			return cmp
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: cmp.Value.Num <= 0},
		}
	case ast.OpBinary_GT:
		cmp := compare(input.Path, operandLeft.Value, operandRight.Value)
		if cmp.Status != EvaluateOutput_OK {
			return cmp
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: cmp.Value.Num > 0},
		}
	case ast.OpBinary_GTE:
		cmp := compare(input.Path, operandLeft.Value, operandRight.Value)
		if cmp.Status != EvaluateOutput_OK {
			return cmp
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: cmp.Value.Num >= 0},
		}
	}
}

func (e *BasicEvaluator) EvaluateOpVariadic(input *EvaluateOpVariadicInput) *EvaluateOutput {
	op := input.OpVariadic
	operands := []*yaml.Value{}
	for _, expr := range op.Operands {
		operand := e.EvaluateExpr(&EvaluateExprInput{Defs: input.Defs, Expr: expr})
		if operand.Status != EvaluateOutput_OK {
			return operand
		}
		operands = append(operands, operand.Value)
	}
	switch op.Operator {
	default:
		return errorUnsupportedOperation(input.Path, op.Operator.String())
	case ast.OpVariadic_ADD:
		v := &yaml.Value{Type: yaml.Type_NUM, Num: 0}
		for _, operand := range operands {
			if operand.Type != yaml.Type_NUM {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Type)
			}
			v.Num += operand.Num
		}
		return &EvaluateOutput{Value: v}
	case ast.OpVariadic_MUL:
		v := &yaml.Value{Type: yaml.Type_NUM, Num: 1}
		for _, operand := range operands {
			if operand.Type != yaml.Type_NUM {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Type)
			}
			v.Num *= operand.Num
		}
		return &EvaluateOutput{Value: v}
	case ast.OpVariadic_AND:
		v := &yaml.Value{Type: yaml.Type_BOOL, Bool: false}
		for _, operand := range operands {
			if operand.Type != yaml.Type_BOOL {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Type)
			}
			v.Bool = v.Bool && operand.Bool
		}
		return &EvaluateOutput{Value: v}
	case ast.OpVariadic_OR:
		v := &yaml.Value{Type: yaml.Type_BOOL, Bool: true}
		for _, operand := range operands {
			if operand.Type != yaml.Type_BOOL {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Type)
			}
			v.Bool = v.Bool || operand.Bool
		}
		return &EvaluateOutput{Value: v}
	case ast.OpVariadic_CAT:
		if len(operands) == 0 {
			return &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_STR, Str: ""}}
		}
		cat := &yaml.Value{Type: yaml.Type_STR, Str: ""}
		for _, o := range operands {
			if o.Type != yaml.Type_STR {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_STR}, o.Type)
			}
			cat.Str += o.Str
		}
		return &EvaluateOutput{Value: cat}
	case ast.OpVariadic_MAX:
		v := &yaml.Value{Type: yaml.Type_NUM, Num: math.Inf(-1)}
		for _, operand := range operands {
			if operand.Type != yaml.Type_NUM {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Type)
			}
			if v.Num < operand.Num {
				v.Num = operand.Num
			}
		}
		return &EvaluateOutput{Value: v}
	case ast.OpVariadic_MIN:
		v := &yaml.Value{Type: yaml.Type_NUM, Num: math.Inf(1)}
		for _, operand := range operands {
			if operand.Type != yaml.Type_NUM {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_NUM}, operand.Type)
			}
			if v.Num > operand.Num {
				v.Num = operand.Num
			}
		}
		return &EvaluateOutput{Value: v}
	case ast.OpVariadic_MERGE:
		v := &yaml.Value{Type: yaml.Type_OBJ, Obj: map[string]*yaml.Value{}}
		for _, o := range operands {
			if o.Type != yaml.Type_OBJ {
				return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_OBJ}, o.Type)
			}
			for k, val := range o.Obj {
				v.Obj[k] = val
			}
		}
		return &EvaluateOutput{Value: v}
	}
}

func equal(path *ast.Path, l, r *yaml.Value) *EvaluateOutput {
	falseValue := &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: false}}
	trueValue := &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: true}}
	switch {
	default:
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_NUM, yaml.Type_BOOL, yaml.Type_STR, yaml.Type_ARR, yaml.Type_OBJ}, l.Type)
	case l.Type != r.Type:
		return falseValue
	case l.Type == yaml.Type_NUM:
		return &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: l.Num == r.Num}}
	case l.Type == yaml.Type_BOOL:
		return &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: l.Bool == r.Bool}}
	case l.Type == yaml.Type_STR:
		return &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_BOOL, Bool: l.Str == r.Str}}
	case l.Type == yaml.Type_ARR:
		if len(l.Arr) != len(r.Arr) {
			return falseValue
		}
		for i, l := range l.Arr {
			r := r.Arr[i]
			eq := equal(path, l, r)
			if eq.Status != EvaluateOutput_OK {
				return eq
			}
			if !eq.Value.Bool {
				return falseValue
			}
		}
		return trueValue
	case l.Type == yaml.Type_OBJ:
		lk, rk := l.Keys(), r.Keys()
		slices.Sort(lk)
		slices.Sort(rk)
		if !slices.Equal(lk, rk) {
			return falseValue
		}
		for k, l := range l.Obj {
			r := r.Obj[k]
			eq := equal(path, l, r)
			if eq.Status != EvaluateOutput_OK {
				return eq
			}
			if !eq.Value.Bool {
				return falseValue
			}
		}
		return trueValue
	}
}
func compare(path *ast.Path, l, r *yaml.Value) *EvaluateOutput {
	ltValue := &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_NUM, Num: -1}}
	gtValue := &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_NUM, Num: 1}}
	eqValue := &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_NUM, Num: 0}}
	switch {
	default:
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_NUM, yaml.Type_BOOL, yaml.Type_STR, yaml.Type_ARR}, l.Type)
	case l.Type == yaml.Type_NUM && r.Type == yaml.Type_NUM:
		if l.Num < r.Num {
			return ltValue
		}
		if l.Num > r.Num {
			return gtValue
		}
		return eqValue
	case l.Type == yaml.Type_BOOL && r.Type == yaml.Type_BOOL:
		if !l.Bool && r.Bool {
			return ltValue
		}
		if l.Bool && !r.Bool {
			return gtValue
		}
		return eqValue
	case l.Type == yaml.Type_STR && r.Type == yaml.Type_STR:
		if l.Str < r.Str {
			return ltValue
		}
		if l.Str > r.Str {
			return gtValue
		}
		return eqValue
	case l.Type == yaml.Type_ARR && r.Type == yaml.Type_ARR:
		for i, l := range l.Arr {
			r := r.Arr[i]
			cmp := compare(path, l, r)
			if cmp.Status != EvaluateOutput_OK {
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
func exprOfValue(v *yaml.Value) *ast.Expr {
	return &ast.Expr{Kind: ast.Expr_JSON, Json: &ast.Json{Json: v}}
}
func errorUnsupportedExpr(expr *ast.Expr) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_UNSUPPORTED_EXPR_ERROR,
		ErrorMessage: fmt.Sprintf("unsupported expr kind: %v", expr.Kind),
		ErrorPath:    expr.Path,
	}
}
func errorUnexpectedType(path *ast.Path, want []yaml.Type, got yaml.Type) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_UNEXPECTED_TYPE_ERROR,
		ErrorMessage: fmt.Sprintf("unexpected type: want %v, got %v", want, got),
		ErrorPath:    path,
	}
}
func errorArithmeticError(path *ast.Path, message string) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_ARITHMETIC_ERROR,
		ErrorMessage: fmt.Sprintf("arithmetic error: %v", message),
		ErrorPath:    path,
	}
}
func errorIndexOutOfBounds(path *ast.Path, begin, end, index int) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_INDEX_OUT_OF_BOUNDS_ERROR,
		ErrorMessage: fmt.Sprintf("index out of bounds: %v not in [%v, %v)", index, begin, end),
		ErrorPath:    path,
	}
}
func errorKeyNotFound(path *ast.Path, keys []string, key string) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_KEY_NOT_FOUND_ERROR,
		ErrorMessage: fmt.Sprintf("key not found: %v not in {%v}", key, strings.Join(keys, ",")),
		ErrorPath:    path,
	}
}
func errorReferenceNotFound(path *ast.Path, ref string) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_REFERENCE_NOT_FOUND_ERROR,
		ErrorMessage: fmt.Sprintf("reference not found: %v", ref),
		ErrorPath:    path,
	}
}
func errorCasesNotExhaustive(path *ast.Path) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_CASES_NOT_EXHAUSTIVE_ERROR,
		ErrorMessage: fmt.Sprintf("cases not exhaustive: %v", path.Format()),
		ErrorPath:    path,
	}
}
func errorUnsupportedOperation(path *ast.Path, op string) *EvaluateOutput {
	return &EvaluateOutput{
		Status:       EvaluateOutput_UNSUPPORTED_OPERATION_ERROR,
		ErrorMessage: fmt.Sprintf("unsupported operation: %v", op),
		ErrorPath:    path,
	}
}
