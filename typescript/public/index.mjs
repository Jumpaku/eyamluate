import {Decoder} from '../dist/esm/yaml/decoder.js';
import {create} from "@bufbuild/protobuf";
import {DecodeInputSchema} from "../dist/esm/yaml/decoder_pb.js";
import {Encoder} from '../dist/esm/yaml/encoder.js';
import {EncodeInputSchema} from "../dist/esm/yaml/encoder_pb.js";
import {DecodeOutputSchema} from "../dist/esm/yaml/decoder_pb.js";

const o = new Decoder().decode(create(DecodeInputSchema, {
    yaml: 'a: b'
}));

console.log(JSON.stringify(o));
const s = new Encoder().encode(create(EncodeInputSchema, {
    value: o.value ?? create(DecodeOutputSchema)
}));
console.log(JSON.stringify(s));