import 'dart:convert';

import 'package:eyamluate/src/yaml.dart';
import 'package:eyamluate/yaml/yaml.dart';
import 'package:yaml_edit/yaml_edit.dart';

class Encoder {
  Encoder();

  EncodeOutput encode(EncodeInput input) {
    final v = convertToDart(input.value);
    final String s;
    switch (input.format) {
      case EncodeFormat.ENCODE_FORMAT_JSON:
        try {
          final encoder = input.pretty ? JsonEncoder.withIndent('  ') : JsonEncoder();
          s = encoder.convert(v);
        } catch (e) {
          return EncodeOutput(isError: true, errorMessage: e.toString());
        }
      case EncodeFormat.ENCODE_FORMAT_YAML:
        try {
          s = (YamlEditor('')..update([], v)).toString();
        } catch (e) {
          return EncodeOutput(isError: true, errorMessage: e.toString());
        }
      default:
        return EncodeOutput(
            isError: true, errorMessage: 'Unsupported format: ${input.format}');
    }
    return EncodeOutput(result: s);
  }
}
