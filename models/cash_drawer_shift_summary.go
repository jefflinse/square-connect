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

// CashDrawerShiftSummary The summary of a closed cash drawer shift.
// This model contains only the money counted to start a cash drawer shift, counted
// at the end of the shift, and the amount that should be in the drawer at shift
// end based on summing all cash drawer shift events.
//
// swagger:model CashDrawerShiftSummary
type CashDrawerShiftSummary struct {

	// The shift close time in ISO 8601 format.
	ClosedAt string `json:"closed_at,omitempty"`

	// The amount of money found in the cash drawer at the end of the shift by
	// an auditing employee. The amount must be greater than or equal to zero.
	ClosedCashMoney *Money `json:"closed_cash_money,omitempty"`

	// An employee free-text description of a cash drawer shift.
	Description string `json:"description,omitempty"`

	// The shift end time in ISO 8601 format.
	EndedAt string `json:"ended_at,omitempty"`

	// The amount of money that should be in the cash drawer at the end of the
	// shift, based on the cash drawer events on the shift.
	// The amount is correct if all shift employees accurately recorded their
	// cash drawer shift events. Unrecorded events and events with the wrong amount
	// result in an incorrect expected_cash_money amount that can be negative.
	ExpectedCashMoney *Money `json:"expected_cash_money,omitempty"`

	// The shift unique ID.
	ID string `json:"id,omitempty"`

	// The shift start time in ISO 8601 format.
	OpenedAt string `json:"opened_at,omitempty"`

	// The amount of money in the cash drawer at the start of the shift. This
	// must be a positive amount.
	OpenedCashMoney *Money `json:"opened_cash_money,omitempty"`

	// The shift current state.
	// See [CashDrawerShiftState](#type-cashdrawershiftstate) for possible values
	State string `json:"state,omitempty"`
}

// Validate validates this cash drawer shift summary
func (m *CashDrawerShiftSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClosedCashMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedCashMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenedCashMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CashDrawerShiftSummary) validateClosedCashMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.ClosedCashMoney) { // not required
		return nil
	}

	if m.ClosedCashMoney != nil {
		if err := m.ClosedCashMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closed_cash_money")
			}
			return err
		}
	}

	return nil
}

func (m *CashDrawerShiftSummary) validateExpectedCashMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpectedCashMoney) { // not required
		return nil
	}

	if m.ExpectedCashMoney != nil {
		if err := m.ExpectedCashMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expected_cash_money")
			}
			return err
		}
	}

	return nil
}

func (m *CashDrawerShiftSummary) validateOpenedCashMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenedCashMoney) { // not required
		return nil
	}

	if m.OpenedCashMoney != nil {
		if err := m.OpenedCashMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opened_cash_money")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cash drawer shift summary based on the context it is used
func (m *CashDrawerShiftSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClosedCashMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpectedCashMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenedCashMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CashDrawerShiftSummary) contextValidateClosedCashMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.ClosedCashMoney != nil {
		if err := m.ClosedCashMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closed_cash_money")
			}
			return err
		}
	}

	return nil
}

func (m *CashDrawerShiftSummary) contextValidateExpectedCashMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.ExpectedCashMoney != nil {
		if err := m.ExpectedCashMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expected_cash_money")
			}
			return err
		}
	}

	return nil
}

func (m *CashDrawerShiftSummary) contextValidateOpenedCashMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.OpenedCashMoney != nil {
		if err := m.OpenedCashMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opened_cash_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CashDrawerShiftSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CashDrawerShiftSummary) UnmarshalBinary(b []byte) error {
	var res CashDrawerShiftSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
