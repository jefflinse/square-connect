// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderQuantityUnit Contains the measurement unit for a quantity and a precision which
// specifies the number of digits after the decimal point for decimal quantities.
//
// swagger:model OrderQuantityUnit
type OrderQuantityUnit struct {

	// A `MeasurementUnit` that represents the
	// unit of measure for the quantity.
	MeasurementUnit *MeasurementUnit `json:"measurement_unit,omitempty"`

	// For non-integer quantities, represents the number of digits after the decimal point that are
	// recorded for this quantity.
	//
	// For example, a precision of 1 allows quantities like `"1.0"` and `"1.1"`, but not `"1.01"`.
	//
	// Min: 0. Max: 5.
	Precision int64 `json:"precision,omitempty"`
}

// Validate validates this order quantity unit
func (m *OrderQuantityUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeasurementUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderQuantityUnit) validateMeasurementUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.MeasurementUnit) { // not required
		return nil
	}

	if m.MeasurementUnit != nil {
		if err := m.MeasurementUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("measurement_unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderQuantityUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderQuantityUnit) UnmarshalBinary(b []byte) error {
	var res OrderQuantityUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}