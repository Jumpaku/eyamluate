package eval

import (
	"context"
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/pb/ast"
	"github.com/Jumpaku/eyamlate/golang/pb/yaml"
	"math"
)

type Evaluator struct{}

var _ EvaluatorServer = (*Evaluator)(nil)

func (e *Evaluator) mustEmbedUnimplementedEvaluatorServer() {
	//TODO implement me
	panic("implement me")
}

func (e *Evaluator) Evaluate(ctx context.Context, input *EvaluateInput) (*EvaluateOutput, error) {
	return e.EvaluateEval(ctx, &EvaluateEvalInput{Defs: input.Defs, Eval: input.Eval})
}

func (e *Evaluator) EvaluateExpr(ctx context.Context, input *EvaluateExprInput) (*EvaluateOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (e *Evaluator) EvaluateEval(ctx context.Context, input *EvaluateEvalInput) (*EvaluateOutput, error) {
	st := input.Defs
	for _, def := range input.Eval.Where {
		st = st.Register(def)
	}
	return e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: st, Expr: input.Eval.Eval})
}

func (e *Evaluator) EvaluateScalar(ctx context.Context, input *EvaluateScalarInput) (*EvaluateOutput, error) {
	v := input.Scalar
	path := v.Path
	switch v.Val.Type {
	default:
		return nil, ast.ErrUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_NUM, yaml.Type_TYPE_BOOL, yaml.Type_TYPE_STR}, v.Val.Type)
	case yaml.Type_TYPE_NUM, yaml.Type_TYPE_BOOL, yaml.Type_TYPE_STR:
		return &EvaluateOutput{Value: v.Val}, nil
	}
}

func (e *Evaluator) EvaluateNewObj(ctx context.Context, input *EvaluateNewObjInput) (*EvaluateOutput, error) {
	v := input.NewObj
	obj := &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: map[string]*yaml.Value{}}
	for k, val := range v.Obj {
		r, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: val})
		if err != nil {
			return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
		}
		obj.Obj[k] = r.Value
	}
	return &EvaluateOutput{Value: obj}, nil
}

func (e *Evaluator) EvaluateNewArr(ctx context.Context, input *EvaluateNewArrInput) (*EvaluateOutput, error) {
	v := input.NewArr
	arr := &yaml.Value{Type: yaml.Type_TYPE_ARR}
	for _, val := range v.Arr {
		r, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: val})
		if err != nil {
			return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
		}
		arr.Arr = append(arr.Arr, r.Value)
	}
	return &EvaluateOutput{Value: arr}, nil
}

func (e *Evaluator) EvaluateValJson(ctx context.Context, input *EvaluateValJsonInput) (*EvaluateOutput, error) {
	return &EvaluateOutput{Value: input.ValJson.Json}, nil
}

