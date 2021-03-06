// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateShiftResponse The response to the request to create a `Shift`. Contains
// the created `Shift` object. May contain a set of `Error` objects if
// the request resulted in errors.
// Example: {"shift":{"breaks":[{"break_type_id":"REGS1EQR1TPZ5","end_at":"2019-01-25T06:16:00-05:00","expected_duration":"PT5M","id":"X7GAQYVVRRG6P","is_paid":true,"name":"Tea Break","start_at":"2019-01-25T06:11:00-05:00"}],"created_at":"2019-02-28T00:39:02Z","employee_id":"ormj0jJJZ5OZIzxrZYJI","end_at":"2019-01-25T13:11:00-05:00","id":"K0YH4CV5462JB","location_id":"PAA1RJZZKXBFG","start_at":"2019-01-25T03:11:00-05:00","status":"CLOSED","team_member_id":"ormj0jJJZ5OZIzxrZYJI","timezone":"America/New_York","updated_at":"2019-02-28T00:39:02Z","version":1,"wage":{"hourly_rate":{"amount":1100,"currency":"USD"},"title":"Barista"}}}
//
// swagger:model CreateShiftResponse
type CreateShiftResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The `Shift` that was created on the request.
	Shift *Shift `json:"shift,omitempty"`
}

// Validate validates this create shift response
func (m *CreateShiftResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShift(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShiftResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *CreateShiftResponse) validateShift(formats strfmt.Registry) error {
	if swag.IsZero(m.Shift) { // not required
		return nil
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

// ContextValidate validate this create shift response based on the context it is used
func (m *CreateShiftResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShift(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShiftResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateShiftResponse) contextValidateShift(ctx context.Context, formats strfmt.Registry) error {

	if m.Shift != nil {
		if err := m.Shift.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shift")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateShiftResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShiftResponse) UnmarshalBinary(b []byte) error {
	var res CreateShiftResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
