<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eval/evaluator.proto

namespace Jumpaku\Eyamluate\Eval_;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>eval.Path</code>
 */
class PBPath extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .eval.Path.Pos pos = 1 [json_name = "pos"];</code>
     */
    private $pos;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\Jumpaku\Eyamluate\Eval_\PBPath\PBPos>|\Google\Protobuf\Internal\RepeatedField $pos
     * }
     */
    public function __construct($data = NULL) {
        \Jumpaku\Eyamluate\Metadata\Eval_\Evaluator::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .eval.Path.Pos pos = 1 [json_name = "pos"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPos()
    {
        return $this->pos;
    }

    /**
     * Generated from protobuf field <code>repeated .eval.Path.Pos pos = 1 [json_name = "pos"];</code>
     * @param array<\Jumpaku\Eyamluate\Eval_\PBPath\PBPos>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPos($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Jumpaku\Eyamluate\Eval_\PBPath\PBPos::class);
        $this->pos = $arr;

        return $this;
    }

}

