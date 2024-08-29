package ast

import (
	context "context"
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/pb/yaml"
	"strings"
)

var _ ParserServer = (*Parser)(nil)

type Parser struct{}

func (p *Parser) mustEmbedUnimplementedParserServer() {
	//TODO implement me
	panic("implement me")
}

func (p *Parser) Parse(ctx context.Context, input *ParseInput) (*ParseOutput, error) {
	o, err := (&yaml.Unmarshaller{}).Unmarshal(ctx, &yaml.UnmarshalInput{Yaml: input.Source})
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal json: %w", err)
	}

	return p.ParseEval(ctx, &ParseExprInput{Path: &Path{}, Value: o.Value})
}

func (p *Parser) ParseExpr(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	switch input.Value.Type {
	default:
		return nil, ErrUnexpectedType(input.Path, []yaml.Type{yaml.Type_TYPE_BOOL, yaml.Type_TYPE_NUM, yaml.Type_TYPE_STR, yaml.Type_TYPE_OBJ}, input.Value.Type)
	case yaml.Type_TYPE_BOOL, yaml.Type_TYPE_NUM, yaml.Type_TYPE_STR:
		return p.ParseScalar(ctx, &ParseExprInput{Value: input.Value, Path: input.Path})
	case yaml.Type_TYPE_OBJ:
		switch {
		default:
			return nil, fmt.Errorf("unsupported keys: %v %v", input.Path.Format(), input.Value.Keys())
		case hasKeys(input.Value, "eval"):
			return p.ParseEval(ctx, input)
		case hasKeys(input.Value, "obj"):
			return p.ParseNewObj(ctx, input)
		case hasKeys(input.Value, "arr"):
			return p.ParseNewArr(ctx, input)
		case hasKeys(input.Value, "json"):
			return p.ParseValJson(ctx, input)
		case hasKeys(input.Value, "for", "in", "do"):
			return p.ParseRangeIter(ctx, input)
		case hasKeys(input.Value, "get", "from"):
			return p.ParseElemAccess(ctx, input)
		case hasKeys(input.Value, "ref"):
			return p.ParseFunCall(ctx, input)
		case hasKeys(input.Value, "cases"):
			return p.ParseCaseBranches(ctx, input)
		case hasKeys(input.Value, OpUnary_OPERATOR_LEN.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_LEN})
		case hasKeys(input.Value, OpUnary_OPERATOR_NOT.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_NOT})
		case hasKeys(input.Value, OpUnary_OPERATOR_HEAD.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_HEAD})
		case hasKeys(input.Value, OpUnary_OPERATOR_TAIL.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_TAIL})
		case hasKeys(input.Value, OpUnary_OPERATOR_LAST.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_LAST})
		case hasKeys(input.Value, OpUnary_OPERATOR_INIT.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_INIT})
		case hasKeys(input.Value, OpUnary_OPERATOR_FLAT.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_FLAT})
		case hasKeys(input.Value, OpUnary_OPERATOR_ERROR.KeyName()):
			return p.ParseOpUnary(ctx, &ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_OPERATOR_ERROR})
		case hasKeys(input.Value, OpBinary_OPERATOR_SUB.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_SUB})
		case hasKeys(input.Value, OpBinary_OPERATOR_DIV.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_DIV})
		case hasKeys(input.Value, OpBinary_OPERATOR_MOD.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_MOD})
		case hasKeys(input.Value, OpBinary_OPERATOR_EQ.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_EQ})
		case hasKeys(input.Value, OpBinary_OPERATOR_NEQ.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_NEQ})
		case hasKeys(input.Value, OpBinary_OPERATOR_LT.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_LT})
		case hasKeys(input.Value, OpBinary_OPERATOR_LTE.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_LTE})
		case hasKeys(input.Value, OpBinary_OPERATOR_GT.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_GT})
		case hasKeys(input.Value, OpBinary_OPERATOR_GTE.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_GTE})
		case hasKeys(input.Value, OpBinary_OPERATOR_CMP.KeyName()):
			return p.ParseOpBinary(ctx, &ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_OPERATOR_CMP})
		case hasKeys(input.Value, OpVariadic_OPERATOR_ADD.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_ADD})
		case hasKeys(input.Value, OpVariadic_OPERATOR_MUL.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_MUL})
		case hasKeys(input.Value, OpVariadic_OPERATOR_AND.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_AND})
		case hasKeys(input.Value, OpVariadic_OPERATOR_OR.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_OR})
		case hasKeys(input.Value, OpVariadic_OPERATOR_CAT.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_CAT})
		case hasKeys(input.Value, OpVariadic_OPERATOR_MIN.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_MIN})
		case hasKeys(input.Value, OpVariadic_OPERATOR_MAX.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_MAX})
		case hasKeys(input.Value, OpVariadic_OPERATOR_MERGE.KeyName()):
			return p.ParseOpVariadic(ctx, &ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OPERATOR_MERGE})
		}
	}
}

