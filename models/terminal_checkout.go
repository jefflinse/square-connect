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

// TerminalCheckout terminal checkout
//
// swagger:model TerminalCheckout
type TerminalCheckout struct {

	// The amount of money (including tax amount) that the Square Terminal device should try to collect.
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// Present if the status is CANCELED.
	// See [TerminalCheckoutCancelReason](#type-terminalcheckoutcancelreason) for possible values
	CancelReason string `json:"cancel_reason,omitempty"`

	// The time when the `TerminalCheckout` was created as an RFC 3339 timestamp.
	CreatedAt string `json:"created_at,omitempty"`

	// The duration as an RFC 3339 duration, after which the checkout will be automatically canceled.
	// TerminalCheckouts that are PENDING will be automatically CANCELED and have a cancellation reason
	// of “TIMED\_OUT”.
	//
	// Default: 5 minutes from creation
	//
	// Maximum: 5 minutes
	DeadlineDuration string `json:"deadline_duration,omitempty"`

	// Options to control the display and behavior of the Square Terminal device.
	// Required: true
	DeviceOptions *DeviceCheckoutOptions `json:"device_options"`

	// A unique ID for this `TerminalCheckout`
	ID string `json:"id,omitempty"`

	// An optional note to associate with the checkout, as well any payments used to complete the checkout.
	// Max Length: 250
	Note string `json:"note,omitempty"`

	// A list of payments created by this `TerminalCheckout`.
	PaymentIds []string `json:"payment_ids"`

	// An optional user-defined reference ID which can be used to associate
	// this TerminalCheckout to another entity in an external system. For example, an order
	// ID generated by a third-party shopping cart. Will also be associated with any payments
	// used to complete the checkout.
	// Max Length: 40
	ReferenceID string `json:"reference_id,omitempty"`

	// The status of the `TerminalCheckout`.
	// Options: PENDING, IN\_PROGRESS, CANCELED, COMPLETED
	Status string `json:"status,omitempty"`

	// The time when the `TerminalCheckout` was last updated as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this terminal checkout
func (m *TerminalCheckout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerminalCheckout) validateAmountMoney(formats strfmt.Registry) error {

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

func (m *TerminalCheckout) validateDeviceOptions(formats strfmt.Registry) error {

	if err := validate.Required("device_options", "body", m.DeviceOptions); err != nil {
		return err
	}

	if m.DeviceOptions != nil {
		if err := m.DeviceOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_options")
			}
			return err
		}
	}

	return nil
}

func (m *TerminalCheckout) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if err := validate.MaxLength("note", "body", string(m.Note), 250); err != nil {
		return err
	}

	return nil
}

func (m *TerminalCheckout) validateReferenceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceID) { // not required
		return nil
	}

	if err := validate.MaxLength("reference_id", "body", string(m.ReferenceID), 40); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerminalCheckout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerminalCheckout) UnmarshalBinary(b []byte) error {
	var res TerminalCheckout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}