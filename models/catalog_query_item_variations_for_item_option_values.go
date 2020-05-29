// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogQueryItemVariationsForItemOptionValues catalog query item variations for item option values
//
// swagger:model CatalogQueryItemVariationsForItemOptionValues
type CatalogQueryItemVariationsForItemOptionValues struct {

	// A set of `CatalogItemOptionValue` IDs to be used to find associated
	// `CatalogItemVariation`s. All ItemVariations that contain all of the given
	// Item Option Values (in any order) will be returned.
	ItemOptionValueIds []string `json:"item_option_value_ids"`
}

// Validate validates this catalog query item variations for item option values
func (m *CatalogQueryItemVariationsForItemOptionValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogQueryItemVariationsForItemOptionValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogQueryItemVariationsForItemOptionValues) UnmarshalBinary(b []byte) error {
	var res CatalogQueryItemVariationsForItemOptionValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
