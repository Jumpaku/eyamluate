<?php

namespace Jumpaku\Eyamluate\Eval_;

use Exception;
use Jumpaku\Eyamluate\Yaml\Decoder;
use Jumpaku\Eyamluate\Yaml\PBDecodeInput;
use Jumpaku\Eyamluate\Yaml\PBType;
use Jumpaku\Eyamluate\Yaml\PBValue;

class BaseEvaluator implements Evaluator
{
    public function evaluate(PBEvaluateInput $input): PBEvaluateOutput
    {
        // Decode input
        $decodeOutput = (new Decoder())->decode((new PBDecodeInput())->setYaml($input->getSource()));
        if ($decodeOutput->getIsError()) {
            return (new PBEvaluateOutput())
                ->setStatus(PBEvaluateOutput\PBStatus::DECODE_ERROR)
                ->setErrorMessage($decodeOutput->getErrorMessage());
        }

        // Validate input
        $validateOutput = (new Validator())->validate((new PBValidateInput())->setSource($input->getSource()));
        if ($validateOutput->getStatus() != PBValidateOutput\PBStatus::OK) {
            return (new PBEvaluateOutput())
                ->setStatus(PBEvaluateOutput\PBStatus::VALIDATE_ERROR)
                ->setErrorMessage($validateOutput->getErrorMessage());
        }

        // Evaluate input
        $evaluateOutput = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(new PBPath())
            ->setDefs(FunDefList::empty())
            ->setExpr($decodeOutput->getValue()));
        if ($evaluateOutput->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return (new PBEvaluateOutput())
                ->setStatus(PBEvaluateOutput\PBStatus::EXPR_ERROR)
                ->setErrorMessage($evaluateOutput->getErrorMessage())
                ->setExprStatus($evaluateOutput->getStatus())
                ->setExprErrorPath($evaluateOutput->getErrorPath());
        }
        return (new PBEvaluateOutput())->setValue($evaluateOutput->getValue());
    }

    public
    function evaluateExpr(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $expr = $input->getExpr();
        switch ($expr->getType()) {
            case PBType::TYPE_BOOL:
            case PBType::TYPE_NUM:
            case PBType::TYPE_STR:
                return $this->evaluateScalar($input);
            case PBType::TYPE_OBJ:
                if ($this->hasKey($expr, "eval")) {
                    return $this->evaluateEval($input);
                } elseif ($this->hasKey($expr, "obj")) {
                    return $this->evaluateObj($input);
                } elseif ($this->hasKey($expr, "arr")) {
                    return $this->evaluateArr($input);
                } elseif ($this->hasKey($expr, "json")) {
                    return $this->evaluateJson($input);
                } elseif ($this->hasKey($expr, "for")) {
                    return $this->evaluateRangeIter($input);
                } elseif ($this->hasKey($expr, "get")) {
                    return $this->evaluateGetElem($input);
                } elseif ($this->hasKey($expr, "ref")) {
                    return $this->evaluateFunCall($input);
                } elseif ($this->hasKey($expr, "cases")) {
                    return $this->evaluateCases($input);
                } elseif ($this->hasKey($expr, "len")
                    || $this->hasKey($expr, "not")
                    || $this->hasKey($expr, "flat")
                    || $this->hasKey($expr, "floor")
                    || $this->hasKey($expr, "ceil")
                    || $this->hasKey($expr, "abort")) {
                    return $this->evaluateOpUnary($input);
                } elseif (
                    $this->hasKey($expr, "sub")
                    || $this->hasKey($expr, "div")
                    || $this->hasKey($expr, "eq")
                    || $this->hasKey($expr, "neq")
                    || $this->hasKey($expr, "lt")
                    || $this->hasKey($expr, "lte")
                    || $this->hasKey($expr, "gt")
                    || $this->hasKey($expr, "gte")) {
                    return $this->evaluateOpBinary($input);
                } elseif (
                    $this->hasKey($expr, "add")
                    || $this->hasKey($expr, "mul")
                    || $this->hasKey($expr, "and")
                    || $this->hasKey($expr, "or")
                    || $this->hasKey($expr, "cat")
                    || $this->hasKey($expr, "min")
                    || $this->hasKey($expr, "max")
                    || $this->hasKey($expr, "merge")) {
                    return $this->evaluateOpVariadic($input);
                }
        }
        return $this->errorUnsupportedExpr($input->getPath(), $input->getExpr());
    }

    public
    function evaluateEval(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $path = $input->getPath();
        $st = $input->getDefs();
        $expr = $input->getExpr();
        if ($this->hasKey($expr, "where")) {
            $pathWhere = Path::append($path, "where");
            $whereExpr = $expr->getObj()["where"];
            foreach ($whereExpr->getArr() as $pos => $whereElem) {
                $funDef = (new PBFunDef())
                    ->setDef($whereElem->getObj()["def"]->getStr())
                    ->setValue($whereElem->getObj()["value"])
                    ->setPath(Path::append($pathWhere, $pos));
                if ($this->hasKey($whereElem, "with")) {
                    $withArr = [];
                    foreach ($whereElem->getObj()["with"]->getArr() as $withElem) {
                        $withArr[] = $withElem->getStr();
                    }
                    $funDef->setWith($withArr);
                }
                $st = FunDefList::register($st, $funDef);
            }
        }
        /** @var PBValue $evalExpr */
        $evalExpr = $expr->getObj()["eval"];
        return $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, "eval"))
            ->setDefs($st)
            ->setExpr($evalExpr));
    }

    public
    function evaluateScalar(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        return (new PBEvaluateExprOutput())->setValue($input->getExpr());
    }

    public
    function evaluateObj(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        /** @var PBValue $obj */
        $obj = $input->getExpr()->getObj()["obj"];
        $path = Path::append($input->getPath(), "obj");
        $v = [];
        foreach ($obj->getObj() as $pos => $val) {
            $expr = $this->evaluateExpr((new PBEvaluateExprInput())
                ->setPath(Path::append($path, $pos))
                ->setDefs($input->getDefs())
                ->setExpr($val));
            if ($expr->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                return $expr;
            }
            $v[$pos] = $expr->getValue();
        }
        return (new PBEvaluateExprOutput())
            ->setValue((new PBValue())
                ->setType(PBType::TYPE_OBJ)
                ->setObj($v));
    }

    public
    function evaluateArr(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $arr = $input->getExpr()->getObj()["arr"];
        $path = Path::append($input->getPath(), "arr");
        $v = [];
        foreach ($arr->getArr() as $pos => $val) {
            $expr = $this->evaluateExpr((new PBEvaluateExprInput())
                ->setPath(Path::append($path, $pos))
                ->setDefs($input->getDefs())
                ->setExpr($val));
            if ($expr->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                return $expr;
            }
            $v[] = $expr->getValue();
        }
        return (new PBEvaluateExprOutput())
            ->setValue((new PBValue())
                ->setType(PBType::TYPE_ARR)
                ->setArr($v));
    }

    public
    function evaluateJson(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        /** @var PBValue $json */
        $json = $input->getExpr()->getObj()["json"];
        return (new PBEvaluateExprOutput())->setValue($json);
    }

    public
    function evaluateRangeIter(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $path = $input->getPath();
        /** @var PBValue $for_ */
        $for_ = $input->getExpr()->getObj()["for"];
        $forPos = $for_->getArr()[0]->getStr();
        $forVal = $for_->getArr()[1]->getStr();
        /** @var PBValue $in */
        $in = $input->getExpr()->getObj()["in"];
        $inValue = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, "in"))
            ->setDefs($input->getDefs())
            ->setExpr($in));
        if ($inValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return $inValue;
        }
        switch ($inValue->getValue()->getType()) {
            case PBType::TYPE_STR:
                $v = [];
                $inStr = $inValue->getValue()->getStr();
                $n = mb_strlen($inStr);
                for ($pos = 0; $pos < $n; $pos++) {
                    $st = $input->getDefs();
                    $st = FunDefList::register($st, (new PBFunDef())
                        ->setDef($forPos)
                        ->setPath(Path::append(Path::append($path, "for"), 0))
                        ->setValue((new PBValue())
                            ->setType(PBType::TYPE_NUM)
                            ->setNum($pos)));
                    $st = FunDefList::register($st, (new PBFunDef())
                        ->setDef($forVal)
                        ->setPath(Path::append(Path::append($path, "for"), 1))
                        ->setValue((new PBValue())
                            ->setType(PBType::TYPE_STR)
                            ->setStr(mb_substr($inStr, $pos, 1))));
                    if ($this->hasKey($input->getExpr(), "if")) {
                        /** @var PBValue $if_ */
                        $if_ = $input->getExpr()->getObj()["if"];
                        $ifValue = $this->evaluateExpr((new PBEvaluateExprInput())
                            ->setPath(Path::append($path, "if"))
                            ->setDefs($st)
                            ->setExpr($if_));
                        if ($ifValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                            return $ifValue;
                        }
                        if ($ifValue->getValue()->getType() != PBType::TYPE_BOOL) {
                            return $this->errorUnexpectedType(Path::append($path, "if"), [PBType::TYPE_BOOL], $ifValue->getValue()->getType());
                        }
                        if (!$ifValue->getValue()->getBool()) {
                            continue;
                        }
                    }
                    /** @var PBValue $do */
                    $do = $input->getExpr()->getObj()["do"];
                    $doValue = $this->evaluateExpr((new PBEvaluateExprInput())
                        ->setPath(Path::append($path, "do"))
                        ->setDefs($st)
                        ->setExpr($do));
                    if ($doValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                        return $doValue;
                    }
                    $v[] = $doValue->getValue();
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_ARR)->setArr($v));
            case PBType::TYPE_ARR:
                $v = [];
                $inArr = $inValue->getValue()->getArr();
                foreach ($inArr as $pos => $val) {
                    $st = $input->getDefs();
                    $st = FunDefList::register($st, (new PBFunDef())
                        ->setDef($forPos)
                        ->setPath(Path::append(Path::append($path, "for"), 0))
                        ->setValue((new PBValue())
                            ->setType(PBType::TYPE_NUM)
                            ->setNum($pos)));
                    $st = FunDefList::register($st, (new PBFunDef())
                        ->setDef($forVal)
                        ->setPath(Path::append(Path::append($path, "for"), 1))
                        ->setValue($val));
                    if ($this->hasKey($input->getExpr(), "if")) {
                        /** @var PBValue $if_ */
                        $if_ = $input->getExpr()->getObj()["if"];
                        $ifValue = $this->evaluateExpr((new PBEvaluateExprInput())
                            ->setPath(Path::append($path, "if"))
                            ->setDefs($st)
                            ->setExpr($if_));
                        if ($ifValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                            return $ifValue;
                        }
                        if ($ifValue->getValue()->getType() != PBType::TYPE_BOOL) {
                            return $this->errorUnexpectedType(Path::append($path, "if"), [PBType::TYPE_BOOL], $ifValue->getValue()->getType());
                        }
                        if (!$ifValue->getValue()->getBool()) {
                            continue;
                        }
                    }
                    /** @var PBValue $do */
                    $do = $input->getExpr()->getObj()["do"];
                    $doValue = $this->evaluateExpr((new PBEvaluateExprInput())
                        ->setPath(Path::append($path, "do"))
                        ->setDefs($st)
                        ->setExpr($do));
                    if ($doValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                        return $doValue;
                    }
                    $v[] = $doValue->getValue();
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_ARR)->setArr($v));
            case PBType::TYPE_OBJ:
                $v = [];
                $inObj = $inValue->getValue()->getObj();
                foreach ($inObj as $pos => $val) {
                    $st = $input->getDefs();
                    $st = FunDefList::register($st, (new PBFunDef())
                        ->setDef($forPos)
                        ->setPath(Path::append(Path::append($path, "for"), 0))
                        ->setValue((new PBValue())
                            ->setType(PBType::TYPE_STR)
                            ->setStr($pos)));
                    $st = FunDefList::register($st, (new PBFunDef())
                        ->setDef($forVal)
                        ->setPath(Path::append(Path::append($path, "for"), 1))
                        ->setValue($val));
                    if ($this->hasKey($input->getExpr(), "if")) {
                        /** @var PBValue $if_ */
                        $if_ = $input->getExpr()->getObj()["if"];
                        $ifValue = $this->evaluateExpr((new PBEvaluateExprInput())
                            ->setPath(Path::append($path, "if"))
                            ->setDefs($st)
                            ->setExpr($if_));
                        if ($ifValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                            return $ifValue;
                        }
                        if ($ifValue->getValue()->getType() != PBType::TYPE_BOOL) {
                            return $this->errorUnexpectedType(Path::append($path, "if"), [PBType::TYPE_BOOL], $ifValue->getValue()->getType());
                        }
                        if (!$ifValue->getValue()->getBool()) {
                            continue;
                        }
                    }
                    /** @var PBValue $do */
                    $do = $input->getExpr()->getObj()["do"];
                    $doValue = $this->evaluateExpr((new PBEvaluateExprInput())
                        ->setPath(Path::append($path, "do"))
                        ->setDefs($st)
                        ->setExpr($do));
                    if ($doValue->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                        return $doValue;
                    }
                    $v[$pos] = $doValue->getValue();
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_OBJ)->setObj($v));
            default:
                return $this->errorUnexpectedType(Path::append($path, "in"), [PBType::TYPE_STR, PBType::TYPE_ARR, PBType::TYPE_OBJ], $inValue->getValue()->getType());
        }
    }

    public
    function evaluateGetElem(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $path = $input->getPath();
        /** @var PBValue $getExpr */
        $getExpr = $input->getExpr()->getObj()["get"];
        $get = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, "get"))
            ->setDefs($input->getDefs())
            ->setExpr($getExpr));
        if ($get->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return $get;
        }
        $getValue = $get->getValue();

        /** @var PBValue $fromExpr */
        $fromExpr = $input->getExpr()->getObj()["from"];
        $from = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, "from"))
            ->setDefs($input->getDefs())
            ->setExpr($fromExpr));
        if ($from->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return $from;
        }
        $fromValue = $from->getValue();
        switch ($fromValue->getType()) {
            case PBType::TYPE_STR:
                if ($getValue->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($path, "get"), [PBType::TYPE_NUM], $getValue->getType());
                }
                if (!$this->isInteger($getValue)) {
                    return $this->errorArithmeticError(Path::append($path, "get"), sprintf("index %f is not an integer", $getValue->getNum()));
                }
                $pos = (int)$getValue->getNum();
                $n = mb_strlen($fromValue->getStr());
                if ($pos < 0 || $pos >= $n) {
                    return $this->errorIndexOutOfBounds(Path::append($path, "get"), 0, $n, $pos);
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_STR)
                        ->setStr(mb_substr($fromValue->getStr(), $pos, 1)));
            case PBType::TYPE_ARR:
                if ($getValue->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($path, "get"), [PBType::TYPE_NUM], $getValue->getType());
                }
                if (!$this->isInteger($getValue)) {
                    return $this->errorArithmeticError(Path::append($path, "get"), sprintf("index %f is not an integer", $getValue->getNum()));
                }
                $pos = (int)$getValue->getNum();
                $n = count($fromValue->getArr());
                if ($pos < 0 || $pos >= $n) {
                    return $this->errorIndexOutOfBounds(Path::append($path, "get"), 0, $n, $pos);
                }
                /** @var PBValue $v */
                $v = $fromValue->getArr()[$pos];
                return (new PBEvaluateExprOutput())->setValue($v);
            case PBType::TYPE_OBJ:
                if ($getValue->getType() != PBType::TYPE_STR) {
                    return $this->errorUnexpectedType(Path::append($path, "get"), [PBType::TYPE_STR], $getValue->getType());
                }
                $pos = $getValue->getStr();
                if (!$this->hasKey($fromValue, $pos)) {
                    return $this->errorKeyNotFound(Path::append($path, "get"), $pos, $this->getKeys($fromValue));
                }
                /** @var PBValue $v */
                $v = $fromValue->getObj()[$pos];
                return (new PBEvaluateExprOutput())->setValue($v);
            default:
                return $this->errorUnexpectedType(Path::append($path, "from"), [PBType::TYPE_STR, PBType::TYPE_ARR, PBType::TYPE_OBJ], $fromValue->getType());
        }
    }

    public
    function evaluateFunCall(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $path = $input->getPath();
        $funCall = $input->getExpr();
        $ref = $funCall->getObj()["ref"];
        $funDef = FunDefList::find($input->getDefs(), $ref->getStr());
        if ($funDef == null) {
            return $this->errorReferenceNotFound(Path::append($path, "ref"), $ref->getStr());
        }
        $st = $funDef;
        $pathWith = Path::append($path, "with");
        /** @var PBValue $with */
        $with = $funCall->getObj()["with"];
        foreach ($funDef->getDef()->getWith() as $argName) {
            if (!$this->hasKey($with, $argName)) {
                return $this->errorKeyNotFound($pathWith, $argName, $this->getKeys($with));
            }
            /** @var PBValue $argExpr */
            $argExpr = $with->getObj()[$argName];
            $arg = $this->evaluateExpr((new PBEvaluateExprInput())
                ->setPath(Path::append($pathWith, $argName))
                ->setDefs($input->getDefs())
                ->setExpr($argExpr));
            if ($arg->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                return $arg;
            }
            $jsonExpr = (new PBValue())
                ->setType(PBType::TYPE_OBJ)
                ->setObj(["json" => $arg->getValue()]);
            $st = FunDefList::register($st, (new PBFunDef)
                ->setPath(Path::append($pathWith, $argName))
                ->setDef($argName)
                ->setValue($jsonExpr));
        }
        return $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, "ref"))
            ->setDefs($st)
            ->setExpr($funDef->getDef()->getValue()));
    }

    public
    function evaluateCases(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $path = Path::append($input->getPath(), "cases");
        /** @var PBValue $cases */
        $cases = $input->getExpr()->getObj()["cases"];
        foreach ($cases->getArr() as $pos => $c) {
            $pathPos = Path::append($path, $pos);
            if ($this->hasKey($c, "when")) {
                $when = $this->evaluateExpr((new PBEvaluateExprInput())
                    ->setPath(Path::append($pathPos, "when"))
                    ->setDefs($input->getDefs())
                    ->setExpr($c->getObj()["when"]));
                if ($when->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                    return $when;
                }
                if ($when->getValue()->getType() != PBType::TYPE_BOOL) {
                    return $this->errorUnexpectedType($pathPos, [PBType::TYPE_BOOL], $when->getValue()->getType());
                }
                if ($when->getValue()->getBool()) {
                    return $this->evaluateExpr((new PBEvaluateExprInput())
                        ->setPath(Path::append($pathPos, "then"))
                        ->setDefs($input->getDefs())
                        ->setExpr($c->getObj()["then"]));
                }
            } elseif ($this->hasKey($c, "otherwise")) {
                return $this->evaluateExpr((new PBEvaluateExprInput())
                    ->setPath(Path::append($pathPos, "otherwise"))
                    ->setDefs($input->getDefs())
                    ->setExpr($c->getObj()["otherwise"]));
            }
        }
        return $this->errorCasesNotExhaustive($path);
    }

    public
    function evaluateOpUnary(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $expr = $input->getExpr();
        $operator = array_keys(iterator_to_array($expr->getObj()))[0];
        $path = Path::append($input->getPath(), $operator);
        /** @var PBValue $o */
        $o = $expr->getObj()[$operator];
        $expr = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, $operator))
            ->setDefs($input->getDefs())
            ->setExpr($o));
        if ($expr->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return $expr;
        }
        /** @var PBValue $operand */
        $operand = $expr->getValue();
        switch ($operator) {
            case OpUnary::keyName(PBOpUnary\PBOperator::LEN):
                switch ($operand->getType()) {
                    case PBType::TYPE_STR:
                        return (new PBEvaluateExprOutput())
                            ->setValue((new PBValue())
                                ->setType(PBType::TYPE_NUM)
                                ->setNum(mb_strlen($operand->getStr())));
                    case PBType::TYPE_ARR:
                        return (new PBEvaluateExprOutput())
                            ->setValue((new PBValue())
                                ->setType(PBType::TYPE_NUM)
                                ->setNum(count($operand->getArr())));
                    case PBType::TYPE_OBJ:
                        return (new PBEvaluateExprOutput())
                            ->setValue((new PBValue())
                                ->setType(PBType::TYPE_NUM)
                                ->setNum(count($operand->getObj())));
                    default:
                        return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_STR, PBType::TYPE_ARR, PBType::TYPE_OBJ], $operand->getType());
                }
            case OpUnary::keyName(PBOpUnary\PBOperator::NOT):
                if ($operand->getType() != PBType::TYPE_BOOL) {
                    return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_BOOL], $operand->getType());
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_BOOL)
                        ->setBool(!$operand->getBool()));
            case OpUnary::keyName(PBOpUnary\PBOperator::FLAT):
                if ($operand->getType() != PBType::TYPE_ARR) {
                    return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_ARR], $operand->getType());
                }
                $v = [];
                foreach ($operand->getArr() as $elem) {
                    if ($elem->getType() != PBType::TYPE_ARR) {
                        return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_ARR], $elem->getType());
                    }
                    $v = array_merge($v, iterator_to_array($elem->getArr()));
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_ARR)
                        ->setArr($v));
            case OpUnary::keyName(PBOpUnary\PBOperator::FLOOR):
                if ($operand->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_NUM], $operand->getType());
                }
                $v = (new PBValue())
                    ->setType(PBType::TYPE_NUM)
                    ->setNum(floor($operand->getNum()));
                if (!$this->isFiniteNumber($v)) {
                    return $this->errorArithmeticError(Path::append($input->getPath(), $operator), sprintf("floor(%f) is not a finite number", $operand->getNum()));
                }
                return (new PBEvaluateExprOutput())->setValue($v);
            case OpUnary::keyName(PBOpUnary\PBOperator::CEIL):
                if ($operand->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_NUM], $operand->getType());
                }
                $v = (new PBValue())
                    ->setType(PBType::TYPE_NUM)
                    ->setNum(ceil($operand->getNum()));
                if (!$this->isFiniteNumber($v)) {
                    return $this->errorArithmeticError(Path::append($input->getPath(), $operator), sprintf("ceil(%f) is not a finite number", $operand->getNum()));
                }
                return (new PBEvaluateExprOutput())->setValue($v);
            case OpUnary::keyName(PBOpUnary\PBOperator::ABORT):
                if ($operand->getType() != PBType::TYPE_STR) {
                    return $this->errorUnexpectedType(Path::append($input->getPath(), $operator), [PBType::TYPE_STR], $operand->getType());
                }
                return (new PBEvaluateExprOutput())
                    ->setStatus(PBEvaluateExprOutput\PBStatus::ABORTED)
                    ->setErrorMessage($operand->getStr());
            default:
                return $this->errorUnsupportedOperation($input->getPath(), $operator);
        }
    }

    public
    function evaluateOpBinary(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $expr = $input->getExpr();
        $operator = array_keys(iterator_to_array($expr->getObj()))[0];
        $path = Path::append($input->getPath(), $operator);
        /** @var PBValue $os */
        $os = $expr->getObj()[$operator];
        $ol = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, 0))
            ->setDefs($input->getDefs())
            ->setExpr($os->getArr()[0]));
        if ($ol->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return $ol;
        }
        $operandL = $ol->getValue();
        $or = $this->evaluateExpr((new PBEvaluateExprInput())
            ->setPath(Path::append($path, 1))
            ->setDefs($input->getDefs())
            ->setExpr($os->getArr()[1]));
        if ($or->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
            return $or;
        }
        $operandR = $or->getValue();
        switch ($operator) {
            case OpBinary::keyName(PBOpBinary\PBOperator::SUB):
                if ($operandL->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($path, $operator), [PBType::TYPE_NUM], $operandL->getType());
                }
                if ($operandR->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($path, $operator), [PBType::TYPE_NUM], $operandR->getType());
                }
                $v = (new PBValue())
                    ->setType(PBType::TYPE_NUM)
                    ->setNum($operandL->getNum() - $operandR->getNum());
                if (!$this->isFiniteNumber($v)) {
                    return $this->errorArithmeticError(Path::append($path, $operator), sprintf("%f-%f is not a finite number", $operandL->getNum(), $operandR->getNum()));
                }
                return (new PBEvaluateExprOutput())->setValue($v);
            case OpBinary::keyName(PBOpBinary\PBOperator::DIV):
                if ($operandL->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($path, $operator), [PBType::TYPE_NUM], $operandL->getType());
                }
                if ($operandR->getType() != PBType::TYPE_NUM) {
                    return $this->errorUnexpectedType(Path::append($path, $operator), [PBType::TYPE_NUM], $operandR->getType());
                }
                $v = (new PBValue())
                    ->setType(PBType::TYPE_NUM)
                    ->setNum($this->compute_div($operandL->getNum(), $operandR->getNum()));
                if (!$this->isFiniteNumber($v)) {
                    return $this->errorArithmeticError(Path::append($path, $operator), sprintf("%f/%f is not a finite number", $operandL->getNum(), $operandR->getNum()));
                }
                return (new PBEvaluateExprOutput())->setValue($v);
            case OpBinary::keyName(PBOpBinary\PBOperator::EQ):
                return $this->equal($path, $operandL, $operandR);
            case OpBinary::keyName(PBOpBinary\PBOperator::NEQ):
                $eq = $this->equal($path, $operandL, $operandR);
                if ($eq->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                    return $eq;
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_BOOL)
                        ->setBool(!$eq->getValue()->getBool()));
            case OpBinary::keyName(PBOpBinary\PBOperator::LT):
                $cmp = $this->compare($path, $operandL, $operandR);
                if ($cmp->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                    return $cmp;
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_BOOL)
                        ->setBool($cmp->getValue()->getNum() < 0));
            case OpBinary::keyName(PBOpBinary\PBOperator::LTE):
                $cmp = $this->compare($path, $operandL, $operandR);
                if ($cmp->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                    return $cmp;
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_BOOL)
                        ->setBool($cmp->getValue()->getNum() <= 0));
            case OpBinary::keyName(PBOpBinary\PBOperator::GT):
                $cmp = $this->compare($path, $operandL, $operandR);
                if ($cmp->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                    return $cmp;
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_BOOL)
                        ->setBool($cmp->getValue()->getNum() > 0));
            case OpBinary::keyName(PBOpBinary\PBOperator::GTE):
                $cmp = $this->compare($path, $operandL, $operandR);
                if ($cmp->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                    return $cmp;
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())
                        ->setType(PBType::TYPE_BOOL)
                        ->setBool($cmp->getValue()->getNum() >= 0));
            default:
                return $this->errorUnsupportedOperation($path, $operator);
        }
    }

    public
    function evaluateOpVariadic(PBEvaluateExprInput $input): PBEvaluateExprOutput
    {
        $expr = $input->getExpr();
        $operator = array_keys(iterator_to_array($expr->getObj()))[0];
        $path = Path::append($input->getPath(), $operator);
        /** @var PBValue $os */
        $os = $expr->getObj()[$operator];
        /** @var PBValue[] $operands */
        $operands = [];
        foreach ($os->getArr() as $pos => $o) {
            $operand = $this->evaluateExpr((new PBEvaluateExprInput())
                ->setPath(Path::append($path, $pos))
                ->setDefs($input->getDefs())
                ->setExpr($o));
            if ($operand->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                return $operand;
            }
            $operands[] = $operand->getValue();
        }
        switch ($operator) {
            case OpVariadic::keyName(PBOpVariadic\PBOperator::ADD):
                $v = 0.0;
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_NUM) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_NUM], $operand->getType());
                    }
                    $v += $operand->getNum();
                }
                if (!$this->isFiniteNumber((new PBValue())->setType(PBType::TYPE_NUM)->setNum($v))) {
                    $v = join(",", array_map(function ($o) {
                        return $o->getNum();
                    }, $operands));
                    return $this->errorArithmeticError($path, sprintf("add(%s) is not a finite number", $v));
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_NUM)->setNum($v));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::MUL):
                $v = 1.0;
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_NUM) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_NUM], $operand->getType());
                    }
                    $v *= $operand->getNum();
                }
                if (!$this->isFiniteNumber((new PBValue())->setType(PBType::TYPE_NUM)->setNum($v))) {
                    $v = join(",", array_map(function ($o) {
                        return $o->getNum();
                    }, $operands));
                    return $this->errorArithmeticError($path, sprintf("mul(%s) is not a finite number", $v));
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_NUM)->setNum($v));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::PBAND):
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_BOOL) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_BOOL], $operand->getType());
                    }
                    if (!$operand->getBool()) {
                        return (new PBEvaluateExprOutput())
                            ->setValue((new PBValue())->setType(PBType::TYPE_BOOL)->setBool(false));
                    }
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_BOOL)->setBool(true));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::PBOR):
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_BOOL) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_BOOL], $operand->getType());
                    }
                    if ($operand->getBool()) {
                        return (new PBEvaluateExprOutput())
                            ->setValue((new PBValue())->setType(PBType::TYPE_BOOL)->setBool(true));
                    }
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_BOOL)->setBool(false));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::CAT):
                $v = [];
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_STR) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_STR], $operand->getType());
                    }
                    $v[] = $operand->getStr();
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_STR)->setStr(join("", $v)));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::MIN):
                $v = INF;
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_NUM) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_NUM], $operand->getType());
                    }
                    $v = min($v, $operand->getNum());
                }
                return (new PBEvaluateExprOutput())->setValue((new PBValue())->setType(PBType::TYPE_NUM)->setNum($v));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::MAX):
                $v = -INF;
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_NUM) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_NUM], $operand->getType());
                    }
                    $v = max($v, $operand->getNum());
                }
                return (new PBEvaluateExprOutput())->setValue((new PBValue())->setType(PBType::TYPE_NUM)->setNum($v));
            case OpVariadic::keyName(PBOpVariadic\PBOperator::MERGE):
                $v = [];
                foreach ($operands as $operand) {
                    if ($operand->getType() != PBType::TYPE_OBJ) {
                        return $this->errorUnexpectedType($path, [PBType::TYPE_OBJ], $operand->getType());
                    }
                    $v = array_merge($v, iterator_to_array($operand->getObj()));
                }
                return (new PBEvaluateExprOutput())
                    ->setValue((new PBValue())->setType(PBType::TYPE_OBJ)->setObj($v));
            default:
                return $this->errorUnsupportedOperation($path, $operator);
        }
    }


    private
    function equal(PBPath $path, PBValue $l, PBValue $r): PBEvaluateExprOutput
    {
        $falseValue = (new PBEvaluateExprOutput)
            ->setValue((new PBValue)
                ->setType(PBType::TYPE_BOOL)
                ->setBool(false));
        $trueValue = (new PBEvaluateExprOutput)
            ->setValue((new PBValue)
                ->setType(PBType::TYPE_BOOL)
                ->setBool(true));
        if ($l->getType() != $r->getType()) {
            return $falseValue;
        }
        switch ($l->getType()) {
            case PBType::TYPE_NUM:
                if ($l->getNum() != $r->getNum()) {
                    return $falseValue;
                }
                return $trueValue;
            case PBType::TYPE_BOOL:
                if ($l->getBool() != $r->getBool()) {
                    return $falseValue;
                }
                return $trueValue;
            case PBType::TYPE_STR:
                if ($l->getStr() != $r->getStr()) {
                    return $falseValue;
                }
                return $trueValue;
            case PBType::TYPE_ARR:
                $lArr = $l->getArr();
                $rArr = $r->getArr();
                if ($lArr->count() != $rArr->count()) {
                    return $falseValue;
                }
                foreach ($lArr as $index => $lItem) {
                    $rItem = $rArr->offsetGet($index);
                    $eq = $this->equal($path, $lItem, $rItem);
                    if ($eq->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                        return $eq;
                    }
                    if (!$eq->getValue()->getBool()) {
                        return $falseValue;
                    }
                }
                return $trueValue;
            case PBType::TYPE_OBJ:
                $lObj = $l->getObj();
                $rObj = $r->getObj();
                $lk = array_keys(iterator_to_array($lObj));
                $rk = array_keys(iterator_to_array($rObj));
                sort($lk);
                sort($rk);
                if ($lk != $rk) {
                    return $falseValue;
                }
                foreach ($lObj as $k => $lItem) {
                    /** @var PBValue $rItem */
                    $rItem = $rObj[$k];
                    $eq = $this->equal($path, $lItem, $rItem);
                    if ($eq->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                        return $eq;
                    }
                    if (!$eq->getValue()->getBool()) {
                        return $falseValue;
                    }
                }
                return $trueValue;
            default:
                return $this->errorUnexpectedType($path, [PBType::TYPE_NUM, PBType::TYPE_BOOL, PBType::TYPE_STR, PBType::TYPE_ARR], $l->getType());
        }
    }

    private
    function compare(PBPath $path, PBValue $l, PBValue $r): PBEvaluateExprOutput
    {
        $ltValue = (new PBEvaluateExprOutput)
            ->setValue((new PBValue)
                ->setType(PBType::TYPE_NUM)
                ->setNum(-1));
        $gtValue = (new PBEvaluateExprOutput)
            ->setValue((new PBValue)
                ->setType(PBType::TYPE_NUM)
                ->setNum(1));
        $eqValue = (new PBEvaluateExprOutput)
            ->setValue((new PBValue)
                ->setType(PBType::TYPE_NUM)
                ->setNum(0));
        if ($l->getType() != $r->getType()) {
            return $this->errorUnexpectedType($path, [PBType::TYPE_NUM, PBType::TYPE_BOOL, PBType::TYPE_STR, PBType::TYPE_ARR], $r->getType());
        }
        switch ($l->getType()) {
            case PBType::TYPE_NUM:
                if ($l->getNum() < $r->getNum()) {
                    return $ltValue;
                }
                if ($l->getNum() > $r->getNum()) {
                    return $gtValue;
                }
                return $eqValue;
            case PBType::TYPE_BOOL:
                if (!$l->getBool() && $r->getBool()) {
                    return $ltValue;
                }
                if ($l->getBool() && !$r->getBool()) {
                    return $gtValue;
                }
                return $eqValue;
            case PBType::TYPE_STR:
                if ($l->getStr() < $r->getStr()) {
                    return $ltValue;
                }
                if ($l->getStr() > $r->getStr()) {
                    return $gtValue;
                }
                return $eqValue;
            case PBType::TYPE_ARR:
                $lArr = $l->getArr();
                $rArr = $r->getArr();
                $n = min($lArr->count(), $rArr->count());
                for ($i = 0; $i < $n; $i++) {
                    $lItem = $lArr[$i];
                    $rItem = $rArr[$i];
                    $cmp = $this->compare($path, $lItem, $rItem);
                    if ($cmp->getStatus() != PBEvaluateExprOutput\PBStatus::OK) {
                        return $cmp;
                    }
                    if ($cmp->getValue()->getNum() != 0) {
                        return $cmp;
                    }
                }
                if (count($l->getArr()) < count($r->getArr())) {
                    return $ltValue;
                }
                if (count($l->getArr()) > count($r->getArr())) {
                    return $gtValue;
                }
                return $eqValue;
            default:
                return $this->errorUnexpectedType($path, [PBType::TYPE_NUM, PBType::TYPE_BOOL, PBType::TYPE_STR, PBType::TYPE_ARR], $l->getType());
        }
    }

    private
    function isFiniteNumber(PBValue $v): bool
    {
        return $v->getType() === PBType::TYPE_NUM && is_finite($v->getNum());
    }

    private
    function isInteger(PBValue $v): bool
    {
        return $v->getType() === PBType::TYPE_NUM && $v->getNum() === (float)((int)($v->getNum()));
    }

    private
    function errorUnsupportedExpr(PBPath $path, PBValue $v): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::UNSUPPORTED_EXPR);
        $err->setErrorMessage(sprintf("unsupported expr: got [%s]", join(',', array_keys(iterator_to_array($v->getObj())))));
        $err->setErrorPath($path);
        return $err;
    }

    /**
     * @param PBPath $path
     * @param int[] $wantTypes
     * @param int $gotType
     * @return PBEvaluateExprOutput
     */
    private
    function errorUnexpectedType(PBPath $path, array $wantTypes, int $gotType): PBEvaluateExprOutput
    {
        $want = join(',', array_map(function ($t) {
            return PBType::name($t);
        }, $wantTypes));
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::UNEXPECTED_TYPE);
        $err->setErrorMessage(sprintf("unexpected type: want [%s], got [%s]", $want, PBType::name($gotType)));
        $err->setErrorPath($path);
        return $err;
    }

    private
    function errorArithmeticError(PBPath $path, string $message): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::ARITHMETIC_ERROR);
        $err->setErrorMessage(sprintf("arithmetic error: %s", $message));
        $err->setErrorPath($path);
        return $err;
    }

    private
    function errorIndexOutOfBounds(PBPath $path, int $begin, int $end, int $index): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::INDEX_OUT_OF_BOUNDS);
        $err->setErrorMessage(sprintf("index out of bounds: %d not in [%d, %d)", $index, $begin, $end));
        $err->setErrorPath($path);
        return $err;
    }

    /**
     * @param PBPath $path
     * @param string $want
     * @param string[] $actual
     * @return PBEvaluateExprOutput
     */
    private
    function errorKeyNotFound(PBPath $path, string $want, array $actual): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::KEY_NOT_FOUND);
        $err->setErrorMessage(sprintf("key not found: %s not in {%s}", $want, join(',', $actual)));
        $err->setErrorPath($path);
        return $err;
    }

    private
    function errorReferenceNotFound(PBPath $path, string $ref): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::REFERENCE_NOT_FOUND);
        $err->setErrorMessage(sprintf("reference not found: '%s'", $ref));
        $err->setErrorPath($path);
        return $err;
    }

    private
    function errorCasesNotExhaustive(PBPath $path): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::CASES_NOT_EXHAUSTIVE);
        $err->setErrorMessage("cases not exhaustive");
        $err->setErrorPath($path);
        return $err;
    }

    private
    function errorUnsupportedOperation(PBPath $path, string $gotOp): PBEvaluateExprOutput
    {
        $err = new PBEvaluateExprOutput;
        $err->setStatus(PBEvaluateExprOutput\PBStatus::UNSUPPORTED_OPERATION);
        $err->setErrorMessage(sprintf("unsupported operation: %s", $gotOp));
        $err->setErrorPath($path);
        return $err;
    }

    private function hasKey(PBValue $expr, string $key): bool
    {
        try {
            return $expr->getType() === PBType::TYPE_OBJ && $expr->getObj()->offsetGet($key);
        } catch (Exception) {
            return false;
        }
    }

    /**
     * @param PBValue $expr
     * @return string[]
     */
    private function getKeys(PBValue $expr): array
    {
        return array_keys(iterator_to_array($expr->getObj()));
    }

    private function compute_div(float $dividend, float $divisor): float
    {
        if (function_exists('fdiv')) {
            return fdiv($dividend, $divisor);
        } else {
            return @($dividend / $divisor);
        }
    }
}