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

// UpdateWageSettingResponse Represents a response from an update request, containing the updated `WageSetting` object
// or error messages.
//
// swagger:model UpdateWageSettingResponse
type UpdateWageSettingResponse struct {

	// The errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The successfully updated `WageSetting` object.
	WageSetting *WageSetting `json:"wage_setting,omitempty"`
}

// Validate validates this update wage setting response
func (m *UpdateWageSettingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWageSetting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateWageSettingResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *UpdateWageSettingResponse) validateWageSetting(formats strfmt.Registry) error {

	if swag.IsZero(m.WageSetting) { // not required
		return nil
	}

	if m.WageSetting != nil {
		if err := m.WageSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wage_setting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateWageSettingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateWageSettingResponse) UnmarshalBinary(b []byte) error {
	var res UpdateWageSettingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
