import 'package:eyamluate/eval/eval.dart';

String opUnaryKeyName(OpUnary_Operator o) {
  return switch (o) {
    OpUnary_Operator.LEN => "len",
    OpUnary_Operator.NOT => "not",
    OpUnary_Operator.FLAT => "flat",
    OpUnary_Operator.FLOOR => "floor",
    OpUnary_Operator.CEIL => "ceil",
    OpUnary_Operator.ABORT => "abort",
    _ => throw "unexpected OperatorUnary $o",
  };
}

String opBinaryKeyName(OpBinary_Operator o) {
  return switch (o) {
    OpBinary_Operator.SUB => "sub",
    OpBinary_Operator.DIV => "div",
    OpBinary_Operator.EQ => "eq",
    OpBinary_Operator.NEQ => "neq",
    OpBinary_Operator.LT => "lt",
    OpBinary_Operator.LTE => "lte",
    OpBinary_Operator.GT => "gt",
    OpBinary_Operator.GTE => "gte",
    _ => throw "unexpected OperatorBinary $o",
  };
}

String opVariadicKeyName(OpVariadic_Operator o) {
  return switch (o) {
    OpVariadic_Operator.ADD => "add",
    OpVariadic_Operator.MUL => "mul",
    OpVariadic_Operator.AND => "and",
    OpVariadic_Operator.OR => "or",
    OpVariadic_Operator.CAT => "cat",
    OpVariadic_Operator.MIN => "min",
    OpVariadic_Operator.MAX => "max",
    OpVariadic_Operator.MERGE => "merge",
    _ => throw "unexpected OperatorVariadic $o",
  };
}
