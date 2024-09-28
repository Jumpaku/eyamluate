import { OpBinary_Operator, OpUnary_Operator, OpVariadic_Operator } from "./operation_pb.js";
export function unaryKeyName(op) {
    switch (op) {
        case OpUnary_Operator.LEN:
            return "len";
        case OpUnary_Operator.NOT:
            return "not";
        case OpUnary_Operator.FLAT:
            return "flat";
        case OpUnary_Operator.FLOOR:
            return "floor";
        case OpUnary_Operator.CEIL:
            return "ceil";
        case OpUnary_Operator.ABORT:
            return "abort";
        default:
            throw new Error(`unexpected OperatorUnary ${op}`);
    }
}
export function binaryKeyName(op) {
    switch (op) {
        case OpBinary_Operator.SUB:
            return "sub";
        case OpBinary_Operator.DIV:
            return "div";
        case OpBinary_Operator.EQ:
            return "eq";
        case OpBinary_Operator.NEQ:
            return "neq";
        case OpBinary_Operator.LT:
            return "lt";
        case OpBinary_Operator.LTE:
            return "lte";
        case OpBinary_Operator.GT:
            return "gt";
        case OpBinary_Operator.GTE:
            return "gte";
        default:
            throw new Error(`unexpected OperatorBinary ${op}`);
    }
}
export function variadicKeyName(op) {
    switch (op) {
        case OpVariadic_Operator.ADD:
            return "add";
        case OpVariadic_Operator.MUL:
            return "mul";
        case OpVariadic_Operator.AND:
            return "and";
        case OpVariadic_Operator.OR:
            return "or";
        case OpVariadic_Operator.CAT:
            return "cat";
        case OpVariadic_Operator.MIN:
            return "min";
        case OpVariadic_Operator.MAX:
            return "max";
        case OpVariadic_Operator.MERGE:
            return "merge";
        default:
            throw new Error(`unexpected OperatorVariadic ${op}`);
    }
}
