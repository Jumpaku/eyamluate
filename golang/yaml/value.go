package yaml

func (v *Value) CanInt() bool {
	return v.Type == Type_NUM && v.Num == float64(int64(v.Num))
}

func (v *Value) Keys() []string {
	if v.Type != Type_OBJ {
		return nil
	}
	keys := []string{}
	for k := range v.Obj {
		keys = append(keys, k)
	}
	return keys
}
