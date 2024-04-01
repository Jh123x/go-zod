package gozod

import (
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
)

type BoolSchema struct {
	constraints ttypes.ValTest[bool]
}

func NewBoolSchema(constraints ...ttypes.ValTest[bool]) *BoolSchema {
	return &BoolSchema{constraints: options.VAnd(constraints...)}
}

func (i *BoolSchema) Compose(fileName string, schema *BoolSchema) *BoolSchema {
	return NewBoolSchema(options.VAnd(i.constraints, schema.Parse))
}

func (i *BoolSchema) Parse(val bool) error {
	return i.constraints(val)
}

func (i *BoolSchema) ToValidate(val bool) ttypes.Validate {
	return func() error { return i.constraints(val) }
}
