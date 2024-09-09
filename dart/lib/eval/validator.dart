import 'dart:convert';

import 'package:embed_annotation/embed_annotation.dart';
import 'package:eyamluate/eval/eval.dart';
import 'package:json_schema/json_schema.dart';
import 'package:yaml/yaml.dart';

part 'validator.g.dart';

@EmbedBinary("../../../schema/eyamluate.schema.yaml")
const languageSchemaYAMLBytes = _$languageSchemaYAMLBytes;

class Validator {
  ValidateOutput validate(ValidateInput input) {
    final csvArmorSchemaYAML = utf8.decode(languageSchemaYAMLBytes);
    final schema = jsonEncode(loadYaml(csvArmorSchemaYAML));
    final validator = JsonSchema.create(schema);

    final dynamic v;
    try {
      v = loadYaml(input.source);
    } catch (e) {
      return ValidateOutput(
        status: ValidateOutput_Status.YAML_ERROR,
        errorMessage: e.toString(),
      );
    }

    final r = validator.validate(v);
    if (!r.isValid) {
      return ValidateOutput(
        status: ValidateOutput_Status.VALIDATION_ERROR,
        errorMessage: r.errors.toString(),
      );
    }

    return ValidateOutput(status: ValidateOutput_Status.OK);
  }
}
