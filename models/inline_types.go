// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// InlineTypes Object types to inline under their respective parent object in certain connect v2 responses
//
// swagger:model InlineTypes
type InlineTypes string

const (

	// InlineTypesINLINENONE captures enum value "INLINE_NONE"
	InlineTypesINLINENONE InlineTypes = "INLINE_NONE"

	// InlineTypesINLINEVARIATIONS captures enum value "INLINE_VARIATIONS"
	InlineTypesINLINEVARIATIONS InlineTypes = "INLINE_VARIATIONS"

	// InlineTypesINLINEALL captures enum value "INLINE_ALL"
	InlineTypesINLINEALL InlineTypes = "INLINE_ALL"
)

// for schema
var inlineTypesEnum []interface{}

func init() {
	var res []InlineTypes
	if err := json.Unmarshal([]byte(`["INLINE_NONE","INLINE_VARIATIONS","INLINE_ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inlineTypesEnum = append(inlineTypesEnum, v)
	}
}

func (m InlineTypes) validateInlineTypesEnum(path, location string, value InlineTypes) error {
	if err := validate.EnumCase(path, location, value, inlineTypesEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this inline types
func (m InlineTypes) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInlineTypesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this inline types based on context it is used
func (m InlineTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
