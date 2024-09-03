package yaml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/goccy/go-yaml"
)

type Encoder interface {
	Encode(*EncodeInput) *EncodeOutput
}

func NewEncoder() Encoder {
	return &encoder{}
}

type encoder struct{}

func (e *encoder) Encode(input *EncodeInput) *EncodeOutput {
	g := convertToGo(input.Value)
	switch input.Format {
	default:
		return &EncodeOutput{IsError: true, ErrorMessage: fmt.Sprintf("unexpected format %v", input.Format)}
	case EncodeFormat_ENCODE_FORMAT_JSON:
		b := bytes.NewBuffer(nil)
		e := json.NewEncoder(b)
		if input.Pretty {
			e.SetIndent("", "  ")
		}
		if err := e.Encode(g); err != nil {
			return &EncodeOutput{
				IsError:      true,
				ErrorMessage: fmt.Sprintf("fail to encode json: %+v", err),
			}
		}
		return &EncodeOutput{Result: b.String()}
	case EncodeFormat_ENCODE_FORMAT_YAML:
		b := bytes.NewBuffer(nil)
		e := yaml.NewEncoder(b)
		if input.Pretty {
			e = yaml.NewEncoder(b, yaml.Indent(2))
		}
		if err := e.Encode(g); err != nil {
			return &EncodeOutput{
				IsError:      true,
				ErrorMessage: fmt.Sprintf("fail to encode yaml: %+v", err),
			}
		}
		return &EncodeOutput{Result: b.String()}
	}
}

func convertToGo(v *Value) any {
	switch v.Type {
	default:
		panic(fmt.Sprintf("unexpected type %v", v.Type))
	case Type_TYPE_NULL:
		return nil
	case Type_TYPE_BOOL:
		return v.Bool
	case Type_TYPE_NUM:
		return v.Num
	case Type_TYPE_STR:
		return v.Str
	case Type_TYPE_ARR:
		arr := []any{}
		for _, elem := range v.Arr {
			arr = append(arr, convertToGo(elem))
		}
		return arr
	case Type_TYPE_OBJ:
		obj := map[string]any{}
		for key, value := range v.Obj {
			obj[key] = convertToGo(value)
		}
		return obj
	}
}
