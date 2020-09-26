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

// V1PaymentSurcharge V1PaymentSurcharge
//
// swagger:model V1PaymentSurcharge
type V1PaymentSurcharge struct {

	// The amount of the surcharge as a Money object. Exactly one of rate or amount_money should be set.
	AmountMoney *V1Money `json:"amount_money,omitempty"`

	// The amount of money applied to the order as a result of the surcharge.
	AppliedMoney *V1Money `json:"applied_money,omitempty"`

	// The name of the surcharge.
	Name string `json:"name,omitempty"`

	// The amount of the surcharge as a percentage. The percentage is provided as a string representing the decimal equivalent of the percentage. For example, "0.7" corresponds to a 7% surcharge. Exactly one of rate or amount_money should be set.
	Rate string `json:"rate,omitempty"`

	// A Square-issued unique identifier associated with the surcharge.
	SurchargeID string `json:"surcharge_id,omitempty"`

	// Indicates whether the surcharge is taxable.
	Taxable bool `json:"taxable,omitempty"`

	// The list of taxes that should be applied to the surcharge.
	Taxes []*V1PaymentTax `json:"taxes"`

	// Indicates the source of the surcharge. For example, if it was applied as an automatic gratuity for a large group.
	// See [V1PaymentSurchargeType](#type-v1paymentsurchargetype) for possible values
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 payment surcharge
func (m *V1PaymentSurcharge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppliedMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PaymentSurcharge) validateAmountMoney(formats strfmt.Registry) error {

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

func (m *V1PaymentSurcharge) validateAppliedMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.AppliedMoney) { // not required
		return nil
	}

	if m.AppliedMoney != nil {
		if err := m.AppliedMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applied_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaymentSurcharge) validateTaxes(formats strfmt.Registry) error {

	if swag.IsZero(m.Taxes) { // not required
		return nil
	}

	for i := 0; i < len(m.Taxes); i++ {
		if swag.IsZero(m.Taxes[i]) { // not required
			continue
		}

		if m.Taxes[i] != nil {
			if err := m.Taxes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PaymentSurcharge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PaymentSurcharge) UnmarshalBinary(b []byte) error {
	var res V1PaymentSurcharge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}