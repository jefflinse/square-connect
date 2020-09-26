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

// SearchOrdersSortField Specifies which timestamp to use to sort SearchOrder results.
//
// swagger:model SearchOrdersSortField
type SearchOrdersSortField string

const (

	// SearchOrdersSortFieldCREATEDAT captures enum value "CREATED_AT"
	SearchOrdersSortFieldCREATEDAT SearchOrdersSortField = "CREATED_AT"

	// SearchOrdersSortFieldUPDATEDAT captures enum value "UPDATED_AT"
	SearchOrdersSortFieldUPDATEDAT SearchOrdersSortField = "UPDATED_AT"

	// SearchOrdersSortFieldCLOSEDAT captures enum value "CLOSED_AT"
	SearchOrdersSortFieldCLOSEDAT SearchOrdersSortField = "CLOSED_AT"
)

// for schema
var searchOrdersSortFieldEnum []interface{}

func init() {
	var res []SearchOrdersSortField
	if err := json.Unmarshal([]byte(`["CREATED_AT","UPDATED_AT","CLOSED_AT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchOrdersSortFieldEnum = append(searchOrdersSortFieldEnum, v)
	}
}

func (m SearchOrdersSortField) validateSearchOrdersSortFieldEnum(path, location string, value SearchOrdersSortField) error {
	if err := validate.Enum(path, location, value, searchOrdersSortFieldEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this search orders sort field
func (m SearchOrdersSortField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSearchOrdersSortFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}