//
//  Generated code. Do not modify.
//  source: yaml/value.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use typeDescriptor instead')
const Type$json = {
  '1': 'Type',
  '2': [
    {'1': 'TYPE_UNSPECIFIED', '2': 0},
    {'1': 'TYPE_NULL', '2': 1},
    {'1': 'TYPE_BOOL', '2': 2},
    {'1': 'TYPE_NUM', '2': 3},
    {'1': 'TYPE_STR', '2': 4},
    {'1': 'TYPE_ARR', '2': 5},
    {'1': 'TYPE_OBJ', '2': 6},
  ],
};

/// Descriptor for `Type`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List typeDescriptor = $convert.base64Decode(
    'CgRUeXBlEhQKEFRZUEVfVU5TUEVDSUZJRUQQABINCglUWVBFX05VTEwQARINCglUWVBFX0JPT0'
    'wQAhIMCghUWVBFX05VTRADEgwKCFRZUEVfU1RSEAQSDAoIVFlQRV9BUlIQBRIMCghUWVBFX09C'
    'ShAG');

@$core.Deprecated('Use valueDescriptor instead')
const Value$json = {
  '1': 'Value',
  '2': [
    {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.yaml.Type', '10': 'type'},
    {'1': 'bool', '3': 2, '4': 1, '5': 8, '10': 'bool'},
    {'1': 'num', '3': 3, '4': 1, '5': 1, '10': 'num'},
    {'1': 'str', '3': 4, '4': 1, '5': 9, '10': 'str'},
    {'1': 'arr', '3': 5, '4': 3, '5': 11, '6': '.yaml.Value', '10': 'arr'},
    {'1': 'obj', '3': 6, '4': 3, '5': 11, '6': '.yaml.Value.ObjEntry', '10': 'obj'},
  ],
  '3': [Value_ObjEntry$json],
};

@$core.Deprecated('Use valueDescriptor instead')
const Value_ObjEntry$json = {
  '1': 'ObjEntry',
  '2': [
    {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.yaml.Value', '10': 'value'},
  ],
  '7': {'7': true},
};

/// Descriptor for `Value`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List valueDescriptor = $convert.base64Decode(
    'CgVWYWx1ZRIeCgR0eXBlGAEgASgOMgoueWFtbC5UeXBlUgR0eXBlEhIKBGJvb2wYAiABKAhSBG'
    'Jvb2wSEAoDbnVtGAMgASgBUgNudW0SEAoDc3RyGAQgASgJUgNzdHISHQoDYXJyGAUgAygLMgsu'
    'eWFtbC5WYWx1ZVIDYXJyEiYKA29iahgGIAMoCzIULnlhbWwuVmFsdWUuT2JqRW50cnlSA29iah'
    'pDCghPYmpFbnRyeRIQCgNrZXkYASABKAlSA2tleRIhCgV2YWx1ZRgCIAEoCzILLnlhbWwuVmFs'
    'dWVSBXZhbHVlOgI4AQ==');

