<?php

namespace Jumpaku\Eyamluate\Eval_;

interface Evaluator
{
    public function evaluate(PBEvaluateInput $input): PBEvaluateOutput;

    public function evaluateExpr(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateEval(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateScalar(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateObj(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateArr(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateJson(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateRangeIter(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateGetElem(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateFunCall(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateCases(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateOpUnary(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateOpBinary(PBEvaluateExprInput $input): PBEvaluateExprOutput;

    public function evaluateOpVariadic(PBEvaluateExprInput $input): PBEvaluateExprOutput;

}