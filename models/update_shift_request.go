// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateShiftRequest A request to update a `Shift` object.
//
// swagger:model UpdateShiftRequest
type UpdateShiftRequest struct {

	// The updated `Shift` object.
	// Required: true
	Shift *Shift `json:"shift"`
}

// Validate validates this update shift request
func (m *UpdateShiftRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShift(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateShiftRequest) validateShift(formats strfmt.Registry) error {

	if err := validate.Required("shift", "body", m.Shift); err != nil {
		return err
	}

	if m.Shift != nil {
		if err := m.Shift.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shift")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateShiftRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateShiftRequest) UnmarshalBinary(b []byte) error {
	var res UpdateShiftRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}