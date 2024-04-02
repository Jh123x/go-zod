package gozod

import (
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
)

type IntSchema struct {
	constraints ttypes.ValTest[int]
}

func NewIntSchema(constraints ...ttypes.ValTest[int]) *IntSchema {
	return &IntSchema{constraints: options.VAnd(constraints...)}
}

func (i *IntSchema) Compose(fileName string, schema *IntSchema) *IntSchema {
	return NewIntSchema(options.VAnd(i.constraints, schema.Parse))
}

func (i *IntSchema) Parse(val int) error {
	return i.constraints(val)
}

func (i *IntSchema) ToValidate(val int) ttypes.Validate {
	return func() error { return i.constraints(val) }
}
