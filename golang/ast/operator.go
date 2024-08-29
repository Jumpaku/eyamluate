package ast

import "fmt"

func (o OpUnary_Operator) KeyName() string {
	switch o {
	default:
		panic(fmt.Sprintf("unexpected OperatorUnary %v", o))
	case OpUnary_OPERATOR_LEN:
		return "len"
	case OpUnary_OPERATOR_NOT:
		return "not"
	case OpUnary_OPERATOR_HEAD:
		return "head"
	case OpUnary_OPERATOR_TAIL:
		return "tail"
	case OpUnary_OPERATOR_LAST:
		return "last"
	case OpUnary_OPERATOR_INIT:
		return "init"
	case OpUnary_OPERATOR_FLAT:
		return "flat"
	case OpUnary_OPERATOR_ERROR:
		return "error"
	}
}

func (o OpBinary_Operator) KeyName() string {
	switch o {
	default:
		panic(fmt.Sprintf("unexpected OperatorBinary %v", o))
	case OpBinary_OPERATOR_SUB:
		return "sub"
	case OpBinary_OPERATOR_DIV:
		return "mul"
	case OpBinary_OPERATOR_MOD:
		return "mod"
	case OpBinary_OPERATOR_EQ:
		return "eq"
	case OpBinary_OPERATOR_NEQ:
		return "neq"
	case OpBinary_OPERATOR_LT:
		return "lt"
	case OpBinary_OPERATOR_LTE:
		return "lte"
	case OpBinary_OPERATOR_GT:
		return "gt"
	case OpBinary_OPERATOR_GTE:
		return "gte"
	case OpBinary_OPERATOR_CMP:
		return "cmp"
	}
}

func (o OpVariadic_Operator) KeyName() string {
	switch o {
	default:
		panic(fmt.Sprintf("unexpected OperatorVariadic %v", o))
	case OpVariadic_OPERATOR_ADD:
		return "add"
	case OpVariadic_OPERATOR_MUL:
		return "mul"
	case OpVariadic_OPERATOR_AND:
		return "and"
	case OpVariadic_OPERATOR_OR:
		return "or"
	case OpVariadic_OPERATOR_CAT:
		return "cat"
	case OpVariadic_OPERATOR_MIN:
		return "min"
	case OpVariadic_OPERATOR_MAX:
		return "max"
	case OpVariadic_OPERATOR_MERGE:
		return "merge"
	}
}