func (e *Evaluator) EvaluateRangeIter(ctx context.Context, input *EvaluateRangeIterInput) (*EvaluateOutput, error) {
	in, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: input.RangeIter.In})
	if err != nil {
		return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
	}
	switch in.Value.Type {
	default:
		return nil, ast.ErrUnexpectedType(input.RangeIter.Path, []yaml.Type{yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, in.Value.Type)
	case yaml.Type_TYPE_ARR:
		arr := &yaml.Value{Type: yaml.Type_TYPE_ARR}
		for i, elem := range in.Value.Arr {
			st := input.Defs
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForPos,
				Value: ExprOfValue(&yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(i)}),
			})
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForVal,
				Value: ExprOfValue(elem),
			})
			if_, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: st, Expr: input.RangeIter.If})
			if err != nil {
				return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
			}
			if if_.Value.Bool {
				do, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: st, Expr: input.RangeIter.Do})
				if err != nil {
					return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
				}
				arr.Arr = append(arr.Arr, do.Value)
			}
		}
		return &EvaluateOutput{Value: arr}, nil
	case yaml.Type_TYPE_OBJ:
		obj := &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: map[string]*yaml.Value{}}
		for k, elem := range in.Value.Obj {
			st := input.Defs
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForPos,
				Value: ExprOfValue(&yaml.Value{Type: yaml.Type_TYPE_STR, Str: k}),
			})
			st = st.Register(&ast.FunDef{
				Def:   input.RangeIter.ForVal,
				Value: ExprOfValue(elem),
			})
			if_, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: st, Expr: input.RangeIter.If})
			if err != nil {
				return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
			}
			if if_.Value.Bool {
				do, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: st, Expr: input.RangeIter.Do})
				if err != nil {
					return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
				}
				obj.Obj[k] = do.Value
			}
		}
		return &EvaluateOutput{Value: obj}, nil
	}
}
func ErrIndexOutOfBounds(path *ast.Path, wantBegin, wantEnd int, index int) error {
	return fmt.Errorf("out of bounds %v: want [%v, %v): got %v", path.Format(), wantBegin, wantEnd, index)
}
func (e *Evaluator) EvaluateElemAccess(ctx context.Context, input *EvaluateElemAccessInput) (*EvaluateOutput, error) {
	get, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: input.ElemAccess.Get})
	if err != nil {
		return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
	}
	from, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: input.ElemAccess.From})
	if err != nil {
		return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
	}

	switch from.Value.Type {
	default:
		return nil, ast.ErrUnexpectedType(input.ElemAccess.Path, []yaml.Type{yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, from.Value.Type)
	case yaml.Type_TYPE_ARR:
		if !get.Value.CanInt() {
			return nil, ast.ErrUnexpectedType(input.ElemAccess.Path, []yaml.Type{yaml.Type_TYPE_NUM}, get.Value.Type)
		}
		index := int(get.Value.Num)
		if index < 0 || index >= len(from.Value.Arr) {
			return nil, ErrIndexOutOfBounds(input.ElemAccess.Path, 0, len(from.Value.Arr), index)
		}
		return &EvaluateOutput{Value: from.Value.Arr[index]}, nil
	case yaml.Type_TYPE_OBJ:
		if get.Value.Type != yaml.Type_TYPE_STR {
			return nil, ast.ErrUnexpectedType(input.ElemAccess.Path, []yaml.Type{yaml.Type_TYPE_STR}, get.Value.Type)
		}
		key := get.Value.Str
		val, ok := from.Value.Obj[key]
		if !ok {
			return nil, ast.ErrKeyNotFound(input.ElemAccess.Path, []string{key})
		}
		return &EvaluateOutput{Value: val}, nil
	}
}

func ErrReferenceNotFound(path *ast.Path, ref string) error {
	return fmt.Errorf("reference not found: %v: %v", path.Format(), ref)
}
func (e *Evaluator) EvaluateFunCall(ctx context.Context, input *EvaluateFunCallInput) (*EvaluateOutput, error) {
	ref := input.FunCall.Ref
	st := input.Defs.Find(ref)
	if st == nil {
		return nil, ErrReferenceNotFound(input.FunCall.Path, ref)
	}
	for k, v := range input.FunCall.With {
		arg, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: v})
		if err != nil {
			return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
		}
		st = st.Register(&ast.FunDef{
			Def:   k,
			Value: ExprOfValue(arg.Value),
		})
	}
	return e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: st, Expr: st.Def.Value})
}

func (e *Evaluator) EvaluateCaseBranches(ctx context.Context, input *EvaluateCaseBranchesInput) (*EvaluateOutput, error) {
	for _, branch := range input.CaseBranches.Branches {
		if branch.IsOtherwise {
			return e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: branch.Otherwise})
		}
		when, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: branch.When})
		if err != nil {
			return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
		}
		if !when.Value.Bool {
			continue
		}
		return e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: branch.Then})
	}
	return nil, fmt.Errorf("unexhaustive cases: %v", input.CaseBranches.Path.Format())
}

