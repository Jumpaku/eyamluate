"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Decoder = void 0;
const protobuf_1 = require("@bufbuild/protobuf");
const decoder_pb_js_1 = require("./decoder_pb.js");
const yaml_1 = __importDefault(require("yaml"));
const value_pb_js_1 = require("./value_pb.js");
class Decoder {
    decode(input) {
        try {
            const o = yaml_1.default.parse(input.yaml);
            const v = this.convertFromJS(o);
            return (0, protobuf_1.create)(decoder_pb_js_1.DecodeOutputSchema, { value: v });
        }
        catch (e) {
            const msg = e instanceof Error ? e.message : JSON.stringify(e);
            return (0, protobuf_1.create)(decoder_pb_js_1.DecodeOutputSchema, {
                isError: true,
                errorMessage: msg,
            });
        }
    }
    convertFromJS(v) {
        if (v === null) {
            return (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NULL });
        }
        else if (typeof v === "boolean") {
            return (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.BOOL, bool: v });
        }
        else if (typeof v === "number") {
            return (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.NUM, num: v });
        }
        else if (typeof v === "string") {
            return (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.STR, str: v });
        }
        else if (Array.isArray(v)) {
            return (0, protobuf_1.create)(value_pb_js_1.ValueSchema, { type: value_pb_js_1.Type.ARR, arr: v.map(this.convertFromJS) });
        }
        else if (typeof v === "object") {
            return (0, protobuf_1.create)(value_pb_js_1.ValueSchema, {
                type: value_pb_js_1.Type.OBJ,
                obj: Object.fromEntries(Object.entries(v).map(([k, v]) => [k, this.convertFromJS(v)]))
            });
        }
        throw new Error("unexpected value type");
    }
}
exports.Decoder = Decoder;
