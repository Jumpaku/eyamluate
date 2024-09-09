import 'package:eyamluate/src/gen/yaml/value.pb.dart';
import 'package:yaml/yaml.dart';

Value convertFromDart(dynamic value) {
  return switch (value.runtimeType) {
    bool => Value(type: Type.TYPE_BOOL, bool_2: value),
    num || int || double => Value(type: Type.TYPE_NUM, num: value.toDouble()),
    String => Value(type: Type.TYPE_STR, str: value),
    YamlList => Value(
        type: Type.TYPE_ARR,
        arr: (value as YamlList).map(convertFromDart).toList(),
      ),
    YamlMap => Value(
        type: Type.TYPE_OBJ,
        obj: (value as YamlMap)
            .map((k, v) => MapEntry(k.toString(), convertFromDart(v))),
      ),
    _ => throw ArgumentError('Unsupported type: ${value.runtimeType}'),
  };
}

dynamic convertToDart(Value value) {
  return switch (value.type) {
    Type.TYPE_BOOL => value.bool_2,
    Type.TYPE_NUM => value.num,
    Type.TYPE_STR => value.str,
    Type.TYPE_ARR => value.arr.map(convertToDart).toList(),
    Type.TYPE_OBJ => value.obj.map((k, v) => MapEntry(k, convertToDart(v))),
    _ => throw ArgumentError('Unsupported type: ${value.type}'),
  };
}
