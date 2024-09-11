<?php

namespace Jumpaku\Eyamluate\Eval_;

class OpBinary
{
    public static function keyName(int $o): string
    {
        switch ($o) {
            case PBOpBinary\PBOperator::SUB:
                return "sub";
            case PBOpBinary\PBOperator::DIV:
                return "div";
            case PBOpBinary\PBOperator::EQ:
                return "eq";
            case PBOpBinary\PBOperator::NEQ:
                return "neq";
            case PBOpBinary\PBOperator::LT:
                return "lt";
            case PBOpBinary\PBOperator::LTE:
                return "lte";
            case PBOpBinary\PBOperator::GT:
                return "gt";
            case PBOpBinary\PBOperator::GTE:
                return "gte";
            default:
                assert(false, sprintf("unexpected OperatorBinary: %s", PBOpBinary\PBOperator::name($o)));
        }
    }
}