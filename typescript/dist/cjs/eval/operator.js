"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.unaryKeyName = unaryKeyName;
exports.binaryKeyName = binaryKeyName;
exports.variadicKeyName = variadicKeyName;
const operation_pb_js_1 = require("./operation_pb.js");
function unaryKeyName(op) {
    switch (op) {
        case operation_pb_js_1.OpUnary_Operator.LEN:
            return "len";
        case operation_pb_js_1.OpUnary_Operator.NOT:
            return "not";
        case operation_pb_js_1.OpUnary_Operator.FLAT:
            return "flat";
        case operation_pb_js_1.OpUnary_Operator.FLOOR:
            return "floor";
        case operation_pb_js_1.OpUnary_Operator.CEIL:
            return "ceil";
        case operation_pb_js_1.OpUnary_Operator.ABORT:
            return "abort";
        default:
            throw new Error(`unexpected OperatorUnary ${op}`);
    }
}
function binaryKeyName(op) {
    switch (op) {
        case operation_pb_js_1.OpBinary_Operator.SUB:
            return "sub";
        case operation_pb_js_1.OpBinary_Operator.DIV:
            return "div";
        case operation_pb_js_1.OpBinary_Operator.EQ:
            return "eq";
        case operation_pb_js_1.OpBinary_Operator.NEQ:
            return "neq";
        case operation_pb_js_1.OpBinary_Operator.LT:
            return "lt";
        case operation_pb_js_1.OpBinary_Operator.LTE:
            return "lte";
        case operation_pb_js_1.OpBinary_Operator.GT:
            return "gt";
        case operation_pb_js_1.OpBinary_Operator.GTE:
            return "gte";
        default:
            throw new Error(`unexpected OperatorBinary ${op}`);
    }
}
function variadicKeyName(op) {
    switch (op) {
        case operation_pb_js_1.OpVariadic_Operator.ADD:
            return "add";
        case operation_pb_js_1.OpVariadic_Operator.MUL:
            return "mul";
        case operation_pb_js_1.OpVariadic_Operator.AND:
            return "and";
        case operation_pb_js_1.OpVariadic_Operator.OR:
            return "or";
        case operation_pb_js_1.OpVariadic_Operator.CAT:
            return "cat";
        case operation_pb_js_1.OpVariadic_Operator.MIN:
            return "min";
        case operation_pb_js_1.OpVariadic_Operator.MAX:
            return "max";
        case operation_pb_js_1.OpVariadic_Operator.MERGE:
            return "merge";
        default:
            throw new Error(`unexpected OperatorVariadic ${op}`);
    }
}
