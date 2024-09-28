"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.append = append;
const evaluator_pb_js_1 = require("./evaluator_pb.js");
const protobuf_1 = require("@bufbuild/protobuf");
function append(path, pos) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.PathSchema, {
        pos: [...path.pos, (0, protobuf_1.create)(evaluator_pb_js_1.Path_PosSchema, {
                index: typeof pos === "number" ? BigInt(pos) : undefined,
                key: typeof pos === "string" ? pos : undefined
            })]
    });
}
