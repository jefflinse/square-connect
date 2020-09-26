// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogV1ID An Items Connect V1 object ID along with its associated location ID.
//
// swagger:model CatalogV1Id
type CatalogV1ID struct {

	// The ID for an object in Connect V1, if different from its Connect V2 ID.
	CatalogV1ID string `json:"catalog_v1_id,omitempty"`

	// The ID of the `Location` this Connect V1 ID is associated with.
	LocationID string `json:"location_id,omitempty"`
}

// Validate validates this catalog v1 Id
func (m *CatalogV1ID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogV1ID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogV1ID) UnmarshalBinary(b []byte) error {
	var res CatalogV1ID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}