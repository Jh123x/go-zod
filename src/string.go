package gozod

import (
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
)

type StringSchema struct {
	constraints ttypes.ValTest[string]
}

func NewStringSchema(constraints ...ttypes.ValTest[string]) *StringSchema {
	return &StringSchema{constraints: options.VAnd(constraints...)}
}

func (s *StringSchema) Compose(fileName string, schema *StringSchema) *StringSchema {
	return NewStringSchema(options.VAnd(s.constraints, schema.Parse))
}

func (s *StringSchema) Parse(val string) error {
	return s.constraints(val)
}

func (s *StringSchema) ToValidate(val string) ttypes.Validate {
	return func() error { return s.constraints(val) }
}
