# Eyamluate API in Go

## Installation

```shell
go get github.com/eyamluate/golang@latest
```

## Example

```go
package main

import (
	"fmt"
	"github.com/Jumpaku/eyamlate/golang/eval"
	"github.com/Jumpaku/eyamlate/golang/yaml"
)

func main() {
	evaluated := eval.NewEvaluator().Evaluate(&eval.EvaluateInput{
		Source: `cat: ["Hello", ", ", "eyamlate", "!"]`,
	})
	encoded := yaml.NewEncoder().Encode(&yaml.EncodeInput{
		Value: evaluated.Value,
	})
	fmt.Println(encoded.Result)
	// Output: Hello, eyamlate!
}
```

## Eyamluate Project

https://github.com/Jumpaku/eyamluate