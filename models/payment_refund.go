// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentRefund Represents a refund of a payment made using Square. Contains information about
// the original payment and the amount of money refunded.
//
// swagger:model PaymentRefund
type PaymentRefund struct {

	// The amount of money refunded. This amount is specified in the smallest denomination
	// of the applicable currency (for example, US dollar amounts are specified in cents).
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// The amount of money the application developer contributed to help cover the refunded amount.
	// This amount is specified in the smallest denomination of the applicable currency (for example,
	// US dollar amounts are specified in cents). For more information, see
	// [Working with Monetary Amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts).
	AppFeeMoney *Money `json:"app_fee_money,omitempty"`

	// The timestamp of when the refund was created, in RFC 3339 format.
	// Max Length: 32
	CreatedAt string `json:"created_at,omitempty"`

	// The unique ID for this refund, generated by Square.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The location ID associated with the payment this refund is attached to.
	// Max Length: 50
	LocationID string `json:"location_id,omitempty"`

	// The ID of the order associated with the refund.
	// Max Length: 192
	OrderID string `json:"order_id,omitempty"`

	// The ID of the payment associated with this refund.
	// Max Length: 192
	PaymentID string `json:"payment_id,omitempty"`

	// Processing fees and fee adjustments assessed by Square for this refund.
	ProcessingFee []*ProcessingFee `json:"processing_fee"`

	// The reason for the refund.
	// Max Length: 192
	Reason string `json:"reason,omitempty"`

	// The refund's status:
	// - `PENDING` - Awaiting approval.
	// - `COMPLETED` - Successfully completed.
	// - `REJECTED` - The refund was rejected.
	// - `FAILED` - An error occurred.
	// Max Length: 50
	Status string `json:"status,omitempty"`

	// The timestamp of when the refund was last updated, in RFC 3339 format.
	// Max Length: 32
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this payment refund
func (m *PaymentRefund) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppFeeMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRefund) validateAmountMoney(formats strfmt.Registry) error {

	if err := validate.Required("amount_money", "body", m.AmountMoney); err != nil {
		return err
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

func (m *PaymentRefund) validateAppFeeMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.AppFeeMoney) { // not required
		return nil
	}

	if m.AppFeeMoney != nil {
		if err := m.AppFeeMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRefund) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MaxLength("created_at", "body", m.CreatedAt, 32); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", *m.ID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", *m.ID, 255); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validateLocationID(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.MaxLength("location_id", "body", m.LocationID, 50); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validateOrderID(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderID) { // not required
		return nil
	}

	if err := validate.MaxLength("order_id", "body", m.OrderID, 192); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validatePaymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentID) { // not required
		return nil
	}

	if err := validate.MaxLength("payment_id", "body", m.PaymentID, 192); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validateProcessingFee(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingFee) { // not required
		return nil
	}

	for i := 0; i < len(m.ProcessingFee); i++ {
		if swag.IsZero(m.ProcessingFee[i]) { // not required
			continue
		}

		if m.ProcessingFee[i] != nil {
			if err := m.ProcessingFee[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processing_fee" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaymentRefund) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := validate.MaxLength("reason", "body", m.Reason, 192); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MaxLength("status", "body", m.Status, 50); err != nil {
		return err
	}

	return nil
}

func (m *PaymentRefund) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MaxLength("updated_at", "body", m.UpdatedAt, 32); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this payment refund based on the context it is used
func (m *PaymentRefund) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppFeeMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessingFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRefund) contextValidateAmountMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.AmountMoney != nil {
		if err := m.AmountMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount_money")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRefund) contextValidateAppFeeMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.AppFeeMoney != nil {
		if err := m.AppFeeMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRefund) contextValidateProcessingFee(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProcessingFee); i++ {

		if m.ProcessingFee[i] != nil {
			if err := m.ProcessingFee[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processing_fee" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRefund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRefund) UnmarshalBinary(b []byte) error {
	var res PaymentRefund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
