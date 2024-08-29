package ast

import (
	"fmt"
	"strings"
)

func (p *Path) AppendIndex(i int) *Path {
	return &Path{
		Pos: append(append([]*Path_Pos{}, p.Pos...), &Path_Pos{Index: int64(i)}),
	}
}

func (p *Path) AppendKey(k string) *Path {
	return &Path{
		Pos: append(append([]*Path_Pos{}, p.Pos...), &Path_Pos{Key: k}),
	}
}

func (p *Path) Format() string {
	var s []string
	for _, pos := range p.Pos {
		if pos.Key == "" {
			s = append(s, fmt.Sprint(pos.Index))
		} else {
			s = append(s, fmt.Sprintf("%v", pos.Key))
		}
	}
	return "/" + strings.Join(s, "/")
}
