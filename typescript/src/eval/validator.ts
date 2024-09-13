/*
import {ValidateInput, ValidateOutput} from "./validator_pb.js";
import schema from "./../schema/eyamluate.schem.json";
export class Validator {
    constructor() {


    }
    validate(input: ValidateInput): ValidateOutput {
        try {
            return create(ValidateOutputSchema, {result: "ok"});
        } catch (e) {
            return create(ValidateOutputSchema, {
                isError: true,
                errorMessage: e instanceof Error ? e.message : JSON.stringify(e),
            });
        }

    }

}

 */