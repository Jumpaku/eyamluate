import 'package:eyamluate/eval/eval.dart';
import 'package:fixnum/fixnum.dart';

Path pathAppendKey(Path path, String key) {
  return Path(pos: List.from(path.pos)..add(Path_Pos(key: key)));
}

Path pathAppendIndex(Path path, int index) {
  return Path(pos: List.from(path.pos)..add(Path_Pos(index: Int64(index))));
}
