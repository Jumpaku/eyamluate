<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eval/evaluator.proto

namespace Jumpaku\Eyamluate\Metadata\Eval_;

class Evaluator
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \Jumpaku\Eyamluate\Metadata\Yaml\Value::initOnce();
        $pool->internalAddGeneratedFile(
            "\x0A\xB4\x12\x0A\x14eval/evaluator.proto\x12\x04eval\"q\x0A\x06FunDef\x12\x10\x0A\x03def\x18\x01 \x01(\x09R\x03def\x12!\x0A\x05value\x18\x02 \x01(\x0B2\x0B.yaml.ValueR\x05value\x12\x12\x0A\x04with\x18\x03 \x03(\x09R\x04with\x12\x1E\x0A\x04path\x18\x0A \x01(\x0B2\x0A.eval.PathR\x04path\"V\x0A\x0AFunDefList\x12(\x0A\x06parent\x18\x01 \x01(\x0B2\x10.eval.FunDefListR\x06parent\x12\x1E\x0A\x03def\x18\x02 \x01(\x0B2\x0C.eval.FunDefR\x03def\"W\x0A\x04Path\x12 \x0A\x03pos\x18\x01 \x03(\x0B2\x0E.eval.Path.PosR\x03pos\x1A-\x0A\x03Pos\x12\x14\x0A\x05index\x18\x01 \x01(\x03R\x05index\x12\x10\x0A\x03key\x18\x02 \x01(\x09R\x03key\"'\x0A\x0DEvaluateInput\x12\x16\x0A\x06source\x18\x01 \x01(\x09R\x06source\"\xCB\x02\x0A\x0EEvaluateOutput\x123\x0A\x06status\x18\x01 \x01(\x0E2\x1B.eval.EvaluateOutput.StatusR\x06status\x12#\x0A\x0Derror_message\x18\x02 \x01(\x09R\x0CerrorMessage\x122\x0A\x0Fexpr_error_path\x18\x03 \x01(\x0B2\x0A.eval.PathR\x0DexprErrorPath\x12@\x0A\x0Bexpr_status\x18\x04 \x01(\x0E2\x1F.eval.EvaluateExprOutput.StatusR\x0AexprStatus\x12!\x0A\x05value\x18\x05 \x01(\x0B2\x0B.yaml.ValueR\x05value\"F\x0A\x06Status\x12\x06\x0A\x02OK\x10\x00\x12\x10\x0A\x0CDECODE_ERROR\x10\x01\x12\x12\x0A\x0EVALIDATE_ERROR\x10\x02\x12\x0E\x0A\x0AEXPR_ERROR\x10\x03\"z\x0A\x11EvaluateExprInput\x12\x1E\x0A\x04path\x18\x0A \x01(\x0B2\x0A.eval.PathR\x04path\x12\$\x0A\x04defs\x18\x01 \x01(\x0B2\x10.eval.FunDefListR\x04defs\x12\x1F\x0A\x04expr\x18\x02 \x01(\x0B2\x0B.yaml.ValueR\x04expr\"\xA8\x03\x0A\x12EvaluateExprOutput\x127\x0A\x06status\x18\x01 \x01(\x0E2\x1F.eval.EvaluateExprOutput.StatusR\x06status\x12#\x0A\x0Derror_message\x18\x02 \x01(\x09R\x0CerrorMessage\x12)\x0A\x0Aerror_path\x18\x03 \x01(\x0B2\x0A.eval.PathR\x09errorPath\x12!\x0A\x05value\x18\x04 \x01(\x0B2\x0B.yaml.ValueR\x05value\"\xE5\x01\x0A\x06Status\x12\x06\x0A\x02OK\x10\x00\x12\x14\x0A\x10UNSUPPORTED_EXPR\x10\x01\x12\x13\x0A\x0FUNEXPECTED_TYPE\x10\x02\x12\x14\x0A\x10ARITHMETIC_ERROR\x10\x03\x12\x17\x0A\x13INDEX_OUT_OF_BOUNDS\x10\x04\x12\x11\x0A\x0DKEY_NOT_FOUND\x10\x05\x12\x17\x0A\x13REFERENCE_NOT_FOUND\x10\x06\x12\x18\x0A\x14CASES_NOT_EXHAUSTIVE\x10\x07\x12\x19\x0A\x15UNSUPPORTED_OPERATION\x10\x08\x12\x0B\x0A\x07ABORTED\x10\x09\x12\x0B\x0A\x07UNKNOWN\x10\x0A2\xDE\x07\x0A\x09Evaluator\x127\x0A\x08Evaluate\x12\x13.eval.EvaluateInput\x1A\x14.eval.EvaluateOutput\"\x00\x12C\x0A\x0CEvaluateExpr\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12C\x0A\x0CEvaluateEval\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12E\x0A\x0EEvaluateScalar\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12B\x0A\x0BEvaluateObj\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12B\x0A\x0BEvaluateArr\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12C\x0A\x0CEvaluateJson\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12H\x0A\x11EvaluateRangeIter\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12F\x0A\x0FEvaluateGetElem\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12F\x0A\x0FEvaluateFunCall\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12D\x0A\x0DEvaluateCases\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12F\x0A\x0FEvaluateOpUnary\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12G\x0A\x10EvaluateOpBinary\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00\x12I\x0A\x12EvaluateOpVariadic\x12\x17.eval.EvaluateExprInput\x1A\x18.eval.EvaluateExprOutput\"\x00BkZ'github.com/Jumpaku/eyamlate/golang/eval\xC2\x02\x02PB\xCA\x02\x17Jumpaku\\Eyamluate\\Eval_\xE2\x02 Jumpaku\\Eyamluate\\Metadata\\Eval_b\x06proto3"
        , true);

        static::$is_initialized = true;
    }
}

