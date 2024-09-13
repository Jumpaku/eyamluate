import {Path, Path_PosSchema, PathSchema} from "./evaluator_pb.js";
import {create} from "@bufbuild/protobuf";

export function append(path: Path, pos: number|string): Path {
    return create(PathSchema,{
        pos: [...path.pos, create(Path_PosSchema,{
            index: typeof pos === "number" ? BigInt(pos) : undefined,
            key: typeof pos === "string" ? pos : undefined
        })]
    });
}
