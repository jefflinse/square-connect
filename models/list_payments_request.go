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

// ListPaymentsRequest Retrieves a list of refunds taken by the account making the request.
//
// Max results per page: 100
//
// swagger:model ListPaymentsRequest
type ListPaymentsRequest struct {

	// Timestamp for the beginning of the reporting period, in RFC 3339 format.
	// Inclusive. Default: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`

	// The brand of `Payment` card. For example, `VISA`
	CardBrand string `json:"card_brand,omitempty"`

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this to retrieve the next set of results for the original query.
	//
	// See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`

	// Timestamp for the end of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time.
	EndTime string `json:"end_time,omitempty"`

	// The last 4 digits of `Payment` card.
	Last4 string `json:"last_4,omitempty"`

	// Limit results to the location supplied. By default, results are returned
	// for all locations associated with the merchant.
	LocationID string `json:"location_id,omitempty"`

	// The order in which results are listed.
	// - `ASC` - oldest to newest
	// - `DESC` - newest to oldest (default).
	SortOrder string `json:"sort_order,omitempty"`

	// The exact amount in the total_money for a `Payment`.
	// Minimum: 0
	Total *int64 `json:"total,omitempty"`
}

// Validate validates this list payments request
func (m *ListPaymentsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListPaymentsRequest) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if err := validate.MinimumInt("total", "body", int64(*m.Total), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListPaymentsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPaymentsRequest) UnmarshalBinary(b []byte) error {
	var res ListPaymentsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
