# Eyamluate API in Go

## Installation

```shell
go get github.com/Jumpaku/eyamluate/golang@latest
```

## Example

```go
package main

import (
	"fmt"
	"github.com/Jumpaku/eyamluate/golang/eval"
	"github.com/Jumpaku/eyamluate/golang/yaml"
)

func main() {
	evaluated := eval.NewEvaluator().Evaluate(&eval.EvaluateInput{
		Source: `cat: ["Hello", ", ", "eyamluate", "!"]`,
	})
	encoded := yaml.NewEncoder().Encode(&yaml.EncodeInput{
		Value: evaluated.Value,
	})
	fmt.Println(encoded.Result)
	// Output: Hello, eyamluate!
}
```

## Eyamluate Project

https://github.com/Jumpaku/eyamluate