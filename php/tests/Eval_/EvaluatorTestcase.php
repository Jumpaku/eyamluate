<?php

namespace Jumpaku\Eyamluate\Tests\Eval_;

use Jumpaku\Eyamluate\Yaml\PBValue;

class EvaluatorTestcase
{
    public string $inputYaml = "";
    public PBValue|null $wantValue = null;
    public bool|null $wantError = null;

}