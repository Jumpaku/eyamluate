"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.empty = empty;
exports.register = register;
exports.find = find;
const evaluator_pb_js_1 = require("./evaluator_pb.js");
const protobuf_1 = require("@bufbuild/protobuf");
function empty() {
    return (0, protobuf_1.create)(evaluator_pb_js_1.FunDefListSchema);
}
function register(l, def) {
    return (0, protobuf_1.create)(evaluator_pb_js_1.FunDefListSchema, { parent: l, def });
}
function find(l, ident) {
    let cur = l;
    while (true) {
        if (cur == null) {
            return null;
        }
        if (cur.def?.def === ident) {
            return cur;
        }
        if (cur.parent == null) {
            return null;
        }
        cur = cur.parent;
    }
}