func ErrUnsupportedOperatorUnary(path *ast.Path, op ast.OpUnary_Operator) error {
	return fmt.Errorf("unsupported unary operator %v: %v", path.Format(), op)
}
func ErrUnsupportedOperatorBinary(path *ast.Path, op ast.OpBinary_Operator) error {
	return fmt.Errorf("unsupported unary operator %v: %v", path.Format(), op)
}
func ErrUnsupportedOperatorVariadic(path *ast.Path, op ast.OpVariadic_Operator) error {
	return fmt.Errorf("unsupported unary operator %v: %v", path.Format(), op)
}
func (e *Evaluator) EvaluateOpUnary(ctx context.Context, input *EvaluateOpUnaryInput) (*EvaluateOutput, error) {
	op := input.OpUnary
	operand, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: op.Operand})
	if err != nil {
		return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
	}
	switch op.Operator {
	default:
		return nil, ErrUnsupportedOperatorUnary(op.Path, op.Operator)
	case ast.OpUnary_OPERATOR_NOT:
		if operand.Value.Type != yaml.Type_TYPE_BOOL {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_BOOL}, operand.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: !operand.Value.Bool},
		}, nil
	case ast.OpUnary_OPERATOR_LEN:
		switch operand.Value.Type {
		default:
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, operand.Value.Type)
		case yaml.Type_TYPE_ARR:
			return &EvaluateOutput{
				Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Value.Arr))},
			}, nil
		case yaml.Type_TYPE_OBJ:
			return &EvaluateOutput{
				Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(len(operand.Value.Obj))},
			}, nil
		}
	case ast.OpUnary_OPERATOR_KEYS:
		switch operand.Value.Type {
		default:
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR, yaml.Type_TYPE_OBJ}, operand.Value.Type)
		case yaml.Type_TYPE_ARR:
			keys := []*yaml.Value{}
			for i := range operand.Value.Arr {
				keys = append(keys, &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(i)})
			}
			return &EvaluateOutput{
				Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: keys},
			}, nil
		case yaml.Type_TYPE_OBJ:
			keys := []*yaml.Value{}
			for k := range operand.Value.Obj {
				keys = append(keys, &yaml.Value{Type: yaml.Type_TYPE_STR, Str: k})
			}
			return &EvaluateOutput{
				Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: keys},
			}, nil
		}
	case ast.OpUnary_OPERATOR_FLAT:
		if operand.Value.Type != yaml.Type_TYPE_ARR {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR}, operand.Value.Type)
		}
		flat := []*yaml.Value{}
		for _, elem := range operand.Value.Arr {
			if elem.Type != yaml.Type_TYPE_ARR {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR}, elem.Type)
			}
			flat = append(flat, elem.Arr...)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_ARR, Arr: flat},
		}, nil
	case ast.OpUnary_OPERATOR_HEAD:
		if operand.Value.Type != yaml.Type_TYPE_ARR {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR}, operand.Value.Type)
		}
		if len(operand.Value.Arr) == 0 {
			return nil, ErrIndexOutOfBounds(op.Path, 0, 1, 0)
		}
		return &EvaluateOutput{Value: operand.Value.Arr[0]}, nil
	case ast.OpUnary_OPERATOR_TAIL:
		if operand.Value.Type != yaml.Type_TYPE_ARR {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR}, operand.Value.Type)
		}
		if len(operand.Value.Arr) == 0 {
			return nil, ErrIndexOutOfBounds(op.Path, 0, 1, 0)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_TYPE_ARR,
				Arr:  operand.Value.Arr[1:],
			},
		}, nil
	case ast.OpUnary_OPERATOR_INIT:
		if operand.Value.Type != yaml.Type_TYPE_ARR {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR}, operand.Value.Type)
		}
		if len(operand.Value.Arr) == 0 {
			return nil, ErrIndexOutOfBounds(op.Path, 0, 1, 0)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_TYPE_ARR,
				Arr:  operand.Value.Arr[:len(operand.Value.Arr)-1],
			},
		}, nil
	case ast.OpUnary_OPERATOR_LAST:
		if operand.Value.Type != yaml.Type_TYPE_ARR {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_ARR}, operand.Value.Type)
		}
		if len(operand.Value.Arr) == 0 {
			return nil, ErrIndexOutOfBounds(op.Path, 0, 1, 0)
		}
		return &EvaluateOutput{Value: operand.Value.Arr[len(operand.Value.Arr)-1]}, nil
	case ast.OpUnary_OPERATOR_ERROR:
		if operand.Value.Type != yaml.Type_TYPE_STR {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_STR}, operand.Value.Type)
		}
		return nil, fmt.Errorf("error: %v", operand.Value.Str)
	}
}

