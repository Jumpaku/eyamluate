//
//  Generated code. Do not modify.
//  source: yaml/decoder.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'value.pb.dart' as $0;

class DecodeInput extends $pb.GeneratedMessage {
  factory DecodeInput({
    $core.String? yaml,
  }) {
    final $result = create();
    if (yaml != null) {
      $result.yaml = yaml;
    }
    return $result;
  }
  DecodeInput._() : super();
  factory DecodeInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DecodeInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DecodeInput', package: const $pb.PackageName(_omitMessageNames ? '' : 'yaml'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'yaml')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DecodeInput clone() => DecodeInput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DecodeInput copyWith(void Function(DecodeInput) updates) => super.copyWith((message) => updates(message as DecodeInput)) as DecodeInput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DecodeInput create() => DecodeInput._();
  DecodeInput createEmptyInstance() => create();
  static $pb.PbList<DecodeInput> createRepeated() => $pb.PbList<DecodeInput>();
  @$core.pragma('dart2js:noInline')
  static DecodeInput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DecodeInput>(create);
  static DecodeInput? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get yaml => $_getSZ(0);
  @$pb.TagNumber(1)
  set yaml($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasYaml() => $_has(0);
  @$pb.TagNumber(1)
  void clearYaml() => clearField(1);
}

class DecodeOutput extends $pb.GeneratedMessage {
  factory DecodeOutput({
    $core.bool? isError,
    $core.String? errorMessage,
    $0.Value? value,
  }) {
    final $result = create();
    if (isError != null) {
      $result.isError = isError;
    }
    if (errorMessage != null) {
      $result.errorMessage = errorMessage;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  DecodeOutput._() : super();
  factory DecodeOutput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DecodeOutput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DecodeOutput', package: const $pb.PackageName(_omitMessageNames ? '' : 'yaml'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isError')
    ..aOS(2, _omitFieldNames ? '' : 'errorMessage')
    ..aOM<$0.Value>(3, _omitFieldNames ? '' : 'value', subBuilder: $0.Value.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DecodeOutput clone() => DecodeOutput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DecodeOutput copyWith(void Function(DecodeOutput) updates) => super.copyWith((message) => updates(message as DecodeOutput)) as DecodeOutput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DecodeOutput create() => DecodeOutput._();
  DecodeOutput createEmptyInstance() => create();
  static $pb.PbList<DecodeOutput> createRepeated() => $pb.PbList<DecodeOutput>();
  @$core.pragma('dart2js:noInline')
  static DecodeOutput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DecodeOutput>(create);
  static DecodeOutput? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isError => $_getBF(0);
  @$pb.TagNumber(1)
  set isError($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsError() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsError() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get errorMessage => $_getSZ(1);
  @$pb.TagNumber(2)
  set errorMessage($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrorMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrorMessage() => clearField(2);

  @$pb.TagNumber(3)
  $0.Value get value => $_getN(2);
  @$pb.TagNumber(3)
  set value($0.Value v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasValue() => $_has(2);
  @$pb.TagNumber(3)
  void clearValue() => clearField(3);
  @$pb.TagNumber(3)
  $0.Value ensureValue() => $_ensure(2);
}

class DecoderApi {
  $pb.RpcClient _client;
  DecoderApi(this._client);

  $async.Future<DecodeOutput> decode($pb.ClientContext? ctx, DecodeInput request) =>
    _client.invoke<DecodeOutput>(ctx, 'Decoder', 'Decode', request, DecodeOutput())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
