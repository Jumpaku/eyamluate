<?php
namespace Jumpaku\Eyamluate\Eval_;

class Path
{
    public static function append(PBPath $path, string|int $pos): PBPath
    {
        $newPos = [];
        foreach ($path->getPos() as $p) {
            $newPos[] = $p;
        }
        $p = new PBPath\PBPos();
        if (is_int($pos)) {
            $p->setIndex($pos);
        } else {
            $p->setKey($pos);
        }
        $newPos[] = $p;

        $newPath = new PBPath();
        $path->setPos($newPos);
        return $newPath;
    }
}