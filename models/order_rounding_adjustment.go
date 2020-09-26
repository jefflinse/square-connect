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

// OrderRoundingAdjustment A rounding adjustment of the money being returned. Commonly used to apply Cash Rounding
// when the minimum unit of account is smaller than the lowest physical denomination of currency.
//
// swagger:model OrderRoundingAdjustment
type OrderRoundingAdjustment struct {

	// Actual rounding adjustment amount.
	AmountMoney *Money `json:"amount_money,omitempty"`

	// The name of the rounding adjustment from the original sale Order.
	Name string `json:"name,omitempty"`

	// Unique ID that identifies the rounding adjustment only within this order.
	// Max Length: 60
	UID string `json:"uid,omitempty"`
}

// Validate validates this order rounding adjustment
func (m *OrderRoundingAdjustment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderRoundingAdjustment) validateAmountMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.AmountMoney) { // not required
		return nil
	}

	if m.AmountMoney != nil {
		if err := m.AmountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderRoundingAdjustment) validateUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.MaxLength("uid", "body", string(m.UID), 60); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderRoundingAdjustment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderRoundingAdjustment) UnmarshalBinary(b []byte) error {
	var res OrderRoundingAdjustment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}