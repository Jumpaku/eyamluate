import 'dart:math';

import 'package:characters/characters.dart';
import 'package:collection/collection.dart';
import 'package:eyamluate/eval/eval.dart';
import 'package:eyamluate/eval/fun_def_list.dart';
import 'package:eyamluate/eval/operation.dart';
import 'package:eyamluate/eval/path.dart';
import 'package:eyamluate/eval/validator.dart';
import 'package:eyamluate/yaml/decoder.dart';
import 'package:eyamluate/yaml/value.dart';
import 'package:eyamluate/yaml/yaml.dart';

class Evaluator {
  EvaluateOutput evaluate(EvaluateInput input) {
    final v = Decoder().decode(DecodeInput(yaml: input.source));
    if (v.isError) {
      return EvaluateOutput(
        status: EvaluateOutput_Status.DECODE_ERROR,
        errorMessage: v.errorMessage,
      );
    }
    {
      final r = Validator().validate(ValidateInput(source: input.source));
      if (r.status != ValidateOutput_Status.OK) {
        return EvaluateOutput(
          status: EvaluateOutput_Status.VALIDATE_ERROR,
          errorMessage: r.errorMessage,
        );
      }
    }
    final e = evaluateExpr(EvaluateExprInput(
        path: Path(), expr: v.value, defs: funDefListEmpty()));
    if (e.status != EvaluateExprOutput_Status.OK) {
      return EvaluateOutput(
        status: EvaluateOutput_Status.EXPR_ERROR,
        errorMessage: e.errorMessage,
        exprStatus: e.status,
        exprErrorPath: e.errorPath,
      );
    }
    return EvaluateOutput(value: e.value);
  }

  EvaluateExprOutput evaluateExpr(EvaluateExprInput input) {
    if ([Type.TYPE_NUM, Type.TYPE_STR, Type.TYPE_BOOL]
        .contains(input.expr.type)) {
      return evaluateScalar(input);
    }
    if (input.expr.type == Type.TYPE_OBJ) {
      final o = input.expr.obj;
      if (o.keys.contains("eval")) {
        return evaluateEval(input);
      } else if (o.keys.contains("obj")) {
        return evaluateObj(input);
      } else if (o.keys.contains("arr")) {
        return evaluateArr(input);
      } else if (o.keys.contains("json")) {
        return evaluateJson(input);
      } else if (o.keys.contains("for")) {
        return evaluateRangeIter(input);
      } else if (o.keys.contains("get")) {
        return evaluateGetElem(input);
      } else if (o.keys.contains("ref")) {
        return evaluateFunCall(input);
      } else if (o.keys.contains("cases")) {
        return evaluateCases(input);
      } else if (OpUnary_Operator.values
          .where((e) => e != OpUnary_Operator.UNSPECIFIED)
          .any((op) => o.keys.contains(opUnaryKeyName(op)))) {
        return evaluateOpUnary(input);
      } else if (OpBinary_Operator.values
          .where((e) => e != OpBinary_Operator.UNSPECIFIED)
          .any((op) => o.keys.contains(opBinaryKeyName(op)))) {
        return evaluateOpBinary(input);
      } else if (OpVariadic_Operator.values
          .where((e) => e != OpVariadic_Operator.UNSPECIFIED)
          .any((op) => o.keys.contains(opVariadicKeyName(op)))) {
        return evaluateOpVariadic(input);
      }
    }
    return _errorUnsupportedExpr(input.path, input.expr);
  }

