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

// CatalogQueryItemsForModifierList The query filter to return the items containing the specified modifier list IDs.
//
// swagger:model CatalogQueryItemsForModifierList
type CatalogQueryItemsForModifierList struct {

	// A set of `CatalogModifierList` IDs to be used to find associated `CatalogItem`s.
	// Required: true
	ModifierListIds []string `json:"modifier_list_ids"`
}

// Validate validates this catalog query items for modifier list
func (m *CatalogQueryItemsForModifierList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifierListIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogQueryItemsForModifierList) validateModifierListIds(formats strfmt.Registry) error {

	if err := validate.Required("modifier_list_ids", "body", m.ModifierListIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this catalog query items for modifier list based on context it is used
func (m *CatalogQueryItemsForModifierList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogQueryItemsForModifierList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogQueryItemsForModifierList) UnmarshalBinary(b []byte) error {
	var res CatalogQueryItemsForModifierList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
