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

// LoyaltyProgramAccrualRule Defines an accrual rule, which is how buyers can earn points.
//
// swagger:model LoyaltyProgramAccrualRule
type LoyaltyProgramAccrualRule struct {

	// The type of the accrual rule that defines how buyers can earn points.
	// See [LoyaltyProgramAccrualRuleType](#type-loyaltyprogramaccrualruletype) for possible values
	// Required: true
	AccrualType *string `json:"accrual_type"`

	// The ID of the `catalog object` to purchase to earn the number of points defined by the
	// rule. This is either an item variation or a category, depending on the type. This is defined on
	// `ITEM_VARIATION` rules and `CATEGORY` rules.
	CatalogObjectID string `json:"catalog_object_id,omitempty"`

	// The number of points that
	// buyers earn based on the `accrual_type`.
	// Minimum: 1
	Points int64 `json:"points,omitempty"`

	// When the accrual rule is spend-based (`accrual_type` is `SPEND`),
	// this field indicates the amount that a buyer must spend
	// to earn the points. For example,
	// suppose the accrual rule is "earn 1 point for every $10 you spend".
	// Then, buyer earns a point for every $10 they spend. If
	// buyer spends $105, the buyer earns 10 points.
	SpendAmountMoney *Money `json:"spend_amount_money,omitempty"`

	// When the accrual rule is visit-based (`accrual_type` is `VISIT`),
	// this field indicates the minimum purchase required during the visit to
	// quality for the reward.
	VisitMinimumAmountMoney *Money `json:"visit_minimum_amount_money,omitempty"`
}

// Validate validates this loyalty program accrual rule
func (m *LoyaltyProgramAccrualRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccrualType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpendAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisitMinimumAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyProgramAccrualRule) validateAccrualType(formats strfmt.Registry) error {

	if err := validate.Required("accrual_type", "body", m.AccrualType); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyProgramAccrualRule) validatePoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if err := validate.MinimumInt("points", "body", int64(m.Points), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyProgramAccrualRule) validateSpendAmountMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.SpendAmountMoney) { // not required
		return nil
	}

	if m.SpendAmountMoney != nil {
		if err := m.SpendAmountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spend_amount_money")
			}
			return err
		}
	}

	return nil
}

func (m *LoyaltyProgramAccrualRule) validateVisitMinimumAmountMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.VisitMinimumAmountMoney) { // not required
		return nil
	}

	if m.VisitMinimumAmountMoney != nil {
		if err := m.VisitMinimumAmountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visit_minimum_amount_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyProgramAccrualRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyProgramAccrualRule) UnmarshalBinary(b []byte) error {
	var res LoyaltyProgramAccrualRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