  EvaluateExprOutput evaluateEval(EvaluateExprInput input) {
    final path = input.path;
    var st = input.defs;
    final where = input.expr.obj["where"];
    if (where != null) {
      final path = pathAppendKey(input.path, "where");
      for (final (pos, w) in where.arr.indexed) {
        final def = w.obj["def"]!;
        final value = w.obj["value"]!;
        final funDef = FunDef(
            def: def.str, value: value, path: pathAppendIndex(path, pos));
        final with_ = w.obj["with"];
        if (with_ != null) {
          for (final w in with_.arr) {
            funDef.with_3.add(w.str);
          }
        }
        st = funDefListRegister(st, funDef);
      }
    }
    return evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(path, "eval"),
        defs: st,
        expr: input.expr.obj["eval"]));
  }

  EvaluateExprOutput evaluateScalar(EvaluateExprInput input) {
    return EvaluateExprOutput(value: input.expr);
  }

  EvaluateExprOutput evaluateObj(EvaluateExprInput input) {
    final obj = input.expr.obj["obj"]!;
    final path = pathAppendKey(input.path, "obj");
    final v = <String, Value>{};
    for (final pos in obj.obj.keys) {
      final expr = evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(path, pos),
        defs: input.defs,
        expr: obj.obj[pos]!,
      ));
      if (expr.status != EvaluateExprOutput_Status.OK) return expr;
      v[pos] = expr.value;
    }
    return EvaluateExprOutput(value: Value(type: Type.TYPE_OBJ, obj: v));
  }

  EvaluateExprOutput evaluateArr(EvaluateExprInput input) {
    final arr = input.expr.obj["arr"]!;
    final path = pathAppendKey(input.path, "arr");
    final v = <Value>[];
    for (final (pos, val) in arr.arr.indexed) {
      final expr = evaluateExpr(EvaluateExprInput(
          path: pathAppendIndex(path, pos), defs: input.defs, expr: val));
      if (expr.status != EvaluateExprOutput_Status.OK) return expr;
      v.add(expr.value);
    }
    return EvaluateExprOutput(value: Value(type: Type.TYPE_ARR, arr: v));
  }

  EvaluateExprOutput evaluateJson(EvaluateExprInput input) {
    return EvaluateExprOutput(value: input.expr.obj["json"]!);
  }

  EvaluateExprOutput evaluateRangeIter(EvaluateExprInput input) {
    final path = input.path;
    final for_ = input.expr.obj["for"]!;
    final forPos = for_.arr[0].str;
    final forVal = for_.arr[1].str;
    final in_ = evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(path, "in"),
        defs: input.defs,
        expr: input.expr.obj["in"]));
    if (in_.status != EvaluateExprOutput_Status.OK) {
      return in_;
    }

    switch (in_.value.type) {
      case Type.TYPE_STR:
        final v = <Value>[];
        for (final (pos, val) in in_.value.str.characters.indexed) {
          var st = input.defs;
          st = funDefListRegister(
              st,
              FunDef(
                  def: forPos,
                  value: Value(type: Type.TYPE_NUM, num: pos.toDouble()),
                  path: pathAppendIndex(pathAppendKey(path, "for"), 0)));
          st = funDefListRegister(
              st,
              FunDef(
                  def: forVal,
                  value: Value(type: Type.TYPE_STR, str: val.toString()),
                  path: pathAppendIndex(pathAppendKey(path, "for"), 1)));
          final ifExpr = input.expr.obj["if"];
          if (ifExpr != null) {
            final if_ = evaluateExpr(EvaluateExprInput(
                path: pathAppendKey(path, "if"), defs: st, expr: ifExpr));
            if (if_.status != EvaluateExprOutput_Status.OK) {
              return if_;
            }
            if (if_.value.type != Type.TYPE_BOOL) {
              return _errorUnexpectedType(
                  pathAppendKey(path, "if"), [Type.TYPE_BOOL], if_.value.type);
            }
            if (!if_.value.bool_2) {
              continue;
            }
          }
          final do_ = evaluateExpr(EvaluateExprInput(
              path: pathAppendKey(path, "do"),
              defs: st,
              expr: input.expr.obj["do"]));
          if (do_.status != EvaluateExprOutput_Status.OK) {
            return do_;
          }
          v.add(do_.value);
        }
        return EvaluateExprOutput(value: Value(type: Type.TYPE_ARR, arr: v));
      case Type.TYPE_ARR:
        final v = <Value>[];
        for (final (pos, val) in in_.value.arr.indexed) {
          var st = input.defs;
          st = funDefListRegister(
              st,
              FunDef(
                  def: forPos,
                  value: Value(type: Type.TYPE_NUM, num: pos.toDouble()),
                  path: pathAppendIndex(pathAppendKey(path, "for"), 0)));
          st = funDefListRegister(
              st,
              FunDef(
                  def: forVal,
                  value: val,
                  path: pathAppendIndex(pathAppendKey(path, "for"), 1)));
          final ifExpr = input.expr.obj["if"];
          if (ifExpr != null) {
            final if_ = evaluateExpr(EvaluateExprInput(
                path: pathAppendKey(path, "if"), defs: st, expr: ifExpr));
            if (if_.status != EvaluateExprOutput_Status.OK) {
              return if_;
            }
            if (if_.value.type != Type.TYPE_BOOL) {
              return _errorUnexpectedType(
                  pathAppendKey(path, "if"), [Type.TYPE_BOOL], if_.value.type);
            }
            if (!if_.value.bool_2) {
              continue;
            }
          }
          final do_ = evaluateExpr(EvaluateExprInput(
              path: pathAppendKey(path, "do"),
              defs: st,
              expr: input.expr.obj["do"]));
          if (do_.status != EvaluateExprOutput_Status.OK) {
            return do_;
          }
          v.add(do_.value);
        }
        return EvaluateExprOutput(value: Value(type: Type.TYPE_ARR, arr: v));
      case Type.TYPE_OBJ:
        final v = <String, Value>{};
        for (final pos in in_.value.obj.keys) {
          var st = input.defs;
          st = funDefListRegister(
              st,
              FunDef(
                def: forPos,
                value: Value(type: Type.TYPE_STR, str: pos),
                path: pathAppendIndex(pathAppendKey(path, "for"), 0),
              ));
          st = funDefListRegister(
              st,
              FunDef(
                def: forVal,
                value: in_.value.obj[pos],
                path: pathAppendIndex(pathAppendKey(path, "for"), 1),
              ));
          final ifExpr = input.expr.obj["if"];
          if (ifExpr != null) {
            final if_ = evaluateExpr(EvaluateExprInput(
                path: pathAppendKey(path, "if"), defs: st, expr: ifExpr));
            if (if_.status != EvaluateExprOutput_Status.OK) {
              return if_;
            }
            if (if_.value.type != Type.TYPE_BOOL) {
              return _errorUnexpectedType(
                  pathAppendKey(path, "if"), [Type.TYPE_BOOL], if_.value.type);
            }
            if (!if_.value.bool_2) {
              continue;
            }
          }
          final do_ = evaluateExpr(EvaluateExprInput(
              path: pathAppendKey(path, "do"),
              defs: st,
              expr: input.expr.obj["do"]));
          if (do_.status != EvaluateExprOutput_Status.OK) {
            return do_;
          }
          v[pos] = do_.value;
        }
        return EvaluateExprOutput(value: Value(type: Type.TYPE_OBJ, obj: v));
      default:
        return _errorUnexpectedType(pathAppendKey(path, "in"),
            [Type.TYPE_ARR, Type.TYPE_OBJ], in_.value.type);
    }
  }

  EvaluateExprOutput evaluateGetElem(EvaluateExprInput input) {
    final path = input.path;
    final get = evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(path, "get"),
        defs: input.defs,
        expr: input.expr.obj["get"]));
    if (get.status != EvaluateExprOutput_Status.OK) {
      return get;
    }
    final from = evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(path, "from"),
        defs: input.defs,
        expr: input.expr.obj["from"]));
    if (from.status != EvaluateExprOutput_Status.OK) {
      return from;
    }

    switch (from.value.type) {
      case Type.TYPE_STR:
        if (get.value.type != Type.TYPE_NUM) {
          return _errorUnexpectedType(
              pathAppendKey(path, "get"), [Type.TYPE_NUM], get.value.type);
        }
        if (!valueCanInt(get.value)) {
          return _errorArithmeticError(pathAppendKey(path, "get"),
              "index ${get.value.num} is not an integer");
        }
        final pos = get.value.num.toInt();
        final fromChars = from.value.str.characters;
        if (pos < 0 || pos >= fromChars.length) {
          return _errorIndexOutOfBounds(
              pathAppendKey(path, "get"), 0, fromChars.length, pos);
        }
        return EvaluateExprOutput(
            value: Value(
                type: Type.TYPE_STR,
                str: fromChars.characterAt(pos).toString()));
      case Type.TYPE_ARR:
        if (get.value.type != Type.TYPE_NUM) {
          return _errorUnexpectedType(
              pathAppendKey(path, "get"), [Type.TYPE_NUM], get.value.type);
        }
        if (!valueCanInt(get.value)) {
          return _errorArithmeticError(pathAppendKey(path, "get"),
              "index ${get.value.num} is not an integer");
        }
        final pos = get.value.num.toInt();
        if (pos < 0 || pos >= from.value.arr.length) {
          return _errorIndexOutOfBounds(
              pathAppendKey(path, "get"), 0, from.value.arr.length, pos);
        }
        return EvaluateExprOutput(value: from.value.arr[pos]);
      case Type.TYPE_OBJ:
        if (get.value.type != Type.TYPE_STR) {
          return _errorUnexpectedType(
              pathAppendKey(path, "get"), [Type.TYPE_STR], get.value.type);
        }
        final pos = get.value.str;
        final val = from.value.obj[pos];
        if (val == null) {
          return _errorKeyNotFound(
              pathAppendKey(path, "get"), pos, from.value.obj.keys.toList());
        }
        return EvaluateExprOutput(value: from.value.obj[pos]);
      default:
        return _errorUnexpectedType(pathAppendKey(path, "from"),
            [Type.TYPE_STR, Type.TYPE_ARR, Type.TYPE_OBJ], from.value.type);
    }
  }

  EvaluateExprOutput evaluateFunCall(EvaluateExprInput input) {
    final path = input.path;
    final funCall = input.expr;
    final ref = funCall.obj["ref"]!;
    final funDef = funDefListFind(input.defs, ref.str);
    if (funDef == null) {
      return _errorReferenceNotFound(pathAppendKey(path, "ref"), ref.str);
    }
    var st = funDef;
    for (final argName in funDef.def.with_3) {
      final with_ = funCall.obj["with"];
      if (with_ == null) {
        return _errorKeyNotFound(path, "with", funCall.obj.keys.toList());
      }
      final argVal = with_.obj[argName];
      if (argVal == null) {
        return _errorKeyNotFound(
            pathAppendKey(path, "with"), argName, with_.obj.keys.toList());
      }
      final arg = evaluateExpr(EvaluateExprInput(
          path: pathAppendKey(pathAppendKey(path, "with"), argName),
          defs: input.defs,
          expr: argVal));
      if (arg.status != EvaluateExprOutput_Status.OK) {
        return arg;
      }
      final jsonExpr =
          Value(type: Type.TYPE_OBJ, obj: <String, Value>{"json": arg.value});
      st = funDefListRegister(
          st,
          FunDef(
              def: argName,
              value: jsonExpr,
              path: pathAppendKey(pathAppendKey(path, "with"), argName)));
    }
    return evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(path, "ref"), defs: st, expr: funDef.def.value));
  }

  EvaluateExprOutput evaluateCases(EvaluateExprInput input) {
    final cases = input.expr.obj["cases"]!;
    for (final (pos, c) in cases.arr.indexed) {
      final path = pathAppendIndex(pathAppendKey(input.path, "cases"), pos);
      if (c.obj.containsKey("when")) {
        final when = evaluateExpr(EvaluateExprInput(
            path: pathAppendKey(path, "when"),
            defs: input.defs,
            expr: c.obj["when"]));
        if (when.status != EvaluateExprOutput_Status.OK) {
          return when;
        }
        if (when.value.type != Type.TYPE_BOOL) {
          return _errorUnexpectedType(
              pathAppendKey(path, "when"), [Type.TYPE_BOOL], when.value.type);
        }
        if (when.value.bool_2) {
          final then = evaluateExpr(EvaluateExprInput(
              path: pathAppendKey(path, "then"),
              defs: input.defs,
              expr: c.obj["then"]));
          if (then.status != EvaluateExprOutput_Status.OK) {
            return then;
          }
          return then;
        }
      } else if (c.obj.containsKey("otherwise")) {
        final otherwise = evaluateExpr(EvaluateExprInput(
            path: pathAppendKey(path, "otherwise"),
            defs: input.defs,
            expr: c.obj["otherwise"]));
        if (otherwise.status != EvaluateExprOutput_Status.OK) {
          return otherwise;
        }
        return otherwise;
      }
    }
    return _errorCasesNotExhaustive(pathAppendKey(input.path, "cases"));
  }

  EvaluateExprOutput evaluateOpUnary(EvaluateExprInput input) {
    opUnaryKeyName(OpUnary_Operator.LEN);
    final operator = input.expr.obj.keys.first;
    final Value operand;
    final o = evaluateExpr(EvaluateExprInput(
        path: pathAppendKey(input.path, operator),
        defs: input.defs,
        expr: input.expr.obj[operator]));
    if (o.status != EvaluateExprOutput_Status.OK) {
      return o;
    }
    operand = o.value;
    if (operator == opUnaryKeyName(OpUnary_Operator.LEN)) {
      switch (operand.type) {
        case Type.TYPE_STR:
          return EvaluateExprOutput(
              value: Value(
                  type: Type.TYPE_NUM,
                  num: operand.str.characters.length.toDouble()));
        case Type.TYPE_ARR:
          return EvaluateExprOutput(
              value: Value(
                  type: Type.TYPE_NUM, num: operand.arr.length.toDouble()));
        case Type.TYPE_OBJ:
          return EvaluateExprOutput(
              value: Value(
                  type: Type.TYPE_NUM, num: operand.obj.length.toDouble()));
        default:
          return _errorUnexpectedType(pathAppendKey(input.path, operator),
              [Type.TYPE_STR, Type.TYPE_ARR, Type.TYPE_OBJ], operand.type);
      }
    } else if (operator == opUnaryKeyName(OpUnary_Operator.NOT)) {
      if (operand.type != Type.TYPE_BOOL) {
        return _errorUnexpectedType(pathAppendKey(input.path, operator),
            [Type.TYPE_BOOL], operand.type);
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: !operand.bool_2));
    } else if (operator == opUnaryKeyName(OpUnary_Operator.FLAT)) {
      if (operand.type != Type.TYPE_ARR) {
        return _errorUnexpectedType(
            pathAppendKey(input.path, operator), [Type.TYPE_ARR], operand.type);
      }
      final v = <Value>[];
      for (final elem in operand.arr) {
        if (elem.type != Type.TYPE_ARR) {
          return _errorUnexpectedType(
              pathAppendKey(input.path, operator), [Type.TYPE_ARR], elem.type);
        }
        v.addAll(elem.arr);
      }
      return EvaluateExprOutput(value: Value(type: Type.TYPE_ARR, arr: v));
    } else if (operator == opUnaryKeyName(OpUnary_Operator.FLOOR)) {
      if (operand.type != Type.TYPE_NUM) {
        return _errorUnexpectedType(
            pathAppendKey(input.path, operator), [Type.TYPE_NUM], operand.type);
      }
      final v = Value(type: Type.TYPE_NUM, num: operand.num.floorToDouble());
      if (!_isFiniteNumber(v)) {
        return _errorArithmeticError(pathAppendKey(input.path, operator),
            "floor(${operand.num}) is not a finite number");
      }
      return EvaluateExprOutput(value: v);
    } else if (operator == opUnaryKeyName(OpUnary_Operator.CEIL)) {
      if (operand.type != Type.TYPE_NUM) {
        return _errorUnexpectedType(
            pathAppendKey(input.path, operator), [Type.TYPE_NUM], operand.type);
      }
      final v = Value(type: Type.TYPE_NUM, num: operand.num.ceilToDouble());
      if (!_isFiniteNumber(v)) {
        return _errorArithmeticError(pathAppendKey(input.path, operator),
            "ceil(${operand.num}) is not a finite number");
      }
      return EvaluateExprOutput(value: v);
    }else if (operator == opUnaryKeyName(OpUnary_Operator.ABORT)) {
      if (operand.type != Type.TYPE_STR) {
        return _errorUnexpectedType(
            pathAppendKey(input.path, operator), [Type.TYPE_STR], operand.type);
      }
      return EvaluateExprOutput(
          status: EvaluateExprOutput_Status.ABORTED, errorMessage: operand.str);
    }
    return _errorUnsupportedOperation(input.path, operator);
  }

  EvaluateExprOutput evaluateOpBinary(EvaluateExprInput input) {
    final operator = input.expr.obj.keys.first;
    final ol = evaluateExpr(EvaluateExprInput(
        path: pathAppendIndex(pathAppendKey(input.path, operator), 0),
        defs: input.defs,
        expr: input.expr.obj[operator]!.arr[0]));
    if (ol.status != EvaluateExprOutput_Status.OK) {
      return ol;
    }
    final operandL = ol.value;
    final or = evaluateExpr(EvaluateExprInput(
        path: pathAppendIndex(pathAppendKey(input.path, operator), 1),
        defs: input.defs,
        expr: input.expr.obj[operator]!.arr[1]));
    if (or.status != EvaluateExprOutput_Status.OK) {
      return or;
    }
    final operandR = or.value;
    final pathOp = pathAppendKey(input.path, operator);
    if (operator == opBinaryKeyName(OpBinary_Operator.SUB)) {
      if (operandL.type != Type.TYPE_NUM) {
        return _errorUnexpectedType(
            pathAppendIndex(pathOp, 0), [Type.TYPE_NUM], operandL.type);
      }
      if (operandR.type != Type.TYPE_NUM) {
        return _errorUnexpectedType(
            pathAppendIndex(pathOp, 1), [Type.TYPE_NUM], operandR.type);
      }
      final v = Value(type: Type.TYPE_NUM, num: operandL.num - operandR.num);
      if (!_isFiniteNumber(v)) {
        return _errorArithmeticError(
            pathOp, "${operandL.num}-${operandR.num} is not a finite number");
      }
      return EvaluateExprOutput(value: v);
    } else if (operator == opBinaryKeyName(OpBinary_Operator.DIV)) {
      if (operandL.type != Type.TYPE_NUM) {
        return _errorUnexpectedType(
            pathAppendIndex(pathOp, 0), [Type.TYPE_NUM], operandL.type);
      }
      if (operandR.type != Type.TYPE_NUM) {
        return _errorUnexpectedType(
            pathAppendIndex(pathOp, 1), [Type.TYPE_NUM], operandR.type);
      }
      final v = Value(type: Type.TYPE_NUM, num: operandL.num / operandR.num);
      if (!_isFiniteNumber(v)) {
        return _errorArithmeticError(
            pathOp, "${operandL.num}/${operandR.num} is not a finite number");
      }
      return EvaluateExprOutput(value: v);
    } else if (operator == opBinaryKeyName(OpBinary_Operator.EQ)) {
      return equal(pathOp, operandL, operandR);
    } else if (operator == opBinaryKeyName(OpBinary_Operator.NEQ)) {
      final eq = equal(pathOp, operandL, operandR);
      if (eq.status != EvaluateExprOutput_Status.OK) {
        return eq;
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: !eq.value.bool_2));
    } else if (operator == opBinaryKeyName(OpBinary_Operator.LT)) {
      final cmp = compare(pathOp, operandL, operandR);
      if (cmp.status != EvaluateExprOutput_Status.OK) {
        return cmp;
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: cmp.value.num < 0));
    } else if (operator == opBinaryKeyName(OpBinary_Operator.LTE)) {
      final cmp = compare(pathOp, operandL, operandR);
      if (cmp.status != EvaluateExprOutput_Status.OK) {
        return cmp;
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: cmp.value.num <= 0));
    } else if (operator == opBinaryKeyName(OpBinary_Operator.GT)) {
      final cmp = compare(pathOp, operandL, operandR);
      if (cmp.status != EvaluateExprOutput_Status.OK) {
        return cmp;
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: cmp.value.num > 0));
    } else if (operator == opBinaryKeyName(OpBinary_Operator.GTE)) {
      final cmp = compare(pathOp, operandL, operandR);
      if (cmp.status != EvaluateExprOutput_Status.OK) {
        return cmp;
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: cmp.value.num >= 0));
    }
    return _errorUnsupportedOperation(input.path, operator);
  }

  EvaluateExprOutput evaluateOpVariadic(EvaluateExprInput input) {
    final operator = input.expr.obj.keys.first;
    final pathOp = pathAppendKey(input.path, operator);
    final operands = <Value>[];
    for (final (pos, o) in input.expr.obj[operator]!.arr.indexed) {
      final op = evaluateExpr(EvaluateExprInput(
          path: pathAppendIndex(pathOp, pos), defs: input.defs, expr: o));
      if (op.status != EvaluateExprOutput_Status.OK) {
        return op;
      }
      operands.add(op.value);
    }
    if (operator == opVariadicKeyName(OpVariadic_Operator.ADD)) {
      var add = 0.0;
      for (final operand in operands) {
        if (operand.type != Type.TYPE_NUM) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_NUM], operand.type);
        }
        add += operand.num;
      }
      final v = Value(type: Type.TYPE_NUM, num: add);
      if (!_isFiniteNumber(v)) {
        return _errorArithmeticError(pathOp,
            "add(${operands.map((o) => o.num.toInt().toString()).join(',')}) is not a finite number");
      }
      return EvaluateExprOutput(value: v);
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.MUL)) {
      var mul = 1.0;
      for (final operand in operands) {
        if (operand.type != Type.TYPE_NUM) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_NUM], operand.type);
        }
        mul *= operand.num;
      }
      final v = Value(type: Type.TYPE_NUM, num: mul);
      if (!_isFiniteNumber(v)) {
        return _errorArithmeticError(pathOp,
            "mul(${operands.map((e) => e.num.toInt().toString()).join(',')}) is not a finite number");
      }
      return EvaluateExprOutput(value: v);
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.AND)) {
      for (final operand in operands) {
        if (operand.type != Type.TYPE_BOOL) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_BOOL], operand.type);
        }
        if (!operand.bool_2) {
          return EvaluateExprOutput(
              value: Value(type: Type.TYPE_BOOL, bool_2: false));
        }
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: true));
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.OR)) {
      for (final operand in operands) {
        if (operand.type != Type.TYPE_BOOL) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_BOOL], operand.type);
        }
        if (operand.bool_2) {
          return EvaluateExprOutput(
              value: Value(type: Type.TYPE_BOOL, bool_2: true));
        }
      }
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: false));
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.CAT)) {
      var cat = '';
      for (final operand in operands) {
        if (operand.type != Type.TYPE_STR) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_STR], operand.type);
        }
        cat += operand.str;
      }
      return EvaluateExprOutput(value: Value(type: Type.TYPE_STR, str: cat));
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.MIN)) {
      var min_ = double.infinity;
      for (final operand in operands) {
        if (operand.type != Type.TYPE_NUM) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_NUM], operand.type);
        }
        min_ = min(min_, operand.num);
      }
      return EvaluateExprOutput(value: Value(type: Type.TYPE_NUM, num: min_));
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.MAX)) {
      var max_ = double.negativeInfinity;
      for (final operand in operands) {
        if (operand.type != Type.TYPE_NUM) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_NUM], operand.type);
        }
        max_ = max(max_, operand.num);
      }
      return EvaluateExprOutput(value: Value(type: Type.TYPE_NUM, num: max_));
    } else if (operator == opVariadicKeyName(OpVariadic_Operator.MERGE)) {
      final merge = <String, Value>{};
      for (final operand in operands) {
        if (operand.type != Type.TYPE_OBJ) {
          return _errorUnexpectedType(pathOp, [Type.TYPE_OBJ], operand.type);
        }
        for (final k in operand.obj.keys) {
          merge[k] = operand.obj[k]!;
        }
      }
      return EvaluateExprOutput(value: Value(type: Type.TYPE_OBJ, obj: merge));
    }
    return _errorUnsupportedOperation(input.path, operator);
  }
}

