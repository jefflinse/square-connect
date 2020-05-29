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

// DisputedPayment The payment the cardholder disputed.
//
// swagger:model DisputedPayment
type DisputedPayment struct {

	// Square-generated unique ID of the payment being disputed.
	// Max Length: 192
	PaymentID string `json:"payment_id,omitempty"`
}

// Validate validates this disputed payment
func (m *DisputedPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisputedPayment) validatePaymentID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentID) { // not required
		return nil
	}

	if err := validate.MaxLength("payment_id", "body", string(m.PaymentID), 192); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DisputedPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisputedPayment) UnmarshalBinary(b []byte) error {
	var res DisputedPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
