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

// LoyaltyAccountMapping Associates a loyalty account with the buyer's phone number.
// For more information, see
// [Loyalty Overview](/docs/loyalty/overview).
//
// swagger:model LoyaltyAccountMapping
type LoyaltyAccountMapping struct {

	// The timestamp when the mapping was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The Square-assigned ID of the mapping.
	// Max Length: 36
	ID string `json:"id,omitempty"`

	// The type of mapping.
	// See [LoyaltyAccountMappingType](#type-loyaltyaccountmappingtype) for possible values
	// Required: true
	Type *string `json:"type"`

	// The phone number, in E.164 format. For example, "+14155551111".
	// Required: true
	// Min Length: 1
	Value *string `json:"value"`
}

// Validate validates this loyalty account mapping
func (m *LoyaltyAccountMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyAccountMapping) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID, 36); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyAccountMapping) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyAccountMapping) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", *m.Value, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this loyalty account mapping based on context it is used
func (m *LoyaltyAccountMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyAccountMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyAccountMapping) UnmarshalBinary(b []byte) error {
	var res LoyaltyAccountMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
