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

// CreateTerminalCheckoutResponse create terminal checkout response
//
// swagger:model CreateTerminalCheckoutResponse
type CreateTerminalCheckoutResponse struct {

	// The created `TerminalCheckout`
	Checkout *TerminalCheckout `json:"checkout,omitempty"`

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this create terminal checkout response
func (m *CreateTerminalCheckoutResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckout(formats); err != nil {
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

func (m *CreateTerminalCheckoutResponse) validateCheckout(formats strfmt.Registry) error {

	if swag.IsZero(m.Checkout) { // not required
		return nil
	}

	if m.Checkout != nil {
		if err := m.Checkout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkout")
			}
			return err
		}
	}

	return nil
}

func (m *CreateTerminalCheckoutResponse) validateErrors(formats strfmt.Registry) error {

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
func (m *CreateTerminalCheckoutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTerminalCheckoutResponse) UnmarshalBinary(b []byte) error {
	var res CreateTerminalCheckoutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
