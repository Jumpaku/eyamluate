import { FunDefListSchema } from "./evaluator_pb.js";
import { create } from "@bufbuild/protobuf";
export function empty() {
    return create(FunDefListSchema);
}
export function register(l, def) {
    return create(FunDefListSchema, { parent: l, def });
}
export function find(l, ident) {
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
