<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: yaml/encoder.proto

namespace Jumpaku\Eyamluate\Metadata\Yaml;

class Encoder
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \Jumpaku\Eyamluate\Metadata\Yaml\Value::initOnce();
        $pool->internalAddGeneratedFile(
            "\x0A\xE9\x03\x0A\x12yaml/encoder.proto\x12\x04yaml\"t\x0A\x0BEncodeInput\x12*\x0A\x06format\x18\x01 \x01(\x0E2\x12.yaml.EncodeFormatR\x06format\x12\x16\x0A\x06pretty\x18\x02 \x01(\x08R\x06pretty\x12!\x0A\x05value\x18\x03 \x01(\x0B2\x0B.yaml.ValueR\x05value\"f\x0A\x0CEncodeOutput\x12\x19\x0A\x08is_error\x18\x01 \x01(\x08R\x07isError\x12#\x0A\x0Derror_message\x18\x02 \x01(\x09R\x0CerrorMessage\x12\x16\x0A\x06result\x18\x03 \x01(\x09R\x06result*>\x0A\x0CEncodeFormat\x12\x16\x0A\x12ENCODE_FORMAT_YAML\x10\x00\x12\x16\x0A\x12ENCODE_FORMAT_JSON\x10\x012<\x0A\x07Encoder\x121\x0A\x06Encode\x12\x11.yaml.EncodeInput\x1A\x12.yaml.EncodeOutput\"\x00BiZ'github.com/Jumpaku/eyamlate/golang/yaml\xC2\x02\x02PB\xCA\x02\x16Jumpaku\\Eyamluate\\Yaml\xE2\x02\x1FJumpaku\\Eyamluate\\Metadata\\Yamlb\x06proto3"
        , true);

        static::$is_initialized = true;
    }
}