func (p *Parser) ParseEval(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "eval") {
		return nil, ErrKeyNotFound(input.Path, []string{"eval"})
	}
	eval, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("eval"), Value: v.Obj["eval"]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	where := []*FunDef{}
	if w, ok := v.Obj["where"]; ok {
		path := path.AppendKey("where")
		if w.Type != yaml.Type_TYPE_ARR {
			return nil, ErrUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_ARR}, w.Type)
		}
		for pos, v := range w.Arr {
			path := path.AppendIndex(pos)
			if !hasKeys(v, "def", "value") {
				return nil, ErrKeyNotFound(path, []string{"def", "value"})
			}
			def := v.Obj["def"]
			if def.Type != yaml.Type_TYPE_STR {
				return nil, ErrUnexpectedType(path.AppendKey("def"), []yaml.Type{yaml.Type_TYPE_STR}, w.Type)
			}
			value, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("value"), Value: v.Obj["value"]})
			if err != nil {
				return nil, fmt.Errorf("fail to ParseExpr: %w", err)
			}
			funDef := &FunDef{Def: def.Str, Value: value.Expr}
			if with, ok := v.Obj["with"]; ok {
				if with.Type != yaml.Type_TYPE_ARR {
					return nil, ErrUnexpectedType(path.AppendKey("with"), []yaml.Type{yaml.Type_TYPE_ARR}, with.Type)
				}
				path := path.AppendKey("with")
				for pos, v := range with.Arr {
					if v.Type != yaml.Type_TYPE_STR {
						return nil, ErrUnexpectedType(path.AppendIndex(pos), []yaml.Type{yaml.Type_TYPE_STR}, w.Type)
					}
					funDef.With = append(funDef.With, v.Str)
				}
			}
			where = append(where)
		}
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind: Expr_KIND_EVAL,
			Eval: &Eval{Path: path, Where: where, Eval: eval.Expr},
		},
	}, nil
}

func (p *Parser) ParseScalar(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	switch input.Value.Type {
	default:
		return nil, ErrUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_BOOL, yaml.Type_TYPE_NUM, yaml.Type_TYPE_STR}, v.Type)
	case yaml.Type_TYPE_BOOL, yaml.Type_TYPE_NUM, yaml.Type_TYPE_STR:
		return &ParseOutput{
			Expr: &Expr{
				Kind:   Expr_KIND_SCALAR,
				Scalar: &Scalar{Path: path, Val: v},
			},
		}, nil
	}
}

func (p *Parser) ParseNewObj(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "obj") {
		return nil, ErrKeyNotFound(path, []string{"obj"})
	}
	obj := map[string]*Expr{}
	for key, value := range v.Obj["obj"].Obj {
		path := path.AppendKey(key)
		v, err := p.ParseExpr(ctx, &ParseExprInput{Path: path, Value: value})
		if err != nil {
			return nil, fmt.Errorf("fail to ParseExpr: %w", err)
		}
		obj[key] = v.Expr
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:   Expr_KIND_NEW_OBJ,
			NewObj: &NewObj{Path: path, Obj: obj},
		},
	}, nil
}

func (p *Parser) ParseNewArr(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "arr") {
		return nil, ErrKeyNotFound(path, []string{"arr"})
	}
	arr := []*Expr{}
	for pos, v := range v.Obj["arr"].Arr {
		path := input.Path.AppendIndex(pos)
		elem, err := p.ParseExpr(ctx, &ParseExprInput{Path: path, Value: v})
		if err != nil {
			return nil, fmt.Errorf("fail to ParseExpr: %w", err)
		}
		arr = append(arr, elem.Expr)
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:   Expr_KIND_NEW_ARR,
			NewArr: &NewArr{Path: path, Arr: arr},
		},
	}, nil
}

