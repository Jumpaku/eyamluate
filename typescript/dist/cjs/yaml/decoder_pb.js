"use strict";
// @generated by protoc-gen-es v2.0.0 with parameter "target=ts,import_extension=js"
// @generated from file yaml/decoder.proto (package yaml, syntax proto3)
/* eslint-disable */
Object.defineProperty(exports, "__esModule", { value: true });
exports.Decoder = exports.DecodeOutputSchema = exports.DecodeInputSchema = exports.file_yaml_decoder = void 0;
const codegenv1_1 = require("@bufbuild/protobuf/codegenv1");
const value_pb_js_1 = require("./value_pb.js");
/**
 * Describes the file yaml/decoder.proto.
 */
exports.file_yaml_decoder = (0, codegenv1_1.fileDesc)("ChJ5YW1sL2RlY29kZXIucHJvdG8SBHlhbWwiGwoLRGVjb2RlSW5wdXQSDAoEeWFtbBgBIAEoCSJTCgxEZWNvZGVPdXRwdXQSEAoIaXNfZXJyb3IYASABKAgSFQoNZXJyb3JfbWVzc2FnZRgCIAEoCRIaCgV2YWx1ZRgDIAEoCzILLnlhbWwuVmFsdWUyPAoHRGVjb2RlchIxCgZEZWNvZGUSES55YW1sLkRlY29kZUlucHV0GhIueWFtbC5EZWNvZGVPdXRwdXQiAEJ2Cghjb20ueWFtbEIMRGVjb2RlclByb3RvUAFaJ2dpdGh1Yi5jb20vSnVtcGFrdS9leWFtbGF0ZS9nb2xhbmcveWFtbKICA1lYWKoCBFlhbWzCAgJQQsoCBFlhbWziAhBZYW1sXEdQQk1ldGFkYXRh6gIEWWFtbGIGcHJvdG8z", [value_pb_js_1.file_yaml_value]);
/**
 * Describes the message yaml.DecodeInput.
 * Use `create(DecodeInputSchema)` to create a new message.
 */
exports.DecodeInputSchema = (0, codegenv1_1.messageDesc)(exports.file_yaml_decoder, 0);
/**
 * Describes the message yaml.DecodeOutput.
 * Use `create(DecodeOutputSchema)` to create a new message.
 */
exports.DecodeOutputSchema = (0, codegenv1_1.messageDesc)(exports.file_yaml_decoder, 1);
/**
 * @generated from service yaml.Decoder
 */
exports.Decoder = (0, codegenv1_1.serviceDesc)(exports.file_yaml_decoder, 0);
