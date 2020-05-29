// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogCategory A category to which a `CatalogItem` belongs in the `Catalog` object model.
//
// swagger:model CatalogCategory
type CatalogCategory struct {

	// The category name. Searchable. This field has max length of 255 Unicode code points.
	Name string `json:"name,omitempty"`
}

// Validate validates this catalog category
func (m *CatalogCategory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogCategory) UnmarshalBinary(b []byte) error {
	var res CatalogCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
