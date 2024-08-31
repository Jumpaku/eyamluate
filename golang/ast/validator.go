package ast

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

type Validator interface {
	Validate(*ValidateInput) *ValidateOutput
}

//go:embed schema.yaml
var schema string

func NewValidator() Validator {
	b, err := yaml.YAMLToJSON([]byte(schema))
	if err != nil {
		panic(fmt.Sprintf("fail to convert yaml to json: %+v", err))
	}

	v, err := jsonschema.UnmarshalJSON(bytes.NewBuffer(b))
	if err != nil {
		panic(fmt.Sprintf("fail to decode json: %+v", err))
	}

	c := jsonschema.NewCompiler()
	if err := c.AddResource("schema", v); err != nil {
		panic(fmt.Sprintf("fail to add schema: %+v", err))
	}

	s, err := c.Compile("schema")
	if err != nil {
		panic(fmt.Sprintf("fail to compile schema: %+v", err))
	}

	return &validator{schema: s}
}

type validator struct {
	schema *jsonschema.Schema
}

var _ Validator = &validator{}

func (v *validator) Validate(input *ValidateInput) *ValidateOutput {
	b, err := yaml.YAMLToJSON([]byte(input.Source))
	if err != nil {
		return &ValidateOutput{
			Status:       ValidateOutput_YAML_ERROR,
			ErrorMessage: fmt.Sprintf("fail to convert yaml to json: %+v", err),
		}
	}

	sourceJSON, err := jsonschema.UnmarshalJSON(bytes.NewBuffer(b))
	if err != nil {
		return &ValidateOutput{
			Status:       ValidateOutput_YAML_ERROR,
			ErrorMessage: fmt.Sprintf("fail to decode json: %+v", err),
		}
	}

	if err := v.schema.Validate(sourceJSON); err != nil {
		return &ValidateOutput{
			Status:       ValidateOutput_VALIDATION_ERROR,
			ErrorMessage: fmt.Sprintf("validation error: %+v", err),
		}
	}

	return &ValidateOutput{Status: ValidateOutput_OK}
}
