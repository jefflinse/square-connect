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

// SearchTerminalRefundsRequest search terminal refunds request
// Example: {"request_body":{"limit":1,"query":{"filter":{"status":"COMPLETED"}}}}
//
// swagger:model SearchTerminalRefundsRequest
type SearchTerminalRefundsRequest struct {

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for the original query.
	Cursor string `json:"cursor,omitempty"`

	// Limit the number of results returned for a single request.
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// Query the terminal refunds based on given conditions and sort order. Calling
	// `SearchTerminalRefunds` without an explicitly query parameter will return all available
	// refunds with the default sort order.
	Query *TerminalRefundQuery `json:"query,omitempty"`
}

// Validate validates this search terminal refunds request
func (m *SearchTerminalRefundsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchTerminalRefundsRequest) validateLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", m.Limit, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", m.Limit, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *SearchTerminalRefundsRequest) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
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

// ContextValidate validate this search terminal refunds request based on the context it is used
func (m *SearchTerminalRefundsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchTerminalRefundsRequest) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {
		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchTerminalRefundsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchTerminalRefundsRequest) UnmarshalBinary(b []byte) error {
	var res SearchTerminalRefundsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
