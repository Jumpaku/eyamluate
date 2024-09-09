//
//  Generated code. Do not modify.
//  source: eval/evaluator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

import '../yaml/value.pbjson.dart' as $0;

@$core.Deprecated('Use funDefDescriptor instead')
const FunDef$json = {
  '1': 'FunDef',
  '2': [
    {'1': 'def', '3': 1, '4': 1, '5': 9, '10': 'def'},
    {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.yaml.Value', '10': 'value'},
    {'1': 'with', '3': 3, '4': 3, '5': 9, '10': 'with'},
    {'1': 'path', '3': 10, '4': 1, '5': 11, '6': '.eval.Path', '10': 'path'},
  ],
};

/// Descriptor for `FunDef`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List funDefDescriptor = $convert.base64Decode(
    'CgZGdW5EZWYSEAoDZGVmGAEgASgJUgNkZWYSIQoFdmFsdWUYAiABKAsyCy55YW1sLlZhbHVlUg'
    'V2YWx1ZRISCgR3aXRoGAMgAygJUgR3aXRoEh4KBHBhdGgYCiABKAsyCi5ldmFsLlBhdGhSBHBh'
    'dGg=');

@$core.Deprecated('Use funDefListDescriptor instead')
const FunDefList$json = {
  '1': 'FunDefList',
  '2': [
    {'1': 'parent', '3': 1, '4': 1, '5': 11, '6': '.eval.FunDefList', '10': 'parent'},
    {'1': 'def', '3': 2, '4': 1, '5': 11, '6': '.eval.FunDef', '10': 'def'},
  ],
};

/// Descriptor for `FunDefList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List funDefListDescriptor = $convert.base64Decode(
    'CgpGdW5EZWZMaXN0EigKBnBhcmVudBgBIAEoCzIQLmV2YWwuRnVuRGVmTGlzdFIGcGFyZW50Eh'
    '4KA2RlZhgCIAEoCzIMLmV2YWwuRnVuRGVmUgNkZWY=');

@$core.Deprecated('Use pathDescriptor instead')
const Path$json = {
  '1': 'Path',
  '2': [
    {'1': 'pos', '3': 1, '4': 3, '5': 11, '6': '.eval.Path.Pos', '10': 'pos'},
  ],
  '3': [Path_Pos$json],
};

@$core.Deprecated('Use pathDescriptor instead')
const Path_Pos$json = {
  '1': 'Pos',
  '2': [
    {'1': 'index', '3': 1, '4': 1, '5': 3, '10': 'index'},
    {'1': 'key', '3': 2, '4': 1, '5': 9, '10': 'key'},
  ],
};

/// Descriptor for `Path`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pathDescriptor = $convert.base64Decode(
    'CgRQYXRoEiAKA3BvcxgBIAMoCzIOLmV2YWwuUGF0aC5Qb3NSA3BvcxotCgNQb3MSFAoFaW5kZX'
    'gYASABKANSBWluZGV4EhAKA2tleRgCIAEoCVIDa2V5');

@$core.Deprecated('Use evaluateInputDescriptor instead')
const EvaluateInput$json = {
  '1': 'EvaluateInput',
  '2': [
    {'1': 'source', '3': 1, '4': 1, '5': 9, '10': 'source'},
  ],
};

/// Descriptor for `EvaluateInput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List evaluateInputDescriptor = $convert.base64Decode(
    'Cg1FdmFsdWF0ZUlucHV0EhYKBnNvdXJjZRgBIAEoCVIGc291cmNl');

@$core.Deprecated('Use evaluateOutputDescriptor instead')
const EvaluateOutput$json = {
  '1': 'EvaluateOutput',
  '2': [
    {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.eval.EvaluateOutput.Status', '10': 'status'},
    {'1': 'error_message', '3': 2, '4': 1, '5': 9, '10': 'errorMessage'},
    {'1': 'expr_error_path', '3': 3, '4': 1, '5': 11, '6': '.eval.Path', '10': 'exprErrorPath'},
    {'1': 'expr_status', '3': 4, '4': 1, '5': 14, '6': '.eval.EvaluateExprOutput.Status', '10': 'exprStatus'},
    {'1': 'value', '3': 5, '4': 1, '5': 11, '6': '.yaml.Value', '10': 'value'},
  ],
  '4': [EvaluateOutput_Status$json],
};

@$core.Deprecated('Use evaluateOutputDescriptor instead')
const EvaluateOutput_Status$json = {
  '1': 'Status',
  '2': [
    {'1': 'OK', '2': 0},
    {'1': 'DECODE_ERROR', '2': 1},
    {'1': 'VALIDATE_ERROR', '2': 2},
    {'1': 'EXPR_ERROR', '2': 3},
  ],
};

/// Descriptor for `EvaluateOutput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List evaluateOutputDescriptor = $convert.base64Decode(
    'Cg5FdmFsdWF0ZU91dHB1dBIzCgZzdGF0dXMYASABKA4yGy5ldmFsLkV2YWx1YXRlT3V0cHV0Ll'
    'N0YXR1c1IGc3RhdHVzEiMKDWVycm9yX21lc3NhZ2UYAiABKAlSDGVycm9yTWVzc2FnZRIyCg9l'
    'eHByX2Vycm9yX3BhdGgYAyABKAsyCi5ldmFsLlBhdGhSDWV4cHJFcnJvclBhdGgSQAoLZXhwcl'
    '9zdGF0dXMYBCABKA4yHy5ldmFsLkV2YWx1YXRlRXhwck91dHB1dC5TdGF0dXNSCmV4cHJTdGF0'
    'dXMSIQoFdmFsdWUYBSABKAsyCy55YW1sLlZhbHVlUgV2YWx1ZSJGCgZTdGF0dXMSBgoCT0sQAB'
    'IQCgxERUNPREVfRVJST1IQARISCg5WQUxJREFURV9FUlJPUhACEg4KCkVYUFJfRVJST1IQAw==');

