# Eyamluate API in Dart

## Installation

```shell
dart pub add eyamluate
```

## Example

```dart
import 'package:eyamluate/eval/eval.dart';
import 'package:eyamluate/eval/evaluator.dart';
import 'package:eyamluate/yaml/encoder.dart';
import 'package:eyamluate/yaml/yaml.dart';

void main() {
  final evaluated = Evaluator().evaluate(EvaluateInput(
    source: '''cat: ["Hello", ", ", "eyamluate", "!"]''',
  ));
  final decoded = Encoder().encode(EncodeInput(
    value: evaluated.value,
  ));
  print(decoded.result);
  // Output: "Hello, eyamluate!"
}
```

## Eyamluate Project

https://github.com/Jumpaku/eyamluate