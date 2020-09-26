// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetEmployeeWageResponse A response to a request to get an `EmployeeWage`. Contains
// the requested `EmployeeWage` objects. May contain a set of `Error` objects if
// the request resulted in errors.
//
// swagger:model GetEmployeeWageResponse
type GetEmployeeWageResponse struct {

	// The requested `EmployeeWage` object.
	EmployeeWage *EmployeeWage `json:"employee_wage,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this get employee wage response
func (m *GetEmployeeWageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployeeWage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetEmployeeWageResponse) validateEmployeeWage(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployeeWage) { // not required
		return nil
	}

	if m.EmployeeWage != nil {
		if err := m.EmployeeWage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee_wage")
			}
			return err
		}
	}

	return nil
}

func (m *GetEmployeeWageResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetEmployeeWageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEmployeeWageResponse) UnmarshalBinary(b []byte) error {
	var res GetEmployeeWageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}