@$core.Deprecated('Use evaluateExprInputDescriptor instead')
const EvaluateExprInput$json = {
  '1': 'EvaluateExprInput',
  '2': [
    {'1': 'path', '3': 10, '4': 1, '5': 11, '6': '.eval.Path', '10': 'path'},
    {'1': 'defs', '3': 1, '4': 1, '5': 11, '6': '.eval.FunDefList', '10': 'defs'},
    {'1': 'expr', '3': 2, '4': 1, '5': 11, '6': '.yaml.Value', '10': 'expr'},
  ],
};

/// Descriptor for `EvaluateExprInput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List evaluateExprInputDescriptor = $convert.base64Decode(
    'ChFFdmFsdWF0ZUV4cHJJbnB1dBIeCgRwYXRoGAogASgLMgouZXZhbC5QYXRoUgRwYXRoEiQKBG'
    'RlZnMYASABKAsyEC5ldmFsLkZ1bkRlZkxpc3RSBGRlZnMSHwoEZXhwchgCIAEoCzILLnlhbWwu'
    'VmFsdWVSBGV4cHI=');

@$core.Deprecated('Use evaluateExprOutputDescriptor instead')
const EvaluateExprOutput$json = {
  '1': 'EvaluateExprOutput',
  '2': [
    {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.eval.EvaluateExprOutput.Status', '10': 'status'},
    {'1': 'error_message', '3': 2, '4': 1, '5': 9, '10': 'errorMessage'},
    {'1': 'error_path', '3': 3, '4': 1, '5': 11, '6': '.eval.Path', '10': 'errorPath'},
    {'1': 'value', '3': 4, '4': 1, '5': 11, '6': '.yaml.Value', '10': 'value'},
  ],
  '4': [EvaluateExprOutput_Status$json],
};

@$core.Deprecated('Use evaluateExprOutputDescriptor instead')
const EvaluateExprOutput_Status$json = {
  '1': 'Status',
  '2': [
    {'1': 'OK', '2': 0},
    {'1': 'UNSUPPORTED_EXPR', '2': 1},
    {'1': 'UNEXPECTED_TYPE', '2': 2},
    {'1': 'ARITHMETIC_ERROR', '2': 3},
    {'1': 'INDEX_OUT_OF_BOUNDS', '2': 4},
    {'1': 'KEY_NOT_FOUND', '2': 5},
    {'1': 'REFERENCE_NOT_FOUND', '2': 6},
    {'1': 'CASES_NOT_EXHAUSTIVE', '2': 7},
    {'1': 'UNSUPPORTED_OPERATION', '2': 8},
    {'1': 'ABORTED', '2': 9},
    {'1': 'UNKNOWN', '2': 10},
  ],
};

/// Descriptor for `EvaluateExprOutput`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List evaluateExprOutputDescriptor = $convert.base64Decode(
    'ChJFdmFsdWF0ZUV4cHJPdXRwdXQSNwoGc3RhdHVzGAEgASgOMh8uZXZhbC5FdmFsdWF0ZUV4cH'
    'JPdXRwdXQuU3RhdHVzUgZzdGF0dXMSIwoNZXJyb3JfbWVzc2FnZRgCIAEoCVIMZXJyb3JNZXNz'
    'YWdlEikKCmVycm9yX3BhdGgYAyABKAsyCi5ldmFsLlBhdGhSCWVycm9yUGF0aBIhCgV2YWx1ZR'
    'gEIAEoCzILLnlhbWwuVmFsdWVSBXZhbHVlIuUBCgZTdGF0dXMSBgoCT0sQABIUChBVTlNVUFBP'
    'UlRFRF9FWFBSEAESEwoPVU5FWFBFQ1RFRF9UWVBFEAISFAoQQVJJVEhNRVRJQ19FUlJPUhADEh'
    'cKE0lOREVYX09VVF9PRl9CT1VORFMQBBIRCg1LRVlfTk9UX0ZPVU5EEAUSFwoTUkVGRVJFTkNF'
    'X05PVF9GT1VORBAGEhgKFENBU0VTX05PVF9FWEhBVVNUSVZFEAcSGQoVVU5TVVBQT1JURURfT1'
    'BFUkFUSU9OEAgSCwoHQUJPUlRFRBAJEgsKB1VOS05PV04QCg==');

