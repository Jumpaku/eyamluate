//
//  Generated code. Do not modify.
//  source: eval/validator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'validator.pbenum.dart';

export 'validator.pbenum.dart';

class ValidateInput extends $pb.GeneratedMessage {
  factory ValidateInput({
    $core.String? source,
  }) {
    final $result = create();
    if (source != null) {
      $result.source = source;
    }
    return $result;
  }
  ValidateInput._() : super();
  factory ValidateInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ValidateInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ValidateInput', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'source')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ValidateInput clone() => ValidateInput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ValidateInput copyWith(void Function(ValidateInput) updates) => super.copyWith((message) => updates(message as ValidateInput)) as ValidateInput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ValidateInput create() => ValidateInput._();
  ValidateInput createEmptyInstance() => create();
  static $pb.PbList<ValidateInput> createRepeated() => $pb.PbList<ValidateInput>();
  @$core.pragma('dart2js:noInline')
  static ValidateInput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ValidateInput>(create);
  static ValidateInput? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get source => $_getSZ(0);
  @$pb.TagNumber(1)
  set source($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSource() => $_has(0);
  @$pb.TagNumber(1)
  void clearSource() => clearField(1);
}

class ValidateOutput extends $pb.GeneratedMessage {
  factory ValidateOutput({
    ValidateOutput_Status? status,
    $core.String? errorMessage,
  }) {
    final $result = create();
    if (status != null) {
      $result.status = status;
    }
    if (errorMessage != null) {
      $result.errorMessage = errorMessage;
    }
    return $result;
  }
  ValidateOutput._() : super();
  factory ValidateOutput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ValidateOutput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ValidateOutput', package: const $pb.PackageName(_omitMessageNames ? '' : 'eval'), createEmptyInstance: create)
    ..e<ValidateOutput_Status>(1, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: ValidateOutput_Status.OK, valueOf: ValidateOutput_Status.valueOf, enumValues: ValidateOutput_Status.values)
    ..aOS(2, _omitFieldNames ? '' : 'errorMessage')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ValidateOutput clone() => ValidateOutput()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ValidateOutput copyWith(void Function(ValidateOutput) updates) => super.copyWith((message) => updates(message as ValidateOutput)) as ValidateOutput;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ValidateOutput create() => ValidateOutput._();
  ValidateOutput createEmptyInstance() => create();
  static $pb.PbList<ValidateOutput> createRepeated() => $pb.PbList<ValidateOutput>();
  @$core.pragma('dart2js:noInline')
  static ValidateOutput getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ValidateOutput>(create);
  static ValidateOutput? _defaultInstance;

  @$pb.TagNumber(1)
  ValidateOutput_Status get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(ValidateOutput_Status v) { setField(1, v); }
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
}

class ValidatorApi {
  $pb.RpcClient _client;
  ValidatorApi(this._client);

  $async.Future<ValidateOutput> validate($pb.ClientContext? ctx, ValidateInput request) =>
    _client.invoke<ValidateOutput>(ctx, 'Validator', 'Validate', request, ValidateOutput())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
