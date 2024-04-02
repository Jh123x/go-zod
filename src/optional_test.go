package gozod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptional_IsPresent(t *testing.T) {
	tests := map[string]struct {
		optional        Optional[int]
		expectedPresent bool
	}{
		"present": {
			optional:        NewOptional[int](1),
			expectedPresent: true,
		},
		"not present": {
			optional:        NewEmptyOptional[int](),
			expectedPresent: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedPresent, tc.optional.IsPresent())
		})
	}
}

func TestOptional_UnWrap(t *testing.T) {
	tests := map[string]struct {
		optional    Optional[int]
		expected    int
		expectedErr bool
	}{
		"present": {
			optional: NewOptional[int](1),
			expected: 1,
		},
		"not present": {
			optional:    NewEmptyOptional[int](),
			expectedErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				assert.Equal(t, tc.expectedErr, recover() != nil)
			}()
			assert.Equal(t, tc.expected, tc.optional.UnWrap())
		})
	}
}

func TestOptional_OrElse(t *testing.T) {
	tests := map[string]struct {
		optional      Optional[int]
		expectedValue int
	}{
		"present": {
			optional:      NewOptional(1),
			expectedValue: 1,
		},
		"not present": {
			optional:      NewEmptyOptional[int](),
			expectedValue: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedValue, tc.optional.OrElse(0))
		})
	}
}
