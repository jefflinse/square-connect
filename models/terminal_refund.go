// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TerminalRefund terminal refund
//
// swagger:model TerminalRefund
type TerminalRefund struct {

	// The amount of money, inclusive of `tax_money`, that the `TerminalRefund` should return.
	// This value is limited to the amount taken in the original payment minus any completed or
	// pending refunds.
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// Present if the status is `CANCELED`.
	// See [ActionCancelReason](#type-actioncancelreason) for possible values
	CancelReason string `json:"cancel_reason,omitempty"`

	// The time when the `TerminalRefund` was created as an RFC 3339 timestamp.
	CreatedAt string `json:"created_at,omitempty"`

	// The duration as an RFC 3339 duration, after which the refund will be automatically canceled.
	// TerminalRefunds that are `PENDING` will be automatically `CANCELED` and have a cancellation reason
	// of `TIMED_OUT`
	//
	// Default: 5 minutes from creation
	//
	// Maximum: 5 minutes
	DeadlineDuration string `json:"deadline_duration,omitempty"`

	// The unique Id of the device intended for this `TerminalRefund`.
	// The Id can be retrieved from /v2/devices api.
	DeviceID string `json:"device_id,omitempty"`

	// A unique ID for this `TerminalRefund`
	// Max Length: 255
	// Min Length: 10
	ID string `json:"id,omitempty"`

	// The reference to the Square order id for the payment identified by the `payment_id`.
	OrderID string `json:"order_id,omitempty"`

	// Unique ID of the payment being refunded.
	// Required: true
	// Min Length: 1
	PaymentID *string `json:"payment_id"`

	// A description of the reason for the refund.
	// Note: maximum 192 characters
	// Max Length: 192
	Reason string `json:"reason,omitempty"`

	// The reference to the payment refund created by completing this `TerminalRefund`.
	RefundID string `json:"refund_id,omitempty"`

	// The status of the `TerminalRefund`.
	// Options: `PENDING`, `IN_PROGRESS`, `CANCELED`, `COMPLETED`
	Status string `json:"status,omitempty"`

	// The time when the `TerminalRefund` was last updated as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this terminal refund
func (m *TerminalRefund) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerminalRefund) validateAmountMoney(formats strfmt.Registry) error {

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

func (m *TerminalRefund) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 10); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", m.ID, 255); err != nil {
		return err
	}

	return nil
}

func (m *TerminalRefund) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.Required("payment_id", "body", m.PaymentID); err != nil {
		return err
	}

	if err := validate.MinLength("payment_id", "body", *m.PaymentID, 1); err != nil {
		return err
	}

	return nil
}

func (m *TerminalRefund) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := validate.MaxLength("reason", "body", m.Reason, 192); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this terminal refund based on the context it is used
func (m *TerminalRefund) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerminalRefund) contextValidateAmountMoney(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *TerminalRefund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerminalRefund) UnmarshalBinary(b []byte) error {
	var res TerminalRefund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
