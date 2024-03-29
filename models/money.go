// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Money Represents an amount of money. `Money` fields can be signed or unsigned.
// Fields that do not explicitly define whether they are signed or unsigned are
// considered unsigned and can only hold positive amounts. For signed fields, the
// sign of the value indicates the purpose of the money transfer. See
// [Working with Monetary Amounts](/build-basics/working-with-monetary-amounts)
// for more information.
//
// swagger:model Money
type Money struct {

	// The amount of money, in the smallest denomination of the currency
	// indicated by `currency`. For example, when `currency` is `USD`, `amount` is
	// in cents. Monetary amounts can be positive or negative. See the specific
	// field description to determine the meaning of the sign in a particular case.
	Amount int64 `json:"amount,omitempty"`

	// The type of currency, in __ISO 4217 format__. For example, the currency
	// code for US dollars is `USD`.
	//
	// See `Currency` for possible values.
	// See [Currency](#type-currency) for possible values
	Currency string `json:"currency,omitempty"`
}

// Validate validates this money
func (m *Money) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this money based on context it is used
func (m *Money) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Money) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Money) UnmarshalBinary(b []byte) error {
	var res Money
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