EvaluateExprOutput equal(Path path, Value l, Value r) {
  final falseValue =
      EvaluateExprOutput(value: Value(type: Type.TYPE_BOOL, bool_2: false));
  final trueValue =
      EvaluateExprOutput(value: Value(type: Type.TYPE_BOOL, bool_2: true));
  if (l.type != r.type) {
    return falseValue;
  }
  switch (l.type) {
    case Type.TYPE_NUM:
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: l.num == r.num));
    case Type.TYPE_BOOL:
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: l.bool_2 == r.bool_2));
    case Type.TYPE_STR:
      return EvaluateExprOutput(
          value: Value(type: Type.TYPE_BOOL, bool_2: l.str == r.str));
    case Type.TYPE_ARR:
      if (l.arr.length != r.arr.length) {
        return falseValue;
      }
      for (final (i, v) in l.arr.indexed) {
        final eq = equal(path, v, r.arr[i]);
        if (eq.value.bool_2 == false) {
          return falseValue;
        }
      }
      return trueValue;
    case Type.TYPE_OBJ:
      final lk = l.obj.keys.toList()..sort();
      final rk = r.obj.keys.toList()..sort();
      if (!ListEquality().equals(lk, rk)) {
        return falseValue;
      }
      for (final k in lk) {
        final eq = equal(path, l.obj[k]!, r.obj[k]!);
        if (eq.value.bool_2 == false) {
          return falseValue;
        }
      }
      return trueValue;
    default:
      return _errorUnexpectedType(
          path,
          [
            Type.TYPE_NUM,
            Type.TYPE_BOOL,
            Type.TYPE_STR,
            Type.TYPE_ARR,
            Type.TYPE_OBJ
          ],
          l.type);
  }
}

