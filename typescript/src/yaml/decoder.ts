import {create} from "@bufbuild/protobuf";
import {DecodeInput, DecodeOutput, DecodeOutputSchema} from "./decoder_pb.js";
import YAML from 'yaml'
import {Type, Value, ValueSchema} from "./value_pb.js";

export class Decoder {
    decode(input: DecodeInput): DecodeOutput {
        try {
            const o = YAML.parse(input.yaml);
            const v = this.convertFromJS(o);
            return create(DecodeOutputSchema, {value: v});
        } catch (e) {
            const msg = e instanceof Error ? e.message : JSON.stringify(e);
            return create(DecodeOutputSchema, {
                isError: true,
                errorMessage: msg,
            });
        }
    }

    private convertFromJS(v: unknown): Value {
        if (v === null) {
            return create(ValueSchema, {type: Type.NULL});
        } else if (typeof v === "boolean") {
            return create(ValueSchema, {type: Type.BOOL, bool: v});
        } else if (typeof v === "number") {
            return create(ValueSchema, {type: Type.NUM, num: v});
        } else if (typeof v === "string") {
            return create(ValueSchema, {type: Type.STR, str: v});
        } else if (Array.isArray(v)) {
            return create(ValueSchema, {type: Type.ARR, arr: v.map(this.convertFromJS)});
        } else if (typeof v === "object") {
            return create(ValueSchema, {
                type: Type.OBJ,
                obj: Object.fromEntries(Object.entries(v).map(([k, v]) => [k, this.convertFromJS(v)]))
            });
        }
        throw new Error("unexpected value type");
    }
}
