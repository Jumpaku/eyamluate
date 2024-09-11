<?php

namespace Jumpaku\Eyamluate\Yaml;

use Exception;
use stdClass;
use Symfony\Component\Yaml\Yaml as SymfonyYaml;

class Encoder
{
    public function encode(PBEncodeInput $input): PBEncodeOutput
    {
        try {
            $y = Encoder::convertToPHP($input->getValue());
            $r = new PBEncodeOutput();
            $flags = SymfonyYaml::DUMP_OBJECT_AS_MAP
                | SymfonyYaml::DUMP_EMPTY_ARRAY_AS_SEQUENCE
                | SymfonyYaml::DUMP_MULTI_LINE_LITERAL_BLOCK
                | SymfonyYaml::DUMP_NUMERIC_KEY_AS_STRING;
            $r->setResult(SymfonyYaml::dump($y, 10, 2, $flags));
            return $r;
        } catch (Exception $e) {
            $r = new PBEncodeOutput();
            $r->setIsError(true);
            $r->setErrorMessage($e->getMessage());
            return $r;
        }
    }

    private static function convertToPHP(PBValue $v): mixed
    {
        switch ($v->getType()) {
            case PBType::TYPE_NULL:
                return null;
            case PBType::TYPE_NUM:
                return $v->getNum();
            case PBType::TYPE_BOOL:
                return $v->getBool();
            case PBType::TYPE_STR:
                return $v->getStr();
            case PBType::TYPE_ARR:
                $arr = [];
                foreach ($v->getArr() as $elem) {
                    $arr[] = Encoder::convertToPHP($elem);
                }
                return $arr;
            case PBType::TYPE_OBJ:
                $obj = new stdClass();
                foreach ($v->getObj() as $key => $val) {
                    $obj->$key = Encoder::convertToPHP($val);
                }
                return $obj;
            default:
                throw new Exception('Invalid type');
        }
    }
}
