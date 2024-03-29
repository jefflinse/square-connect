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

// CardPaymentDetails Reflects the current status of a card payment.
//
// swagger:model CardPaymentDetails
type CardPaymentDetails struct {

	// For EMV payments, the cryptogram generated for the payment.
	// Max Length: 16
	ApplicationCryptogram string `json:"application_cryptogram,omitempty"`

	// For EMV payments, the application ID identifies the EMV application used for the payment.
	// Max Length: 32
	ApplicationIdentifier string `json:"application_identifier,omitempty"`

	// For EMV payments, the human-readable name of the EMV application used for the payment.
	// Max Length: 16
	ApplicationName string `json:"application_name,omitempty"`

	// The status code returned by the card issuer that describes the payment's
	// authorization status.
	// Max Length: 10
	AuthResultCode string `json:"auth_result_code,omitempty"`

	// The status code returned from the Address Verification System (AVS) check. The code can be
	// `AVS_ACCEPTED`, `AVS_REJECTED`, or `AVS_NOT_CHECKED`.
	// Max Length: 50
	AvsStatus string `json:"avs_status,omitempty"`

	// The credit card's non-confidential details.
	Card *Card `json:"card,omitempty"`

	// The timeline for card payments.
	CardPaymentTimeline *CardPaymentTimeline `json:"card_payment_timeline,omitempty"`

	// The status code returned from the Card Verification Value (CVV) check. The code can be
	// `CVV_ACCEPTED`, `CVV_REJECTED`, or `CVV_NOT_CHECKED`.
	// Max Length: 50
	CvvStatus string `json:"cvv_status,omitempty"`

	// Details about the device that took the payment.
	DeviceDetails *DeviceDetails `json:"device_details,omitempty"`

	// The method used to enter the card's details for the payment. The method can be
	// `KEYED`, `SWIPED`, `EMV`, `ON_FILE`, or `CONTACTLESS`.
	// Max Length: 50
	EntryMethod string `json:"entry_method,omitempty"`

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// Whether the card must be physically present for the payment to
	// be refunded.  If set to `true`, the card must be present.
	RefundRequiresCardPresence bool `json:"refund_requires_card_presence,omitempty"`

	// The statement description sent to the card networks.
	//
	// Note: The actual statement description varies and is likely to be truncated and appended with
	// additional information on a per issuer basis.
	// Max Length: 50
	StatementDescription string `json:"statement_description,omitempty"`

	// The card payment's current state. The state can be AUTHORIZED, CAPTURED, VOIDED, or
	// FAILED.
	// Max Length: 50
	Status string `json:"status,omitempty"`

	// For EMV payments, the method used to verify the cardholder's identity. The method can be
	// `PIN`, `SIGNATURE`, `PIN_AND_SIGNATURE`, `ON_DEVICE`, or `NONE`.
	// Max Length: 50
	VerificationMethod string `json:"verification_method,omitempty"`

	// For EMV payments, the results of the cardholder verification. The result can be
	// `SUCCESS`, `FAILURE`, or `UNKNOWN`.
	// Max Length: 50
	VerificationResults string `json:"verification_results,omitempty"`
}

// Validate validates this card payment details
func (m *CardPaymentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationCryptogram(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthResultCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvsStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardPaymentTimeline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCvvStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntryMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerificationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerificationResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardPaymentDetails) validateApplicationCryptogram(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationCryptogram) { // not required
		return nil
	}

	if err := validate.MaxLength("application_cryptogram", "body", m.ApplicationCryptogram, 16); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateApplicationIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationIdentifier) { // not required
		return nil
	}

	if err := validate.MaxLength("application_identifier", "body", m.ApplicationIdentifier, 32); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateApplicationName(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationName) { // not required
		return nil
	}

	if err := validate.MaxLength("application_name", "body", m.ApplicationName, 16); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateAuthResultCode(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthResultCode) { // not required
		return nil
	}

	if err := validate.MaxLength("auth_result_code", "body", m.AuthResultCode, 10); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateAvsStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AvsStatus) { // not required
		return nil
	}

	if err := validate.MaxLength("avs_status", "body", m.AvsStatus, 50); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateCard(formats strfmt.Registry) error {
	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {
		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

func (m *CardPaymentDetails) validateCardPaymentTimeline(formats strfmt.Registry) error {
	if swag.IsZero(m.CardPaymentTimeline) { // not required
		return nil
	}

	if m.CardPaymentTimeline != nil {
		if err := m.CardPaymentTimeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_payment_timeline")
			}
			return err
		}
	}

	return nil
}

func (m *CardPaymentDetails) validateCvvStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.CvvStatus) { // not required
		return nil
	}

	if err := validate.MaxLength("cvv_status", "body", m.CvvStatus, 50); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateDeviceDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceDetails) { // not required
		return nil
	}

	if m.DeviceDetails != nil {
		if err := m.DeviceDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_details")
			}
			return err
		}
	}

	return nil
}

func (m *CardPaymentDetails) validateEntryMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.EntryMethod) { // not required
		return nil
	}

	if err := validate.MaxLength("entry_method", "body", m.EntryMethod, 50); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CardPaymentDetails) validateStatementDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.StatementDescription) { // not required
		return nil
	}

	if err := validate.MaxLength("statement_description", "body", m.StatementDescription, 50); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MaxLength("status", "body", m.Status, 50); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateVerificationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.VerificationMethod) { // not required
		return nil
	}

	if err := validate.MaxLength("verification_method", "body", m.VerificationMethod, 50); err != nil {
		return err
	}

	return nil
}

func (m *CardPaymentDetails) validateVerificationResults(formats strfmt.Registry) error {
	if swag.IsZero(m.VerificationResults) { // not required
		return nil
	}

	if err := validate.MaxLength("verification_results", "body", m.VerificationResults, 50); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this card payment details based on the context it is used
func (m *CardPaymentDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCardPaymentTimeline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardPaymentDetails) contextValidateCard(ctx context.Context, formats strfmt.Registry) error {

	if m.Card != nil {
		if err := m.Card.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

func (m *CardPaymentDetails) contextValidateCardPaymentTimeline(ctx context.Context, formats strfmt.Registry) error {

	if m.CardPaymentTimeline != nil {
		if err := m.CardPaymentTimeline.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card_payment_timeline")
			}
			return err
		}
	}

	return nil
}

func (m *CardPaymentDetails) contextValidateDeviceDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceDetails != nil {
		if err := m.DeviceDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_details")
			}
			return err
		}
	}

	return nil
}

func (m *CardPaymentDetails) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardPaymentDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardPaymentDetails) UnmarshalBinary(b []byte) error {
	var res CardPaymentDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