EvaluateExprOutput compare(Path path, Value l, Value r) {
  final ltValue =
      EvaluateExprOutput(value: Value(type: Type.TYPE_NUM, num: -1));
  final gtValue = EvaluateExprOutput(value: Value(type: Type.TYPE_NUM, num: 1));
  final eqValue = EvaluateExprOutput(value: Value(type: Type.TYPE_NUM, num: 0));
  switch ((l.type, r.type)) {
    case (Type.TYPE_NUM, Type.TYPE_NUM):
      if (l.num < r.num) {
        return ltValue;
      }
      if (l.num > r.num) {
        return gtValue;
      }
      return eqValue;
    case (Type.TYPE_BOOL, Type.TYPE_BOOL):
      if (!l.bool_2 && r.bool_2) {
        return ltValue;
      }
      if (l.bool_2 && !r.bool_2) {
        return gtValue;
      }
      return eqValue;
    case (Type.TYPE_STR, Type.TYPE_STR):
      if (l.str.compareTo(r.str) < 0) {
        return ltValue;
      }
      if (l.str.compareTo(r.str) > 0) {
        return gtValue;
      }
      return eqValue;
    case (Type.TYPE_ARR, Type.TYPE_ARR):
      final n = min(l.arr.length, r.arr.length);
      for (var i = 0; i < n; i++) {
        final cmp = compare(path, l.arr[i], r.arr[i]);
        if (cmp.status != EvaluateExprOutput_Status.OK) {
          return cmp;
        }
        if (cmp.value.num != 0) {
          return cmp;
        }
      }
      if (l.arr.length < r.arr.length) {
        return ltValue;
      }
      if (l.arr.length > r.arr.length) {
        return gtValue;
      }
      return eqValue;
    default:
      return _errorUnexpectedType(
          path,
          [
            Type.TYPE_NUM,
            Type.TYPE_BOOL,
            Type.TYPE_STR,
            Type.TYPE_ARR,
            Type.TYPE_OBJ
          ],
          l.type);
  }
}

