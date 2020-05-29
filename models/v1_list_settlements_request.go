// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListSettlementsRequest v1 list settlements request
//
// swagger:model V1ListSettlementsRequest
type V1ListSettlementsRequest struct {

	// A pagination cursor to retrieve the next set of results for your
	// original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`

	// The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`

	// The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time.
	EndTime string `json:"end_time,omitempty"`

	// The maximum number of settlements to return in a single response. This value cannot exceed 200.
	Limit int64 `json:"limit,omitempty"`

	// The order in which settlements are listed in the response.
	// See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`

	// Provide this parameter to retrieve only settlements with a particular status (SENT or FAILED).
	// See [V1ListSettlementsRequestStatus](#type-v1listsettlementsrequeststatus) for possible values
	Status string `json:"status,omitempty"`
}

// Validate validates this v1 list settlements request
func (m *V1ListSettlementsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListSettlementsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListSettlementsRequest) UnmarshalBinary(b []byte) error {
	var res V1ListSettlementsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
