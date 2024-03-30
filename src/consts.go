package gozod

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidField  = errors.New("invalid field")
	ErrInvalidSchema = errors.New("invalid schema")

	genericTypeRegex = regexp.MustCompile(`.*\[(.*)\]`)
)
