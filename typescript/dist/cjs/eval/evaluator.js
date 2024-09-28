"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.BaseEvaluator = void 0;
const evaluator_pb_js_1 = require("./evaluator_pb.js");
const value_pb_js_1 = require("../yaml/value_pb.js");
const protobuf_1 = require("@bufbuild/protobuf");
const decoder_js_1 = require("../yaml/decoder.js");
const decoder_pb_js_1 = require("../yaml/decoder_pb.js");
const validator_js_1 = require("./validator.js");
const validator_pb_js_1 = require("./validator_pb.js");
const fun_def_list_js_1 = require("./fun_def_list.js");
const operation_pb_js_1 = require("./operation_pb.js");
const path_js_1 = require("./path.js");
class BaseEvaluator {
    evaluate(input) {
        // Decode input
        const v = new decoder_js_1.Decoder().decode((0, protobuf_1.create)(decoder_pb_js_1.DecodeInputSchema, { yaml: input.source }));
        if (v.isError) {
            return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateOutputSchema, {
                status: evaluator_pb_js_1.EvaluateOutput_Status.DECODE_ERROR,
                errorMessage: v.errorMessage,
            });
        }
        // Validate input
        {
            const v = new validator_js_1.Validator().validate((0, protobuf_1.create)(validator_pb_js_1.ValidateInputSchema, {
                source: input.source
            }));
            if (v.status != validator_pb_js_1.ValidateOutput_Status.OK) {
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateOutputSchema, {
                    status: evaluator_pb_js_1.EvaluateOutput_Status.VALIDATE_ERROR,
                    errorMessage: v.errorMessage,
                });
            }
        }
        // Evaluate input
        const e = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, protobuf_1.create)(evaluator_pb_js_1.PathSchema, { pos: [] }),
            defs: (0, fun_def_list_js_1.empty)(),
            expr: v.value,
        }));
        if (e.status != evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateOutputSchema, {
                status: evaluator_pb_js_1.EvaluateOutput_Status.EXPR_ERROR,
                exprStatus: e.status,
                errorMessage: e.errorMessage,
            });
        }
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateOutputSchema, { value: e.value });
    }
    evaluateExpr(input) {
        switch (input.expr?.type) {
            case value_pb_js_1.Type.BOOL:
            case value_pb_js_1.Type.NUM:
            case value_pb_js_1.Type.STR:
                return this.evaluateScalar(input);
            case value_pb_js_1.Type.OBJ:
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
                if (operation_pb_js_1.OpUnary_OperatorSchema.values.some(value => value.name in (input.expr?.obj ?? {}))) {
                    return this.evaluateOpUnary(input);
                }
                if (operation_pb_js_1.OpBinary_OperatorSchema.values.some(value => value.name in (input.expr?.obj ?? {}))) {
                    return this.evaluateOpBinary(input);
                }
                if (operation_pb_js_1.OpVariadic_OperatorSchema.values.some(value => value.name in (input.expr?.obj ?? {}))) {
                    return this.evaluateOpVariadic(input);
                }
        }
        return errorUnsupportedExpr(input.path ?? (0, protobuf_1.create)(evaluator_pb_js_1.PathSchema), input.expr);
    }
    evaluateEval(input) {
        const path = input.path;
        const st = input.defs;
        if ('where' in input.expr.obj["where"]) {
            const where = input.expr.obj["where"];
            for (let pos = 0; pos < where.arr.length; pos++) {
                const w = where.arr[pos];
                const [def, value] = [w.obj["def"], w.obj["value"]];
                const funDef = (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                    def: def.str,
                    value: value,
                    path: (0, path_js_1.append)((0, path_js_1.append)(path, "where"), pos)
                });
                if ("with" in w.obj) {
                    const ws = w.obj["with"].arr;
                    for (let pos = 0; pos < ws.length; pos++) {
                        funDef.with.push(ws[pos].str);
                    }
                }
                (0, fun_def_list_js_1.register)(st, funDef);
            }
        }
        return this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, "eval"),
            defs: st,
            expr: input.expr.obj["eval"],
        }));
    }
    evaluateScalar(input) {
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: input.expr });
    }
    evaluateObj(input) {
        const obj = input.expr.obj["obj"];
        const path = (0, path_js_1.append)(input.path, "obj");
        let v = {};
        for (const k in obj.obj) {
            const val = obj.obj[k];
            const expr = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                path: (0, path_js_1.append)(path, k), defs: input.defs, expr: val,
            }));
            if (expr.status != evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                return expr;
            }
            v[k] = expr.value;
        }
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.OBJ, obj: v }) });
    }
    evaluateArr(input) {
        const arr = input.expr.obj["arr"];
        const path = (0, path_js_1.append)(input.path, "arr");
        const v = [];
        for (const pos in arr.arr) {
            const val = arr.arr[pos];
            const expr = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                path: (0, path_js_1.append)(path, pos),
                defs: input.defs,
                expr: val,
            }));
            if (expr.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                return expr;
            }
            v.push(expr.value);
        }
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.ARR, arr: v }) });
    }
    evaluateJson(input) {
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: input.expr.obj["json"] });
    }
    evaluateRangeIter(input) {
        const path = input.path;
        const for_ = input.expr.obj["for"];
        const [forPos, forVal] = [for_.arr[0].str, for_.arr[1].str];
        const in_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, "in"),
            defs: input.defs,
            expr: input.expr.obj["in"]
        }));
        if (in_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return in_;
        }
        switch (in_.value.type) {
            case value_pb_js_1.Type.STR: {
                const v = [];
                const chars = [...in_.value.str];
                for (const idx in chars) {
                    let st = input.defs;
                    st = (0, fun_def_list_js_1.register)(input.defs, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                        def: forPos,
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: Number.parseInt(idx) }),
                        path: (0, path_js_1.append)((0, path_js_1.append)(path, "for"), 0),
                    }));
                    st = (0, fun_def_list_js_1.register)(input.defs, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                        def: forVal,
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.STR, str: chars[idx] }),
                    }));
                    if ("if" in input.expr.obj) {
                        const if_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                            path: (0, path_js_1.append)(path, "if"),
                            defs: st,
                            expr: input.expr.obj["if"]
                        }));
                        if (if_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                            return if_;
                        }
                        if (if_.value.type !== value_pb_js_1.Type.BOOL) {
                            return errorUnexpectedType((0, path_js_1.append)(path, "if"), [value_pb_js_1.Type.BOOL], if_.value.type);
                        }
                        if (!if_.value.bool) {
                            continue;
                        }
                    }
                    const do_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                        path: (0, path_js_1.append)(path, "do"),
                        defs: st,
                        expr: input.expr.obj["do"]
                    }));
                    if (do_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                        return do_;
                    }
                    v.push(do_.value);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.ARR, arr: v }) });
            }
            case value_pb_js_1.Type.ARR: {
                const v = [];
                for (const idx in in_.value.arr) {
                    let st = input.defs;
                    st = (0, fun_def_list_js_1.register)(input.defs, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                        def: forPos,
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: Number.parseInt(idx) }),
                        path: (0, path_js_1.append)((0, path_js_1.append)(path, "for"), 0),
                    }));
                    st = (0, fun_def_list_js_1.register)(input.defs, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                        def: forVal,
                        value: in_.value.arr[Number.parseInt(idx)],
                        path: (0, path_js_1.append)((0, path_js_1.append)(path, "for"), 1),
                    }));
                    if ("if" in input.expr.obj) {
                        const if_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                            path: (0, path_js_1.append)(path, "if"),
                            defs: st,
                            expr: input.expr.obj["if"]
                        }));
                        if (if_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                            return if_;
                        }
                        if (if_.value.type !== value_pb_js_1.Type.BOOL) {
                            return errorUnexpectedType((0, path_js_1.append)(path, "if"), [value_pb_js_1.Type.BOOL], if_.value.type);
                        }
                        if (!if_.value.bool) {
                            continue;
                        }
                    }
                    const do_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                        path: (0, path_js_1.append)(path, "do"),
                        defs: st,
                        expr: input.expr.obj["do"]
                    }));
                    if (do_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                        return do_;
                    }
                    v.push(do_.value);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.ARR, arr: v }) });
            }
            case value_pb_js_1.Type.OBJ: {
                const v = {};
                for (const idx in in_.value.obj) {
                    let st = input.defs;
                    st = (0, fun_def_list_js_1.register)(input.defs, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                        def: forPos,
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.STR, str: idx }),
                        path: (0, path_js_1.append)((0, path_js_1.append)(path, "for"), 0),
                    }));
                    st = (0, fun_def_list_js_1.register)(input.defs, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                        def: forVal,
                        value: in_.value.obj[idx],
                        path: (0, path_js_1.append)((0, path_js_1.append)(path, "for"), 1),
                    }));
                    if ("if" in input.expr.obj) {
                        const if_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                            path: (0, path_js_1.append)(path, "if"),
                            defs: st,
                            expr: input.expr.obj["if"],
                        }));
                        if (if_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                            return if_;
                        }
                        if (if_.value.type !== value_pb_js_1.Type.BOOL) {
                            return errorUnexpectedType((0, path_js_1.append)(path, "if"), [value_pb_js_1.Type.BOOL], if_.value.type);
                        }
                        if (!if_.value.bool) {
                            continue;
                        }
                    }
                    const do_ = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                        path: (0, path_js_1.append)(path, "do"),
                        defs: st,
                        expr: input.expr.obj["do"]
                    }));
                    if (do_.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                        return do_;
                    }
                    v[idx] = do_.value;
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.OBJ, obj: v }) });
            }
        }
        return errorUnexpectedType(path, [value_pb_js_1.Type.STR, value_pb_js_1.Type.ARR, value_pb_js_1.Type.OBJ], in_.value.type);
    }
    evaluateGetElem(input) {
        const path = input.path;
        const get = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, "get"),
            defs: input.defs,
            expr: input.expr.obj["get"]
        }));
        if (get.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return get;
        }
        const from = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, "from"),
            defs: input.defs,
            expr: input.expr.obj["from"],
        }));
        if (from.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return from;
        }
        switch (from.value.type) {
            case value_pb_js_1.Type.STR: {
                const chars = [...from.value.str];
                if (get.value.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType((0, path_js_1.append)(path, "get"), [value_pb_js_1.Type.NUM], get.value.type);
                }
                if (Number.isInteger(!get.value.num)) {
                    return errorArithmeticError((0, path_js_1.append)(path, "get"), `index ${get.value.num} is not an integer`);
                }
                const pos = get.value.num;
                if (pos < 0 || pos >= from.value.str.length) {
                    return errorIndexOutOfBounds((0, path_js_1.append)(path, "get"), 0, from.value.arr.length, pos);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, {
                        type: value_pb_js_1.Type.STR,
                        str: chars[pos]
                    })
                });
            }
            case value_pb_js_1.Type.ARR: {
                if (get.value.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType((0, path_js_1.append)(path, "get"), [value_pb_js_1.Type.NUM], get.value.type);
                }
                if (Number.isInteger(!get.value.num)) {
                    return errorArithmeticError((0, path_js_1.append)(path, "get"), `index ${get.value.num} is not an integer`);
                }
                const pos = get.value.num;
                if (pos < 0 || pos >= from.value.arr.length) {
                    return errorIndexOutOfBounds((0, path_js_1.append)(path, "get"), 0, from.value.arr.length, pos);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: from.value.arr[pos] });
            }
            case value_pb_js_1.Type.OBJ: {
                if (get.value.type !== value_pb_js_1.Type.STR) {
                    return errorUnexpectedType((0, path_js_1.append)(path, "get"), [value_pb_js_1.Type.STR], get.value.type);
                }
                const pos = get.value.str;
                if (!(pos in from.value.obj)) {
                    return errorKeyNotFound((0, path_js_1.append)(path, "get"), pos, Object.keys(from.value.obj));
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, { value: from.value.obj[pos] });
            }
            default:
                return errorUnexpectedType((0, path_js_1.append)(path, "from"), [value_pb_js_1.Type.STR, value_pb_js_1.Type.ARR, value_pb_js_1.Type.OBJ], from.value.type);
        }
    }
    evaluateFunCall(input) {
        const path = input.path;
        const funCall = input.expr.obj["ref"];
        const ref = funCall.obj["ref"];
        const funDef = (0, fun_def_list_js_1.find)(input.defs, ref.str);
        if (funDef === null) {
            return errorReferenceNotFound((0, path_js_1.append)(path, "ref"), ref.str);
        }
        let st = funDef;
        for (const argName of funDef.def.with) {
            if (!("with" in funCall.obj)) {
                return errorKeyNotFound(path, "with", Object.keys(funCall));
            }
            const with_ = funCall.obj["with"];
            if (!(argName in with_.obj)) {
                return errorKeyNotFound((0, path_js_1.append)(path, "with"), argName, Object.keys(with_.obj));
            }
            const argVal = with_.obj[argName];
            const arg = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                path: (0, path_js_1.append)((0, path_js_1.append)(path, "with"), argName),
                defs: input.defs,
                expr: argVal,
            }));
            if (arg.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                return arg;
            }
            const jsonExpr = (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.OBJ, obj: { json: arg.value } });
            st = (0, fun_def_list_js_1.register)(st, (0, protobuf_1.create)(evaluator_pb_js_1.FunDefSchema, {
                def: argName,
                value: jsonExpr,
                path: (0, path_js_1.append)((0, path_js_1.append)(path, "with"), argName),
            }));
        }
        return this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, "ref"),
            defs: st,
            expr: funDef.def.value,
        }));
    }
    evaluateCases(input) {
        const path = input.path;
        const cases = input.expr.obj["cases"];
        const pathCases = (0, path_js_1.append)(path, "cases");
        for (const pos in cases.arr) {
            const path = (0, path_js_1.append)(pathCases, pos);
            const c = cases.arr[pos];
            if ("when" in c.obj) {
                const when = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                    path: (0, path_js_1.append)(path, "when"),
                    defs: input.defs,
                    expr: c.obj["when"],
                }));
                if (when.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                    return when;
                }
                if (when.value.type !== value_pb_js_1.Type.BOOL) {
                    return errorUnexpectedType((0, path_js_1.append)(path, "when"), [value_pb_js_1.Type.BOOL], when.value.type);
                }
                if (when.value.bool) {
                    const then = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                        path: (0, path_js_1.append)(path, "then"),
                        defs: input.defs,
                        expr: c.obj["then"],
                    }));
                    if (then.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                        return then;
                    }
                    return then;
                }
            }
            else if ("otherwise" in c.obj) {
                const otherwise = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                    path: (0, path_js_1.append)(path, "otherwise"),
                    defs: input.defs,
                    expr: c.obj["otherwise"],
                }));
                if (otherwise.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
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
        const o = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, operator),
            defs: input.defs,
            expr: input.expr.obj[operator],
        }));
        if (o.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return o;
        }
        const operand = o.value;
        switch (operator) {
            case "len": {
                if (operand.type === value_pb_js_1.Type.STR) {
                    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: [...operand.str].length }),
                    });
                }
                if (operand.type === value_pb_js_1.Type.ARR) {
                    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: operand.arr.length }),
                    });
                }
                if (operand.type === value_pb_js_1.Type.OBJ) {
                    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: Object.keys(operand.obj).length }),
                    });
                }
                return errorUnexpectedType(path, [value_pb_js_1.Type.STR, value_pb_js_1.Type.ARR, value_pb_js_1.Type.OBJ], operand.type);
            }
            case "not": {
                if (operand.type !== value_pb_js_1.Type.BOOL) {
                    return errorUnexpectedType((0, path_js_1.append)(path, "not"), [value_pb_js_1.Type.BOOL], operand.type);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: !operand.bool }),
                });
            }
            case "flat": {
                if (operand.type !== value_pb_js_1.Type.ARR) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.ARR], operand.type);
                }
                const v = [];
                for (const el of operand.arr) {
                    if (el.type !== value_pb_js_1.Type.ARR) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.ARR], el.type);
                    }
                    for (const el2 of el.arr) {
                        v.push(el2);
                    }
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.ARR, arr: v }),
                });
            }
            case "floor": {
                if (operand.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operand.type);
                }
                const v = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: Math.floor(operand.num) }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `floor(${operand.num}) is not finite`);
                }
                return v;
            }
            case "ceil": {
                if (operand.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operand.type);
                }
                const v = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: Math.ceil(operand.num) }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `ceil(${operand.num}) is not finite`);
                }
                return v;
            }
            case "abort": {
                if (operand.type !== value_pb_js_1.Type.STR) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.STR], operand.type);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    status: evaluator_pb_js_1.EvaluateExprOutput_Status.ABORTED,
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
        const ol = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, operator),
            defs: input.defs,
            expr: input.expr.obj[operator].arr[0],
        }));
        if (ol.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return ol;
        }
        const or = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
            path: (0, path_js_1.append)(path, operator),
            defs: input.defs,
            expr: input.expr.obj[operator].arr[1],
        }));
        if (or.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
            return or;
        }
        const [operandL, operandR] = [ol.value, or.value];
        switch (operator) {
            case "sub": {
                if (operandL.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operandL.type);
                }
                if (operandR.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operandR.type);
                }
                const v = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: operandL.num - operandR.num }),
                });
                if (!Number.isFinite(v.value.num)) {
                    return errorArithmeticError(path, `${operandL.num} - ${operandR.num} is not finite`);
                }
                return v;
            }
            case "div": {
                if (operandL.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operandL.type);
                }
                if (operandR.type !== value_pb_js_1.Type.NUM) {
                    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operandR.type);
                }
                const v = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: operandL.num / operandR.num }),
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
                if (eq.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                    return eq;
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: !eq.value.bool }),
                });
            }
            case "lt": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: cmp.value.num < 0 }),
                });
            }
            case "lte": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: cmp.value.num <= 0 }),
                });
            }
            case "gt": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: cmp.value.num > 0 }),
                });
            }
            case "gte": {
                const cmp = compare(path, operandL, operandR);
                if (cmp.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                    return cmp;
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: cmp.value.num >= 0 }),
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
            const o = this.evaluateExpr((0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprInputSchema, {
                path: (0, path_js_1.append)((0, path_js_1.append)(path, operator), pos),
                defs: input.defs,
                expr: os[pos],
            }));
            if (o.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                return o;
            }
            operands.push(o.value);
        }
        switch (operator) {
            case "add": {
                let add = 0.0;
                for (const operand of operands) {
                    if (operand.type !== value_pb_js_1.Type.NUM) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operand.type);
                    }
                    add += operand.num;
                }
                if (!Number.isFinite(add)) {
                    return errorArithmeticError(path, `add(${operands.map(o => `${o}`).join(",")}) is not a finite number`);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: add }),
                });
            }
            case "mul": {
                let mul = 1.0;
                for (const operand of operands) {
                    if (operand.type !== value_pb_js_1.Type.NUM) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.NUM], operand.type);
                    }
                    mul *= operand.num;
                }
                if (!Number.isFinite(mul)) {
                    return errorArithmeticError(path, `mul(${operands.map(o => `${o}`).join(",")}) is not a finite number`);
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: mul }),
                });
            }
            case "and": {
                for (const operand of operands) {
                    if (operand.type !== value_pb_js_1.Type.BOOL) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.BOOL], operand.type);
                    }
                    if (!operand.bool) {
                        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                            value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: false }),
                        });
                    }
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: true }),
                });
            }
            case "or": {
                for (const operand of operands) {
                    if (operand.type !== value_pb_js_1.Type.BOOL) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.BOOL], operand.type);
                    }
                    if (operand.bool) {
                        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                            value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: true }),
                        });
                    }
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: false }),
                });
            }
            case "cat": {
                let cat = "";
                for (const operand of operands) {
                    if (operand.type !== value_pb_js_1.Type.STR) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.STR], operand.type);
                    }
                    for (const el of operand.arr) {
                        cat += el.str;
                    }
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.STR, str: cat }),
                });
            }
            case "merge": {
                const merge = {};
                for (const operand of operands) {
                    if (operand.type !== value_pb_js_1.Type.OBJ) {
                        return errorUnexpectedType(path, [value_pb_js_1.Type.OBJ], operand.type);
                    }
                    for (const k in operand.obj) {
                        merge[k] = operand.obj[k];
                    }
                }
                return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
                    value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.OBJ, obj: merge }),
                });
            }
            default:
                return errorUnsupportedOperation(path, operator);
        }
    }
}
exports.BaseEvaluator = BaseEvaluator;
function equal(path, l, r) {
    const falseValue = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: false }),
    });
    const trueValue = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: true }),
    });
    if (l.type !== r.type) {
        return falseValue;
    }
    if (l.type === value_pb_js_1.Type.NUM) {
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
            value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: l.num === r.num }),
        });
    }
    if (l.type === value_pb_js_1.Type.BOOL) {
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
            value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: l.bool === r.bool }),
        });
    }
    if (l.type === value_pb_js_1.Type.STR) {
        return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
            value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: l.str === r.str }),
        });
    }
    if (l.type === value_pb_js_1.Type.ARR) {
        if (l.arr.length !== r.arr.length) {
            return falseValue;
        }
        for (let i = 0; i < l.arr.length; i++) {
            const el = l.arr[i];
            const er = r.arr[i];
            const eq = equal(path, el, er);
            if (eq.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                return eq;
            }
            if (!(eq.value?.bool ?? false)) {
                return falseValue;
            }
        }
        return trueValue;
    }
    if (l.type === value_pb_js_1.Type.OBJ) {
        const lk = Object.keys(l.obj);
        const rk = Object.keys(r.obj);
        lk.sort();
        rk.sort();
        for (const k of lk) {
            if (!(k in r.obj)) {
                return falseValue;
            }
            const eq = equal(path, l.obj[k], r.obj[k]);
            if (eq.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
                return eq;
            }
            if (!(eq.value?.bool ?? false)) {
                return falseValue;
            }
        }
        return trueValue;
    }
    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM, value_pb_js_1.Type.BOOL, value_pb_js_1.Type.STR, value_pb_js_1.Type.ARR, value_pb_js_1.Type.OBJ], l.type);
}
function compare(path, l, r) {
    const ltValue = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: -1 }),
    });
    const gtValue = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: 1 }),
    });
    const eqValue = (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        value: (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: 0 }),
    });
    if (l.type === value_pb_js_1.Type.NUM && r.type === value_pb_js_1.Type.NUM) {
        if (l.num < r.num) {
            return ltValue;
        }
        if (l.num > r.num) {
            return gtValue;
        }
        return eqValue;
    }
    if (l.type === value_pb_js_1.Type.BOOL && r.type === value_pb_js_1.Type.BOOL) {
        if (!l.bool && r.bool) {
            return ltValue;
        }
        if (l.bool && !r.bool) {
            return gtValue;
        }
        return eqValue;
    }
    if (l.type === value_pb_js_1.Type.STR && r.type === value_pb_js_1.Type.STR) {
        if (l.str < r.str) {
            return ltValue;
        }
        if (l.str > r.str) {
            return gtValue;
        }
        return eqValue;
    }
    if (l.type === value_pb_js_1.Type.ARR && r.type === value_pb_js_1.Type.ARR) {
        const n = Math.min(l.arr.length, r.arr.length);
        for (let i = 0; i < n; i++) {
            const cmp = compare(path, l.arr[i], r.arr[i]);
            if (cmp.status !== evaluator_pb_js_1.EvaluateExprOutput_Status.OK) {
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
    return errorUnexpectedType(path, [value_pb_js_1.Type.NUM, value_pb_js_1.Type.BOOL, value_pb_js_1.Type.STR, value_pb_js_1.Type.ARR], l.type);
}
function errorUnsupportedExpr(path, v) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.UNSUPPORTED_EXPR,
        errorMessage: `unsupported expr: got ${Object.keys(v.obj)}`,
        errorPath: path,
    });
}
function errorUnexpectedType(path, want, got) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.UNEXPECTED_TYPE,
        errorMessage: `unexpected type: want ${want}, got ${got}`,
        errorPath: path,
    });
}
function errorArithmeticError(path, message) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.ARITHMETIC_ERROR,
        errorMessage: `arithmetic error: ${message}`,
        errorPath: path,
    });
}
function errorIndexOutOfBounds(path, begin, end, index) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.INDEX_OUT_OF_BOUNDS,
        errorMessage: `index out of bounds: ${index} not in [${begin}, ${end})`,
        errorPath: path,
    });
}
function errorKeyNotFound(path, want, actual) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.KEY_NOT_FOUND,
        errorMessage: `key not found: ${want} not in {${actual.join(",")}}`,
        errorPath: path,
    });
}
function errorReferenceNotFound(path, ref) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.REFERENCE_NOT_FOUND,
        errorMessage: `reference not found: ${ref}`,
        errorPath: path,
    });
}
function errorCasesNotExhaustive(path) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.CASES_NOT_EXHAUSTIVE,
        errorMessage: "cases not exhaustive",
        errorPath: path,
    });
}
function errorUnsupportedOperation(path, gotOp) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.EvaluateExprOutputSchema, {
        status: evaluator_pb_js_1.EvaluateExprOutput_Status.UNSUPPORTED_OPERATION,
        errorMessage: `unsupported operation: ${gotOp}`,
        errorPath: path,
    });
}
