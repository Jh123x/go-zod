package gozod

import (
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestIntSchema(t *testing.T) {
	tests := map[string]struct {
		constraints []ttypes.ValTest[int]
		value       int
		expectedErr error
	}{
		"valid int": {
			constraints: []ttypes.ValTest[int]{options.VIsDefault[int]()},
			value:       0,
			expectedErr: nil,
		},
		"invalid int": {
			constraints: []ttypes.ValTest[int]{options.VIsDefault[int]()},
			value:       1,
			expectedErr: errs.IsDefaultErr,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			schema := NewIntSchema(tc.constraints...)
			assert.Equal(t, tc.expectedErr, schema.Parse(tc.value))
		})
	}
}
