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

// CreateTerminalCheckoutRequest create terminal checkout request
//
// swagger:model CreateTerminalCheckoutRequest
type CreateTerminalCheckoutRequest struct {

	// The checkout to create.
	// Required: true
	Checkout *TerminalCheckout `json:"checkout"`

	// A unique string that identifies this CreateCheckout request. Keys can be any valid string but
	// must be unique for every CreateCheckout request.
	//
	// See [Idempotency keys](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	// Required: true
	// Max Length: 64
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`
}

// Validate validates this create terminal checkout request
func (m *CreateTerminalCheckoutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTerminalCheckoutRequest) validateCheckout(formats strfmt.Registry) error {

	if err := validate.Required("checkout", "body", m.Checkout); err != nil {
		return err
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

func (m *CreateTerminalCheckoutRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", string(*m.IdempotencyKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", string(*m.IdempotencyKey), 64); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTerminalCheckoutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTerminalCheckoutRequest) UnmarshalBinary(b []byte) error {
	var res CreateTerminalCheckoutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