func (p *Parser) ParseValJson(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "json") {
		return nil, ErrKeyNotFound(path, []string{"json"})
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:    Expr_KIND_VAL_JSON,
			ValJson: &ValJson{Path: path, Json: v.Obj["json"]},
		},
	}, nil
}

func (p *Parser) ParseRangeIter(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "for", "in", "do") {
		return nil, ErrKeyNotFound(path, []string{"for", "in", "do"})
	}
	rangeIter := &RangeIter{Path: path}
	{
		v := v.Obj["for"]
		path := path.AppendKey("for")
		if v.Type != yaml.Type_TYPE_ARR {
			return nil, ErrUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_ARR}, v.Type)
		}
		if len(v.Arr) != 2 {
			return nil, ErrUnexpectedLength(path, "= 2", len(v.Arr))
		}
		if v.Arr[0].Type != yaml.Type_TYPE_STR {
			return nil, ErrUnexpectedType(path.AppendIndex(0), []yaml.Type{yaml.Type_TYPE_STR}, v.Type)
		}
		if v.Arr[1].Type != yaml.Type_TYPE_STR {
			return nil, ErrUnexpectedType(path.AppendIndex(1), []yaml.Type{yaml.Type_TYPE_STR}, v.Type)
		}
		rangeIter.ForPos, rangeIter.ForVal = v.Arr[0].Str, v.Arr[1].Str
	}
	in, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("in"), Value: v.Obj["in"]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	rangeIter.In = in.Expr
	do, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("do"), Value: v.Obj["do"]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	rangeIter.Do = do.Expr
	if v, ok := v.Obj["if"]; ok {
		if_, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("if"), Value: v})
		if err != nil {
			return nil, fmt.Errorf("fail to ParseExpr: %w", err)
		}
		rangeIter.If = if_.Expr
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:      Expr_KIND_RANGE_ITER,
			RangeIter: rangeIter,
		},
	}, nil
}

func (p *Parser) ParseElemAccess(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "get", "from") {
		return nil, ErrKeyNotFound(path, []string{"get", "from"})
	}
	var err error
	get, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("get"), Value: v.Obj["get"]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	from, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("from"), Value: v.Obj["from"]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:       Expr_KIND_ELEM_ACCESS,
			ElemAccess: &ElemAccess{Path: path, Get: get.Expr, From: from.Expr},
		},
	}, nil
}

func (p *Parser) ParseFunCall(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "ref") {
		return nil, ErrKeyNotFound(path, []string{"ref"})
	}
	ref := v.Obj["ref"]
	if ref.Type != yaml.Type_TYPE_STR {
		return nil, ErrUnexpectedType(path.AppendKey("ref"), []yaml.Type{yaml.Type_TYPE_STR}, ref.Type)
	}
	funCall := &FunCall{Path: path, Ref: ref.Str}
	if with, ok := v.Obj["with"]; ok {
		path := path.AppendKey("with")
		if with.Type != yaml.Type_TYPE_OBJ {
			return nil, ErrUnexpectedType(path, []yaml.Type{yaml.Type_TYPE_OBJ}, with.Type)
		}
		funCall.With = map[string]*Expr{}
		for k, v := range v.Obj {
			expr, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey(k), Value: v})
			if err != nil {
				return nil, fmt.Errorf("fail to ParseExpr: %w", err)
			}
			funCall.With[k] = expr.Expr
		}
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:    Expr_KIND_FUN_CALL,
			FunCall: funCall,
		},
	}, nil
}

func (p *Parser) ParseCaseBranches(ctx context.Context, input *ParseExprInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	if !hasKeys(v, "cases") {
		return nil, ErrKeyNotFound(input.Path, []string{"cases"})
	}
	cases := v.Obj["cases"]
	if cases.Type != yaml.Type_TYPE_ARR {
		return nil, ErrUnexpectedType(path.AppendKey("cases"), []yaml.Type{yaml.Type_TYPE_ARR}, cases.Type)
	}
	casesBranches := &CaseBranches{Path: path}
	for i, v := range cases.Arr {
		path := path.AppendIndex(i)
		switch {
		default:
			return nil, fmt.Errorf("unsupported keys: %v: want [when,then] or [otherwise]: got [%v]", input.Path.AppendIndex(i), strings.Join(v.Keys(), ","))
		case hasKeys(v, "otherwise"):
			otherwise, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("otherwise"), Value: v.Obj["otherwise"]})
			if err != nil {
				return nil, fmt.Errorf("fail to ParseExpr: %w", err)
			}
			casesBranches.Branches = append(casesBranches.Branches, &CaseBranches_Branch{Otherwise: otherwise.Expr})
		case hasKeys(v, "when", "then"):
			when, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("when"), Value: v.Obj["when"]})
			if err != nil {
				return nil, fmt.Errorf("fail to ParseExpr: %w", err)
			}
			then, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey("then"), Value: v.Obj["then"]})
			if err != nil {
				return nil, fmt.Errorf("fail to ParseExpr: %w", err)
			}
			casesBranches.Branches = append(casesBranches.Branches, &CaseBranches_Branch{When: when.Expr, Then: then.Expr})
		}
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:         Expr_KIND_CASE_BRANCHES,
			CaseBranches: casesBranches,
		},
	}, nil
}

