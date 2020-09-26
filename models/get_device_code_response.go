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

// GetDeviceCodeResponse get device code response
//
// swagger:model GetDeviceCodeResponse
type GetDeviceCodeResponse struct {

	// The queried DeviceCode.
	DeviceCode *DeviceCode `json:"device_code,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this get device code response
func (m *GetDeviceCodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceCode(formats); err != nil {
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

func (m *GetDeviceCodeResponse) validateDeviceCode(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceCode) { // not required
		return nil
	}

	if m.DeviceCode != nil {
		if err := m.DeviceCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_code")
			}
			return err
		}
	}

	return nil
}

func (m *GetDeviceCodeResponse) validateErrors(formats strfmt.Registry) error {

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
func (m *GetDeviceCodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDeviceCodeResponse) UnmarshalBinary(b []byte) error {
	var res GetDeviceCodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}