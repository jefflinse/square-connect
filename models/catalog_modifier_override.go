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

// CatalogModifierOverride catalog modifier override
//
// swagger:model CatalogModifierOverride
type CatalogModifierOverride struct {

	// The ID of the `CatalogModifier` whose default behavior is being overridden.
	// Required: true
	// Min Length: 1
	ModifierID *string `json:"modifier_id"`

	// If `true`, this `CatalogModifier` should be selected by default for this `CatalogItem`.
	OnByDefault bool `json:"on_by_default,omitempty"`
}

// Validate validates this catalog modifier override
func (m *CatalogModifierOverride) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifierID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogModifierOverride) validateModifierID(formats strfmt.Registry) error {

	if err := validate.Required("modifier_id", "body", m.ModifierID); err != nil {
		return err
	}

	if err := validate.MinLength("modifier_id", "body", string(*m.ModifierID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogModifierOverride) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogModifierOverride) UnmarshalBinary(b []byte) error {
	var res CatalogModifierOverride
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}