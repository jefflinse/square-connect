// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListPaymentRefundsRequest Retrieves a list of refunds for the account making the request.
//
// The maximum results per page is 100.
// Example: {"request_body":{}}
//
// swagger:model ListPaymentRefundsRequest
type ListPaymentRefundsRequest struct {

	// The timestamp for the beginning of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`

	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor string `json:"cursor,omitempty"`

	// The timestamp for the end of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time.
	EndTime string `json:"end_time,omitempty"`

	// The maximum number of results to be returned in a single page.
	//
	// It is possible to receive fewer results than the specified limit on a given page.
	//
	// If the supplied value is greater than 100, no more than 100 results are returned.
	//
	// Default: 100
	Limit int64 `json:"limit,omitempty"`

	// Limit results to the location supplied. By default, results are returned
	// for all locations associated with the seller.
	LocationID string `json:"location_id,omitempty"`

	// The order in which results are listed:
	// - `ASC` - Oldest to newest.
	// - `DESC` - Newest to oldest (default).
	SortOrder string `json:"sort_order,omitempty"`

	// If provided, only refunds with the given source type are returned.
	// - `CARD` - List refunds only for payments where `CARD` was specified as the payment
	// source.
	//
	// Default: If omitted, refunds are returned regardless of the source type.
	SourceType string `json:"source_type,omitempty"`

	// If provided, only refunds with the given status are returned.
	// For a list of refund status values, see `PaymentRefund`.
	//
	// Default: If omitted, refunds are returned regardless of their status.
	Status string `json:"status,omitempty"`
}

// Validate validates this list payment refunds request
func (m *ListPaymentRefundsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list payment refunds request based on context it is used
func (m *ListPaymentRefundsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListPaymentRefundsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPaymentRefundsRequest) UnmarshalBinary(b []byte) error {
	var res ListPaymentRefundsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
