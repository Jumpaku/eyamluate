<?php

namespace Jumpaku\Eyamluate;

use Jumpaku\Eyamluate\Yaml\PBDecodeInput;
use Jumpaku\Eyamluate\Yaml\PBDecodeOutput;
use Symfony\Component\Yaml\Yaml as SymfonyYaml;

class Decoder
{
    public function decode(PBDecodeInput $input): PBDecodeOutput
    {
        $y = SymfonyYaml::parse($input->getYaml(), SymfonyYaml::PARSE_OBJECT_FOR_MAP);

        var_dump($y);

        return new PBDecodeOutput;
    }
}

function convertFromYaml()
{
    
}