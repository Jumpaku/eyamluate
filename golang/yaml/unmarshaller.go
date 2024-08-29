package yaml

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goccy/go-yaml"
)

type Unmarshaller struct{}

var _ UnmarshallerServer = (*Unmarshaller)(nil)

func (u *Unmarshaller) mustEmbedUnimplementedUnmarshallerServer() {
	//TODO implement me
	panic("implement me")
}

func (u *Unmarshaller) Unmarshal(ctx context.Context, input *UnmarshalInput) (*UnmarshalOutput, error) {
	b, err := yaml.YAMLToJSON([]byte(input.Yaml))
	if err != nil {
		return nil, fmt.Errorf("fail to convert yaml to json: %w", err)
	}
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, fmt.Errorf("fail to unmarshal: %w", err)
	}
	return &UnmarshalOutput{Value: convert(v)}, nil
}

func convert(v any) *Value {
	switch v := v.(type) {
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	case nil:
		return &Value{Type: Type_TYPE_NULL}
	case bool:
		return &Value{Type: Type_TYPE_BOOL, Bool: v}
	case float64:
		return &Value{Type: Type_TYPE_NUM, Num: v}
	case string:
		return &Value{Type: Type_TYPE_STR, Str: v}
	case []interface{}:
		arr := []*Value{}
		for _, elem := range v {
			arr = append(arr, convert(elem))
		}
		return &Value{Type: Type_TYPE_ARR, Arr: arr}
	case map[string]interface{}:
		obj := map[string]*Value{}
		for key, value := range v {
			obj[key] = convert(value)
		}
		return &Value{Type: Type_TYPE_OBJ, Obj: obj}
	}
}
