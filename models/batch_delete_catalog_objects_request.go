// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BatchDeleteCatalogObjectsRequest batch delete catalog objects request
//
// swagger:model BatchDeleteCatalogObjectsRequest
type BatchDeleteCatalogObjectsRequest struct {

	// The IDs of the CatalogObjects to be deleted. When an object is deleted, other objects
	// in the graph that depend on that object will be deleted as well (for example, deleting a
	// CatalogItem will delete its CatalogItemVariation.
	ObjectIds []string `json:"object_ids"`
}

// Validate validates this batch delete catalog objects request
func (m *BatchDeleteCatalogObjectsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BatchDeleteCatalogObjectsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchDeleteCatalogObjectsRequest) UnmarshalBinary(b []byte) error {
	var res BatchDeleteCatalogObjectsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
