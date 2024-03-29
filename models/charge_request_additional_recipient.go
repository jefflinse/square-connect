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

// ChargeRequestAdditionalRecipient Represents an additional recipient (other than the merchant) entitled to a portion of the tender.
// Support is currently limited to USD, CAD and GBP currencies
//
// swagger:model ChargeRequestAdditionalRecipient
type ChargeRequestAdditionalRecipient struct {

	// The amount of money distributed to the recipient.
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// The description of the additional recipient.
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Description *string `json:"description"`

	// The location ID for a recipient (other than the merchant) receiving a portion of the tender.
	// Required: true
	// Max Length: 50
	// Min Length: 1
	LocationID *string `json:"location_id"`
}

// Validate validates this charge request additional recipient
func (m *ChargeRequestAdditionalRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChargeRequestAdditionalRecipient) validateAmountMoney(formats strfmt.Registry) error {

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

func (m *ChargeRequestAdditionalRecipient) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 100); err != nil {
		return err
	}

	return nil
}

func (m *ChargeRequestAdditionalRecipient) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	if err := validate.MinLength("location_id", "body", *m.LocationID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("location_id", "body", *m.LocationID, 50); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this charge request additional recipient based on the context it is used
func (m *ChargeRequestAdditionalRecipient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChargeRequestAdditionalRecipient) contextValidateAmountMoney(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ChargeRequestAdditionalRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChargeRequestAdditionalRecipient) UnmarshalBinary(b []byte) error {
	var res ChargeRequestAdditionalRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
