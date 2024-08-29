package eval

import "github.com/Jumpaku/eyamlate/golang/pb/ast"

func (l *FunDefList) Register(def *ast.FunDef) *FunDefList {
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
