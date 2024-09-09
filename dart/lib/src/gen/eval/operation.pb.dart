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

export 'operation.pbenum.dart';

class OpUnary extends $pb.GeneratedMessage {
  factory OpUnary() => create();
  OpUnary._() : super();
  factory OpUnary.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OpUnary.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'OpUnary', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  OpUnary clone() => OpUnary()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  OpUnary copyWith(void Function(OpUnary) updates) => super.copyWith((message) => updates(message as OpUnary)) as OpUnary;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static OpUnary create() => OpUnary._();
  OpUnary createEmptyInstance() => create();
  static $pb.PbList<OpUnary> createRepeated() => $pb.PbList<OpUnary>();
  @$core.pragma('dart2js:noInline')
  static OpUnary getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OpUnary>(create);
  static OpUnary? _defaultInstance;
}

class OpBinary extends $pb.GeneratedMessage {
  factory OpBinary() => create();
  OpBinary._() : super();
  factory OpBinary.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OpBinary.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'OpBinary', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  OpBinary clone() => OpBinary()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  OpBinary copyWith(void Function(OpBinary) updates) => super.copyWith((message) => updates(message as OpBinary)) as OpBinary;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static OpBinary create() => OpBinary._();
  OpBinary createEmptyInstance() => create();
  static $pb.PbList<OpBinary> createRepeated() => $pb.PbList<OpBinary>();
  @$core.pragma('dart2js:noInline')
  static OpBinary getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OpBinary>(create);
  static OpBinary? _defaultInstance;
}

class OpVariadic extends $pb.GeneratedMessage {
  factory OpVariadic() => create();
  OpVariadic._() : super();
  factory OpVariadic.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OpVariadic.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'OpVariadic', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  OpVariadic clone() => OpVariadic()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  OpVariadic copyWith(void Function(OpVariadic) updates) => super.copyWith((message) => updates(message as OpVariadic)) as OpVariadic;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static OpVariadic create() => OpVariadic._();
  OpVariadic createEmptyInstance() => create();
  static $pb.PbList<OpVariadic> createRepeated() => $pb.PbList<OpVariadic>();
  @$core.pragma('dart2js:noInline')
  static OpVariadic getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OpVariadic>(create);
  static OpVariadic? _defaultInstance;
}


const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
