<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eval/validator.proto

namespace Jumpaku\Eyamluate\Metadata\Eval_;

class Validator
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            "\x0A\xA5\x03\x0A\x14eval/validator.proto\x12\x04eval\"'\x0A\x0DValidateInput\x12\x16\x0A\x06source\x18\x01 \x01(\x09R\x06source\"\xA2\x01\x0A\x0EValidateOutput\x123\x0A\x06status\x18\x01 \x01(\x0E2\x1B.eval.ValidateOutput.StatusR\x06status\x12#\x0A\x0Derror_message\x18\x02 \x01(\x09R\x0CerrorMessage\"6\x0A\x06Status\x12\x06\x0A\x02OK\x10\x00\x12\x0E\x0A\x0AYAML_ERROR\x10\x01\x12\x14\x0A\x10VALIDATION_ERROR\x10\x022D\x0A\x09Validator\x127\x0A\x08Validate\x12\x13.eval.ValidateInput\x1A\x14.eval.ValidateOutput\"\x00BkZ'github.com/Jumpaku/eyamlate/golang/eval\xC2\x02\x02PB\xCA\x02\x17Jumpaku\\Eyamluate\\Eval_\xE2\x02 Jumpaku\\Eyamluate\\Metadata\\Eval_b\x06proto3"
        , true);

        static::$is_initialized = true;
    }
}

