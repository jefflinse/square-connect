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

// ListSubscriptionEventsRequest Defines parameters in a
// [ListSubscriptionEvents](#endpoint-subscriptions-listsubscriptionevents)
// endpoint request.
//
// swagger:model ListSubscriptionEventsRequest
type ListSubscriptionEventsRequest struct {

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for the original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`

	// The upper limit on the number of subscription events to return
	// in the response.
	//
	// Default: `200`
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`
}

// Validate validates this list subscription events request
func (m *ListSubscriptionEventsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListSubscriptionEventsRequest) validateLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", m.Limit, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list subscription events request based on context it is used
func (m *ListSubscriptionEventsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListSubscriptionEventsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListSubscriptionEventsRequest) UnmarshalBinary(b []byte) error {
	var res ListSubscriptionEventsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
