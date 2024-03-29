// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListDisputesRequest Defines the request parameters for the `ListDisputes` endpoint.
// Example: {"request_body":{}}
//
// swagger:model ListDisputesRequest
type ListDisputesRequest struct {

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	// For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor string `json:"cursor,omitempty"`

	// The ID of the location for which to return a list of disputes. If not specified, the endpoint returns
	// all open disputes (the dispute status is not `INQUIRY_CLOSED`, `WON`, or `LOST`) associated with all locations.
	// Max Length: 40
	// Min Length: 1
	LocationID string `json:"location_id,omitempty"`

	// The dispute states to filter the result.
	// If not specified, the endpoint returns all open disputes (the dispute status is not `INQUIRY_CLOSED`, `WON`,
	// or `LOST`).
	// See [DisputeState](#type-disputestate) for possible values
	States []string `json:"states"`
}

// Validate validates this list disputes request
func (m *ListDisputesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDisputesRequest) validateLocationID(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.MinLength("location_id", "body", m.LocationID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("location_id", "body", m.LocationID, 40); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list disputes request based on context it is used
func (m *ListDisputesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListDisputesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDisputesRequest) UnmarshalBinary(b []byte) error {
	var res ListDisputesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
