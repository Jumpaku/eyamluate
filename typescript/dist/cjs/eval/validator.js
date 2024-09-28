"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Validator = void 0;
const validator_pb_js_1 = require("./validator_pb.js");
const eyamluate_schema_js_1 = __importDefault(require("../schema/eyamluate.schema.js"));
const protobuf_1 = require("@bufbuild/protobuf");
const ajv_1 = require("ajv");
const yaml_1 = __importDefault(require("yaml"));
class Validator {
    constructor() {
        const ajv = new ajv_1.Ajv(); // options can be passed, e.g. {allErrors: true}
        this.validator = ajv.compile(eyamluate_schema_js_1.default);
    }
    validator;
    validate(input) {
        let data;
        try {
            data = yaml_1.default.parse(input.source);
        }
        catch (e) {
            return (0, protobuf_1.create)(validator_pb_js_1.ValidateOutputSchema, {
                status: validator_pb_js_1.ValidateOutput_Status.YAML_ERROR,
                errorMessage: e instanceof Error ? e.message : JSON.stringify(e),
            });
        }
        const valid = this.validator(data);
        if (!valid) {
            return (0, protobuf_1.create)(validator_pb_js_1.ValidateOutputSchema, {
                status: validator_pb_js_1.ValidateOutput_Status.VALIDATION_ERROR,
                errorMessage: (this.validator.errors ?? []).map((e) => (JSON.stringify({
                    instancePath: e.instancePath,
                    message: e.message,
                    schemaPath: e.schemaPath,
                    propertyName: e.propertyName,
                    keyword: e.keyword,
                }))).join("\n"),
            });
        }
        return (0, protobuf_1.create)(validator_pb_js_1.ValidateOutputSchema, {
            status: validator_pb_js_1.ValidateOutput_Status.OK,
        });
    }
}
exports.Validator = Validator;
