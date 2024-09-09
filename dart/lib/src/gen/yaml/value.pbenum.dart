//
//  Generated code. Do not modify.
//  source: yaml/value.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Type extends $pb.ProtobufEnum {
  static const Type TYPE_UNSPECIFIED = Type._(0, _omitEnumNames ? '' : 'TYPE_UNSPECIFIED');
  static const Type TYPE_NULL = Type._(1, _omitEnumNames ? '' : 'TYPE_NULL');
  static const Type TYPE_BOOL = Type._(2, _omitEnumNames ? '' : 'TYPE_BOOL');
  static const Type TYPE_NUM = Type._(3, _omitEnumNames ? '' : 'TYPE_NUM');
  static const Type TYPE_STR = Type._(4, _omitEnumNames ? '' : 'TYPE_STR');
  static const Type TYPE_ARR = Type._(5, _omitEnumNames ? '' : 'TYPE_ARR');
  static const Type TYPE_OBJ = Type._(6, _omitEnumNames ? '' : 'TYPE_OBJ');

  static const $core.List<Type> values = <Type> [
    TYPE_UNSPECIFIED,
    TYPE_NULL,
    TYPE_BOOL,
    TYPE_NUM,
    TYPE_STR,
    TYPE_ARR,
    TYPE_OBJ,
  ];

  static final $core.Map<$core.int, Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Type? valueOf($core.int value) => _byValue[value];

  const Type._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
