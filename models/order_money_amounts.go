// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderMoneyAmounts A collection of various money amounts.
//
// swagger:model OrderMoneyAmounts
type OrderMoneyAmounts struct {

	// Money associated with discounts.
	DiscountMoney *Money `json:"discount_money,omitempty"`

	// Money associated with service charges.
	ServiceChargeMoney *Money `json:"service_charge_money,omitempty"`

	// Money associated with taxes.
	TaxMoney *Money `json:"tax_money,omitempty"`

	// Money associated with tips.
	TipMoney *Money `json:"tip_money,omitempty"`

	// Total money.
	TotalMoney *Money `json:"total_money,omitempty"`
}

// Validate validates this order money amounts
func (m *OrderMoneyAmounts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceChargeMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTipMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderMoneyAmounts) validateDiscountMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscountMoney) { // not required
		return nil
	}

	if m.DiscountMoney != nil {
		if err := m.DiscountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discount_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) validateServiceChargeMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceChargeMoney) { // not required
		return nil
	}

	if m.ServiceChargeMoney != nil {
		if err := m.ServiceChargeMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_charge_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) validateTaxMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxMoney) { // not required
		return nil
	}

	if m.TaxMoney != nil {
		if err := m.TaxMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) validateTipMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TipMoney) { // not required
		return nil
	}

	if m.TipMoney != nil {
		if err := m.TipMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tip_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) validateTotalMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalMoney) { // not required
		return nil
	}

	if m.TotalMoney != nil {
		if err := m.TotalMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_money")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order money amounts based on the context it is used
func (m *OrderMoneyAmounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceChargeMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTipMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderMoneyAmounts) contextValidateDiscountMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.DiscountMoney != nil {
		if err := m.DiscountMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discount_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) contextValidateServiceChargeMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceChargeMoney != nil {
		if err := m.ServiceChargeMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_charge_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) contextValidateTaxMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxMoney != nil {
		if err := m.TaxMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) contextValidateTipMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TipMoney != nil {
		if err := m.TipMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tip_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderMoneyAmounts) contextValidateTotalMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalMoney != nil {
		if err := m.TotalMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderMoneyAmounts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderMoneyAmounts) UnmarshalBinary(b []byte) error {
	var res OrderMoneyAmounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
