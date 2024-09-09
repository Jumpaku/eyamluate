import 'package:eyamluate/yaml/yaml.dart';

bool valueCanInt(Value value) {
  return value.type == Type.TYPE_NUM &&
      value.num == value.num.toInt().toDouble();
}
