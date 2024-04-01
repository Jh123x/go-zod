package gozod

import (
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestBoolSchema(t *testing.T) {
	tests := map[string]struct {
		constraints []ttypes.ValTest[bool]
		value       bool
		expectedErr error
	}{
		"valid bool": {
			constraints: []ttypes.ValTest[bool]{
				options.VIsDefault[bool](),
			},
			value:       false,
			expectedErr: nil,
		},
		"invalid bool": {
			constraints: []ttypes.ValTest[bool]{
				options.VIsDefault[bool](),
			},
			value:       true,
			expectedErr: errs.IsDefaultErr,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			schema := NewBoolSchema(tc.constraints...)
			assert.Equal(t, tc.expectedErr, schema.Parse(tc.value))
		})
	}
}
