import 'package:eyamluate/src/yaml.dart';
import 'package:eyamluate/yaml/yaml.dart';
import 'package:yaml/yaml.dart';

class Decoder {
  Decoder();

  DecodeOutput decode(DecodeInput input) {
    final dynamic decoded;
    try {
      decoded = loadYaml(input.yaml);
    } catch (e) {
      return DecodeOutput(isError: true, errorMessage: e.toString());
    }
    return DecodeOutput(
      value: convertFromDart(decoded),
    );
  }
}
