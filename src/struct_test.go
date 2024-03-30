package gozod

import (
	"testing"

	"github.com/Jh123x/go-validate/errs"
	"github.com/Jh123x/go-validate/options"
	"github.com/stretchr/testify/assert"
)

type TestStruct2 struct {
	EmailVal string
}

type TestStruct struct {
	JsonVal string
	Struct2 TestStruct2
}

func TestNewStructSchema(t *testing.T) {
	testsS, err := NewStructSchema[TestStruct2](
		map[string]any{
			"EmailVal": NewStringSchema(
				options.VIsNotDefault[string](),
				options.VIsValidEmail,
			),
		},
	)
	assert.Nil(t, err)
	tests := map[string]struct {
		objSchema   map[string]any
		expectedErr error
	}{
		"invalid struct validator should error": {
			objSchema: map[string]any{
				"JsonVal": NewStringSchema(
					options.VIsNotDefault[string](),
					options.VIsValidJson,
				),
				"Struct2": NewStringSchema(
					options.VIsNotDefault[string](),
					options.VIsValidEmail,
				),
			},
			expectedErr: ErrInvalidSchema,
		},
		"valid struct": {
			objSchema: map[string]any{
				"JsonVal": NewStringSchema(
					options.VIsNotDefault[string](),
					options.VIsValidJson,
				),
				"Struct2": testsS,
			},
			expectedErr: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewStructSchema[TestStruct](tc.objSchema)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestStructSchema(t *testing.T) {
	testsS, err := NewStructSchema[TestStruct2](
		map[string]any{
			"EmailVal": NewStringSchema(
				options.VIsNotDefault[string](),
				options.VIsValidEmail,
			),
		},
	)
	assert.Nil(t, err)
	tests := map[string]struct {
		objSchema   map[string]any
		value       TestStruct
		expectedErr error
	}{
		"valid schema match should throw no errors": {
			objSchema: map[string]any{
				"JsonVal": NewStringSchema(
					options.VIsNotDefault[string](),
					options.VIsValidJson,
				),
				"Struct2": testsS,
			},
			value: TestStruct{
				JsonVal: `{"key": "value"}`,
				Struct2: TestStruct2{EmailVal: "email@email.com"},
			},
			expectedErr: nil,
		},
		"invalid string schema match should throw error": {
			objSchema: map[string]any{
				"JsonVal": NewStringSchema(
					options.VIsNotDefault[string](),
					options.VIsValidJson,
				),
				"Struct2": testsS,
			},
			value: TestStruct{
				JsonVal: `{"key": "value"`, // Invalid json
				Struct2: TestStruct2{EmailVal: "email@email.com"},
			},
			expectedErr: errs.InvalidJsonError,
		},
		"invalid struct error match should throw error": {
			objSchema: map[string]any{
				"JsonVal": NewStringSchema(
					options.VIsNotDefault[string](),
					options.VIsValidJson,
				),
				"Struct2": testsS,
			},
			value: TestStruct{
				JsonVal: `{"key": "value"}`,
				Struct2: TestStruct2{EmailVal: "not an email"}, // Invalid email
			},
			expectedErr: errs.InvalidEmailError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			schema, err := NewStructSchema[TestStruct](tc.objSchema)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedErr, schema.Parse(tc.value))
		})
	}
}
