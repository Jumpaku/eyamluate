package eval

import (
	"fmt"
	"github.com/Jumpaku/eyamluate/golang/yaml"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestEvaluator_Evaluate(t *testing.T) {
	type Testcase struct {
		YamlInput string
		WantValue *yaml.Value
		WantError bool
	}
	testcases := map[string]*Testcase{}
	err := filepath.Walk(filepath.Join("..", "..", "testdata"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".in.yaml") {
			key := strings.TrimSuffix(path, ".in.yaml")
			b, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("fail to read file: %+v", err)
			}
			if _, ok := testcases[key]; !ok {
				testcases[key] = &Testcase{}
			}
			testcases[key].YamlInput = string(b)
			return nil
		}
		if strings.HasSuffix(path, ".want.yaml") {
			key := strings.TrimSuffix(path, ".want.yaml")
			b, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("fail to read file: %+v", err)
			}
			want := yaml.NewDecoder().Decode(&yaml.DecodeInput{Yaml: string(b)})
			if want.IsError {
				return fmt.Errorf("fail to decode yaml: %+v", want.ErrorMessage)
			}
			if _, ok := testcases[key]; !ok {
				testcases[key] = &Testcase{}
			}
			if v, ok := want.Value.Obj["want_value"]; ok {
				testcases[key].WantValue = v
			}
			if v, ok := want.Value.Obj["want_error"]; ok {
				testcases[key].WantError = v.Bool
			}
			if testcases[key].WantValue == nil && !testcases[key].WantError {
				return fmt.Errorf("want_value or want_error is not found in %v", path)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	testcaseKeys := []string{}
	for key := range testcases {
		testcaseKeys = append(testcaseKeys, key)
	}
	slices.Sort(testcaseKeys)
	for _, name := range testcaseKeys {
		testcase := testcases[name]
		t.Run(name, func(t *testing.T) {
			got := NewEvaluator().Evaluate(&EvaluateInput{Source: testcase.YamlInput})
			if testcase.WantError {
				require.NotEqual(t, got.Status, EvaluateOutput_OK)
			} else {
				require.Equal(t, got.Status, EvaluateOutput_OK)
				require.Nil(t, checkEqual([]string{}, testcase.WantValue, got.Value))
			}
		})
	}
}

func checkEqual(path []string, want, got *yaml.Value) error {
	p := "/" + strings.Join(path, "/")
	if want.Type != got.Type {
		return fmt.Errorf("type mismatch: path=%v, got=%v, want=%v", p, got.Type, want.Type)
	}
	switch want.Type {
	default:
		return fmt.Errorf("unexpected type: path=%v, got=%v, want=%v", p, got.Type, want.Type)
	case yaml.Type_TYPE_NULL:
		return nil
	case yaml.Type_TYPE_BOOL:
		if want.Bool != got.Bool {
			return fmt.Errorf("boolean mismatch: path=%v, got=%v, want=%v", p, got.Bool, want.Bool)
		}
		return nil
	case yaml.Type_TYPE_NUM:
		if want.Num != got.Num {
			return fmt.Errorf("number mismatch: path=%v, got=%v, want=%v", p, got.Num, want.Num)
		}
		return nil
	case yaml.Type_TYPE_STR:
		if want.Str != got.Str {
			return fmt.Errorf("string mismatch: path=%v, got=%v, want=%v", p, got.Str, want.Str)
		}
		return nil
	case yaml.Type_TYPE_ARR:
		if len(want.Arr) != len(got.Arr) {
			return fmt.Errorf("array length mismatch: path=%v, got=%v, want=%v", p, len(got.Arr), len(want.Arr))
		}
		for i := 0; i < len(want.Arr); i++ {
			path := append([]string{}, path...)
			if err := checkEqual(append(path, fmt.Sprintf("%v", i)), want.Arr[i], got.Arr[i]); err != nil {
				return err
			}
		}
		return nil
	case yaml.Type_TYPE_OBJ:
		wk := want.Keys()
		slices.Sort(wk)
		gk := got.Keys()
		slices.Sort(gk)
		if !slices.Equal(wk, gk) {
			return fmt.Errorf("object keys mismatch: path=%v, got=[%v], want=[%v]", p, strings.Join(wk, ","), strings.Join(gk, ","))
		}
		for _, k := range wk {
			path := append([]string{}, path...)
			if err := checkEqual(append(path, k), want.Obj[k], got.Obj[k]); err != nil {
				return err
			}
		}
		return nil
	}
}
