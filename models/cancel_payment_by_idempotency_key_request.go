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

// CancelPaymentByIdempotencyKeyRequest Specifies idempotency key of a payment to cancel.
//
// swagger:model CancelPaymentByIdempotencyKeyRequest
type CancelPaymentByIdempotencyKeyRequest struct {

	// `idempotency_key` identifying the payment to be canceled.
	// Required: true
	// Max Length: 45
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`
}

// Validate validates this cancel payment by idempotency key request
func (m *CancelPaymentByIdempotencyKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelPaymentByIdempotencyKeyRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", string(*m.IdempotencyKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", string(*m.IdempotencyKey), 45); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelPaymentByIdempotencyKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelPaymentByIdempotencyKeyRequest) UnmarshalBinary(b []byte) error {
	var res CancelPaymentByIdempotencyKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
