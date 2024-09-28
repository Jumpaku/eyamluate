package golang_test

import (
	"fmt"
	"github.com/Jumpaku/eyamluate/golang/eval"
	"github.com/Jumpaku/eyamluate/golang/yaml"
)

func ExampleEvaluator_Evaluate() {
	evaluated := eval.NewEvaluator().Evaluate(&eval.EvaluateInput{
		Source: `cat: ["Hello", ", ", "eyamluate", "!"]`,
	})
	encoded := yaml.NewEncoder().Encode(&yaml.EncodeInput{
		Value: evaluated.Value,
	})
	fmt.Println(encoded.Result)
	// Output: Hello, eyamluate!
}
