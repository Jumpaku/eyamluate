//
//  Generated code. Do not modify.
//  source: eval/operation.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class OpUnary_Operator extends $pb.ProtobufEnum {
  static const OpUnary_Operator UNSPECIFIED = OpUnary_Operator._(0, _omitEnumNames ? '' : 'UNSPECIFIED');
  static const OpUnary_Operator LEN = OpUnary_Operator._(1, _omitEnumNames ? '' : 'LEN');
  static const OpUnary_Operator NOT = OpUnary_Operator._(2, _omitEnumNames ? '' : 'NOT');
  static const OpUnary_Operator FLAT = OpUnary_Operator._(3, _omitEnumNames ? '' : 'FLAT');
  static const OpUnary_Operator FLOOR = OpUnary_Operator._(4, _omitEnumNames ? '' : 'FLOOR');
  static const OpUnary_Operator CEIL = OpUnary_Operator._(5, _omitEnumNames ? '' : 'CEIL');
  static const OpUnary_Operator ABORT = OpUnary_Operator._(6, _omitEnumNames ? '' : 'ABORT');

  static const $core.List<OpUnary_Operator> values = <OpUnary_Operator> [
    UNSPECIFIED,
    LEN,
    NOT,
    FLAT,
    FLOOR,
    CEIL,
    ABORT,
  ];

  static final $core.Map<$core.int, OpUnary_Operator> _byValue = $pb.ProtobufEnum.initByValue(values);
  static OpUnary_Operator? valueOf($core.int value) => _byValue[value];

  const OpUnary_Operator._($core.int v, $core.String n) : super(v, n);
}

class OpBinary_Operator extends $pb.ProtobufEnum {
  static const OpBinary_Operator UNSPECIFIED = OpBinary_Operator._(0, _omitEnumNames ? '' : 'UNSPECIFIED');
  static const OpBinary_Operator SUB = OpBinary_Operator._(1, _omitEnumNames ? '' : 'SUB');
  static const OpBinary_Operator DIV = OpBinary_Operator._(2, _omitEnumNames ? '' : 'DIV');
  static const OpBinary_Operator EQ = OpBinary_Operator._(4, _omitEnumNames ? '' : 'EQ');
  static const OpBinary_Operator NEQ = OpBinary_Operator._(5, _omitEnumNames ? '' : 'NEQ');
  static const OpBinary_Operator LT = OpBinary_Operator._(6, _omitEnumNames ? '' : 'LT');
  static const OpBinary_Operator LTE = OpBinary_Operator._(7, _omitEnumNames ? '' : 'LTE');
  static const OpBinary_Operator GT = OpBinary_Operator._(8, _omitEnumNames ? '' : 'GT');
  static const OpBinary_Operator GTE = OpBinary_Operator._(9, _omitEnumNames ? '' : 'GTE');

  static const $core.List<OpBinary_Operator> values = <OpBinary_Operator> [
    UNSPECIFIED,
    SUB,
    DIV,
    EQ,
    NEQ,
    LT,
    LTE,
    GT,
    GTE,
  ];

  static final $core.Map<$core.int, OpBinary_Operator> _byValue = $pb.ProtobufEnum.initByValue(values);
  static OpBinary_Operator? valueOf($core.int value) => _byValue[value];

  const OpBinary_Operator._($core.int v, $core.String n) : super(v, n);
}

class OpVariadic_Operator extends $pb.ProtobufEnum {
  static const OpVariadic_Operator UNSPECIFIED = OpVariadic_Operator._(0, _omitEnumNames ? '' : 'UNSPECIFIED');
  static const OpVariadic_Operator ADD = OpVariadic_Operator._(1, _omitEnumNames ? '' : 'ADD');
  static const OpVariadic_Operator MUL = OpVariadic_Operator._(2, _omitEnumNames ? '' : 'MUL');
  static const OpVariadic_Operator AND = OpVariadic_Operator._(3, _omitEnumNames ? '' : 'AND');
  static const OpVariadic_Operator OR = OpVariadic_Operator._(4, _omitEnumNames ? '' : 'OR');
  static const OpVariadic_Operator CAT = OpVariadic_Operator._(5, _omitEnumNames ? '' : 'CAT');
  static const OpVariadic_Operator MIN = OpVariadic_Operator._(6, _omitEnumNames ? '' : 'MIN');
  static const OpVariadic_Operator MAX = OpVariadic_Operator._(7, _omitEnumNames ? '' : 'MAX');
  static const OpVariadic_Operator MERGE = OpVariadic_Operator._(8, _omitEnumNames ? '' : 'MERGE');

  static const $core.List<OpVariadic_Operator> values = <OpVariadic_Operator> [
    UNSPECIFIED,
    ADD,
    MUL,
    AND,
    OR,
    CAT,
    MIN,
    MAX,
    MERGE,
  ];

  static final $core.Map<$core.int, OpVariadic_Operator> _byValue = $pb.ProtobufEnum.initByValue(values);
  static OpVariadic_Operator? valueOf($core.int value) => _byValue[value];

  const OpVariadic_Operator._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
