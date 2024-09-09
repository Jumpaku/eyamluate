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

import 'value.pbenum.dart';

export 'value.pbenum.dart';

class Value extends $pb.GeneratedMessage {
  factory Value({
    Type? type,
    $core.bool? bool_2,
    $core.double? num,
    $core.String? str,
    $core.Iterable<Value>? arr,
    $core.Map<$core.String, Value>? obj,
  }) {
    final $result = create();
    if (type != null) {
      $result.type = type;
    }
    if (bool_2 != null) {
      $result.bool_2 = bool_2;
    }
    if (num != null) {
      $result.num = num;
    }
    if (str != null) {
      $result.str = str;
    }
    if (arr != null) {
      $result.arr.addAll(arr);
    }
    if (obj != null) {
      $result.obj.addAll(obj);
    }
    return $result;
  }
  Value._() : super();
  factory Value.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Value.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Value', package: const $pb.PackageName(_omitMessageNames ? '' : 'yaml'), createEmptyInstance: create)
    ..e<Type>(1, _omitFieldNames ? '' : 'type', $pb.PbFieldType.OE, defaultOrMaker: Type.TYPE_UNSPECIFIED, valueOf: Type.valueOf, enumValues: Type.values)
    ..aOB(2, _omitFieldNames ? '' : 'bool')
    ..a<$core.double>(3, _omitFieldNames ? '' : 'num', $pb.PbFieldType.OD)
    ..aOS(4, _omitFieldNames ? '' : 'str')
    ..pc<Value>(5, _omitFieldNames ? '' : 'arr', $pb.PbFieldType.PM, subBuilder: Value.create)
    ..m<$core.String, Value>(6, _omitFieldNames ? '' : 'obj', entryClassName: 'Value.ObjEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: Value.create, valueDefaultOrMaker: Value.getDefault, packageName: const $pb.PackageName('yaml'))
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Value clone() => Value()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Value copyWith(void Function(Value) updates) => super.copyWith((message) => updates(message as Value)) as Value;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Value create() => Value._();
  Value createEmptyInstance() => create();
  static $pb.PbList<Value> createRepeated() => $pb.PbList<Value>();
  @$core.pragma('dart2js:noInline')
  static Value getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Value>(create);
  static Value? _defaultInstance;

  @$pb.TagNumber(1)
  Type get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(Type v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get bool_2 => $_getBF(1);
  @$pb.TagNumber(2)
  set bool_2($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasBool_2() => $_has(1);
  @$pb.TagNumber(2)
  void clearBool_2() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get num => $_getN(2);
  @$pb.TagNumber(3)
  set num($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNum() => $_has(2);
  @$pb.TagNumber(3)
  void clearNum() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get str => $_getSZ(3);
  @$pb.TagNumber(4)
  set str($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasStr() => $_has(3);
  @$pb.TagNumber(4)
  void clearStr() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<Value> get arr => $_getList(4);

  @$pb.TagNumber(6)
  $core.Map<$core.String, Value> get obj => $_getMap(5);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
