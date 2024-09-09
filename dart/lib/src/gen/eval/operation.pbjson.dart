//
//  Generated code. Do not modify.
//  source: eval/operation.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use opUnaryDescriptor instead')
const OpUnary$json = {
  '1': 'OpUnary',
  '4': [OpUnary_Operator$json],
};

@$core.Deprecated('Use opUnaryDescriptor instead')
const OpUnary_Operator$json = {
  '1': 'Operator',
  '2': [
    {'1': 'UNSPECIFIED', '2': 0},
    {'1': 'LEN', '2': 1},
    {'1': 'NOT', '2': 2},
    {'1': 'FLAT', '2': 3},
    {'1': 'FLOOR', '2': 4},
    {'1': 'CEIL', '2': 5},
    {'1': 'ABORT', '2': 6},
  ],
};

/// Descriptor for `OpUnary`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List opUnaryDescriptor = $convert.base64Decode(
    'CgdPcFVuYXJ5IlcKCE9wZXJhdG9yEg8KC1VOU1BFQ0lGSUVEEAASBwoDTEVOEAESBwoDTk9UEA'
    'ISCAoERkxBVBADEgkKBUZMT09SEAQSCAoEQ0VJTBAFEgkKBUFCT1JUEAY=');

@$core.Deprecated('Use opBinaryDescriptor instead')
const OpBinary$json = {
  '1': 'OpBinary',
  '4': [OpBinary_Operator$json],
};

@$core.Deprecated('Use opBinaryDescriptor instead')
const OpBinary_Operator$json = {
  '1': 'Operator',
  '2': [
    {'1': 'UNSPECIFIED', '2': 0},
    {'1': 'SUB', '2': 1},
    {'1': 'DIV', '2': 2},
    {'1': 'EQ', '2': 4},
    {'1': 'NEQ', '2': 5},
    {'1': 'LT', '2': 6},
    {'1': 'LTE', '2': 7},
    {'1': 'GT', '2': 8},
    {'1': 'GTE', '2': 9},
  ],
};

/// Descriptor for `OpBinary`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List opBinaryDescriptor = $convert.base64Decode(
    'CghPcEJpbmFyeSJgCghPcGVyYXRvchIPCgtVTlNQRUNJRklFRBAAEgcKA1NVQhABEgcKA0RJVh'
    'ACEgYKAkVREAQSBwoDTkVREAUSBgoCTFQQBhIHCgNMVEUQBxIGCgJHVBAIEgcKA0dURRAJ');

@$core.Deprecated('Use opVariadicDescriptor instead')
const OpVariadic$json = {
  '1': 'OpVariadic',
  '4': [OpVariadic_Operator$json],
};

@$core.Deprecated('Use opVariadicDescriptor instead')
const OpVariadic_Operator$json = {
  '1': 'Operator',
  '2': [
    {'1': 'UNSPECIFIED', '2': 0},
    {'1': 'ADD', '2': 1},
    {'1': 'MUL', '2': 2},
    {'1': 'AND', '2': 3},
    {'1': 'OR', '2': 4},
    {'1': 'CAT', '2': 5},
    {'1': 'MIN', '2': 6},
    {'1': 'MAX', '2': 7},
    {'1': 'MERGE', '2': 8},
  ],
};

/// Descriptor for `OpVariadic`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List opVariadicDescriptor = $convert.base64Decode(
    'CgpPcFZhcmlhZGljImQKCE9wZXJhdG9yEg8KC1VOU1BFQ0lGSUVEEAASBwoDQUREEAESBwoDTV'
    'VMEAISBwoDQU5EEAMSBgoCT1IQBBIHCgNDQVQQBRIHCgNNSU4QBhIHCgNNQVgQBxIJCgVNRVJH'
    'RRAI');

