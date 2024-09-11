<?php

namespace Jumpaku\Eyamluate\Eval_;

class OpVariadic
{
    public static function keyName(int $o): string
    {
        switch ($o) {
            case PBOpVariadic\PBOperator::ADD:
                return "add";
            case PBOpVariadic\PBOperator::MUL:
                return "mul";
            case PBOpVariadic\PBOperator::PBAND:
                return "and";
            case PBOpVariadic\PBOperator::PBOR:
                return "or";
            case PBOpVariadic\PBOperator::CAT:
                return "cat";
            case PBOpVariadic\PBOperator::MIN:
                return "min";
            case PBOpVariadic\PBOperator::MAX:
                return "max";
            case PBOpVariadic\PBOperator::MERGE:
                return "merge";
            default:
                assert(false, sprintf("unexpected OperatorVariadic: %s", PBOpVariadic\PBOperator::name($o)));
        }
    }

}