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

// CreateRefundRequest Defines the body parameters that can be included in
// a request to the [CreateRefund](#endpoint-createrefund) endpoint.
//
// Deprecated - recommend using [RefundPayment](#endpoint-refunds-refundpayment)
//
// swagger:model CreateRefundRequest
type CreateRefundRequest struct {

	// The amount of money to refund.
	//
	// Note that you specify the amount in the
	// __smallest denomination of the applicable currency__. For example, US dollar
	// amounts are specified in cents. See
	// [Working with monetary amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts) for details.
	//
	// This amount cannot exceed the amount that was originally charged to the
	// tender that corresponds to `tender_id`.
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// A value you specify that uniquely identifies this
	// refund among refunds you've created for the tender.
	//
	// If you're unsure whether a particular refund succeeded,
	// you can reattempt it with the same idempotency key without
	// worrying about duplicating the refund.
	//
	// See [Idempotency keys](#idempotencykeys) for more information.
	// Required: true
	// Max Length: 192
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`

	// A description of the reason for the refund.
	//
	// Default value: `Refund via API`
	// Max Length: 192
	Reason string `json:"reason,omitempty"`

	// The ID of the tender to refund.
	//
	// A ``Transaction`` has one or more `tenders` (i.e., methods
	// of payment) associated with it, and you refund each tender separately with
	// the Connect API.
	// Required: true
	// Max Length: 192
	// Min Length: 1
	TenderID *string `json:"tender_id"`
}

// Validate validates this create refund request
func (m *CreateRefundRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenderID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRefundRequest) validateAmountMoney(formats strfmt.Registry) error {

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

func (m *CreateRefundRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", string(*m.IdempotencyKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", string(*m.IdempotencyKey), 192); err != nil {
		return err
	}

	return nil
}

func (m *CreateRefundRequest) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := validate.MaxLength("reason", "body", string(m.Reason), 192); err != nil {
		return err
	}

	return nil
}

func (m *CreateRefundRequest) validateTenderID(formats strfmt.Registry) error {

	if err := validate.Required("tender_id", "body", m.TenderID); err != nil {
		return err
	}

	if err := validate.MinLength("tender_id", "body", string(*m.TenderID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("tender_id", "body", string(*m.TenderID), 192); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateRefundRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRefundRequest) UnmarshalBinary(b []byte) error {
	var res CreateRefundRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
