// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogItemOptionValue An enumerated value that can link a
// `CatalogItemVariation` to an item option as one of
// its item option values.
//
// swagger:model CatalogItemOptionValue
type CatalogItemOptionValue struct {

	// The HTML-supported hex color for the item option (e.g., "#ff8d4e85").
	// Only displayed if `show_colors` is enabled on the parent `ItemOption`. When
	// left unset, `color` defaults to white ("#ffffff") when `show_colors` is
	// enabled on the parent `ItemOption`.
	Color string `json:"color,omitempty"`

	// A human-readable description for the option value.
	Description string `json:"description,omitempty"`

	// Unique ID of the associated item option.
	ItemOptionID string `json:"item_option_id,omitempty"`

	// The number of `CatalogItemVariation`s that
	// currently make use of this Item Option value. Present only if `retrieve_counts`
	// was specified on the request used to retrieve the parent Item Option of this
	// value.
	//
	// Maximum: 100 counts.
	ItemVariationCount int64 `json:"item_variation_count,omitempty"`

	// Name of this item option value. Searchable.
	Name string `json:"name,omitempty"`

	// Determines where this option value appears in a list of option values.
	Ordinal int64 `json:"ordinal,omitempty"`
}

// Validate validates this catalog item option value
func (m *CatalogItemOptionValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItemOptionValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemOptionValue) UnmarshalBinary(b []byte) error {
	var res CatalogItemOptionValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}