package ast

import (
	_ "embed"
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/yaml"
	"slices"
	"strings"
)

type Parser interface {
	Parse(*ParseInput) *ParseOutput
	ParseExpr(*ParseExprInput) *ParseExprOutput
	ParseEval(*ParseExprInput) *ParseExprOutput
	ParseScalar(*ParseExprInput) *ParseExprOutput
	ParseObj(*ParseExprInput) *ParseExprOutput
	ParseArr(*ParseExprInput) *ParseExprOutput
	ParseJson(*ParseExprInput) *ParseExprOutput
	ParseRangeIter(*ParseExprInput) *ParseExprOutput
	ParseGetElem(*ParseExprInput) *ParseExprOutput
	ParseFunCall(*ParseExprInput) *ParseExprOutput
	ParseCases(*ParseExprInput) *ParseExprOutput
	ParseOpUnary(*ParseOpUnaryInput) *ParseExprOutput
	ParseOpBinary(*ParseOpBinaryInput) *ParseExprOutput
	ParseOpVariadic(*ParseOpVariadicInput) *ParseExprOutput
}

func NewParser() Parser {
	return &parser{}
}

type parser struct{}

var _ Parser = &parser{}

func (p *parser) Parse(input *ParseInput) *ParseOutput {
	// decode
	o := yaml.NewDecoder().Decode(&yaml.DecodeInput{Yaml: input.Source})
	if o.IsError {
		return &ParseOutput{
			Status:       ParseOutput_DECODE_ERROR,
			ErrorMessage: o.ErrorMessage,
		}
	}

	// validate
	if o := NewValidator().Validate(&ValidateInput{Source: input.Source}); o.Status != ValidateOutput_OK {
		return &ParseOutput{
			Status:       ParseOutput_VALIDATE_ERROR,
			ErrorMessage: "validation error",
		}
	}

	// parse
	result := p.ParseExpr(&ParseExprInput{Path: &Path{}, Value: o.Value})
	if result.Code != ParseErrorCode_OK {
		return &ParseOutput{
			Status:         ParseOutput_PARSE_ERROR,
			ErrorMessage:   result.ErrorMessage,
			ParseErrorCode: result.Code,
			ParseErrorPath: result.ErrorPath,
		}
	}
	return &ParseOutput{Expr: result.Expr}
}
func (p *parser) ParseExpr(input *ParseExprInput) *ParseExprOutput {
	switch input.Value.Type {
	default:
		return errorUnexpectedType(input.Path, []yaml.Type{yaml.Type_BOOL, yaml.Type_NUM, yaml.Type_STR, yaml.Type_OBJ}, input.Value.Type)
	case yaml.Type_BOOL, yaml.Type_NUM, yaml.Type_STR:
		return p.ParseScalar(&ParseExprInput{Value: input.Value, Path: input.Path})
	case yaml.Type_OBJ:
		switch {
		default:
			return errorUnsupportedKeys(input.Path, "suitable expression not found", input.Value.Keys(), nil, nil)
		case keysAreMatched(input.Value, []string{"eval"}, []string{"where"}):
			return p.ParseEval(input)
		case keysAreMatched(input.Value, []string{"obj"}, nil):
			return p.ParseObj(input)
		case keysAreMatched(input.Value, []string{"arr"}, nil):
			return p.ParseArr(input)
		case keysAreMatched(input.Value, []string{"json"}, nil):
			return p.ParseJson(input)
		case keysAreMatched(input.Value, []string{"for", "in", "do"}, []string{"if"}):
			return p.ParseRangeIter(input)
		case keysAreMatched(input.Value, []string{"get", "from"}, nil):
			return p.ParseGetElem(input)
		case keysAreMatched(input.Value, []string{"ref"}, nil):
			return p.ParseFunCall(input)
		case keysAreMatched(input.Value, []string{"cases"}, nil):
			return p.ParseCases(input)
		case keysAreMatched(input.Value, []string{OpUnary_LEN.KeyName()}, nil):
			return p.ParseOpUnary(&ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_LEN})
		case keysAreMatched(input.Value, []string{OpUnary_NOT.KeyName()}, nil):
			return p.ParseOpUnary(&ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_NOT})
		case keysAreMatched(input.Value, []string{OpUnary_FLAT.KeyName()}, nil):
			return p.ParseOpUnary(&ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_FLAT})
		case keysAreMatched(input.Value, []string{OpUnary_FLOOR.KeyName()}, nil):
			return p.ParseOpUnary(&ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_FLOOR})
		case keysAreMatched(input.Value, []string{OpUnary_CEIL.KeyName()}, nil):
			return p.ParseOpUnary(&ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_CEIL})
		case keysAreMatched(input.Value, []string{OpUnary_ABORT.KeyName()}, nil):
			return p.ParseOpUnary(&ParseOpUnaryInput{Path: input.Path, Value: input.Value, Operator: OpUnary_ABORT})
		case keysAreMatched(input.Value, []string{OpBinary_SUB.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_SUB})
		case keysAreMatched(input.Value, []string{OpBinary_DIV.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_DIV})
		case keysAreMatched(input.Value, []string{OpBinary_MOD.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_MOD})
		case keysAreMatched(input.Value, []string{OpBinary_EQ.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_EQ})
		case keysAreMatched(input.Value, []string{OpBinary_NEQ.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_NEQ})
		case keysAreMatched(input.Value, []string{OpBinary_LT.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_LT})
		case keysAreMatched(input.Value, []string{OpBinary_LTE.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_LTE})
		case keysAreMatched(input.Value, []string{OpBinary_GT.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_GT})
		case keysAreMatched(input.Value, []string{OpBinary_GTE.KeyName()}, nil):
			return p.ParseOpBinary(&ParseOpBinaryInput{Path: input.Path, Value: input.Value, Operator: OpBinary_GTE})
		case keysAreMatched(input.Value, []string{OpVariadic_ADD.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_ADD})
		case keysAreMatched(input.Value, []string{OpVariadic_MUL.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_MUL})
		case keysAreMatched(input.Value, []string{OpVariadic_AND.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_AND})
		case keysAreMatched(input.Value, []string{OpVariadic_OR.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_OR})
		case keysAreMatched(input.Value, []string{OpVariadic_CAT.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_CAT})
		case keysAreMatched(input.Value, []string{OpVariadic_MIN.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_MIN})
		case keysAreMatched(input.Value, []string{OpVariadic_MAX.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_MAX})
		case keysAreMatched(input.Value, []string{OpVariadic_MERGE.KeyName()}, nil):
			return p.ParseOpVariadic(&ParseOpVariadicInput{Path: input.Path, Value: input.Value, Operator: OpVariadic_MERGE})
		}
	}
}
func (p *parser) ParseEval(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if !keysAreMatched(v, []string{"eval"}, []string{"where"}) {
		return errorUnsupportedKeys(input.Path, "Eval", v.Keys(), []string{"eval"}, []string{"where"})
	}
	eval := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("eval"), Value: v.Obj["eval"]})
	if eval.Code != ParseErrorCode_OK {
		return eval
	}
	where := []*FunDef{}
	if w, ok := v.Obj["where"]; ok {
		path := path.AppendKey("where")
		if w.Type != yaml.Type_ARR {
			return errorUnexpectedType(path, []yaml.Type{yaml.Type_ARR}, w.Type)
		}
		for pos, v := range w.Arr {
			path := path.AppendIndex(pos)
			if !keysAreMatched(v, []string{"def", "value"}, []string{"with"}) {
				return errorUnsupportedKeys(path, "FunDef", v.Keys(), []string{"def", "value"}, []string{"with"})
			}
			def := v.Obj["def"]
			if def.Type != yaml.Type_STR {
				return errorUnexpectedType(path.AppendKey("def"), []yaml.Type{yaml.Type_STR}, def.Type)
			}
			value := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("value"), Value: v.Obj["value"]})
			if value.Code != ParseErrorCode_OK {
				return value
			}
			funDef := &FunDef{Def: def.Str, Value: value.Expr}
			if with, ok := v.Obj["with"]; ok {
				if with.Type != yaml.Type_ARR {
					return errorUnexpectedType(path.AppendKey("with"), []yaml.Type{yaml.Type_ARR}, with.Type)
				}
				path := path.AppendKey("with")
				for pos, v := range with.Arr {
					if v.Type != yaml.Type_STR {
						return errorUnexpectedType(path.AppendIndex(pos), []yaml.Type{yaml.Type_STR}, w.Type)
					}
					funDef.With = append(funDef.With, v.Str)
				}
			}
			where = append(where)
		}
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path: path,
			Kind: Expr_EVAL,
			Eval: &Eval{Where: where, Eval: eval.Expr},
		},
	}
}
func (p *parser) ParseScalar(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	switch input.Value.Type {
	default:
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_BOOL, yaml.Type_NUM, yaml.Type_STR}, v.Type)
	case yaml.Type_BOOL, yaml.Type_NUM, yaml.Type_STR:
		return &ParseExprOutput{
			Expr: &Expr{
				Path:   path,
				Kind:   Expr_SCALAR,
				Scalar: &Scalar{Val: v},
			},
		}
	}
}
func (p *parser) ParseObj(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"obj"}, nil) {
		return errorUnsupportedKeys(input.Path, "NewObj", v.Keys(), []string{"obj"}, nil)
	}
	if v.Obj["obj"].Type != yaml.Type_OBJ {
		return errorUnexpectedType(path.AppendKey("obj"), []yaml.Type{yaml.Type_OBJ}, v.Obj["obj"].Type)
	}
	obj := map[string]*Expr{}
	for key, value := range v.Obj["obj"].Obj {
		path := path.AppendKey("obj").AppendKey(key)
		expr := p.ParseExpr(&ParseExprInput{Path: path, Value: value})
		if expr.Code != ParseErrorCode_OK {
			return expr
		}
		obj[key] = expr.Expr
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path: path,
			Kind: Expr_OBJ,
			Obj:  &Obj{Obj: obj},
		},
	}
}
func (p *parser) ParseArr(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"arr"}, nil) {
		return errorUnsupportedKeys(input.Path, "NewArr", v.Keys(), []string{"arr"}, nil)
	}
	if v.Obj["arr"].Type != yaml.Type_ARR {
		return errorUnexpectedType(path.AppendKey("arr"), []yaml.Type{yaml.Type_ARR}, v.Obj["arr"].Type)
	}
	arr := []*Expr{}
	for pos, v := range v.Obj["arr"].Arr {
		path := input.Path.AppendKey("arr").AppendIndex(pos)
		elem := p.ParseExpr(&ParseExprInput{Path: path, Value: v})
		if elem.Code != ParseErrorCode_OK {
			return elem
		}
		arr = append(arr, elem.Expr)
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path: path,
			Kind: Expr_ARR,
			Arr:  &Arr{Arr: arr},
		},
	}
}
func (p *parser) ParseJson(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"json"}, nil) {
		return errorUnsupportedKeys(input.Path, "ValJson", v.Keys(), []string{"json"}, nil)
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path: path,
			Kind: Expr_JSON,
			Json: &Json{Json: v.Obj["json"]},
		},
	}
}
func (p *parser) ParseRangeIter(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"for", "in", "do"}, []string{"if"}) {
		return errorUnsupportedKeys(input.Path, "RangeIter", v.Keys(), []string{"for", "in", "do"}, []string{"if"})
	}
	rangeIter := &RangeIter{}
	{
		v := v.Obj["for"]
		path := path.AppendKey("for")
		if v.Type != yaml.Type_ARR {
			return errorUnexpectedType(path, []yaml.Type{yaml.Type_ARR}, v.Type)
		}
		if len(v.Arr) != 2 {
			return errorUnexpectedLength(path, len(v.Arr), "= 2")
		}
		if v.Arr[0].Type != yaml.Type_STR {
			return errorUnexpectedType(path.AppendIndex(0), []yaml.Type{yaml.Type_STR}, v.Type)
		}
		if v.Arr[1].Type != yaml.Type_STR {
			return errorUnexpectedType(path.AppendIndex(1), []yaml.Type{yaml.Type_STR}, v.Type)
		}
		rangeIter.ForPos, rangeIter.ForVal = v.Arr[0].Str, v.Arr[1].Str
	}
	{
		in := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("in"), Value: v.Obj["in"]})
		if in.Code != ParseErrorCode_OK {
			return in
		}
		rangeIter.In = in.Expr
	}
	{
		do := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("do"), Value: v.Obj["do"]})
		if do.Code != ParseErrorCode_OK {
			return do
		}
		rangeIter.Do = do.Expr
	}
	{
		if v, ok := v.Obj["if"]; ok {
			if_ := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("if"), Value: v})
			if if_.Code != ParseErrorCode_OK {
				return if_
			}
			rangeIter.If = if_.Expr
		}
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:      path,
			Kind:      Expr_RANGE_ITER,
			RangeIter: rangeIter,
		},
	}
}
func (p *parser) ParseGetElem(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"get", "from"}, nil) {
		return errorUnsupportedKeys(input.Path, "ElemAccess", v.Keys(), []string{"get", "from"}, nil)
	}
	get := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("get"), Value: v.Obj["get"]})
	if get.Code != ParseErrorCode_OK {
		return get
	}
	from := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("from"), Value: v.Obj["from"]})
	if from.Code != ParseErrorCode_OK {
		return from
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:    path,
			Kind:    Expr_GET_ELEM,
			GetElem: &GetElem{Get: get.Expr, From: from.Expr},
		},
	}
}
func (p *parser) ParseFunCall(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"ref"}, []string{"with"}) {
		return errorUnsupportedKeys(input.Path, "FunCall", v.Keys(), []string{"ref"}, []string{"with"})
	}
	ref := v.Obj["ref"]
	if ref.Type != yaml.Type_STR {
		return errorUnexpectedType(path.AppendKey("ref"), []yaml.Type{yaml.Type_STR}, ref.Type)
	}
	funCall := &FunCall{Ref: ref.Str}
	if with, ok := v.Obj["with"]; ok {
		path := path.AppendKey("with")
		if with.Type != yaml.Type_OBJ {
			return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, with.Type)
		}
		funCall.With = map[string]*Expr{}
		for k, v := range v.Obj {
			expr := p.ParseExpr(&ParseExprInput{Path: path.AppendKey(k), Value: v})
			if expr.Code != ParseErrorCode_OK {
				return expr
			}
			funCall.With[k] = expr.Expr
		}
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:    path,
			Kind:    Expr_FUN_CALL,
			FunCall: funCall,
		},
	}
}
func (p *parser) ParseCases(input *ParseExprInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	if !keysAreMatched(v, []string{"cases"}, nil) {
		return errorUnsupportedKeys(input.Path, "CaseBranches", v.Keys(), []string{"cases"}, nil)
	}
	if v.Obj["cases"].Type != yaml.Type_ARR {
		return errorUnexpectedType(path.AppendKey("cases"), []yaml.Type{yaml.Type_ARR}, v.Obj["cases"].Type)
	}
	cases := &Cases{}
	for i, v := range v.Obj["cases"].Arr {
		path := path.AppendIndex(i)
		switch {
		default:
			if _, ok := v.Obj["otherwise"]; !ok {
				if !keysAreMatched(v, []string{"otherwise"}, nil) {
					return errorUnsupportedKeys(path, "Cases", v.Keys(), []string{"otherwise"}, nil)
				}
			}
			return errorUnsupportedKeys(path, "Cases", v.Keys(), []string{"when", "then"}, nil)
		case !keysAreMatched(v, []string{"otherwise"}, nil):
			otherwise := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("otherwise"), Value: v.Obj["otherwise"]})
			if otherwise.Code != ParseErrorCode_OK {
				return otherwise
			}
			cases.Branches = append(cases.Branches, &Cases_Branch{IsOtherwise: true, Otherwise: otherwise.Expr})
		case !keysAreMatched(v, []string{"when", "then"}, nil):
			when := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("when"), Value: v.Obj["when"]})
			if when.Code != ParseErrorCode_OK {
				return when
			}
			then := p.ParseExpr(&ParseExprInput{Path: path.AppendKey("then"), Value: v.Obj["then"]})
			if then.Code != ParseErrorCode_OK {
				return then
			}
			cases.Branches = append(cases.Branches, &Cases_Branch{When: when.Expr, Then: then.Expr})
		}
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:  path,
			Kind:  Expr_CASES,
			Cases: cases,
		},
	}
}
func (p *parser) ParseOpUnary(input *ParseOpUnaryInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	operator := input.Operator.KeyName()
	if !keysAreMatched(v, []string{operator}, nil) {
		return errorUnsupportedKeys(input.Path, fmt.Sprintf("OpUnary<%v>", operator), v.Keys(), []string{operator}, nil)
	}
	operand := p.ParseExpr(&ParseExprInput{Path: path.AppendKey(operator), Value: v.Obj[operator]})
	if operand.Code != ParseErrorCode_OK {
		return operand
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:    path,
			Kind:    Expr_OP_UNARY,
			OpUnary: &OpUnary{Operator: input.Operator, Operand: operand.Expr},
		},
	}
}
func (p *parser) ParseOpBinary(input *ParseOpBinaryInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	operator := input.Operator.KeyName()
	if !keysAreMatched(v, []string{operator}, nil) {
		return errorUnsupportedKeys(input.Path, fmt.Sprintf("OpBinary<%v>", operator), v.Keys(), []string{operator}, nil)
	}
	os := v.Obj[operator]
	if os.Type != yaml.Type_ARR {
		return errorUnexpectedType(path.AppendKey(operator), []yaml.Type{yaml.Type_ARR}, os.Type)
	}
	if len(os.Arr) != 2 {
		return errorUnexpectedLength(path.AppendKey(operator), len(os.Arr), "= 2")
	}
	l := p.ParseExpr(&ParseExprInput{Path: path.AppendKey(operator).AppendIndex(0), Value: os.Arr[0]})
	if l.Code != ParseErrorCode_OK {
		return l
	}
	r := p.ParseExpr(&ParseExprInput{Path: path.AppendKey(operator).AppendIndex(1), Value: os.Arr[1]})
	if r.Code != ParseErrorCode_OK {
		return r
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:     path,
			Kind:     Expr_OP_BINARY,
			OpBinary: &OpBinary{Operator: input.Operator, OperandLeft: l.Expr, OperandRight: r.Expr},
		},
	}
}
func (p *parser) ParseOpVariadic(input *ParseOpVariadicInput) *ParseExprOutput {
	v := input.Value
	path := input.Path
	if v.Type != yaml.Type_OBJ {
		return errorUnexpectedType(path, []yaml.Type{yaml.Type_OBJ}, v.Type)
	}
	operator := input.Operator.KeyName()
	if !keysAreMatched(v, []string{operator}, nil) {
		return errorUnsupportedKeys(input.Path, fmt.Sprintf("OpVariadic<%v>", operator), v.Keys(), []string{operator}, nil)
	}
	os := v.Obj[operator]
	if os.Type != yaml.Type_ARR {
		return errorUnexpectedType(path.AppendKey(operator), []yaml.Type{yaml.Type_ARR}, os.Type)
	}
	operation := &OpVariadic{Operator: input.Operator}
	for i, v := range os.Arr {
		path := input.Path.AppendKey(operator)
		operand := p.ParseExpr(&ParseExprInput{Path: path.AppendIndex(i), Value: v})
		if operand.Code != ParseErrorCode_OK {
			return operand
		}
		operation.Operands = append(operation.Operands, operand.Expr)
	}
	return &ParseExprOutput{
		Expr: &Expr{
			Path:       path,
			Kind:       Expr_OP_VARIADIC,
			OpVariadic: operation,
		},
	}
}

