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

// ShiftSortField Enumerates the `Shift` fields to sort on.
//
// swagger:model ShiftSortField
type ShiftSortField string

const (

	// ShiftSortFieldSTARTAT captures enum value "START_AT"
	ShiftSortFieldSTARTAT ShiftSortField = "START_AT"

	// ShiftSortFieldENDAT captures enum value "END_AT"
	ShiftSortFieldENDAT ShiftSortField = "END_AT"

	// ShiftSortFieldCREATEDAT captures enum value "CREATED_AT"
	ShiftSortFieldCREATEDAT ShiftSortField = "CREATED_AT"

	// ShiftSortFieldUPDATEDAT captures enum value "UPDATED_AT"
	ShiftSortFieldUPDATEDAT ShiftSortField = "UPDATED_AT"
)

// for schema
var shiftSortFieldEnum []interface{}

func init() {
	var res []ShiftSortField
	if err := json.Unmarshal([]byte(`["START_AT","END_AT","CREATED_AT","UPDATED_AT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shiftSortFieldEnum = append(shiftSortFieldEnum, v)
	}
}

func (m ShiftSortField) validateShiftSortFieldEnum(path, location string, value ShiftSortField) error {
	if err := validate.Enum(path, location, value, shiftSortFieldEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this shift sort field
func (m ShiftSortField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateShiftSortFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}