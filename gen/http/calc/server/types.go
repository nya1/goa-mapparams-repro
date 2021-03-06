// Code generated by goa v3.7.2, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	calc "calcsvc/gen/calc"
)

// NewMultiplyPayload builds a calc service multiply endpoint payload.
func NewMultiplyPayload(type_ *int, metadata map[string]string) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{}
	v.Type = type_
	v.Metadata = metadata

	return v
}
