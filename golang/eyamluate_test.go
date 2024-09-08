package golang_test

import (
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/eval"
	"github.com/Jumpaku/eyamlate/golang/yaml"
)

func ExampleEyamluate() {
	evaluated := eval.NewEvaluator().Evaluate(&eval.EvaluateInput{
		Source: `cat: ["Hello", ", ", "eyamlate", "!"]`,
	})
	decoded := yaml.NewEncoder().Encode(&yaml.EncodeInput{
		Value: evaluated.Value,
	})
	fmt.Println(decoded.Result)
	// Output: Hello, eyamlate!
}
