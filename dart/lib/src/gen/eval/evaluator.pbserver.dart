//
//  Generated code. Do not modify.
//  source: eval/evaluator.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'evaluator.pb.dart' as $1;
import 'evaluator.pbjson.dart';

export 'evaluator.pb.dart';

abstract class EvaluatorServiceBase extends $pb.GeneratedService {
  $async.Future<$1.EvaluateOutput> evaluate($pb.ServerContext ctx, $1.EvaluateInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateExpr($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateEval($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateScalar($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateObj($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateArr($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateJson($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateRangeIter($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateGetElem($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateFunCall($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateCases($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateOpUnary($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateOpBinary($pb.ServerContext ctx, $1.EvaluateExprInput request);
  $async.Future<$1.EvaluateExprOutput> evaluateOpVariadic($pb.ServerContext ctx, $1.EvaluateExprInput request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'Evaluate': return $1.EvaluateInput();
      case 'EvaluateExpr': return $1.EvaluateExprInput();
      case 'EvaluateEval': return $1.EvaluateExprInput();
      case 'EvaluateScalar': return $1.EvaluateExprInput();
      case 'EvaluateObj': return $1.EvaluateExprInput();
      case 'EvaluateArr': return $1.EvaluateExprInput();
      case 'EvaluateJson': return $1.EvaluateExprInput();
      case 'EvaluateRangeIter': return $1.EvaluateExprInput();
      case 'EvaluateGetElem': return $1.EvaluateExprInput();
      case 'EvaluateFunCall': return $1.EvaluateExprInput();
      case 'EvaluateCases': return $1.EvaluateExprInput();
      case 'EvaluateOpUnary': return $1.EvaluateExprInput();
      case 'EvaluateOpBinary': return $1.EvaluateExprInput();
      case 'EvaluateOpVariadic': return $1.EvaluateExprInput();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'Evaluate': return this.evaluate(ctx, request as $1.EvaluateInput);
      case 'EvaluateExpr': return this.evaluateExpr(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateEval': return this.evaluateEval(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateScalar': return this.evaluateScalar(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateObj': return this.evaluateObj(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateArr': return this.evaluateArr(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateJson': return this.evaluateJson(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateRangeIter': return this.evaluateRangeIter(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateGetElem': return this.evaluateGetElem(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateFunCall': return this.evaluateFunCall(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateCases': return this.evaluateCases(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateOpUnary': return this.evaluateOpUnary(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateOpBinary': return this.evaluateOpBinary(ctx, request as $1.EvaluateExprInput);
      case 'EvaluateOpVariadic': return this.evaluateOpVariadic(ctx, request as $1.EvaluateExprInput);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => EvaluatorServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => EvaluatorServiceBase$messageJson;
}