func (e *Evaluator) EvaluateOpBinary(ctx context.Context, input *EvaluateOpBinaryInput) (*EvaluateOutput, error) {
	op := input.OpBinary
	operandLeft, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: op.OperandLeft})
	if err != nil {
		return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
	}
	operandRight, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: op.OperandRight})
	if err != nil {
		return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
	}
	switch op.Operator {
	default:
		return nil, ErrUnsupportedOperatorBinary(op.Path, op.Operator)
	case ast.OpBinary_OPERATOR_SUB:
		if operandLeft.Value.Type != yaml.Type_TYPE_NUM {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operandLeft.Value.Type)
		}
		if operandRight.Value.Type != yaml.Type_TYPE_NUM {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operandRight.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_TYPE_NUM,
				Num:  operandLeft.Value.Num - operandRight.Value.Num,
			},
		}, nil
	case ast.OpBinary_OPERATOR_DIV:
		if operandLeft.Value.Type != yaml.Type_TYPE_NUM {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operandLeft.Value.Type)
		}
		if operandRight.Value.Type != yaml.Type_TYPE_NUM {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operandRight.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_TYPE_NUM,
				Num:  operandLeft.Value.Num / operandRight.Value.Num,
			},
		}, nil
	case ast.OpBinary_OPERATOR_MOD:
		if !operandLeft.Value.CanInt() {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operandLeft.Value.Type)
		}
		if !operandRight.Value.CanInt() {
			return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operandRight.Value.Type)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{
				Type: yaml.Type_TYPE_NUM,
				Num:  float64(int64(operandLeft.Value.Num) % int64(operandRight.Value.Num)),
			},
		}, nil
	case ast.OpBinary_OPERATOR_EQ:
		eq, err := equal(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to equal: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: eq},
		}, nil
	case ast.OpBinary_OPERATOR_NEQ:
		eq, err := equal(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to equal: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: !eq},
		}, nil
	case ast.OpBinary_OPERATOR_LT:
		cmp, err := compare(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to compare: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp < 0},
		}, nil
	case ast.OpBinary_OPERATOR_LTE:
		cmp, err := compare(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to compare: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp <= 0},
		}, nil
	case ast.OpBinary_OPERATOR_GT:
		cmp, err := compare(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to compare: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp > 0},
		}, nil
	case ast.OpBinary_OPERATOR_GTE:
		cmp, err := compare(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to compare: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: cmp >= 0},
		}, nil
	case ast.OpBinary_OPERATOR_CMP:
		cmp, err := compare(op.Path, operandLeft.Value, operandRight.Value)
		if err != nil {
			return nil, fmt.Errorf("fail to compare: %w", err)
		}
		return &EvaluateOutput{
			Value: &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: float64(cmp)},
		}, nil
	}
}

func (e *Evaluator) EvaluateOpVariadic(ctx context.Context, input *EvaluateOpVariadicInput) (*EvaluateOutput, error) {
	op := input.OpVariadic
	operands := []*yaml.Value{}
	for _, expr := range op.Operands {
		operand, err := e.EvaluateExpr(ctx, &EvaluateExprInput{Defs: input.Defs, Expr: expr})
		if err != nil {
			return nil, fmt.Errorf("fail to EvaluateExpr: %w", err)
		}
		operands = append(operands, operand.Value)
	}
	switch op.Operator {
	default:
		return nil, ErrUnsupportedOperatorVariadic(op.Path, op.Operator)
	case ast.OpVariadic_OPERATOR_ADD:
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: 0}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			v.Num += operand.Num
		}
		return &EvaluateOutput{Value: v}, nil
	case ast.OpVariadic_OPERATOR_MUL:
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: 1}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			v.Num *= operand.Num
		}
		return &EvaluateOutput{Value: v}, nil
	case ast.OpVariadic_OPERATOR_AND:
		v := &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: false}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_BOOL {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			v.Bool = v.Bool && operand.Bool
		}
		return &EvaluateOutput{Value: v}, nil
	case ast.OpVariadic_OPERATOR_OR:
		v := &yaml.Value{Type: yaml.Type_TYPE_BOOL, Bool: true}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_BOOL {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			v.Bool = v.Bool || operand.Bool
		}
		return &EvaluateOutput{Value: v}, nil
	case ast.OpVariadic_OPERATOR_CAT:
		if len(operands) == 0 {
			return &EvaluateOutput{Value: &yaml.Value{Type: yaml.Type_TYPE_STR, Str: ""}}, nil
		}
		cat := &yaml.Value{Type: yaml.Type_TYPE_STR, Str: ""}
		for _, o := range operands {
			if o.Type != yaml.Type_TYPE_STR {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_STR}, o.Type)
			}
			cat.Str += o.Str
		}
		return &EvaluateOutput{Value: cat}, nil
	case ast.OpVariadic_OPERATOR_MAX:
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: math.Inf(-1)}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			if v.Num < operand.Num {
				v.Num = operand.Num
			}
		}
		return &EvaluateOutput{Value: v}, nil
	case ast.OpVariadic_OPERATOR_MIN:
		v := &yaml.Value{Type: yaml.Type_TYPE_NUM, Num: math.Inf(1)}
		for _, operand := range operands {
			if operand.Type != yaml.Type_TYPE_NUM {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_NUM}, operand.Type)
			}
			if v.Num > operand.Num {
				v.Num = operand.Num
			}
		}
		return &EvaluateOutput{Value: v}, nil
	case ast.OpVariadic_OPERATOR_MERGE:
		v := &yaml.Value{Type: yaml.Type_TYPE_OBJ, Obj: map[string]*yaml.Value{}}
		for _, o := range operands {
			if o.Type != yaml.Type_TYPE_OBJ {
				return nil, ast.ErrUnexpectedType(op.Path, []yaml.Type{yaml.Type_TYPE_OBJ}, o.Type)
			}
			for k, val := range o.Obj {
				v.Obj[k] = val
			}
		}
		return &EvaluateOutput{Value: v}, nil
	}
}

