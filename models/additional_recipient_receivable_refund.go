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

// AdditionalRecipientReceivableRefund A refund of an [AdditionalRecipientReceivable](#type-additionalrecipientreceivable). This includes the ID of the additional recipient receivable associated to this object, as well as a reference to the [Refund](#type-refund) that created this receivable refund.
//
// swagger:model AdditionalRecipientReceivableRefund
type AdditionalRecipientReceivableRefund struct {

	// The amount of the refund. This will always be non-negative.
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// The time when the refund was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The receivable refund's unique ID, issued by Square payments servers.
	// Required: true
	// Min Length: 1
	ID *string `json:"id"`

	// The ID of the receivable that the refund was applied to.
	// Required: true
	// Min Length: 1
	ReceivableID *string `json:"receivable_id"`

	// The ID of the refund that is associated to this receivable refund.
	// Required: true
	// Min Length: 1
	RefundID *string `json:"refund_id"`

	// The ID of the location that created the receivable. This is the location ID on the associated transaction.
	// Required: true
	// Min Length: 1
	TransactionLocationID *string `json:"transaction_location_id"`
}

// Validate validates this additional recipient receivable refund
func (m *AdditionalRecipientReceivableRefund) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivableID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefundID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdditionalRecipientReceivableRefund) validateAmountMoney(formats strfmt.Registry) error {

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

func (m *AdditionalRecipientReceivableRefund) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalRecipientReceivableRefund) validateReceivableID(formats strfmt.Registry) error {

	if err := validate.Required("receivable_id", "body", m.ReceivableID); err != nil {
		return err
	}

	if err := validate.MinLength("receivable_id", "body", string(*m.ReceivableID), 1); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalRecipientReceivableRefund) validateRefundID(formats strfmt.Registry) error {

	if err := validate.Required("refund_id", "body", m.RefundID); err != nil {
		return err
	}

	if err := validate.MinLength("refund_id", "body", string(*m.RefundID), 1); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalRecipientReceivableRefund) validateTransactionLocationID(formats strfmt.Registry) error {

	if err := validate.Required("transaction_location_id", "body", m.TransactionLocationID); err != nil {
		return err
	}

	if err := validate.MinLength("transaction_location_id", "body", string(*m.TransactionLocationID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdditionalRecipientReceivableRefund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdditionalRecipientReceivableRefund) UnmarshalBinary(b []byte) error {
	var res AdditionalRecipientReceivableRefund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}