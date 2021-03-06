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

// V1CreateRefundRequest V1CreateRefundRequest
//
// swagger:model V1CreateRefundRequest
type V1CreateRefundRequest struct {

	// The ID of the payment to refund. If you are creating a `PARTIAL`
	// refund for a split tender payment, instead provide the id of the
	// particular tender you want to refund.
	// Required: true
	PaymentID *string `json:"payment_id"`

	// The reason for the refund.
	// Required: true
	Reason *string `json:"reason"`

	// The amount of money to refund. Required only for PARTIAL refunds.
	RefundedMoney *V1Money `json:"refunded_money,omitempty"`

	// An optional key to ensure idempotence if you issue the same PARTIAL refund request more than once.
	RequestIdempotenceKey string `json:"request_idempotence_key,omitempty"`

	// TThe type of refund (FULL or PARTIAL).
	// See [V1CreateRefundRequestType](#type-v1createrefundrequesttype) for possible values
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this v1 create refund request
func (m *V1CreateRefundRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefundedMoney(formats); err != nil {
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

func (m *V1CreateRefundRequest) validatePaymentID(formats strfmt.Registry) error {

	if err := validate.Required("payment_id", "body", m.PaymentID); err != nil {
		return err
	}

	return nil
}

func (m *V1CreateRefundRequest) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *V1CreateRefundRequest) validateRefundedMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.RefundedMoney) { // not required
		return nil
	}

	if m.RefundedMoney != nil {
		if err := m.RefundedMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refunded_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateRefundRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateRefundRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateRefundRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateRefundRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
