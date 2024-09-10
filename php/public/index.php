<?php

use Jumpaku\Eyamluate\Yaml\PBDecodeInput;
use Jumpaku\Eyamluate\Decoder;

require_once __DIR__ . '/../vendor/autoload.php';

$d = new Decoder();
$i = new PBDecodeInput;
$i->setYaml('{}');
$d->decode($i);
