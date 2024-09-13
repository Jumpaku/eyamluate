import {FunDef, FunDefList, FunDefListSchema} from "./evaluator_pb.js";
import {create} from "@bufbuild/protobuf";


export function empty(): FunDefList {
    return create(FunDefListSchema);
}

export function register(l: FunDefList, def: FunDef): FunDefList {
    return create(FunDefListSchema, {parent: l, def});
}

export function find(l: FunDefList, ident: string): FunDefList | null {
    let cur: FunDefList | null = l;
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

