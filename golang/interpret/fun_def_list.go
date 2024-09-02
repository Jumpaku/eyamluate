package interpret

func EmptyFunDefList() *FunDefList {
	return nil
}

func (l *FunDefList) Register(def *FunDef) *FunDefList {
	return &FunDefList{
		Parent: l,
		Def:    def,
	}
}

func (l *FunDefList) Find(ident string) *FunDefList {
	cur := l
	for {
		if cur == nil {
			return nil
		}
		if cur.Def.Def == ident {
			return cur
		}
		cur = cur.Parent
	}
}
