//
//  Generated code. Do not modify.
//  source: eval/evaluator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../yaml/value.pb.dart' as $0;
import 'evaluator.pbenum.dart';

export 'evaluator.pbenum.dart';

class FunDef extends $pb.GeneratedMessage {
  factory FunDef({
    $core.String? def,
    $0.Value? value,
    $core.Iterable<$core.String>? with_3,
    Path? path,
  }) {
    final $result = create();
    if (def != null) {
      $result.def = def;
    }
    if (value != null) {
      $result.value = value;
    }
    if (with_3 != null) {
      $result.with_3.addAll(with_3);
    }
    if (path != null) {
      $result.path = path;
    }
    return $result;
  }
  FunDef._() : super();
  factory FunDef.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FunDef.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'FunDef', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'def')
    ..aOM<$0.Value>(2, _omitFieldNames ? '' : 'value', subBuilder: $0.Value.create)
    ..pPS(3, _omitFieldNames ? '' : 'with')
    ..aOM<Path>(10, _omitFieldNames ? '' : 'path', subBuilder: Path.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FunDef clone() => FunDef()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FunDef copyWith(void Function(FunDef) updates) => super.copyWith((message) => updates(message as FunDef)) as FunDef;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static FunDef create() => FunDef._();
  FunDef createEmptyInstance() => create();
  static $pb.PbList<FunDef> createRepeated() => $pb.PbList<FunDef>();
  @$core.pragma('dart2js:noInline')
  static FunDef getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FunDef>(create);
  static FunDef? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get def => $_getSZ(0);
  @$pb.TagNumber(1)
  set def($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDef() => $_has(0);
  @$pb.TagNumber(1)
  void clearDef() => clearField(1);

  @$pb.TagNumber(2)
  $0.Value get value => $_getN(1);
  @$pb.TagNumber(2)
  set value($0.Value v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
  @$pb.TagNumber(2)
  $0.Value ensureValue() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.List<$core.String> get with_3 => $_getList(2);

  @$pb.TagNumber(10)
  Path get path => $_getN(3);
  @$pb.TagNumber(10)
  set path(Path v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasPath() => $_has(3);
  @$pb.TagNumber(10)
  void clearPath() => clearField(10);
  @$pb.TagNumber(10)
  Path ensurePath() => $_ensure(3);
}

class FunDefList extends $pb.GeneratedMessage {
  factory FunDefList({
    FunDefList? parent,
    FunDef? def,
  }) {
    final $result = create();
    if (parent != null) {
      $result.parent = parent;
    }
    if (def != null) {
      $result.def = def;
    }
    return $result;
  }
  FunDefList._() : super();
  factory FunDefList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FunDefList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'FunDefList', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..aOM<FunDefList>(1, _omitFieldNames ? '' : 'parent', subBuilder: FunDefList.create)
    ..aOM<FunDef>(2, _omitFieldNames ? '' : 'def', subBuilder: FunDef.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FunDefList clone() => FunDefList()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FunDefList copyWith(void Function(FunDefList) updates) => super.copyWith((message) => updates(message as FunDefList)) as FunDefList;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static FunDefList create() => FunDefList._();
  FunDefList createEmptyInstance() => create();
  static $pb.PbList<FunDefList> createRepeated() => $pb.PbList<FunDefList>();
  @$core.pragma('dart2js:noInline')
  static FunDefList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FunDefList>(create);
  static FunDefList? _defaultInstance;

  @$pb.TagNumber(1)
  FunDefList get parent => $_getN(0);
  @$pb.TagNumber(1)
  set parent(FunDefList v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasParent() => $_has(0);
  @$pb.TagNumber(1)
  void clearParent() => clearField(1);
  @$pb.TagNumber(1)
  FunDefList ensureParent() => $_ensure(0);

  @$pb.TagNumber(2)
  FunDef get def => $_getN(1);
  @$pb.TagNumber(2)
  set def(FunDef v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasDef() => $_has(1);
  @$pb.TagNumber(2)
  void clearDef() => clearField(2);
  @$pb.TagNumber(2)
  FunDef ensureDef() => $_ensure(1);
}

class Path_Pos extends $pb.GeneratedMessage {
  factory Path_Pos({
    $fixnum.Int64? index,
    $core.String? key,
  }) {
    final $result = create();
    if (index != null) {
      $result.index = index;
    }
    if (key != null) {
      $result.key = key;
    }
    return $result;
  }
  Path_Pos._() : super();
  factory Path_Pos.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Path_Pos.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Path.Pos', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..aInt64(1, _omitFieldNames ? '' : 'index')
    ..aOS(2, _omitFieldNames ? '' : 'key')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Path_Pos clone() => Path_Pos()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Path_Pos copyWith(void Function(Path_Pos) updates) => super.copyWith((message) => updates(message as Path_Pos)) as Path_Pos;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Path_Pos create() => Path_Pos._();
  Path_Pos createEmptyInstance() => create();
  static $pb.PbList<Path_Pos> createRepeated() => $pb.PbList<Path_Pos>();
  @$core.pragma('dart2js:noInline')
  static Path_Pos getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Path_Pos>(create);
  static Path_Pos? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get index => $_getI64(0);
  @$pb.TagNumber(1)
  set index($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIndex() => $_has(0);
  @$pb.TagNumber(1)
  void clearIndex() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get key => $_getSZ(1);
  @$pb.TagNumber(2)
  set key($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasKey() => $_has(1);
  @$pb.TagNumber(2)
  void clearKey() => clearField(2);
}

class Path extends $pb.GeneratedMessage {
  factory Path({
    $core.Iterable<Path_Pos>? pos,
  }) {
    final $result = create();
    if (pos != null) {
      $result.pos.addAll(pos);
    }
    return $result;
  }
  Path._() : super();
  factory Path.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Path.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Path', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..pc<Path_Pos>(1, _omitFieldNames ? '' : 'pos', $pb.PbFieldType.PM, subBuilder: Path_Pos.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Path clone() => Path()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Path copyWith(void Function(Path) updates) => super.copyWith((message) => updates(message as Path)) as Path;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Path create() => Path._();
  Path createEmptyInstance() => create();
  static $pb.PbList<Path> createRepeated() => $pb.PbList<Path>();
  @$core.pragma('dart2js:noInline')
  static Path getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Path>(create);
  static Path? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Path_Pos> get pos => $_getList(0);
}

class EvaluateInput extends $pb.GeneratedMessage {
  factory EvaluateInput({
    $core.String? source,
  }) {
    final $result = create();
    if (source != null) {
      $result.source = source;
    }
    return $result;
  }
  EvaluateInput._() : super();
  factory EvaluateInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EvaluateInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'EvaluateInput', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'source')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EvaluateInput clone() => EvaluateInput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EvaluateInput copyWith(void Function(EvaluateInput) updates) => super.copyWith((message) => updates(message as EvaluateInput)) as EvaluateInput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EvaluateInput create() => EvaluateInput._();
  EvaluateInput createEmptyInstance() => create();
  static $pb.PbList<EvaluateInput> createRepeated() => $pb.PbList<EvaluateInput>();
  @$core.pragma('dart2js:noInline')
  static EvaluateInput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EvaluateInput>(create);
  static EvaluateInput? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get source => $_getSZ(0);
  @$pb.TagNumber(1)
  set source($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSource() => $_has(0);
  @$pb.TagNumber(1)
  void clearSource() => clearField(1);
}

class EvaluateOutput extends $pb.GeneratedMessage {
  factory EvaluateOutput({
    EvaluateOutput_Status? status,
    $core.String? errorMessage,
    Path? exprErrorPath,
    EvaluateExprOutput_Status? exprStatus,
    $0.Value? value,
  }) {
    final $result = create();
    if (status != null) {
      $result.status = status;
    }
    if (errorMessage != null) {
      $result.errorMessage = errorMessage;
    }
    if (exprErrorPath != null) {
      $result.exprErrorPath = exprErrorPath;
    }
    if (exprStatus != null) {
      $result.exprStatus = exprStatus;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  EvaluateOutput._() : super();
  factory EvaluateOutput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EvaluateOutput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'EvaluateOutput', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..e<EvaluateOutput_Status>(1, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: EvaluateOutput_Status.OK, valueOf: EvaluateOutput_Status.valueOf, enumValues: EvaluateOutput_Status.values)
    ..aOS(2, _omitFieldNames ? '' : 'errorMessage')
    ..aOM<Path>(3, _omitFieldNames ? '' : 'exprErrorPath', subBuilder: Path.create)
    ..e<EvaluateExprOutput_Status>(4, _omitFieldNames ? '' : 'exprStatus', $pb.PbFieldType.OE, defaultOrMaker: EvaluateExprOutput_Status.OK, valueOf: EvaluateExprOutput_Status.valueOf, enumValues: EvaluateExprOutput_Status.values)
    ..aOM<$0.Value>(5, _omitFieldNames ? '' : 'value', subBuilder: $0.Value.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EvaluateOutput clone() => EvaluateOutput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EvaluateOutput copyWith(void Function(EvaluateOutput) updates) => super.copyWith((message) => updates(message as EvaluateOutput)) as EvaluateOutput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EvaluateOutput create() => EvaluateOutput._();
  EvaluateOutput createEmptyInstance() => create();
  static $pb.PbList<EvaluateOutput> createRepeated() => $pb.PbList<EvaluateOutput>();
  @$core.pragma('dart2js:noInline')
  static EvaluateOutput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EvaluateOutput>(create);
  static EvaluateOutput? _defaultInstance;

  @$pb.TagNumber(1)
  EvaluateOutput_Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(EvaluateOutput_Status v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get errorMessage => $_getSZ(1);
  @$pb.TagNumber(2)
  set errorMessage($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrorMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrorMessage() => clearField(2);

  @$pb.TagNumber(3)
  Path get exprErrorPath => $_getN(2);
  @$pb.TagNumber(3)
  set exprErrorPath(Path v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasExprErrorPath() => $_has(2);
  @$pb.TagNumber(3)
  void clearExprErrorPath() => clearField(3);
  @$pb.TagNumber(3)
  Path ensureExprErrorPath() => $_ensure(2);

  @$pb.TagNumber(4)
  EvaluateExprOutput_Status get exprStatus => $_getN(3);
  @$pb.TagNumber(4)
  set exprStatus(EvaluateExprOutput_Status v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasExprStatus() => $_has(3);
  @$pb.TagNumber(4)
  void clearExprStatus() => clearField(4);

  @$pb.TagNumber(5)
  $0.Value get value => $_getN(4);
  @$pb.TagNumber(5)
  set value($0.Value v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasValue() => $_has(4);
  @$pb.TagNumber(5)
  void clearValue() => clearField(5);
  @$pb.TagNumber(5)
  $0.Value ensureValue() => $_ensure(4);
}

class EvaluateExprInput extends $pb.GeneratedMessage {
  factory EvaluateExprInput({
    FunDefList? defs,
    $0.Value? expr,
    Path? path,
  }) {
    final $result = create();
    if (defs != null) {
      $result.defs = defs;
    }
    if (expr != null) {
      $result.expr = expr;
    }
    if (path != null) {
      $result.path = path;
    }
    return $result;
  }
  EvaluateExprInput._() : super();
  factory EvaluateExprInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EvaluateExprInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'EvaluateExprInput', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..aOM<FunDefList>(1, _omitFieldNames ? '' : 'defs', subBuilder: FunDefList.create)
    ..aOM<$0.Value>(2, _omitFieldNames ? '' : 'expr', subBuilder: $0.Value.create)
    ..aOM<Path>(10, _omitFieldNames ? '' : 'path', subBuilder: Path.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EvaluateExprInput clone() => EvaluateExprInput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EvaluateExprInput copyWith(void Function(EvaluateExprInput) updates) => super.copyWith((message) => updates(message as EvaluateExprInput)) as EvaluateExprInput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EvaluateExprInput create() => EvaluateExprInput._();
  EvaluateExprInput createEmptyInstance() => create();
  static $pb.PbList<EvaluateExprInput> createRepeated() => $pb.PbList<EvaluateExprInput>();
  @$core.pragma('dart2js:noInline')
  static EvaluateExprInput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EvaluateExprInput>(create);
  static EvaluateExprInput? _defaultInstance;

  @$pb.TagNumber(1)
  FunDefList get defs => $_getN(0);
  @$pb.TagNumber(1)
  set defs(FunDefList v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasDefs() => $_has(0);
  @$pb.TagNumber(1)
  void clearDefs() => clearField(1);
  @$pb.TagNumber(1)
  FunDefList ensureDefs() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.Value get expr => $_getN(1);
  @$pb.TagNumber(2)
  set expr($0.Value v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasExpr() => $_has(1);
  @$pb.TagNumber(2)
  void clearExpr() => clearField(2);
  @$pb.TagNumber(2)
  $0.Value ensureExpr() => $_ensure(1);

  @$pb.TagNumber(10)
  Path get path => $_getN(2);
  @$pb.TagNumber(10)
  set path(Path v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasPath() => $_has(2);
  @$pb.TagNumber(10)
  void clearPath() => clearField(10);
  @$pb.TagNumber(10)
  Path ensurePath() => $_ensure(2);
}

class EvaluateExprOutput extends $pb.GeneratedMessage {
  factory EvaluateExprOutput({
    EvaluateExprOutput_Status? status,
    $core.String? errorMessage,
    Path? errorPath,
    $0.Value? value,
  }) {
    final $result = create();
    if (status != null) {
      $result.status = status;
    }
    if (errorMessage != null) {
      $result.errorMessage = errorMessage;
    }
    if (errorPath != null) {
      $result.errorPath = errorPath;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  EvaluateExprOutput._() : super();
  factory EvaluateExprOutput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory EvaluateExprOutput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'EvaluateExprOutput', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..e<EvaluateExprOutput_Status>(1, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: EvaluateExprOutput_Status.OK, valueOf: EvaluateExprOutput_Status.valueOf, enumValues: EvaluateExprOutput_Status.values)
    ..aOS(2, _omitFieldNames ? '' : 'errorMessage')
    ..aOM<Path>(3, _omitFieldNames ? '' : 'errorPath', subBuilder: Path.create)
    ..aOM<$0.Value>(4, _omitFieldNames ? '' : 'value', subBuilder: $0.Value.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  EvaluateExprOutput clone() => EvaluateExprOutput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  EvaluateExprOutput copyWith(void Function(EvaluateExprOutput) updates) => super.copyWith((message) => updates(message as EvaluateExprOutput)) as EvaluateExprOutput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static EvaluateExprOutput create() => EvaluateExprOutput._();
  EvaluateExprOutput createEmptyInstance() => create();
  static $pb.PbList<EvaluateExprOutput> createRepeated() => $pb.PbList<EvaluateExprOutput>();
  @$core.pragma('dart2js:noInline')
  static EvaluateExprOutput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<EvaluateExprOutput>(create);
  static EvaluateExprOutput? _defaultInstance;

  @$pb.TagNumber(1)
  EvaluateExprOutput_Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(EvaluateExprOutput_Status v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get errorMessage => $_getSZ(1);
  @$pb.TagNumber(2)
  set errorMessage($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrorMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrorMessage() => clearField(2);

  @$pb.TagNumber(3)
  Path get errorPath => $_getN(2);
  @$pb.TagNumber(3)
  set errorPath(Path v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrorPath() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrorPath() => clearField(3);
  @$pb.TagNumber(3)
  Path ensureErrorPath() => $_ensure(2);

  @$pb.TagNumber(4)
  $0.Value get value => $_getN(3);
  @$pb.TagNumber(4)
  set value($0.Value v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasValue() => $_has(3);
  @$pb.TagNumber(4)
  void clearValue() => clearField(4);
  @$pb.TagNumber(4)
  $0.Value ensureValue() => $_ensure(3);
}

class EvaluatorApi {
  $pb.RpcClient _client;
  EvaluatorApi(this._client);

  $async.Future<EvaluateOutput> evaluate($pb.ClientContext? ctx, EvaluateInput request) =>
    _client.invoke<EvaluateOutput>(ctx, 'Evaluator', 'Evaluate', request, EvaluateOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateExpr($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateExpr', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateEval($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateEval', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateScalar($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateScalar', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateObj($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateObj', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateArr($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateArr', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateJson($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateJson', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateRangeIter($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateRangeIter', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateGetElem($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateGetElem', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateFunCall($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateFunCall', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateCases($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateCases', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateOpUnary($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateOpUnary', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateOpBinary($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateOpBinary', request, EvaluateExprOutput())
  ;
  $async.Future<EvaluateExprOutput> evaluateOpVariadic($pb.ClientContext? ctx, EvaluateExprInput request) =>
    _client.invoke<EvaluateExprOutput>(ctx, 'Evaluator', 'EvaluateOpVariadic', request, EvaluateExprOutput())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
