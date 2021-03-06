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

// CalculateLoyaltyPointsRequest A request to calculate the points that a buyer can earn from
// a specified purchase.
// Example: {"request_body":{"order_id":"RFZfrdtm3mhO1oGzf5Cx7fEMsmGZY"},"request_params":"?program_id=d619f755-2d17-41f3-990d-c04ecedd64dd"}
//
// swagger:model CalculateLoyaltyPointsRequest
type CalculateLoyaltyPointsRequest struct {

	// The `order` ID for which to calculate the points.
	// Specify this field if your application uses the Orders API to process orders.
	// Otherwise, specify the `transaction_amount`.
	OrderID string `json:"order_id,omitempty"`

	// The purchase amount for which to calculate the points.
	// Specify this field if your application does not use the Orders API to process orders.
	// Otherwise, specify the `order_id`.
	TransactionAmountMoney *Money `json:"transaction_amount_money,omitempty"`
}

// Validate validates this calculate loyalty points request
func (m *CalculateLoyaltyPointsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransactionAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateLoyaltyPointsRequest) validateTransactionAmountMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionAmountMoney) { // not required
		return nil
	}

	if m.TransactionAmountMoney != nil {
		if err := m.TransactionAmountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction_amount_money")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this calculate loyalty points request based on the context it is used
func (m *CalculateLoyaltyPointsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransactionAmountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateLoyaltyPointsRequest) contextValidateTransactionAmountMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TransactionAmountMoney != nil {
		if err := m.TransactionAmountMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction_amount_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CalculateLoyaltyPointsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculateLoyaltyPointsRequest) UnmarshalBinary(b []byte) error {
	var res CalculateLoyaltyPointsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
