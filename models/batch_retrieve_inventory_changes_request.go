// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BatchRetrieveInventoryChangesRequest batch retrieve inventory changes request
//
// swagger:model BatchRetrieveInventoryChangesRequest
type BatchRetrieveInventoryChangesRequest struct {

	// Filters results by `CatalogObject` ID.
	// Only applied when set. Default: unset.
	CatalogObjectIds []string `json:"catalog_object_ids"`

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for the original query.
	//
	// See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`

	// Filters results by `Location` ID. Only
	// applied when set. Default: unset.
	LocationIds []string `json:"location_ids"`

	// Filters `ADJUSTMENT` query results by
	// `InventoryState`. Only applied when set.
	// Default: unset.
	// See [InventoryState](#type-inventorystate) for possible values
	States []string `json:"states"`

	// Filters results by `InventoryChangeType`.
	// Default: [`PHYSICAL_COUNT`, `ADJUSTMENT`]. `TRANSFER` is not supported as
	// a filter.
	// See [InventoryChangeType](#type-inventorychangetype) for possible values
	Types []string `json:"types"`

	// Provided as an RFC 3339 timestamp. Returns results whose
	// `created_at` or `calculated_at` value is after the given time.
	// Default: UNIX epoch (`1970-01-01T00:00:00Z`).
	UpdatedAfter string `json:"updated_after,omitempty"`

	// Provided as an RFC 3339 timestamp. Returns results whose
	// `created_at` or `calculated_at` value is strictly before the given time.
	// Default: UNIX epoch (`1970-01-01T00:00:00Z`).
	UpdatedBefore string `json:"updated_before,omitempty"`
}

// Validate validates this batch retrieve inventory changes request
func (m *BatchRetrieveInventoryChangesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BatchRetrieveInventoryChangesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchRetrieveInventoryChangesRequest) UnmarshalBinary(b []byte) error {
	var res BatchRetrieveInventoryChangesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}