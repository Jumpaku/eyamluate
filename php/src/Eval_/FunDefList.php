<?php

namespace Jumpaku\Eyamluate\Eval_;

class FunDefList
{
    public static function empty(): PBFunDefList
    {
        return new PBFunDefList();
    }

    public static function register(PBFunDefList $l, PBFunDef $fd): PBFunDefList
    {
        $newList = new PBFunDefList();
        $newList->setParent($l);
        $newList->setDef($fd);
        return $newList;
    }

    public static function find(PBFunDefList $l, string $ident): PBFunDefList|null
    {
        $cur = $l;
        while (true) {
            if ($cur->hasDef() === false) {
                return null;
            }
            if ($cur->getDef()->getDef() === $ident) {
                return $cur;
            }
            if ($cur->hasParent() === false) {
                return null;
            }
            $cur = $cur->getParent();
        }
    }
}