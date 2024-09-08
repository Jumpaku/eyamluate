# eyamluate

`eyamluate` is a YAML expression evaluator that is small, integratable, and extensible.
This repository provides:

- The `eyamluate` language: the programming language to represent and evaluate expressions written in the YAML format, which is defined as a JSON Schema file.
- The `eyamluate` API: the API specification to process the code written in the `eyamluate` language, which is defined as Protocol Buffers Schema files.
- The `eyamluate` libraries: a collection of exemplar implementations for the `eyamluate` API, which is currently available in only Go.
- The `eyamluate` testcases: comprehensive testcases to check libraries implementing the `eyamluate` API. 

## Concepts

### Small

The `eyamluate` language is designed to be small.
The language definition is written in a JSON Schema of about 250 lines.
The core algorithm of the implementation of the `eyamluate` API is written within 1000 lines in Go.

### Integratable

Expressions written in the`eyamluate` language can be integrated into YAML file.
The JSON Schema of the `eyamluate` language enables editors and IDEs to utilize functionalities such as syntax highlighting, code completion, and static analysis.
The `eyamluate` libraries enables to evaluate the expression in the YAML file

### Extensible

The functionalities of the YAML expressions can be extended as necessary.
User-defined functions can be plugged-in to the `eyamluate` API.

## Eyamluate API Usage

### Go

```shell
go get github.com/eyamluate/golang@latest
```

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
	decoded := yaml.NewEncoder().Encode(&yaml.EncodeInput{
		Value: evaluated.Value,
	})
	fmt.Println(decoded.Result)
	// Output: Hello, eyamlate!
}

```

## CLI

### Installation

```shell
go install github.com/eyamluate/golang/cmd/eyamluate@latest
```

### Usage

The CLI document is available at https://github.com/Jumpaku/eyamluate/main/golang/cmd/eyamluate/README.md

#### Eaxmple

```shell
echo 'cat: ["Hello", ", ", "eyamlate", "!"]' | eyamluate eval
# => 'Hello, eyamlate!'
```

## Related Work

- https://murano.readthedocs.io/en/stable-liberty/appdev-guide/murano_pl.html
  - Murano Programming Language is an object-oriented language in represented in YAML format. 
- https://yamlscript.org
  - YAMLScript is a programming language based on YAML format.
- https://docs.racket-lang.org/yaml-exp/index.html
  - yaml-exp a variation of the Racket Language in YAML format.
- https://github.com/google/cel-spec
  - Common Expression Language (CEL) is a language to represent and evaluate expressions.