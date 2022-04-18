// Code generated by goa v3.7.2, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	calc "calcsvc/gen/calc"
	"encoding/json"
	"fmt"
	"strconv"
)

// BuildMultiplyPayload builds the payload for the calc multiply endpoint from
// CLI flags.
func BuildMultiplyPayload(calcMultiplyType string, calcMultiplyMetadata string) (*calc.MultiplyPayload, error) {
	var err error
	var type_ *int
	{
		if calcMultiplyType != "" {
			var v int64
			v, err = strconv.ParseInt(calcMultiplyType, 10, strconv.IntSize)
			val := int(v)
			type_ = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for type_, must be INT")
			}
		}
	}
	var metadata map[string]string
	{
		if calcMultiplyMetadata != "" {
			err = json.Unmarshal([]byte(calcMultiplyMetadata), &metadata)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for metadata, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"testKey\": \"val\"\n   }'")
			}
		}
	}
	v := &calc.MultiplyPayload{}
	v.Type = type_
	v.Metadata = metadata

	return v, nil
}