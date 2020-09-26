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

// BatchRetrieveCatalogObjectsRequest batch retrieve catalog objects request
//
// swagger:model BatchRetrieveCatalogObjectsRequest
type BatchRetrieveCatalogObjectsRequest struct {

	// If `true`, the response will include additional objects that are related to the
	// requested objects, as follows:
	//
	// If the `objects` field of the response contains a CatalogItem, its associated
	// CatalogCategory objects, CatalogTax objects, CatalogImage objects and
	// CatalogModifierLists will be returned in the `related_objects` field of the
	// response. If the `objects` field of the response contains a CatalogItemVariation,
	// its parent CatalogItem will be returned in the `related_objects` field of
	// the response.
	IncludeRelatedObjects bool `json:"include_related_objects,omitempty"`

	// The IDs of the CatalogObjects to be retrieved.
	// Required: true
	ObjectIds []string `json:"object_ids"`
}

// Validate validates this batch retrieve catalog objects request
func (m *BatchRetrieveCatalogObjectsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRetrieveCatalogObjectsRequest) validateObjectIds(formats strfmt.Registry) error {

	if err := validate.Required("object_ids", "body", m.ObjectIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchRetrieveCatalogObjectsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchRetrieveCatalogObjectsRequest) UnmarshalBinary(b []byte) error {
	var res BatchRetrieveCatalogObjectsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}