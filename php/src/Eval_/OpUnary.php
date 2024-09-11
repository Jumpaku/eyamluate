<?php

namespace Jumpaku\Eyamluate\Eval_;

class OpUnary
{
    public static function keyName(int $o): string
    {
        switch ($o) {
            case PBOpUnary\PBOperator::LEN:
                return "len";
            case PBOpUnary\PBOperator::NOT:
                return "not";
            case PBOpUnary\PBOperator::FLAT:
                return "flat";
            case PBOpUnary\PBOperator::FLOOR:
                return "floor";
            case PBOpUnary\PBOperator::CEIL:
                return "ceil";
            case PBOpUnary\PBOperator::ABORT:
                return "abort";
            default:
                assert(false, sprintf("unexpected OperatorUnary: %s", PBOpUnary\PBOperator::name($o)));
        }
    }
}