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

// SearchInvoicesRequest Describes a `SearchInvoices` request.
//
// swagger:model SearchInvoicesRequest
type SearchInvoicesRequest struct {

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for your original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`

	// The maximum number of invoices to return (200 is the maximum `limit`).
	// If not provided, the server
	// uses a default limit of 100 invoices.
	Limit int64 `json:"limit,omitempty"`

	// Describes the query criteria for searching invoices.
	// Required: true
	Query *InvoiceQuery `json:"query"`
}

// Validate validates this search invoices request
func (m *SearchInvoicesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchInvoicesRequest) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchInvoicesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchInvoicesRequest) UnmarshalBinary(b []byte) error {
	var res SearchInvoicesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}