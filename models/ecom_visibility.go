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

// EcomVisibility Determines item visibility in Ecom (Online Store) and Online Checkout.
//
// swagger:model EcomVisibility
type EcomVisibility string

const (

	// EcomVisibilityUNINDEXED captures enum value "UNINDEXED"
	EcomVisibilityUNINDEXED EcomVisibility = "UNINDEXED"

	// EcomVisibilityUNAVAILABLE captures enum value "UNAVAILABLE"
	EcomVisibilityUNAVAILABLE EcomVisibility = "UNAVAILABLE"

	// EcomVisibilityHIDDEN captures enum value "HIDDEN"
	EcomVisibilityHIDDEN EcomVisibility = "HIDDEN"

	// EcomVisibilityVISIBLE captures enum value "VISIBLE"
	EcomVisibilityVISIBLE EcomVisibility = "VISIBLE"
)

// for schema
var ecomVisibilityEnum []interface{}

func init() {
	var res []EcomVisibility
	if err := json.Unmarshal([]byte(`["UNINDEXED","UNAVAILABLE","HIDDEN","VISIBLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ecomVisibilityEnum = append(ecomVisibilityEnum, v)
	}
}

func (m EcomVisibility) validateEcomVisibilityEnum(path, location string, value EcomVisibility) error {
	if err := validate.EnumCase(path, location, value, ecomVisibilityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ecom visibility
func (m EcomVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEcomVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ecom visibility based on context it is used
func (m EcomVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
