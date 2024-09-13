import { EncodeOutputSchema } from "./encoder_pb.js";
import { Type, ValueSchema } from "./value_pb.js";
import { create } from "@bufbuild/protobuf";
import { stringify } from "yaml";
export class Encoder {
    encode(input) {
        try {
            const o = this.convertToJS(input.value ?? create(ValueSchema));
            return create(EncodeOutputSchema, { result: stringify(o) });
        }
        catch (e) {
            return create(EncodeOutputSchema, {
                isError: true,
                errorMessage: e instanceof Error ? e.message : JSON.stringify(e),
            });
        }
    }
    convertToJS(v) {
        switch (v.type) {
            case Type.NULL:
                return null;
            case Type.BOOL:
                return v.bool;
            case Type.NUM:
                return v.num;
            case Type.STR:
                return v.str;
            case Type.ARR:
                return v.arr.map(this.convertToJS);
            case Type.OBJ:
                return Object.fromEntries(Object.entries(v.obj).map(([k, v]) => [k, this.convertToJS(v)]));
            default:
                throw new Error("unexpected value type: " + Type[v.type]);
        }
    }
}