func (p *Parser) ParseOpUnary(ctx context.Context, input *ParseOpUnaryInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	operator := input.Operator.KeyName()
	if !hasKeys(v, operator) {
		return nil, ErrKeyNotFound(path, []string{operator})
	}
	operand, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey(operator), Value: v.Obj[operator]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:    Expr_KIND_OP_UNARY,
			OpUnary: &OpUnary{Path: path, Operator: input.Operator, Operand: operand.Expr},
		},
	}, nil
}

func (p *Parser) ParseOpBinary(ctx context.Context, input *ParseOpBinaryInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	operator := input.Operator.KeyName()
	if !hasKeys(v, operator) {
		return nil, ErrKeyNotFound(path, []string{operator})
	}
	os := v.Obj[operator]
	if os.Type != yaml.Type_TYPE_ARR {
		return nil, ErrUnexpectedType(path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_ARR}, os.Type)
	}
	if len(os.Arr) != 2 {
		return nil, ErrUnexpectedLength(path.AppendKey(operator), "= 2", len(os.Arr))
	}
	l, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey(operator).AppendIndex(0), Value: os.Arr[0]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	r, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendKey(operator).AppendIndex(1), Value: os.Arr[1]})
	if err != nil {
		return nil, fmt.Errorf("fail to ParseExpr: %w", err)
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:     Expr_KIND_OP_BINARY,
			OpBinary: &OpBinary{Path: path, Operator: input.Operator, OperandLeft: l.Expr, OperandRight: r.Expr},
		},
	}, nil
}

func (p *Parser) ParseOpVariadic(ctx context.Context, input *ParseOpVariadicInput) (*ParseOutput, error) {
	v := input.Value
	path := input.Path
	operator := input.Operator.KeyName()
	if !hasKeys(v, operator) {
		return nil, ErrKeyNotFound(path, []string{operator})
	}
	os := v.Obj[operator]
	if os.Type != yaml.Type_TYPE_ARR {
		return nil, ErrUnexpectedType(path.AppendKey(operator), []yaml.Type{yaml.Type_TYPE_ARR}, os.Type)
	}
	operation := &OpVariadic{Path: path, Operator: input.Operator}
	for i, v := range os.Arr {
		path := input.Path.AppendKey(operator)
		operand, err := p.ParseExpr(ctx, &ParseExprInput{Path: path.AppendIndex(i), Value: v})
		if err != nil {
			return nil, fmt.Errorf("fail to ParseExpr: %w", err)
		}
		operation.Operands = append(operation.Operands, operand.Expr)
	}
	return &ParseOutput{
		Expr: &Expr{
			Kind:       Expr_KIND_OP_VARIADIC,
			OpVariadic: operation,
		},
	}, nil
}

func hasKeys(v *yaml.Value, keys ...string) bool {
	if v.Type != yaml.Type_TYPE_OBJ {
		return false
	}
	for _, k := range keys {
		if _, ok := v.Obj[k]; !ok {
			return false
		}
	}
	return true
}

func ErrUnexpectedType(path *Path, want []yaml.Type, got yaml.Type) error {
	w := []string{}
	for _, t := range want {
		w = append(w, t.String())
	}
	return fmt.Errorf("unexpected type: %q: want [%q]: got %q", path.Format(), strings.Join(w, ","), got)
}

func ErrKeyNotFound(path *Path, want []string) error {
	return fmt.Errorf("key not found: %q: want [%v]", path.Format(), strings.Join(want, ","))
}

func ErrUnexpectedLength(path *Path, want string, got int) error {
	return fmt.Errorf("unexpected length: %q: want %q: got %v", path.Format(), want, got)
}
