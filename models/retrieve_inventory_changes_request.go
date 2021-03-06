// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrieveInventoryChangesRequest retrieve inventory changes request
// Example: {"request_params":"?location_ids=\u0026cursor="}
//
// swagger:model RetrieveInventoryChangesRequest
type RetrieveInventoryChangesRequest struct {

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for the original query.
	//
	// See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`

	// The `Location` IDs to look up as a comma-separated
	// list. An empty list queries all locations.
	LocationIds string `json:"location_ids,omitempty"`
}

// Validate validates this retrieve inventory changes request
func (m *RetrieveInventoryChangesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this retrieve inventory changes request based on context it is used
func (m *RetrieveInventoryChangesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetrieveInventoryChangesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveInventoryChangesRequest) UnmarshalBinary(b []byte) error {
	var res RetrieveInventoryChangesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
