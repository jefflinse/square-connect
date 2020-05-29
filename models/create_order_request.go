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

// CreateOrderRequest create order request
//
// swagger:model CreateOrderRequest
type CreateOrderRequest struct {

	// A value you specify that uniquely identifies this
	// order among orders you've created.
	//
	// If you're unsure whether a particular order was created successfully,
	// you can reattempt it with the same idempotency key without
	// worrying about creating duplicate orders.
	//
	// See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	// Max Length: 192
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The order to create. If this field is set, then the only other top-level field that can be
	// set is the idempotency_key.
	Order *Order `json:"order,omitempty"`
}

// Validate validates this create order request
func (m *CreateOrderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrderRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if swag.IsZero(m.IdempotencyKey) { // not required
		return nil
	}

	if err := validate.MaxLength("idempotency_key", "body", string(m.IdempotencyKey), 192); err != nil {
		return err
	}

	return nil
}

func (m *CreateOrderRequest) validateOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.Order) { // not required
		return nil
	}

	if m.Order != nil {
		if err := m.Order.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrderRequest) UnmarshalBinary(b []byte) error {
	var res CreateOrderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
