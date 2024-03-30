package gozod

import (
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/Jh123x/go-validate/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestStringSchema(t *testing.T) {
	tests := map[string]struct {
		constraints []ttypes.ValTest[string]
		value       string
		expectedErr error
	}{
		"valid json string": {
			constraints: []ttypes.ValTest[string]{
				options.VIsNotDefault[string](),
				options.VIsValidJson,
			},
			value:       `{"name": "John"}`,
			expectedErr: nil,
		},
		"invalid json string": {
			constraints: []ttypes.ValTest[string]{
				options.VIsNotDefault[string](),
				options.VIsValidJson,
			},
			value:       `{"name": "John"`,
			expectedErr: errs.InvalidJsonError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			schema := NewStringSchema(tc.constraints...)
			assert.Equal(t, tc.expectedErr, schema.Parse(tc.value))
		})
	}
}
