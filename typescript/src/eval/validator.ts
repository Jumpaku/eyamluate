import {ValidateInput, ValidateOutput, ValidateOutput_Status, ValidateOutputSchema} from "./validator_pb.js";
import schema from "../schema/eyamluate.schema.js";
import {create} from "@bufbuild/protobuf";
import {Ajv, ValidateFunction} from "ajv"
import YAML from "yaml";

export class Validator {
    constructor() {
        const ajv = new Ajv();// options can be passed, e.g. {allErrors: true}
        this.validator = ajv.compile(schema);
    }

    private readonly validator: ValidateFunction<unknown>;

    validate(input: ValidateInput): ValidateOutput {
        let data: string;
        try {
            data = YAML.parse(input.source);
        } catch (e) {
            return create(ValidateOutputSchema, {
                status: ValidateOutput_Status.YAML_ERROR,
                errorMessage: e instanceof Error ? e.message : JSON.stringify(e),
            });
        }
        const valid = this.validator(data);
        if (!valid) {
            return create(ValidateOutputSchema, {
                status: ValidateOutput_Status.VALIDATION_ERROR,
                errorMessage: (this.validator.errors ?? []).map((e) => (JSON.stringify({
                    instancePath: e.instancePath,
                    message: e.message,
                    schemaPath: e.schemaPath,
                    propertyName: e.propertyName,
                    keyword: e.keyword,
                }))).join("\n"),
            });
        }
        return create(ValidateOutputSchema, {
            status: ValidateOutput_Status.OK,
        });

    }

}