func equal(path *ast.Path, l, r *yaml.Value) (bool, error) {
	switch {
	default:
		return false, fmt.Errorf("unsupported eq types: %v: %v and %v", path.Format(), l.Type, r.Type)
	case l.Type != r.Type:
		return false, nil
	case l.Type == yaml.Type_TYPE_NUM:
		return l.Num == r.Num, nil
	case l.Type == yaml.Type_TYPE_BOOL:
		return l.Bool == r.Bool, nil
	case l.Type == yaml.Type_TYPE_STR:
		return l.Str == r.Str, nil
	case l.Type == yaml.Type_TYPE_ARR:
		if len(l.Arr) != len(r.Arr) {
			return false, nil
		}
		for i, l := range l.Arr {
			r := r.Arr[i]
			eq, err := equal(path, l, r)
			if err != nil {
				return false, err
			}
			return eq, nil
		}
		return true, nil
	case l.Type == yaml.Type_TYPE_OBJ:
		for k := range l.Obj {
			if _, ok := r.Obj[k]; !ok {
				return false, nil
			}
		}
		for k := range r.Obj {
			if _, ok := l.Obj[k]; !ok {
				return false, nil
			}
		}
		for k, l := range l.Obj {
			r := r.Obj[k]
			eq, err := equal(path, l, r)
			if err != nil {
				return false, err
			}
			return eq, nil
		}
		return true, nil
	}
}
func compare(path *ast.Path, l, r *yaml.Value) (int, error) {
	switch {
	default:
		return 0, fmt.Errorf("unsupported cmp types: %v: %v and %v", path.Format(), l.Type, r.Type)
	case l.Type == yaml.Type_TYPE_NUM && r.Type == yaml.Type_TYPE_NUM:
		if l.Num < r.Num {
			return -1, nil
		}
		if l.Num > r.Num {
			return 1, nil
		}
		return 0, nil
	case l.Type == yaml.Type_TYPE_BOOL && r.Type == yaml.Type_TYPE_BOOL:
		if !l.Bool && r.Bool {
			return -1, nil
		}
		if l.Bool && !r.Bool {
			return 1, nil
		}
		return 0, nil
	case l.Type == yaml.Type_TYPE_STR && r.Type == yaml.Type_TYPE_STR:
		if l.Str < r.Str {
			return -1, nil
		}
		if l.Str > r.Str {
			return 1, nil
		}
		return 0, nil
	case l.Type == yaml.Type_TYPE_ARR && r.Type == yaml.Type_TYPE_ARR:
		for i, l := range l.Arr {
			r := r.Arr[i]
			cmp, err := compare(path, l, r)
			if err != nil {
				return 0, err
			}
			if cmp != 0 {
				return cmp, nil
			}
		}
		if len(l.Arr) < len(r.Arr) {
			return -1, nil
		}
		if len(l.Arr) > len(r.Arr) {
			return 1, nil
		}
		return 0, nil
	}
}

func ExprOfValue(v *yaml.Value) *ast.Expr {
	return &ast.Expr{Kind: ast.Expr_KIND_VAL_JSON, ValJson: &ast.ValJson{Json: v}}
}
