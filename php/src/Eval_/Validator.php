<?php

namespace Jumpaku\Eyamluate\Eval_;

use Exception;
use Jumpaku\Eyamluate\Schema\EyamluateSchemaYaml;
use Symfony\Component\Yaml\Yaml;

class Validator
{
    private mixed $schema;

    function __construct()
    {
        $this->schema = Yaml::parse(EyamluateSchemaYaml::CONTENT, Yaml::PARSE_OBJECT_FOR_MAP);
    }

    function validate(PBValidateInput $input): PBValidateOutput
    {
        $r = new PBValidateOutput();
        try {
            $data = Yaml::parse($input->getSource(), Yaml::PARSE_OBJECT_FOR_MAP);
        } catch (Exception $e) {
            $r->setStatus(PBValidateOutput\PBStatus::YAML_ERROR);
            $r->setErrorMessage($e->getMessage());
            return $r;
        }

        $validator = new \JsonSchema\Validator;
        $validator->validate($data, $this->schema);
        if (!$validator->isValid()) {
            $r->setStatus(PBValidateOutput\PBStatus::VALIDATION_ERROR);
            $msg = '';
            foreach ($validator->getErrors() as $error) {
                $msg .= sprintf("[%s] %s\n", $error['property'], $error['message']);
            }
            $r->setErrorMessage($msg);
            return $r;
        }

        $r->setStatus(PBValidateOutput\PBStatus::OK);
        return $r;
    }
}