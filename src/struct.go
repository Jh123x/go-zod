package gozod

import (
	"reflect"

	"github.com/Jh123x/go-validate/ttypes"
)

type StructSchema[T any] struct {
	fieldSchema map[string]any
}

func NewStructSchema[T any](objSchema map[string]any) (*StructSchema[T], error) {
	// Validate objSchema is valid
	var testStruct T
	for fieldName, fieldSchema := range objSchema {
		structType := reflect.TypeOf(testStruct)
		field, ok := structType.FieldByName(fieldName)
		if !ok {
			return nil, ErrInvalidField
		}

		if !IsValidSchema(fieldSchema, field.Type.Name()) && !GenericIsOfType(reflect.TypeOf(fieldSchema), field.Type) {
			return nil, ErrInvalidSchema
		}
	}
	return &StructSchema[T]{fieldSchema: objSchema}, nil
}

func (s *StructSchema[T]) Parse(val T) error {
	for fieldName, fieldSchema := range s.fieldSchema {
		fieldValue := reflect.ValueOf(val).FieldByName(fieldName)
		switch schema := fieldSchema.(type) {
		case *StringSchema:
			if err := schema.Parse(fieldValue.Interface().(string)); err != nil {
				return err
			}
		default:
			if err := reflect.ValueOf(schema).MethodByName("Parse").Call([]reflect.Value{fieldValue}); err != nil {
				val := err[0]
				if val.IsNil() {
					return nil
				}
				return val.Interface().(error)
			}
		}
	}
	return nil
}

func (s *StructSchema[T]) ToValidate() ttypes.Validate {
	return func() error { return nil }
}
