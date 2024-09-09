# eyamluate

`eyamluate` is a YAML expression evaluator that is small, integratable, and extensible.
This repository provides:

- The `eyamluate` language: the programming language to represent and evaluate expressions written in the YAML format, which is defined as a JSON Schema file.
- The `eyamluate` API: the API specification to process the code written in the `eyamluate` language, which is defined as Protocol Buffers Schema files.
- The `eyamluate` libraries: a collection of exemplar implementations for the `eyamluate` API, which are currently available in Go and Dart.
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

https://github.com/Jumpaku/eyamluate/blob/main/golang/README.md

## CLI

### Installation

```shell
go install github.com/eyamluate/golang/cmd/eyamluate@latest
```

### Usage

The CLI document is available at https://github.com/Jumpaku/eyamluate/main/golang/cmd/eyamluate/README.md

### Eaxmple

```shell
echo 'cat: ["Hello", ", ", "eyamlate", "!"]' | eyamluate eval
# => 'Hello, eyamlate!'
```

## Related Work

- https://murano.readthedocs.io/en/stable-liberty/appdev-guide/murano_pl.html
  - Murano Programming Language is an object-oriented language represented in the YAML format. 
- https://yamlscript.org
  - YAMLScript is a programming language based on the YAML format.
- https://docs.racket-lang.org/yaml-exp/index.html
  - yaml-exp is a variation of the Racket Language in the YAML format.
- https://github.com/google/cel-spec
  - Common Expression Language (CEL) is a language to represent and evaluate expressions.