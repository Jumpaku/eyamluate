//
//  Generated code. Do not modify.
//  source: eval/validator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use validateInputDescriptor instead')
const ValidateInput$json = {
  '1': 'ValidateInput',
  '2': [
    {'1': 'source', '3': 1, '4': 1, '5': 9, '10': 'source'},
  ],
};

/// Descriptor for `ValidateInput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List validateInputDescriptor = $convert.base64Decode(
    'Cg1WYWxpZGF0ZUlucHV0EhYKBnNvdXJjZRgBIAEoCVIGc291cmNl');

@$core.Deprecated('Use validateOutputDescriptor instead')
const ValidateOutput$json = {
  '1': 'ValidateOutput',
  '2': [
    {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.eval.ValidateOutput.Status', '10': 'status'},
    {'1': 'error_message', '3': 2, '4': 1, '5': 9, '10': 'errorMessage'},
  ],
  '4': [ValidateOutput_Status$json],
};

@$core.Deprecated('Use validateOutputDescriptor instead')
const ValidateOutput_Status$json = {
  '1': 'Status',
  '2': [
    {'1': 'OK', '2': 0},
    {'1': 'YAML_ERROR', '2': 1},
    {'1': 'VALIDATION_ERROR', '2': 2},
  ],
};

/// Descriptor for `ValidateOutput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List validateOutputDescriptor = $convert.base64Decode(
    'Cg5WYWxpZGF0ZU91dHB1dBIzCgZzdGF0dXMYASABKA4yGy5ldmFsLlZhbGlkYXRlT3V0cHV0Ll'
    'N0YXR1c1IGc3RhdHVzEiMKDWVycm9yX21lc3NhZ2UYAiABKAlSDGVycm9yTWVzc2FnZSI2CgZT'
    'dGF0dXMSBgoCT0sQABIOCgpZQU1MX0VSUk9SEAESFAoQVkFMSURBVElPTl9FUlJPUhAC');

const $core.Map<$core.String, $core.dynamic> ValidatorServiceBase$json = {
  '1': 'Validator',
  '2': [
    {'1': 'Validate', '2': '.eval.ValidateInput', '3': '.eval.ValidateOutput', '4': {}},
  ],
};

@$core.Deprecated('Use validatorServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> ValidatorServiceBase$messageJson = {
  '.eval.ValidateInput': ValidateInput$json,
  '.eval.ValidateOutput': ValidateOutput$json,
};

/// Descriptor for `Validator`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List validatorServiceDescriptor = $convert.base64Decode(
    'CglWYWxpZGF0b3ISNwoIVmFsaWRhdGUSEy5ldmFsLlZhbGlkYXRlSW5wdXQaFC5ldmFsLlZhbG'
    'lkYXRlT3V0cHV0IgA=');

