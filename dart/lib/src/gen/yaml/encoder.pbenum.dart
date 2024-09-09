//
//  Generated code. Do not modify.
//  source: yaml/encoder.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class EncodeFormat extends $pb.ProtobufEnum {
  static const EncodeFormat ENCODE_FORMAT_YAML = EncodeFormat._(0, _omitEnumNames ? '' : 'ENCODE_FORMAT_YAML');
  static const EncodeFormat ENCODE_FORMAT_JSON = EncodeFormat._(1, _omitEnumNames ? '' : 'ENCODE_FORMAT_JSON');

  static const $core.List<EncodeFormat> values = <EncodeFormat> [
    ENCODE_FORMAT_YAML,
    ENCODE_FORMAT_JSON,
  ];

  static final $core.Map<$core.int, EncodeFormat> _byValue = $pb.ProtobufEnum.initByValue(values);
  static EncodeFormat? valueOf($core.int value) => _byValue[value];

  const EncodeFormat._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
