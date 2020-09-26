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

// LoyaltyEventOther Provides metadata when the event `type` is `OTHER`.
//
// swagger:model LoyaltyEventOther
type LoyaltyEventOther struct {

	// The Square-assigned ID of the `loyalty program`.
	// Required: true
	// Max Length: 36
	// Min Length: 1
	LoyaltyProgramID *string `json:"loyalty_program_id"`

	// The number of points added or removed.
	// Required: true
	Points *int64 `json:"points"`
}

// Validate validates this loyalty event other
func (m *LoyaltyEventOther) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoyaltyProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyEventOther) validateLoyaltyProgramID(formats strfmt.Registry) error {

	if err := validate.Required("loyalty_program_id", "body", m.LoyaltyProgramID); err != nil {
		return err
	}

	if err := validate.MinLength("loyalty_program_id", "body", string(*m.LoyaltyProgramID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("loyalty_program_id", "body", string(*m.LoyaltyProgramID), 36); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyEventOther) validatePoints(formats strfmt.Registry) error {

	if err := validate.Required("points", "body", m.Points); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyEventOther) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyEventOther) UnmarshalBinary(b []byte) error {
	var res LoyaltyEventOther
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