const $core.Map<$core.String, $core.dynamic> EvaluatorServiceBase$json = {
  '1': 'Evaluator',
  '2': [
    {'1': 'Evaluate', '2': '.eval.EvaluateInput', '3': '.eval.EvaluateOutput', '4': {}},
    {'1': 'EvaluateExpr', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateEval', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateScalar', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateObj', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateArr', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateJson', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateRangeIter', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateGetElem', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateFunCall', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateCases', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateOpUnary', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateOpBinary', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
    {'1': 'EvaluateOpVariadic', '2': '.eval.EvaluateExprInput', '3': '.eval.EvaluateExprOutput', '4': {}},
  ],
};

@$core.Deprecated('Use evaluatorServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> EvaluatorServiceBase$messageJson = {
  '.eval.EvaluateInput': EvaluateInput$json,
  '.eval.EvaluateOutput': EvaluateOutput$json,
  '.eval.Path': Path$json,
  '.eval.Path.Pos': Path_Pos$json,
  '.yaml.Value': $0.Value$json,
  '.yaml.Value.ObjEntry': $0.Value_ObjEntry$json,
  '.eval.EvaluateExprInput': EvaluateExprInput$json,
  '.eval.FunDefList': FunDefList$json,
  '.eval.FunDef': FunDef$json,
  '.eval.EvaluateExprOutput': EvaluateExprOutput$json,
};

/// Descriptor for `Evaluator`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List evaluatorServiceDescriptor = $convert.base64Decode(
    'CglFdmFsdWF0b3ISNwoIRXZhbHVhdGUSEy5ldmFsLkV2YWx1YXRlSW5wdXQaFC5ldmFsLkV2YW'
    'x1YXRlT3V0cHV0IgASQwoMRXZhbHVhdGVFeHByEhcuZXZhbC5FdmFsdWF0ZUV4cHJJbnB1dBoY'
    'LmV2YWwuRXZhbHVhdGVFeHByT3V0cHV0IgASQwoMRXZhbHVhdGVFdmFsEhcuZXZhbC5FdmFsdW'
    'F0ZUV4cHJJbnB1dBoYLmV2YWwuRXZhbHVhdGVFeHByT3V0cHV0IgASRQoORXZhbHVhdGVTY2Fs'
    'YXISFy5ldmFsLkV2YWx1YXRlRXhwcklucHV0GhguZXZhbC5FdmFsdWF0ZUV4cHJPdXRwdXQiAB'
    'JCCgtFdmFsdWF0ZU9iahIXLmV2YWwuRXZhbHVhdGVFeHBySW5wdXQaGC5ldmFsLkV2YWx1YXRl'
    'RXhwck91dHB1dCIAEkIKC0V2YWx1YXRlQXJyEhcuZXZhbC5FdmFsdWF0ZUV4cHJJbnB1dBoYLm'
    'V2YWwuRXZhbHVhdGVFeHByT3V0cHV0IgASQwoMRXZhbHVhdGVKc29uEhcuZXZhbC5FdmFsdWF0'
    'ZUV4cHJJbnB1dBoYLmV2YWwuRXZhbHVhdGVFeHByT3V0cHV0IgASSAoRRXZhbHVhdGVSYW5nZU'
    'l0ZXISFy5ldmFsLkV2YWx1YXRlRXhwcklucHV0GhguZXZhbC5FdmFsdWF0ZUV4cHJPdXRwdXQi'
    'ABJGCg9FdmFsdWF0ZUdldEVsZW0SFy5ldmFsLkV2YWx1YXRlRXhwcklucHV0GhguZXZhbC5Fdm'
    'FsdWF0ZUV4cHJPdXRwdXQiABJGCg9FdmFsdWF0ZUZ1bkNhbGwSFy5ldmFsLkV2YWx1YXRlRXhw'
    'cklucHV0GhguZXZhbC5FdmFsdWF0ZUV4cHJPdXRwdXQiABJECg1FdmFsdWF0ZUNhc2VzEhcuZX'
    'ZhbC5FdmFsdWF0ZUV4cHJJbnB1dBoYLmV2YWwuRXZhbHVhdGVFeHByT3V0cHV0IgASRgoPRXZh'
    'bHVhdGVPcFVuYXJ5EhcuZXZhbC5FdmFsdWF0ZUV4cHJJbnB1dBoYLmV2YWwuRXZhbHVhdGVFeH'
    'ByT3V0cHV0IgASRwoQRXZhbHVhdGVPcEJpbmFyeRIXLmV2YWwuRXZhbHVhdGVFeHBySW5wdXQa'
    'GC5ldmFsLkV2YWx1YXRlRXhwck91dHB1dCIAEkkKEkV2YWx1YXRlT3BWYXJpYWRpYxIXLmV2YW'
    'wuRXZhbHVhdGVFeHBySW5wdXQaGC5ldmFsLkV2YWx1YXRlRXhwck91dHB1dCIA');

