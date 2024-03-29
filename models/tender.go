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

// Tender Represents a tender (i.e., a method of payment) used in a Square transaction.
//
// swagger:model Tender
type Tender struct {

	// Additional recipients (other than the merchant) receiving a portion of this tender.
	// For example, fees assessed on the purchase by a third party integration.
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients"`

	// The total amount of the tender, including `tip_money`. If the tender has a `payment_id`,
	// the `total_money` of the corresponding `Payment` will be equal to the
	// `amount_money` of the tender.
	AmountMoney *Money `json:"amount_money,omitempty"`

	// The details of the card tender.
	//
	// This value is present only if the value of `type` is `CARD`.
	CardDetails *TenderCardDetails `json:"card_details,omitempty"`

	// The details of the cash tender.
	//
	// This value is present only if the value of `type` is `CASH`.
	CashDetails *TenderCashDetails `json:"cash_details,omitempty"`

	// The timestamp for when the tender was created, in RFC 3339 format.
	// Max Length: 32
	CreatedAt string `json:"created_at,omitempty"`

	// If the tender is associated with a customer or represents a customer's card on file,
	// this is the ID of the associated customer.
	// Max Length: 191
	CustomerID string `json:"customer_id,omitempty"`

	// The tender's unique ID.
	// Max Length: 192
	ID string `json:"id,omitempty"`

	// The ID of the transaction's associated location.
	// Max Length: 50
	LocationID string `json:"location_id,omitempty"`

	// An optional note associated with the tender at the time of payment.
	// Max Length: 500
	Note string `json:"note,omitempty"`

	// The ID of the `Payment` that corresponds to this tender.
	// This value is only present for payments created with the v2 Payments API.
	// Max Length: 192
	PaymentID string `json:"payment_id,omitempty"`

	// The amount of any Square processing fees applied to the tender.
	//
	// This field is not immediately populated when a new transaction is created.
	// It is usually available after about ten seconds.
	ProcessingFeeMoney *Money `json:"processing_fee_money,omitempty"`

	// The tip's amount of the tender.
	TipMoney *Money `json:"tip_money,omitempty"`

	// The ID of the tender's associated transaction.
	// Max Length: 192
	TransactionID string `json:"transaction_id,omitempty"`

	// The type of tender, such as `CARD` or `CASH`.
	// See [TenderType](#type-tendertype) for possible values
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this tender
func (m *Tender) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingFeeMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTipMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tender) validateAdditionalRecipients(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalRecipients) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalRecipients); i++ {
		if swag.IsZero(m.AdditionalRecipients[i]) { // not required
			continue
		}

		if m.AdditionalRecipients[i] != nil {
			if err := m.AdditionalRecipients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tender) validateAmountMoney(formats strfmt.Registry) error {
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

func (m *Tender) validateCardDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.CardDetails) { // not required
		return nil
	}

	if m.CardDetails != nil {
		if err := m.CardDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_details")
			}
			return err
		}
	}

	return nil
}

func (m *Tender) validateCashDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.CashDetails) { // not required
		return nil
	}

	if m.CashDetails != nil {
		if err := m.CashDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cash_details")
			}
			return err
		}
	}

	return nil
}

func (m *Tender) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.MaxLength("created_at", "body", m.CreatedAt, 32); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validateCustomerID(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.MaxLength("customer_id", "body", m.CustomerID, 191); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID, 192); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validateLocationID(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.MaxLength("location_id", "body", m.LocationID, 50); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if err := validate.MaxLength("note", "body", m.Note, 500); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validatePaymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentID) { // not required
		return nil
	}

	if err := validate.MaxLength("payment_id", "body", m.PaymentID, 192); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validateProcessingFeeMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingFeeMoney) { // not required
		return nil
	}

	if m.ProcessingFeeMoney != nil {
		if err := m.ProcessingFeeMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processing_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *Tender) validateTipMoney(formats strfmt.Registry) error {
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

func (m *Tender) validateTransactionID(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionID) { // not required
		return nil
	}

	if err := validate.MaxLength("transaction_id", "body", m.TransactionID, 192); err != nil {
		return err
	}

	return nil
}

func (m *Tender) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tender based on the context it is used
func (m *Tender) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalRecipients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAmountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCardDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCashDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessingFeeMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTipMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tender) contextValidateAdditionalRecipients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalRecipients); i++ {

		if m.AdditionalRecipients[i] != nil {
			if err := m.AdditionalRecipients[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tender) contextValidateAmountMoney(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Tender) contextValidateCardDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.CardDetails != nil {
		if err := m.CardDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_details")
			}
			return err
		}
	}

	return nil
}

func (m *Tender) contextValidateCashDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.CashDetails != nil {
		if err := m.CashDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cash_details")
			}
			return err
		}
	}

	return nil
}

func (m *Tender) contextValidateProcessingFeeMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessingFeeMoney != nil {
		if err := m.ProcessingFeeMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processing_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *Tender) contextValidateTipMoney(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Tender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tender) UnmarshalBinary(b []byte) error {
	var res Tender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
