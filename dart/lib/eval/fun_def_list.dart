import 'package:eyamluate/eval/eval.dart';

FunDefList funDefListEmpty() {
  return FunDefList();
}

FunDefList funDefListRegister(FunDefList list, FunDef def) {
  return FunDefList(def: def, parent: list);
}

FunDefList? funDefListFind(FunDefList list, String ident) {
  var cur = list;
  while (true) {
    if (!cur.hasDef()) {
      return null;
    }
    if (cur.def.def == ident) {
      return cur;
    }
    if (!cur.hasParent()) {
      return null;
    }
    cur = cur.parent;
  }
}
