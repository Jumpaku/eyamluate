<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eval/evaluator.proto

namespace Jumpaku\Eyamluate\Eval_;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>eval.EvaluateInput</code>
 */
class PBEvaluateInput extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string source = 1 [json_name = "source"];</code>
     */
    protected $source = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $source
     * }
     */
    public function __construct($data = NULL) {
        \Jumpaku\Eyamluate\Metadata\Eval_\Evaluator::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string source = 1 [json_name = "source"];</code>
     * @return string
     */
    public function getSource()
    {
        return $this->source;
    }

    /**
     * Generated from protobuf field <code>string source = 1 [json_name = "source"];</code>
     * @param string $var
     * @return $this
     */
    public function setSource($var)
    {
        GPBUtil::checkString($var, True);
        $this->source = $var;

        return $this;
    }

}

