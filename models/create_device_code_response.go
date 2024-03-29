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

// CreateDeviceCodeResponse create device code response
// Example: {"device_code":{"code":"EBCARJ","created_at":"2020-02-06T18:44:33.000Z","id":"B3Z6NAMYQSMTM","location_id":"B5E4484SHHNYH","name":"Counter 1","pair_by":"2020-02-06T18:49:33.000Z","product_type":"TERMINAL_API","status":"UNPAIRED","status_changed_at":"2020-02-06T18:44:33.000Z"}}
//
// swagger:model CreateDeviceCodeResponse
type CreateDeviceCodeResponse struct {

	// The created DeviceCode object containing the device code string.
	DeviceCode *DeviceCode `json:"device_code,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this create device code response
func (m *CreateDeviceCodeResponse) Validate(formats strfmt.Registry) error {
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

func (m *CreateDeviceCodeResponse) validateDeviceCode(formats strfmt.Registry) error {
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

func (m *CreateDeviceCodeResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this create device code response based on the context it is used
func (m *CreateDeviceCodeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDeviceCodeResponse) contextValidateDeviceCode(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceCode != nil {
		if err := m.DeviceCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_code")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDeviceCodeResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CreateDeviceCodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDeviceCodeResponse) UnmarshalBinary(b []byte) error {
	var res CreateDeviceCodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
