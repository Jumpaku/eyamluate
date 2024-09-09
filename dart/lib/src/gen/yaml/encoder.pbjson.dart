//
//  Generated code. Do not modify.
//  source: yaml/encoder.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import 'value.pbjson.dart' as $0;

@$core.Deprecated('Use encodeFormatDescriptor instead')
const EncodeFormat$json = {
  '1': 'EncodeFormat',
  '2': [
    {'1': 'ENCODE_FORMAT_YAML', '2': 0},
    {'1': 'ENCODE_FORMAT_JSON', '2': 1},
  ],
};

/// Descriptor for `EncodeFormat`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List encodeFormatDescriptor = $convert.base64Decode(
    'CgxFbmNvZGVGb3JtYXQSFgoSRU5DT0RFX0ZPUk1BVF9ZQU1MEAASFgoSRU5DT0RFX0ZPUk1BVF'
    '9KU09OEAE=');

@$core.Deprecated('Use encodeInputDescriptor instead')
const EncodeInput$json = {
  '1': 'EncodeInput',
  '2': [
    {'1': 'format', '3': 1, '4': 1, '5': 14, '6': '.yaml.EncodeFormat', '10': 'format'},
    {'1': 'pretty', '3': 2, '4': 1, '5': 8, '10': 'pretty'},
    {'1': 'value', '3': 3, '4': 1, '5': 11, '6': '.yaml.Value', '10': 'value'},
  ],
};

/// Descriptor for `EncodeInput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List encodeInputDescriptor = $convert.base64Decode(
    'CgtFbmNvZGVJbnB1dBIqCgZmb3JtYXQYASABKA4yEi55YW1sLkVuY29kZUZvcm1hdFIGZm9ybW'
    'F0EhYKBnByZXR0eRgCIAEoCFIGcHJldHR5EiEKBXZhbHVlGAMgASgLMgsueWFtbC5WYWx1ZVIF'
    'dmFsdWU=');

@$core.Deprecated('Use encodeOutputDescriptor instead')
const EncodeOutput$json = {
  '1': 'EncodeOutput',
  '2': [
    {'1': 'is_error', '3': 1, '4': 1, '5': 8, '10': 'isError'},
    {'1': 'error_message', '3': 2, '4': 1, '5': 9, '10': 'errorMessage'},
    {'1': 'result', '3': 3, '4': 1, '5': 9, '10': 'result'},
  ],
};

/// Descriptor for `EncodeOutput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List encodeOutputDescriptor = $convert.base64Decode(
    'CgxFbmNvZGVPdXRwdXQSGQoIaXNfZXJyb3IYASABKAhSB2lzRXJyb3ISIwoNZXJyb3JfbWVzc2'
    'FnZRgCIAEoCVIMZXJyb3JNZXNzYWdlEhYKBnJlc3VsdBgDIAEoCVIGcmVzdWx0');

const $core.Map<$core.String, $core.dynamic> EncoderServiceBase$json = {
  '1': 'Encoder',
  '2': [
    {'1': 'Encode', '2': '.yaml.EncodeInput', '3': '.yaml.EncodeOutput', '4': {}},
  ],
};

@$core.Deprecated('Use encoderServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> EncoderServiceBase$messageJson = {
  '.yaml.EncodeInput': EncodeInput$json,
  '.yaml.Value': $0.Value$json,
  '.yaml.Value.ObjEntry': $0.Value_ObjEntry$json,
  '.yaml.EncodeOutput': EncodeOutput$json,
};

/// Descriptor for `Encoder`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List encoderServiceDescriptor = $convert.base64Decode(
    'CgdFbmNvZGVyEjEKBkVuY29kZRIRLnlhbWwuRW5jb2RlSW5wdXQaEi55YW1sLkVuY29kZU91dH'
    'B1dCIA');

