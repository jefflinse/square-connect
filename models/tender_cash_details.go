// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenderCashDetails Represents the details of a tender with `type` `CASH`.
//
// swagger:model TenderCashDetails
type TenderCashDetails struct {

	// The total amount of cash provided by the buyer, before change is given.
	BuyerTenderedMoney *Money `json:"buyer_tendered_money,omitempty"`

	// The amount of change returned to the buyer.
	ChangeBackMoney *Money `json:"change_back_money,omitempty"`
}

// Validate validates this tender cash details
func (m *TenderCashDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuyerTenderedMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeBackMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenderCashDetails) validateBuyerTenderedMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.BuyerTenderedMoney) { // not required
		return nil
	}

	if m.BuyerTenderedMoney != nil {
		if err := m.BuyerTenderedMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buyer_tendered_money")
			}
			return err
		}
	}

	return nil
}

func (m *TenderCashDetails) validateChangeBackMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangeBackMoney) { // not required
		return nil
	}

	if m.ChangeBackMoney != nil {
		if err := m.ChangeBackMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("change_back_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenderCashDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenderCashDetails) UnmarshalBinary(b []byte) error {
	var res TenderCashDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}