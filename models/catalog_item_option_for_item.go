// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogItemOptionForItem  A list of item option values that can be assigned to item variations.
// For example, a t-shirt item may offer a color option or a size option.
//
// swagger:model CatalogItemOptionForItem
type CatalogItemOptionForItem struct {

	// The unique id of the item option, used to form the dimensions of the item option matrix in a specified order.
	ItemOptionID string `json:"item_option_id,omitempty"`
}

// Validate validates this catalog item option for item
func (m *CatalogItemOptionForItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItemOptionForItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemOptionForItem) UnmarshalBinary(b []byte) error {
	var res CatalogItemOptionForItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
