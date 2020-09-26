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

// ListInvoicesRequest Describes a `ListInvoice` request.
//
// swagger:model ListInvoicesRequest
type ListInvoicesRequest struct {

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`

	// The maximum number of invoices to return (200 is the maximum `limit`).
	// If not provided, the server
	// uses a default limit of 100 invoices.
	Limit int64 `json:"limit,omitempty"`

	// The ID of the location for which to list invoices.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	LocationID *string `json:"location_id"`
}

// Validate validates this list invoices request
func (m *ListInvoicesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListInvoicesRequest) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	if err := validate.MinLength("location_id", "body", string(*m.LocationID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("location_id", "body", string(*m.LocationID), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListInvoicesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListInvoicesRequest) UnmarshalBinary(b []byte) error {
	var res ListInvoicesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}