package ast

import "fmt"

func (o OpUnary_Operator) KeyName() string {
	switch o {
	default:
		panic(fmt.Sprintf("unexpected OperatorUnary %v", o))
	case OpUnary_LEN:
		return "len"
	case OpUnary_NOT:
		return "not"
	case OpUnary_HEAD:
		return "head"
	case OpUnary_TAIL:
		return "tail"
	case OpUnary_LAST:
		return "last"
	case OpUnary_INIT:
		return "init"
	case OpUnary_FLAT:
		return "flat"
	case OpUnary_ABORT:
		return "abort"
	}
}

func (o OpBinary_Operator) KeyName() string {
	switch o {
	default:
		panic(fmt.Sprintf("unexpected OperatorBinary %v", o))
	case OpBinary_SUB:
		return "sub"
	case OpBinary_DIV:
		return "mul"
	case OpBinary_MOD:
		return "mod"
	case OpBinary_EQ:
		return "eq"
	case OpBinary_NEQ:
		return "neq"
	case OpBinary_LT:
		return "lt"
	case OpBinary_LTE:
		return "lte"
	case OpBinary_GT:
		return "gt"
	case OpBinary_GTE:
		return "gte"
	case OpBinary_CMP:
		return "cmp"
	}
}

func (o OpVariadic_Operator) KeyName() string {
	switch o {
	default:
		panic(fmt.Sprintf("unexpected OperatorVariadic %v", o))
	case OpVariadic_ADD:
		return "add"
	case OpVariadic_MUL:
		return "mul"
	case OpVariadic_AND:
		return "and"
	case OpVariadic_OR:
		return "or"
	case OpVariadic_CAT:
		return "cat"
	case OpVariadic_MIN:
		return "min"
	case OpVariadic_MAX:
		return "max"
	case OpVariadic_MERGE:
		return "merge"
	}
}
