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

// CatalogItemModifierListInfo Options to control the properties of a `CatalogModifierList` applied to a `CatalogItem` instance.
//
// swagger:model CatalogItemModifierListInfo
type CatalogItemModifierListInfo struct {

	// If `true`, enable this `CatalogModifierList`. The default value is `true`.
	Enabled bool `json:"enabled,omitempty"`

	// If 0 or larger, the largest number of `CatalogModifier`s that can be selected from this `CatalogModifierList`.
	MaxSelectedModifiers int64 `json:"max_selected_modifiers,omitempty"`

	// If 0 or larger, the smallest number of `CatalogModifier`s that must be selected from this `CatalogModifierList`.
	MinSelectedModifiers int64 `json:"min_selected_modifiers,omitempty"`

	// The ID of the `CatalogModifierList` controlled by this `CatalogModifierListInfo`.
	// Required: true
	// Min Length: 1
	ModifierListID *string `json:"modifier_list_id"`

	// A set of `CatalogModifierOverride` objects that override whether a given `CatalogModifier` is enabled by default.
	ModifierOverrides []*CatalogModifierOverride `json:"modifier_overrides"`
}

// Validate validates this catalog item modifier list info
func (m *CatalogItemModifierListInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifierListID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifierOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItemModifierListInfo) validateModifierListID(formats strfmt.Registry) error {

	if err := validate.Required("modifier_list_id", "body", m.ModifierListID); err != nil {
		return err
	}

	if err := validate.MinLength("modifier_list_id", "body", *m.ModifierListID, 1); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemModifierListInfo) validateModifierOverrides(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifierOverrides) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifierOverrides); i++ {
		if swag.IsZero(m.ModifierOverrides[i]) { // not required
			continue
		}

		if m.ModifierOverrides[i] != nil {
			if err := m.ModifierOverrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifier_overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this catalog item modifier list info based on the context it is used
func (m *CatalogItemModifierListInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModifierOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItemModifierListInfo) contextValidateModifierOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifierOverrides); i++ {

		if m.ModifierOverrides[i] != nil {
			if err := m.ModifierOverrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifier_overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItemModifierListInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemModifierListInfo) UnmarshalBinary(b []byte) error {
	var res CatalogItemModifierListInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
