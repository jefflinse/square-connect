// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListRefundsRequest v1 list refunds request
//
// swagger:model V1ListRefundsRequest
type V1ListRefundsRequest struct {

	// A pagination cursor to retrieve the next set of results for your
	// original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`

	// The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`

	// The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time.
	EndTime string `json:"end_time,omitempty"`

	// The approximate number of refunds to return in a single response. Default: 100. Max: 200. Response may contain more results than the prescribed limit when refunds are made simultaneously to multiple tenders in a payment or when refunds are generated in an exchange to account for the value of returned goods.
	Limit int64 `json:"limit,omitempty"`

	// The order in which payments are listed in the response.
	// See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`
}

// Validate validates this v1 list refunds request
func (m *V1ListRefundsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 list refunds request based on context it is used
func (m *V1ListRefundsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListRefundsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListRefundsRequest) UnmarshalBinary(b []byte) error {
	var res V1ListRefundsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
