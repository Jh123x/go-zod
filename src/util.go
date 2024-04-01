package gozod

import (
	"fmt"
	"reflect"
	"strings"
)

func IsValidSchema(schema any, fieldType string) bool {
	if schema == nil {
		return false
	}
	switch schema.(type) {
	case *StringSchema:
		return fieldType == "string"
	case *IntSchema:
		return fieldType == "int"
	case *BoolSchema:
		return fieldType == "bool"
	default:
		return false
	}
}

func GenericIsOfType(genericType, fieldType reflect.Type) bool {
	genericTypeVal := genericType.String()
	fieldTypeVal := fieldType.Name()
	res := genericTypeRegex.FindAllString(genericTypeVal, -1)
	if len(res) == 0 {
		return false
	}

	res2 := genericTypeRegex.FindStringSubmatch(genericTypeVal)
	if len(res2) != 2 {
		return false
	}
	typeMatch := res2[1]
	return len(typeMatch) > 0 && strings.HasSuffix(typeMatch, fmt.Sprintf(".%s", fieldTypeVal))
}
