//
//  Generated code. Do not modify.
//  source: eval/evaluator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class EvaluateOutput_Status extends $pb.ProtobufEnum {
  static const EvaluateOutput_Status OK = EvaluateOutput_Status._(0, _omitEnumNames ? '' : 'OK');
  static const EvaluateOutput_Status DECODE_ERROR = EvaluateOutput_Status._(1, _omitEnumNames ? '' : 'DECODE_ERROR');
  static const EvaluateOutput_Status VALIDATE_ERROR = EvaluateOutput_Status._(2, _omitEnumNames ? '' : 'VALIDATE_ERROR');
  static const EvaluateOutput_Status EXPR_ERROR = EvaluateOutput_Status._(3, _omitEnumNames ? '' : 'EXPR_ERROR');

  static const $core.List<EvaluateOutput_Status> values = <EvaluateOutput_Status> [
    OK,
    DECODE_ERROR,
    VALIDATE_ERROR,
    EXPR_ERROR,
  ];

  static final $core.Map<$core.int, EvaluateOutput_Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static EvaluateOutput_Status? valueOf($core.int value) => _byValue[value];

  const EvaluateOutput_Status._($core.int v, $core.String n) : super(v, n);
}

class EvaluateExprOutput_Status extends $pb.ProtobufEnum {
  static const EvaluateExprOutput_Status OK = EvaluateExprOutput_Status._(0, _omitEnumNames ? '' : 'OK');
  static const EvaluateExprOutput_Status UNSUPPORTED_EXPR = EvaluateExprOutput_Status._(1, _omitEnumNames ? '' : 'UNSUPPORTED_EXPR');
  static const EvaluateExprOutput_Status UNEXPECTED_TYPE = EvaluateExprOutput_Status._(2, _omitEnumNames ? '' : 'UNEXPECTED_TYPE');
  static const EvaluateExprOutput_Status ARITHMETIC_ERROR = EvaluateExprOutput_Status._(3, _omitEnumNames ? '' : 'ARITHMETIC_ERROR');
  static const EvaluateExprOutput_Status INDEX_OUT_OF_BOUNDS = EvaluateExprOutput_Status._(4, _omitEnumNames ? '' : 'INDEX_OUT_OF_BOUNDS');
  static const EvaluateExprOutput_Status KEY_NOT_FOUND = EvaluateExprOutput_Status._(5, _omitEnumNames ? '' : 'KEY_NOT_FOUND');
  static const EvaluateExprOutput_Status REFERENCE_NOT_FOUND = EvaluateExprOutput_Status._(6, _omitEnumNames ? '' : 'REFERENCE_NOT_FOUND');
  static const EvaluateExprOutput_Status CASES_NOT_EXHAUSTIVE = EvaluateExprOutput_Status._(7, _omitEnumNames ? '' : 'CASES_NOT_EXHAUSTIVE');
  static const EvaluateExprOutput_Status UNSUPPORTED_OPERATION = EvaluateExprOutput_Status._(8, _omitEnumNames ? '' : 'UNSUPPORTED_OPERATION');
  static const EvaluateExprOutput_Status ABORTED = EvaluateExprOutput_Status._(9, _omitEnumNames ? '' : 'ABORTED');
  static const EvaluateExprOutput_Status UNKNOWN = EvaluateExprOutput_Status._(10, _omitEnumNames ? '' : 'UNKNOWN');

  static const $core.List<EvaluateExprOutput_Status> values = <EvaluateExprOutput_Status> [
    OK,
    UNSUPPORTED_EXPR,
    UNEXPECTED_TYPE,
    ARITHMETIC_ERROR,
    INDEX_OUT_OF_BOUNDS,
    KEY_NOT_FOUND,
    REFERENCE_NOT_FOUND,
    CASES_NOT_EXHAUSTIVE,
    UNSUPPORTED_OPERATION,
    ABORTED,
    UNKNOWN,
  ];

  static final $core.Map<$core.int, EvaluateExprOutput_Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static EvaluateExprOutput_Status? valueOf($core.int value) => _byValue[value];

  const EvaluateExprOutput_Status._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
