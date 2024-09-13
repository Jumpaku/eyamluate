"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Encoder = void 0;
const encoder_pb_js_1 = require("./encoder_pb.js");
const value_pb_js_1 = require("./value_pb.js");
const protobuf_1 = require("@bufbuild/protobuf");
const yaml_1 = require("yaml");
class Encoder {
    encode(input) {
        try {
            const o = this.convertToJS(input.value ?? (0, protobuf_1.create)(value_pb_js_1.ValueSchema));
            return (0, protobuf_1.create)(encoder_pb_js_1.EncodeOutputSchema, { result: (0, yaml_1.stringify)(o) });
        }
        catch (e) {
            return (0, protobuf_1.create)(encoder_pb_js_1.EncodeOutputSchema, {
                isError: true,
                errorMessage: e instanceof Error ? e.message : JSON.stringify(e),
            });
        }
    }
    convertToJS(v) {
        switch (v.type) {
            case value_pb_js_1.Type.NULL:
                return null;
            case value_pb_js_1.Type.BOOL:
                return v.bool;
            case value_pb_js_1.Type.NUM:
                return v.num;
            case value_pb_js_1.Type.STR:
                return v.str;
            case value_pb_js_1.Type.ARR:
                return v.arr.map(this.convertToJS);
            case value_pb_js_1.Type.OBJ:
                return Object.fromEntries(Object.entries(v.obj).map(([k, v]) => [k, this.convertToJS(v)]));
            default:
                throw new Error("unexpected value type: " + value_pb_js_1.Type[v.type]);
        }
    }
}
exports.Encoder = Encoder;
