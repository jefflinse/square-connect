// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ProductType product type
//
// swagger:model ProductType
type ProductType string

const (

	// ProductTypeTERMINALAPI captures enum value "TERMINAL_API"
	ProductTypeTERMINALAPI ProductType = "TERMINAL_API"
)

// for schema
var productTypeEnum []interface{}

func init() {
	var res []ProductType
	if err := json.Unmarshal([]byte(`["TERMINAL_API"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productTypeEnum = append(productTypeEnum, v)
	}
}

func (m ProductType) validateProductTypeEnum(path, location string, value ProductType) error {
	if err := validate.Enum(path, location, value, productTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this product type
func (m ProductType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProductTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
