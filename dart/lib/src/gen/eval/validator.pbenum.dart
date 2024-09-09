//
//  Generated code. Do not modify.
//  source: eval/validator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class ValidateOutput_Status extends $pb.ProtobufEnum {
  static const ValidateOutput_Status OK = ValidateOutput_Status._(0, _omitEnumNames ? '' : 'OK');
  static const ValidateOutput_Status YAML_ERROR = ValidateOutput_Status._(1, _omitEnumNames ? '' : 'YAML_ERROR');
  static const ValidateOutput_Status VALIDATION_ERROR = ValidateOutput_Status._(2, _omitEnumNames ? '' : 'VALIDATION_ERROR');

  static const $core.List<ValidateOutput_Status> values = <ValidateOutput_Status> [
    OK,
    YAML_ERROR,
    VALIDATION_ERROR,
  ];

  static final $core.Map<$core.int, ValidateOutput_Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ValidateOutput_Status? valueOf($core.int value) => _byValue[value];

  const ValidateOutput_Status._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
