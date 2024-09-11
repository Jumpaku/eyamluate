<?php

namespace Jumpaku\Eyamluate\Tests\Eval_;

use Jumpaku\Eyamluate\Eval_\BaseEvaluator;
use Jumpaku\Eyamluate\Eval_\Evaluator;
use Jumpaku\Eyamluate\Eval_\PBEvaluateInput;
use Jumpaku\Eyamluate\Eval_\PBEvaluateOutput;
use Jumpaku\Eyamluate\Yaml\Decoder;
use Jumpaku\Eyamluate\Yaml\PBDecodeInput;
use Jumpaku\Eyamluate\Yaml\PBType;
use Jumpaku\Eyamluate\Yaml\PBValue;
use PHPUnit\Framework\TestCase;
use Symfony\Component\Yaml\Yaml;


class EvaluatorTest extends TestCase
{
    /** @var array */
    private $testcases = [];

    protected function setUp(): void
    {
        $dataDir = join(DIRECTORY_SEPARATOR, [__DIR__, "..", "..", "..", "testdata"]);
        $filePaths = [$dataDir];
        while (count($filePaths) > 0) {
            $path = array_shift($filePaths);
            if ($path == "." || $path == "..") {
                continue;
            }
            if (is_dir($path)) {
                foreach (scandir($path) as $child) {
                    if ($child == "." || $child == "..") {
                        continue;
                    }
                    $filePaths[] = join(DIRECTORY_SEPARATOR, [$path, $child]);
                }
                continue;
            }
            if (str_ends_with($path, ".in.yaml")) {
                $key = mb_strimwidth($path, 0, strlen($path) - strlen(".in.yaml"));
                if (!array_key_exists($key, $this->testcases)) {
                    $this->testcases[$key] = new EvaluatorTestcase;
                }
                $this->testcases[$key]->inputYaml = file_get_contents($path);
            } elseif (str_ends_with($path, ".want.yaml")) {
                $key = mb_strimwidth($path, 0, strlen($path) - strlen(".want.yaml"));
                if (!array_key_exists($key, $this->testcases)) {
                    $this->testcases[$key] = new EvaluatorTestcase;
                }
                $yaml = (new Decoder)->decode((new PBDecodeInput)->setYaml(file_get_contents($path)));
                if ($yaml->getIsError()) {
                    $this->fail($yaml->getErrorMessage());
                }
                /** @var PBValue $wantValue */
                $wantValue = $yaml->getValue()->getObj()["want_value"];
                if ($wantValue != null) {
                    $this->testcases[$key]->wantValue = $wantValue;
                }
                /** @var PBValue $wantError */
                $wantError = $yaml->getValue()->getObj()["want_error"];
                if ($wantError != null) {
                    $this->testcases[$key]->wantError = $wantError->getBool();
                }
            }
        }
    }

    public function testEvaluate()
    {
        $names = array_keys($this->testcases);
        foreach ($names as $name) {
            print_r($name . "\n");
            /** @var EvaluatorTestcase $t */
            $t = $this->testcases[$name];
            $sut = new BaseEvaluator();
            $got = $sut->evaluate((new PBEvaluateInput)->setSource($t->inputYaml));

            if ($t->wantError) {
                $this->assertNotEquals(PBEvaluateOutput\PBStatus::OK, $got->getStatus());
            } else {
                $this->assertEquals(PBEvaluateOutput\PBStatus::OK, $got->getStatus());
                $msg = $this->checkValueEqual([], $t->wantValue, $got->getValue());
                if ($msg != null) {
                    $this->fail($msg);
                }
            }

        }
    }

    /**
     * @param string[] $path
     * @param PBValue $want
     * @param PBValue $got
     * @return string|null
     */
    private function checkValueEqual(array $path, PBValue $want, PBValue $got): string|null
    {
        if ($want->getType() != $got->getType()) {
            return sprintf("type mismatch: /%s", join("/", $path));
        }
        switch ($want->getType()) {
            case PBType::TYPE_NULL:
                return null;
            case PBType::TYPE_BOOL:
                if ($want->getBool() != $got->getBool()) {
                    return sprintf("boolean mismatch want %s, got %s: /%s", $want->getBool(), $got->getBool(), join("/", $path));
                }
                return null;
            case PBType::TYPE_NUM:
                if ($want->getNum() != $got->getNum()) {
                    return sprintf("number mismatch want %s, got %s: /%s", $want->getNum(), $got->getNum(), join("/", $path));
                }
                return null;
            case PBType::TYPE_STR:
                if ($want->getStr() != $got->getStr()) {
                    return sprintf("string mismatch want %s, got %s: /%s", $want->getStr(), $got->getStr(), join("/", $path));
                }
                return null;
            case PBType::TYPE_ARR:
                if ($want->getArr()->count() != $got->getArr()->count()) {
                    return sprintf("array length mismatch want %d, got %d: /%s", $want->getArr()->count(), $got->getArr()->count(), join("/", $path));
                }
                foreach ($want->getArr() as $index => $wantItem) {
                    $gotItem = $got->getArr()[$index];
                    $msg = $this->checkValueEqual([...$path, (string)$index], $wantItem, $gotItem);
                    if ($msg != null) {
                        return $msg;
                    }
                }
                return null;
            case PBType::TYPE_OBJ:
                $wk = array_keys(iterator_to_array($want->getObj()));
                $gk = array_keys(iterator_to_array($got->getObj()));
                sort($wk);
                sort($gk);
                if ($wk != $gk) {
                    return sprintf("object keys mismatch want [%s], got [%s]: /%s", join(",", $wk), join(",", $gk), join("/", $path));
                }
                foreach ($want->getObj() as $key => $wantItem) {
                    /** @var PBValue $gotItem */
                    $gotItem = $got->getObj()[$key];
                    $msg = $this->checkValueEqual([...$path, $key], $wantItem, $gotItem);
                    if ($msg != null) {
                        return $msg;
                    }
                }
                return null;
            default:
                assert(false, sprintf("unexpected type: %s", PBType::name($want->getType())));
        }
    }
}