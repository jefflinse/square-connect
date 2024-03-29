// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogObjectReference A reference to a Catalog object at a specific version. In general this is
// used as an entry point into a graph of catalog objects, where the objects exist
// at a specific version.
//
// swagger:model CatalogObjectReference
type CatalogObjectReference struct {

	// The version of the object.
	CatalogVersion int64 `json:"catalog_version,omitempty"`

	// The ID of the referenced object.
	ObjectID string `json:"object_id,omitempty"`
}

// Validate validates this catalog object reference
func (m *CatalogObjectReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog object reference based on context it is used
func (m *CatalogObjectReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogObjectReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogObjectReference) UnmarshalBinary(b []byte) error {
	var res CatalogObjectReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
