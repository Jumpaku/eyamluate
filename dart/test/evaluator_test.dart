import 'dart:io';

import 'package:collection/collection.dart';
import 'package:eyamluate/eval/eval.dart';
import 'package:eyamluate/eval/evaluator.dart';
import 'package:eyamluate/yaml/decoder.dart';
import 'package:eyamluate/yaml/yaml.dart';
import 'package:path/path.dart' as p;
import 'package:test/test.dart';

class _Testcase {
  String? yamlInput;
  Value? wantValue;
  bool? wantError;
}

void main() {
  final testcases = <String, _Testcase>{};
  final testdataDir =
      Directory(p.join(Directory.current.path, "test", "testdata"));
  testdataDir.listSync(recursive: true, followLinks: true).forEach((e) {
    if (e is! File) return;
    final path = p.canonicalize(e.path);
    if (e.path.endsWith(".in.yaml")) {
      final key = path.substring(0, path.length - ".in.yaml".length);
      if (!testcases.containsKey(key)) {
        testcases[key] = _Testcase();
      }
      testcases[key]!.yamlInput = e.readAsStringSync();
      return;
    }
    if (e.path.endsWith(".want.yaml")) {
      final key = path.substring(0, path.length - ".want.yaml".length);
      if (!testcases.containsKey(key)) {
        testcases[key] = _Testcase();
      }
      final d = Decoder().decode(DecodeInput(yaml: e.readAsStringSync()));
      if (d.isError) {
        fail("fail to decode want yaml file: ${e.path}: ${d.errorMessage}");
      }
      final v = d.value;
      if (v.type != Type.TYPE_OBJ) {
        fail("want yaml file is not a map: ${e.path}");
      }
      final wantError = v.obj["want_error"];
      if (wantError != null) {
        if (wantError.type != Type.TYPE_BOOL) {
          fail("want_error is not a bool: ${e.path}");
        }
        testcases[key]!.wantError = wantError.bool_2;
        return;
      }
      final wantValue = v.obj["want_value"];
      if (wantValue != null) {
        testcases[key]!.wantValue = wantValue;
        return;
      }
    }
  });

  final names = testcases.keys.toList()..sort();
  for (final name in names) {
    final testcase = testcases[name]!;
    test(name, () {
      final sut = Evaluator();
      final got = sut.evaluate(EvaluateInput(source: testcase.yamlInput));
      if (testcase.wantError ?? false) {
        expect(got.status, isNot(equals(EvaluateOutput_Status.OK)));
      } else {
        expect(got.status, (EvaluateOutput_Status.OK));
        final msg = _checkEqual([], testcase.wantValue!, got.value);
        if (msg != null) {
          fail(msg);
        }
      }
    });
  }
}

String? _checkEqual(List<String> path, Value want, Value got) {
  if (want.type != got.type) {
    return "type mismatch: /${path.join("/")}: ${want.type} != ${got.type}";
  }
  switch (want.type) {
    case Type.TYPE_NULL:
      return null;
    case Type.TYPE_BOOL:
      if (want.bool_2 != got.bool_2) {
        return "bool mismatch: /${path.join("/")}: ${want.bool_2} != ${got.bool_2}";
      }
      return null;
    case Type.TYPE_NUM:
      if (want.num != got.num) {
        return "num mismatch: /${path.join("/")}: ${want.num} != ${got.num}";
      }
      return null;
    case Type.TYPE_STR:
      if (want.str != got.str) {
        return "str mismatch: /${path.join("/")}: ${want.str} != ${got.str}";
      }
      return null;
    case Type.TYPE_ARR:
      if (want.arr.length != got.arr.length) {
        return "arr length mismatch: /${path.join("/")}: ${want.arr.length} != ${got.arr.length}";
      }
      for (var i = 0; i < want.arr.length; i++) {
        final msg =
            _checkEqual([...path, i.toString()], want.arr[i], got.arr[i]);
        if (msg != null) {
          return msg;
        }
      }
      return null;
    case Type.TYPE_OBJ:
      if (!ListEquality().equals(
          want.obj.keys.toList()..sort(), got.obj.keys.toList()..sort())) {
        return "obj key mismatch: /${path.join("/")}: ${want.obj.keys.toList()..sort()} != ${got.obj.keys.toList()..sort()}";
      }
      for (var key in want.obj.keys) {
        expect(got.obj.containsKey(key), isTrue);
        final msg = _checkEqual(
            [...path, key.toString()], want.obj[key]!, got.obj[key]!);
        if (msg != null) {
          return msg;
        }
      }
      return null;
    default:
      return "unknown type: /${path.join("/")}: ${want.type}";
  }
}
