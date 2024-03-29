// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PayOrderRequest Defines the fields that are included in requests to the
// [PayOrder](#endpoint-payorder) endpoint.
// Example: {"request_body":{"idempotency_key":"c043a359-7ad9-4136-82a9-c3f1d66dcbff","payment_ids":["EnZdNAlWCmfh6Mt5FMNST1o7taB","0LRiVlbXVwe8ozu4KbZxd12mvaB"]}}
//
// swagger:model PayOrderRequest
type PayOrderRequest struct {

	// A value you specify that uniquely identifies this request among requests you've sent. If
	// you're unsure whether a particular payment request was completed successfully, you can reattempt
	// it with the same idempotency key without worrying about duplicate payments.
	//
	// See [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency) for more information.
	// Required: true
	// Max Length: 192
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`

	// The version of the order being paid. If not supplied, the latest version will be paid.
	OrderVersion int64 `json:"order_version,omitempty"`

	// The IDs of the `payments` to collect.
	// The payment total must match the order total.
	PaymentIds []string `json:"payment_ids"`
}

// Validate validates this pay order request
func (m *PayOrderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayOrderRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", *m.IdempotencyKey, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", *m.IdempotencyKey, 192); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pay order request based on context it is used
func (m *PayOrderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PayOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayOrderRequest) UnmarshalBinary(b []byte) error {
	var res PayOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
