import { EvaluateExprInputSchema, EvaluateExprOutput_Status, EvaluateExprOutputSchema, EvaluateOutput_Status, EvaluateOutputSchema, FunDefSchema, PathSchema } from "./evaluator_pb.js";
import { Type, ValueSchema } from "../yaml/value_pb.js";
import { create } from "@bufbuild/protobuf";
import { Decoder } from "../yaml/decoder.js";
import { DecodeInputSchema } from "../yaml/decoder_pb.js";
import { Validator } from "./validator.js";
import { ValidateInputSchema, ValidateOutput_Status } from "./validator_pb.js";
import { empty, find, register } from "./fun_def_list.js";
import { OpBinary_OperatorSchema, OpUnary_OperatorSchema, OpVariadic_OperatorSchema } from "./operation_pb.js";
import { append } from "./path.js";
export class BaseEvaluator {
    evaluate(input) {
        // Decode input
        const v = new Decoder().decode(create(DecodeInputSchema, { yaml: input.source }));
        if (v.isError) {
            return create(EvaluateOutputSchema, {
                status: EvaluateOutput_Status.DECODE_ERROR,
                errorMessage: v.errorMessage,
            });
        }
        // Validate input
        {
            const v = new Validator().validate(create(ValidateInputSchema, {
                source: input.source
            }));
            if (v.status != ValidateOutput_Status.OK) {
                return create(EvaluateOutputSchema, {
                    status: EvaluateOutput_Status.VALIDATE_ERROR,
                    errorMessage: v.errorMessage,
                });
            }
        }
        // Evaluate input
        const e = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: create(PathSchema, { pos: [] }),
            defs: empty(),
            expr: v.value,
        }));
        if (e.status != EvaluateExprOutput_Status.OK) {
            return create(EvaluateOutputSchema, {
                status: EvaluateOutput_Status.EXPR_ERROR,
                exprStatus: e.status,
                errorMessage: e.errorMessage,
            });
        }
        return create(EvaluateOutputSchema, { value: e.value });
    }
    evaluateExpr(input) {
        switch (input.expr?.type) {
            case Type.BOOL:
            case Type.NUM:
            case Type.STR:
                return this.evaluateScalar(input);
            case Type.OBJ:
                if ('eval' in input.expr.obj) {
                    return this.evaluateEval(input);
                }
                if ('obj' in input.expr.obj) {
                    return this.evaluateObj(input);
                }
                if ('arr' in input.expr.obj) {
                    return this.evaluateArr(input);
                }
                if ('json' in input.expr.obj) {
                    return this.evaluateJson(input);
                }
                if ('for' in input.expr.obj) {
                    return this.evaluateRangeIter(input);
                }
                if ('get' in input.expr.obj) {
                    return this.evaluateGetElem(input);
                }
                if ('ref' in input.expr.obj) {
                    return this.evaluateFunCall(input);
                }
                if ('cases' in input.expr.obj) {
                    return this.evaluateCases(input);
                }
                if (OpUnary_OperatorSchema.values.some(value => value.name in (input.expr?.obj ?? {}))) {
                    return this.evaluateOpUnary(input);
                }
                if (OpBinary_OperatorSchema.values.some(value => value.name in (input.expr?.obj ?? {}))) {
                    return this.evaluateOpBinary(input);
                }
                if (OpVariadic_OperatorSchema.values.some(value => value.name in (input.expr?.obj ?? {}))) {
                    return this.evaluateOpVariadic(input);
                }
        }
        return errorUnsupportedExpr(input.path ?? create(PathSchema), input.expr);
    }
    evaluateEval(input) {
        const path = input.path;
        const st = input.defs;
        if ('where' in input.expr.obj["where"]) {
            const where = input.expr.obj["where"];
            for (let pos = 0; pos < where.arr.length; pos++) {
                const w = where.arr[pos];
                const [def, value] = [w.obj["def"], w.obj["value"]];
                const funDef = create(FunDefSchema, {
                    def: def.str,
                    value: value,
                    path: append(append(path, "where"), pos)
                });
                if ("with" in w.obj) {
                    const ws = w.obj["with"].arr;
                    for (let pos = 0; pos < ws.length; pos++) {
                        funDef.with.push(ws[pos].str);
                    }
                }
                register(st, funDef);
            }
        }
        return this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, "eval"),
            defs: st,
            expr: input.expr.obj["eval"],
        }));
    }
    evaluateScalar(input) {
        return create(EvaluateExprOutputSchema, { value: input.expr });
    }
    evaluateObj(input) {
        const obj = input.expr.obj["obj"];
        const path = append(input.path, "obj");
        let v = {};
        for (const k in obj.obj) {
            const val = obj.obj[k];
            const expr = this.evaluateExpr(create(EvaluateExprInputSchema, {
                path: append(path, k), defs: input.defs, expr: val,
            }));
            if (expr.status != EvaluateExprOutput_Status.OK) {
                return expr;
            }
            v[k] = expr.value;
        }
        return create(EvaluateExprOutputSchema, { value: create(ValueSchema, { type: Type.OBJ, obj: v }) });
    }
    evaluateArr(input) {
        const arr = input.expr.obj["arr"];
        const path = append(input.path, "arr");
        const v = [];
        for (const pos in arr.arr) {
            const val = arr.arr[pos];
            const expr = this.evaluateExpr(create(EvaluateExprInputSchema, {
                path: append(path, pos),
                defs: input.defs,
                expr: val,
            }));
            if (expr.status !== EvaluateExprOutput_Status.OK) {
                return expr;
            }
            v.push(expr.value);
        }
        return create(EvaluateExprOutputSchema, { value: create(ValueSchema, { type: Type.ARR, arr: v }) });
    }
    evaluateJson(input) {
        return create(EvaluateExprOutputSchema, { value: input.expr.obj["json"] });
    }
    evaluateRangeIter(input) {
        const path = input.path;
        const for_ = input.expr.obj["for"];
        const [forPos, forVal] = [for_.arr[0].str, for_.arr[1].str];
        const in_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, "in"),
            defs: input.defs,
            expr: input.expr.obj["in"]
        }));
        if (in_.status !== EvaluateExprOutput_Status.OK) {
            return in_;
        }
        switch (in_.value.type) {
            case Type.STR: {
                const v = [];
                const chars = [...in_.value.str];
                for (const idx in chars) {
                    let st = input.defs;
                    st = register(input.defs, create(FunDefSchema, {
                        def: forPos,
                        value: create(ValueSchema, { type: Type.NUM, num: Number.parseInt(idx) }),
                        path: append(append(path, "for"), 0),
                    }));
                    st = register(input.defs, create(FunDefSchema, {
                        def: forVal,
                        value: create(ValueSchema, { type: Type.STR, str: chars[idx] }),
                    }));
                    if ("if" in input.expr.obj) {
                        const if_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
                            path: append(path, "if"),
                            defs: st,
                            expr: input.expr.obj["if"]
                        }));
                        if (if_.status !== EvaluateExprOutput_Status.OK) {
                            return if_;
                        }
                        if (if_.value.type !== Type.BOOL) {
                            return errorUnexpectedType(append(path, "if"), [Type.BOOL], if_.value.type);
                        }
                        if (!if_.value.bool) {
                            continue;
                        }
                    }
                    const do_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
                        path: append(path, "do"),
                        defs: st,
                        expr: input.expr.obj["do"]
                    }));
                    if (do_.status !== EvaluateExprOutput_Status.OK) {
                        return do_;
                    }
                    v.push(do_.value);
                }
                return create(EvaluateExprOutputSchema, { value: create(ValueSchema, { type: Type.ARR, arr: v }) });
            }
            case Type.ARR: {
                const v = [];
                for (const idx in in_.value.arr) {
                    let st = input.defs;
                    st = register(input.defs, create(FunDefSchema, {
                        def: forPos,
                        value: create(ValueSchema, { type: Type.NUM, num: Number.parseInt(idx) }),
                        path: append(append(path, "for"), 0),
                    }));
                    st = register(input.defs, create(FunDefSchema, {
                        def: forVal,
                        value: in_.value.arr[Number.parseInt(idx)],
                        path: append(append(path, "for"), 1),
                    }));
                    if ("if" in input.expr.obj) {
                        const if_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
                            path: append(path, "if"),
                            defs: st,
                            expr: input.expr.obj["if"]
                        }));
                        if (if_.status !== EvaluateExprOutput_Status.OK) {
                            return if_;
                        }
                        if (if_.value.type !== Type.BOOL) {
                            return errorUnexpectedType(append(path, "if"), [Type.BOOL], if_.value.type);
                        }
                        if (!if_.value.bool) {
                            continue;
                        }
                    }
                    const do_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
                        path: append(path, "do"),
                        defs: st,
                        expr: input.expr.obj["do"]
                    }));
                    if (do_.status !== EvaluateExprOutput_Status.OK) {
                        return do_;
                    }
                    v.push(do_.value);
                }
                return create(EvaluateExprOutputSchema, { value: create(ValueSchema, { type: Type.ARR, arr: v }) });
            }
            case Type.OBJ: {
                const v = {};
                for (const idx in in_.value.obj) {
                    let st = input.defs;
                    st = register(input.defs, create(FunDefSchema, {
                        def: forPos,
                        value: create(ValueSchema, { type: Type.STR, str: idx }),
                        path: append(append(path, "for"), 0),
                    }));
                    st = register(input.defs, create(FunDefSchema, {
                        def: forVal,
                        value: in_.value.obj[idx],
                        path: append(append(path, "for"), 1),
                    }));
                    if ("if" in input.expr.obj) {
                        const if_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
                            path: append(path, "if"),
                            defs: st,
                            expr: input.expr.obj["if"],
                        }));
                        if (if_.status !== EvaluateExprOutput_Status.OK) {
                            return if_;
                        }
                        if (if_.value.type !== Type.BOOL) {
                            return errorUnexpectedType(append(path, "if"), [Type.BOOL], if_.value.type);
                        }
                        if (!if_.value.bool) {
                            continue;
                        }
                    }
                    const do_ = this.evaluateExpr(create(EvaluateExprInputSchema, {
                        path: append(path, "do"),
                        defs: st,
                        expr: input.expr.obj["do"]
                    }));
                    if (do_.status !== EvaluateExprOutput_Status.OK) {
                        return do_;
                    }
                    v[idx] = do_.value;
                }
                return create(EvaluateExprOutputSchema, { value: create(ValueSchema, { type: Type.OBJ, obj: v }) });
            }
        }
        return errorUnexpectedType(path, [Type.STR, Type.ARR, Type.OBJ], in_.value.type);
    }
    evaluateGetElem(input) {
        const path = input.path;
        const get = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, "get"),
            defs: input.defs,
            expr: input.expr.obj["get"]
        }));
        if (get.status !== EvaluateExprOutput_Status.OK) {
            return get;
        }
        const from = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, "from"),
            defs: input.defs,
            expr: input.expr.obj["from"],
        }));
        if (from.status !== EvaluateExprOutput_Status.OK) {
            return from;
        }
        switch (from.value.type) {
            case Type.STR: {
                const chars = [...from.value.str];
                if (get.value.type !== Type.NUM) {
                    return errorUnexpectedType(append(path, "get"), [Type.NUM], get.value.type);
                }
                if (Number.isInteger(!get.value.num)) {
                    return errorArithmeticError(append(path, "get"), `index ${get.value.num} is not an integer`);
                }
                const pos = get.value.num;
                if (pos < 0 || pos >= from.value.str.length) {
                    return errorIndexOutOfBounds(append(path, "get"), 0, from.value.arr.length, pos);
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, {
                        type: Type.STR,
                        str: chars[pos]
                    })
                });
            }
            case Type.ARR: {
                if (get.value.type !== Type.NUM) {
                    return errorUnexpectedType(append(path, "get"), [Type.NUM], get.value.type);
                }
                if (Number.isInteger(!get.value.num)) {
                    return errorArithmeticError(append(path, "get"), `index ${get.value.num} is not an integer`);
                }
                const pos = get.value.num;
                if (pos < 0 || pos >= from.value.arr.length) {
                    return errorIndexOutOfBounds(append(path, "get"), 0, from.value.arr.length, pos);
                }
                return create(EvaluateExprOutputSchema, { value: from.value.arr[pos] });
            }
            case Type.OBJ: {
                if (get.value.type !== Type.STR) {
                    return errorUnexpectedType(append(path, "get"), [Type.STR], get.value.type);
                }
                const pos = get.value.str;
                if (!(pos in from.value.obj)) {
                    return errorKeyNotFound(append(path, "get"), pos, Object.keys(from.value.obj));
                }
                return create(EvaluateExprOutputSchema, { value: from.value.obj[pos] });
            }
            default:
                return errorUnexpectedType(append(path, "from"), [Type.STR, Type.ARR, Type.OBJ], from.value.type);
        }
    }
    evaluateFunCall(input) {
        const path = input.path;
        const funCall = input.expr.obj["ref"];
        const ref = funCall.obj["ref"];
        const funDef = find(input.defs, ref.str);
        if (funDef === null) {
            return errorReferenceNotFound(append(path, "ref"), ref.str);
        }
        let st = funDef;
        for (const argName of funDef.def.with) {
            if (!("with" in funCall.obj)) {
                return errorKeyNotFound(path, "with", Object.keys(funCall));
            }
            const with_ = funCall.obj["with"];
            if (!(argName in with_.obj)) {
                return errorKeyNotFound(append(path, "with"), argName, Object.keys(with_.obj));
            }
            const argVal = with_.obj[argName];
            const arg = this.evaluateExpr(create(EvaluateExprInputSchema, {
                path: append(append(path, "with"), argName),
                defs: input.defs,
                expr: argVal,
            }));
            if (arg.status !== EvaluateExprOutput_Status.OK) {
                return arg;
            }
            const jsonExpr = create(ValueSchema, { type: Type.OBJ, obj: { json: arg.value } });
            st = register(st, create(FunDefSchema, {
                def: argName,
                value: jsonExpr,
                path: append(append(path, "with"), argName),
            }));
        }
        return this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, "ref"),
            defs: st,
            expr: funDef.def.value,
        }));
    }
    evaluateCases(input) {
        const path = input.path;
        const cases = input.expr.obj["cases"];
        const pathCases = append(path, "cases");
        for (const pos in cases.arr) {
            const path = append(pathCases, pos);
            const c = cases.arr[pos];
            if ("when" in c.obj) {
                const when = this.evaluateExpr(create(EvaluateExprInputSchema, {
                    path: append(path, "when"),
                    defs: input.defs,
                    expr: c.obj["when"],
                }));
                if (when.status !== EvaluateExprOutput_Status.OK) {
                    return when;
                }
                if (when.value.type !== Type.BOOL) {
                    return errorUnexpectedType(append(path, "when"), [Type.BOOL], when.value.type);
                }
                if (when.value.bool) {
                    const then = this.evaluateExpr(create(EvaluateExprInputSchema, {
                        path: append(path, "then"),
                        defs: input.defs,
                        expr: c.obj["then"],
                    }));
                    if (then.status !== EvaluateExprOutput_Status.OK) {
                        return then;
                    }
                    return then;
                }
            }
            else if ("otherwise" in c.obj) {
                const otherwise = this.evaluateExpr(create(EvaluateExprInputSchema, {
                    path: append(path, "otherwise"),
                    defs: input.defs,
                    expr: c.obj["otherwise"],
                }));
                if (otherwise.status !== EvaluateExprOutput_Status.OK) {
                    return otherwise;
                }
                return otherwise;
            }
        }
        return errorCasesNotExhaustive(path);
    }
    evaluateOpUnary(input) {
        const path = input.path;
        let operator = Object.keys(input.expr.obj)[0];
        const o = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, operator),
            defs: input.defs,
            expr: input.expr.obj[operator],
        }));
        if (o.status !== EvaluateExprOutput_Status.OK) {
            return o;
        }
        const operand = o.value;
        switch (operator) {
            case "len": {
                if (operand.type === Type.STR) {
                    return create(EvaluateExprOutputSchema, {
                        value: create(ValueSchema, { type: Type.NUM, num: [...operand.str].length }),
                    });
                }
                if (operand.type === Type.ARR) {
                    return create(EvaluateExprOutputSchema, {
                        value: create(ValueSchema, { type: Type.NUM, num: operand.arr.length }),
                    });
                }
                if (operand.type === Type.OBJ) {
                    return create(EvaluateExprOutputSchema, {
                        value: create(ValueSchema, { type: Type.NUM, num: Object.keys(operand.obj).length }),
                    });
                }
                return errorUnexpectedType(path, [Type.STR, Type.ARR, Type.OBJ], operand.type);
            }
            case "not": {
                if (operand.type !== Type.BOOL) {
                    return errorUnexpectedType(append(path, "not"), [Type.BOOL], operand.type);
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: !operand.bool }),
                });
            }
            case "flat": {
                if (operand.type !== Type.ARR) {
                    return errorUnexpectedType(path, [Type.ARR], operand.type);
                }
                const v = [];
                for (const el of operand.arr) {
                    if (el.type !== Type.ARR) {
                        return errorUnexpectedType(path, [Type.ARR], el.type);
                    }
                    for (const el2 of el.arr) {
                        v.push(el2);
                    }
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.ARR, arr: v }),
                });
            }
            case "floor": {
                if (operand.type !== Type.NUM) {
                    return errorUnexpectedType(path, [Type.NUM], operand.type);
                }
                const v = create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.NUM, num: Math.floor(operand.num) }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `floor(${operand.num}) is not finite`);
                }
                return v;
            }
            case "ceil": {
                if (operand.type !== Type.NUM) {
                    return errorUnexpectedType(path, [Type.NUM], operand.type);
                }
                const v = create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.NUM, num: Math.ceil(operand.num) }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `ceil(${operand.num}) is not finite`);
                }
                return v;
            }
            case "abort": {
                if (operand.type !== Type.STR) {
                    return errorUnexpectedType(path, [Type.STR], operand.type);
                }
                return create(EvaluateExprOutputSchema, {
                    status: EvaluateExprOutput_Status.ABORTED,
                    errorMessage: operand.str,
                    errorPath: path,
                });
            }
            default:
                return errorUnsupportedOperation(path, operator);
        }
    }
    evaluateOpBinary(input) {
        const path = input.path;
        const operator = Object.keys(input.expr.obj)[0];
        const ol = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, operator),
            defs: input.defs,
            expr: input.expr.obj[operator].arr[0],
        }));
        if (ol.status !== EvaluateExprOutput_Status.OK) {
            return ol;
        }
        const or = this.evaluateExpr(create(EvaluateExprInputSchema, {
            path: append(path, operator),
            defs: input.defs,
            expr: input.expr.obj[operator].arr[1],
        }));
        if (or.status !== EvaluateExprOutput_Status.OK) {
            return or;
        }
        const [operandL, operandR] = [ol.value, or.value];
        switch (operator) {
            case "sub": {
                if (operandL.type !== Type.NUM) {
                    return errorUnexpectedType(path, [Type.NUM], operandL.type);
                }
                if (operandR.type !== Type.NUM) {
                    return errorUnexpectedType(path, [Type.NUM], operandR.type);
                }
                const v = create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.NUM, num: operandL.num - operandR.num }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `${operandL.num} - ${operandR.num} is not finite`);
                }
                return v;
            }
            case "div": {
                if (operandL.type !== Type.NUM) {
                    return errorUnexpectedType(path, [Type.NUM], operandL.type);
                }
                if (operandR.type !== Type.NUM) {
                    return errorUnexpectedType(path, [Type.NUM], operandR.type);
                }
                const v = create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.NUM, num: operandL.num / operandR.num }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `${operandL.num} / ${operandR.num} is not finite`);
                }
                return v;
            }
            case "eq": {
                return equal(path, operandL, operandR);
            }
            case "neq": {
                const eq = equal(path, operandL, operandR);
                if (eq.status !== EvaluateExprOutput_Status.OK) {
                    return eq;
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: !eq.value.bool }),
                });
            }
            case "lt": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: cmp.value.num < 0 }),
                });
            }
            case "lte": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: cmp.value.num <= 0 }),
                });
            }
            case "gt": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: cmp.value.num > 0 }),
                });
            }
            case "gte": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: cmp.value.num >= 0 }),
                });
            }
            default:
                return errorUnsupportedOperation(path, operator);
        }
    }
    evaluateOpVariadic(input) {
        const path = input.path;
        const operator = Object.keys(input.expr.obj)[0];
        const os = input.expr.obj[operator].arr;
        const operands = [];
        for (const pos in os) {
            const o = this.evaluateExpr(create(EvaluateExprInputSchema, {
                path: append(append(path, operator), pos),
                defs: input.defs,
                expr: os[pos],
            }));
            if (o.status !== EvaluateExprOutput_Status.OK) {
                return o;
            }
            operands.push(o.value);
        }
        switch (operator) {
            case "add": {
                let add = 0.0;
                for (const operand of operands) {
                    if (operand.type !== Type.NUM) {
                        return errorUnexpectedType(path, [Type.NUM], operand.type);
                    }
                    add += operand.num;
                }
                if (!Number.isFinite(add)) {
                    return errorArithmeticError(path, `add(${operands.map(o => `${o}`).join(",")}) is not a finite number`);
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.NUM, num: add }),
                });
            }
            case "mul": {
                let mul = 1.0;
                for (const operand of operands) {
                    if (operand.type !== Type.NUM) {
                        return errorUnexpectedType(path, [Type.NUM], operand.type);
                    }
                    mul *= operand.num;
                }
                if (!Number.isFinite(mul)) {
                    return errorArithmeticError(path, `mul(${operands.map(o => `${o}`).join(",")}) is not a finite number`);
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.NUM, num: mul }),
                });
            }
            case "and": {
                for (const operand of operands) {
                    if (operand.type !== Type.BOOL) {
                        return errorUnexpectedType(path, [Type.BOOL], operand.type);
                    }
                    if (!operand.bool) {
                        return create(EvaluateExprOutputSchema, {
                            value: create(ValueSchema, { type: Type.BOOL, bool: false }),
                        });
                    }
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: true }),
                });
            }
            case "or": {
                for (const operand of operands) {
                    if (operand.type !== Type.BOOL) {
                        return errorUnexpectedType(path, [Type.BOOL], operand.type);
                    }
                    if (operand.bool) {
                        return create(EvaluateExprOutputSchema, {
                            value: create(ValueSchema, { type: Type.BOOL, bool: true }),
                        });
                    }
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.BOOL, bool: false }),
                });
            }
            case "cat": {
                let cat = "";
                for (const operand of operands) {
                    if (operand.type !== Type.STR) {
                        return errorUnexpectedType(path, [Type.STR], operand.type);
                    }
                    for (const el of operand.arr) {
                        cat += el.str;
                    }
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.STR, str: cat }),
                });
            }
            case "merge": {
                const merge = {};
                for (const operand of operands) {
                    if (operand.type !== Type.OBJ) {
                        return errorUnexpectedType(path, [Type.OBJ], operand.type);
                    }
                    for (const k in operand.obj) {
                        merge[k] = operand.obj[k];
                    }
                }
                return create(EvaluateExprOutputSchema, {
                    value: create(ValueSchema, { type: Type.OBJ, obj: merge }),
                });
            }
            default:
                return errorUnsupportedOperation(path, operator);
        }
    }
}
function equal(path, l, r) {
    const falseValue = create(EvaluateExprOutputSchema, {
        value: create(ValueSchema, { type: Type.BOOL, bool: false }),
    });
    const trueValue = create(EvaluateExprOutputSchema, {
        value: create(ValueSchema, { type: Type.BOOL, bool: true }),
    });
    if (l.type !== r.type) {
        return falseValue;
    }
    if (l.type === Type.NUM) {
        return create(EvaluateExprOutputSchema, {
            value: create(ValueSchema, { type: Type.BOOL, bool: l.num === r.num }),
        });
    }
    if (l.type === Type.BOOL) {
        return create(EvaluateExprOutputSchema, {
            value: create(ValueSchema, { type: Type.BOOL, bool: l.bool === r.bool }),
        });
    }
    if (l.type === Type.STR) {
        return create(EvaluateExprOutputSchema, {
            value: create(ValueSchema, { type: Type.BOOL, bool: l.str === r.str }),
        });
    }
    if (l.type === Type.ARR) {
        if (l.arr.length !== r.arr.length) {
            return falseValue;
        }
        for (let i = 0; i < l.arr.length; i++) {
            const el = l.arr[i];
            const er = r.arr[i];
            const eq = equal(path, el, er);
            if (eq.status !== EvaluateExprOutput_Status.OK) {
                return eq;
            }
            if (!(eq.value?.bool ?? false)) {
                return falseValue;
            }
        }
        return trueValue;
    }
    if (l.type === Type.OBJ) {
        const lk = Object.keys(l.obj);
        const rk = Object.keys(r.obj);
        lk.sort();
        rk.sort();
        for (const k of lk) {
            if (!(k in r.obj)) {
                return falseValue;
            }
            const eq = equal(path, l.obj[k], r.obj[k]);
            if (eq.status !== EvaluateExprOutput_Status.OK) {
                return eq;
            }
            if (!(eq.value?.bool ?? false)) {
                return falseValue;
            }
        }
        return trueValue;
    }
    return errorUnexpectedType(path, [Type.NUM, Type.BOOL, Type.STR, Type.ARR, Type.OBJ], l.type);
}
function compare(path, l, r) {
    const ltValue = create(EvaluateExprOutputSchema, {
        value: create(ValueSchema, { type: Type.NUM, num: -1 }),
    });
    const gtValue = create(EvaluateExprOutputSchema, {
        value: create(ValueSchema, { type: Type.NUM, num: 1 }),
    });
    const eqValue = create(EvaluateExprOutputSchema, {
        value: create(ValueSchema, { type: Type.NUM, num: 0 }),
    });
    if (l.type === Type.NUM && r.type === Type.NUM) {
        if (l.num < r.num) {
            return ltValue;
        }
        if (l.num > r.num) {
            return gtValue;
        }
        return eqValue;
    }
    if (l.type === Type.BOOL && r.type === Type.BOOL) {
        if (!l.bool && r.bool) {
            return ltValue;
        }
        if (l.bool && !r.bool) {
            return gtValue;
        }
        return eqValue;
    }
    if (l.type === Type.STR && r.type === Type.STR) {
        if (l.str < r.str) {
            return ltValue;
        }
        if (l.str > r.str) {
            return gtValue;
        }
        return eqValue;
    }
    if (l.type === Type.ARR && r.type === Type.ARR) {
        const n = Math.min(l.arr.length, r.arr.length);
        for (let i = 0; i < n; i++) {
            const cmp = compare(path, l.arr[i], r.arr[i]);
            if (cmp.status !== EvaluateExprOutput_Status.OK) {
                return cmp;
            }
            if ((cmp.value?.num ?? 0) === 0) {
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
    }
    if (l.type !== r.type) {
        return errorUnexpectedType(path, [l.type], r.type);
    }
    return errorUnexpectedType(path, [Type.NUM, Type.BOOL, Type.STR, Type.ARR], l.type);
}
function errorUnsupportedExpr(path, v) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.UNSUPPORTED_EXPR,
        errorMessage: `unsupported expr: got ${Object.keys(v.obj)}`,
        errorPath: path,
    });
}
function errorUnexpectedType(path, want, got) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.UNEXPECTED_TYPE,
        errorMessage: `unexpected type: want ${want}, got ${got}`,
        errorPath: path,
    });
}
function errorArithmeticError(path, message) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.ARITHMETIC_ERROR,
        errorMessage: `arithmetic error: ${message}`,
        errorPath: path,
    });
}
function errorIndexOutOfBounds(path, begin, end, index) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.INDEX_OUT_OF_BOUNDS,
        errorMessage: `index out of bounds: ${index} not in [${begin}, ${end})`,
        errorPath: path,
    });
}
function errorKeyNotFound(path, want, actual) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.KEY_NOT_FOUND,
        errorMessage: `key not found: ${want} not in {${actual.join(",")}}`,
        errorPath: path,
    });
}
function errorReferenceNotFound(path, ref) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.REFERENCE_NOT_FOUND,
        errorMessage: `reference not found: ${ref}`,
        errorPath: path,
    });
}
function errorCasesNotExhaustive(path) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.CASES_NOT_EXHAUSTIVE,
        errorMessage: "cases not exhaustive",
        errorPath: path,
    });
}
function errorUnsupportedOperation(path, gotOp) {
    return create(EvaluateExprOutputSchema, {
        status: EvaluateExprOutput_Status.UNSUPPORTED_OPERATION,
        errorMessage: `unsupported operation: ${gotOp}`,
        errorPath: path,
    });
}