func errorUnexpectedType(path *Path, want []yaml.Type, got yaml.Type) *ParseExprOutput {
	return &ParseExprOutput{
		Code:         ParseErrorCode_UNEXPECTED_TYPE,
		ErrorPath:    path,
		ErrorMessage: fmt.Sprintf("unexpected type: want %v, got %v", want, got),
	}
}
func errorUnsupportedKeys(path *Path, message string, got []string, required []string, optional []string) *ParseExprOutput {
	return &ParseExprOutput{
		Code:      ParseErrorCode_UNSUPORTED_KEYS,
		ErrorPath: path,
		ErrorMessage: fmt.Sprintf("unexpected key: %v: required [%v], optional [%v], got [%v]",
			message,
			strings.Join(required, ","),
			strings.Join(optional, ","),
			strings.Join(got, ","),
		),
	}
}
func errorUnexpectedLength(path *Path, got int, want string) *ParseExprOutput {
	return &ParseExprOutput{
		Code:         ParseErrorCode_UNEXPECTED_LENGTH,
		ErrorPath:    path,
		ErrorMessage: fmt.Sprintf("unexpected length: want %q, got %v", want, got),
	}
}
func keysAreMatched(v *yaml.Value, required []string, optional []string) bool {
	keys := v.Keys()
	for _, k := range required {
		if !slices.Contains(keys, k) {
			return false
		}
	}
	for _, k := range keys {
		if !slices.Contains(append(append([]string{}, required...), optional...), k) {
			return false
		}
	}
	return true
}
