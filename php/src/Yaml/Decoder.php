<?php

namespace Jumpaku\Eyamluate\Yaml;

use Exception;
use Symfony\Component\Yaml\Yaml as SymfonyYaml;

class Decoder
{
    public function decode(PBDecodeInput $input): PBDecodeOutput
    {
        try {
            $y = SymfonyYaml::parse($input->getYaml(), SymfonyYaml::PARSE_OBJECT_FOR_MAP);
            $v = Decoder::convertFromPHP($y);
            return (new PBDecodeOutput)->setValue($v);
        } catch (Exception $e) {
            $r = new PBDecodeOutput();
            $r->setIsError(true);
            $r->setErrorMessage($e->getMessage());
            return $r;
        }
    }


    /**
     * @throws Exception
     */
    private static function convertFromPHP(mixed $yaml): PBValue
    {
        $v = new PBValue();
        switch (gettype($yaml)) {
            case 'NULL':
                $v->setType(PBType::TYPE_NULL);
                return $v;
            case 'integer':
            case 'double':
                $v->setType(PBType::TYPE_NUM);
                $v->setNum((float)$yaml);
                return $v;
            case 'boolean':
                $v->setType(PBType::TYPE_BOOL);
                $v->setBool((bool)$yaml);
                return $v;
            case 'string':
                $v->setType(PBType::TYPE_STR);
                $v->setStr((string)$yaml);
                return $v;
            case 'array':
                $v->setType(PBType::TYPE_ARR);
                $arr = [];
                foreach ($yaml as $elem) {
                    $arr[] = Decoder::convertFromPHP($elem);
                }
                $v->setArr($arr);
                return $v;
            case 'object':
                $v->setType(PBType::TYPE_OBJ);
                $obj = [];
                foreach ($yaml as $key => $val) {
                    $obj[$key] = Decoder::convertFromPHP($val);
                }
                $v->setObj($obj);
                return $v;
            default:
                throw new Exception('Unsupported type: ' . gettype($yaml));
        }
    }
}