bool _isFiniteNumber(Value v) {
  return v.type == Type.TYPE_NUM && v.num.isFinite;
}

EvaluateExprOutput _errorUnsupportedExpr(Path path, Value v) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.UNSUPPORTED_EXPR,
    errorMessage: "unsupported expr: got ${v.obj.keys}",
    errorPath: path,
  );
}

EvaluateExprOutput _errorUnexpectedType(Path path, List<Type> want, Type got) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.UNEXPECTED_TYPE,
    errorMessage: "unexpected type: want ${want}, got ${got}",
    errorPath: path,
  );
}

EvaluateExprOutput _errorArithmeticError(Path path, String message) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.ARITHMETIC_ERROR,
    errorMessage: "arithmetic error: ${message}",
    errorPath: path,
  );
}

EvaluateExprOutput _errorIndexOutOfBounds(
    Path path, int begin, int end, int index) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.INDEX_OUT_OF_BOUNDS,
    errorMessage: "index out of bounds: ${index} not in [${begin}, ${end})",
    errorPath: path,
  );
}

EvaluateExprOutput _errorKeyNotFound(Path path, String want, List<String> actual) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.KEY_NOT_FOUND,
    errorMessage: "key not found: ${want} not in {${actual.join(",")}}",
    errorPath: path,
  );
}

EvaluateExprOutput _errorReferenceNotFound(Path path, String ref) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.REFERENCE_NOT_FOUND,
    errorMessage: "reference not found: ${ref}",
    errorPath: path,
  );
}

EvaluateExprOutput _errorCasesNotExhaustive(Path path) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.CASES_NOT_EXHAUSTIVE,
    errorMessage: "cases not exhaustive",
    errorPath: path,
  );
}

EvaluateExprOutput _errorUnsupportedOperation(Path path, String gotOp) {
  return EvaluateExprOutput(
    status: EvaluateExprOutput_Status.UNSUPPORTED_OPERATION,
    errorMessage: "unsupported operation: ${gotOp}",
    errorPath: path,
  );
}
