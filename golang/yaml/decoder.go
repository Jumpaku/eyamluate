package yaml

import (
	"encoding/json"
	"fmt"
	"github.com/goccy/go-yaml"
)

type Decoder interface {
	Decode(*DecodeInput) *DecodeOutput
}
type decoder struct{}

func NewDecoder() Decoder {
	return &decoder{}
}
func (u *decoder) Decode(input *DecodeInput) *DecodeOutput {
	b, err := yaml.YAMLToJSON([]byte(input.Yaml))
	if err != nil {
		return &DecodeOutput{
			IsError:      true,
			ErrorMessage: fmt.Sprintf("fail to convert yaml to json: %+v", err),
		}
	}
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return &DecodeOutput{
			IsError:      true,
			ErrorMessage: fmt.Sprintf("fail to unmarshal json: %+v", err),
		}
	}
	return &DecodeOutput{Value: convert(v)}
}

func convert(v any) *Value {
	switch v := v.(type) {
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	case nil:
		return &Value{Type: Type_NULL}
	case bool:
		return &Value{Type: Type_BOOL, Bool: v}
	case float64:
		return &Value{Type: Type_NUM, Num: v}
	case string:
		return &Value{Type: Type_STR, Str: v}
	case []interface{}:
		arr := []*Value{}
		for _, elem := range v {
			arr = append(arr, convert(elem))
		}
		return &Value{Type: Type_ARR, Arr: arr}
	case map[string]interface{}:
		obj := map[string]*Value{}
		for key, value := range v {
			obj[key] = convert(value)
		}
		return &Value{Type: Type_OBJ, Obj: obj}
	}
}
