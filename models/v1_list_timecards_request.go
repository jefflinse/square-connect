// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListTimecardsRequest v1 list timecards request
//
// swagger:model V1ListTimecardsRequest
type V1ListTimecardsRequest struct {

	// A pagination cursor to retrieve the next set of results for your
	// original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`

	// If filtering results by their clockin_time field, the beginning of the requested reporting period, in ISO 8601 format.
	BeginClockinTime string `json:"begin_clockin_time,omitempty"`

	// If filtering results by their clockout_time field, the beginning of the requested reporting period, in ISO 8601 format.
	BeginClockoutTime string `json:"begin_clockout_time,omitempty"`

	// If filtering results by their updated_at field, the beginning of the requested reporting period, in ISO 8601 format.
	BeginUpdatedAt string `json:"begin_updated_at,omitempty"`

	// If true, only deleted timecards are returned. If false, only valid timecards are returned.If you don't provide this parameter, both valid and deleted timecards are returned.
	Deleted bool `json:"deleted,omitempty"`

	// If provided, the endpoint returns only timecards for the employee with the specified ID.
	EmployeeID string `json:"employee_id,omitempty"`

	// If filtering results by their clockin_time field, the end of the requested reporting period, in ISO 8601 format.
	EndClockinTime string `json:"end_clockin_time,omitempty"`

	// If filtering results by their clockout_time field, the end of the requested reporting period, in ISO 8601 format.
	EndClockoutTime string `json:"end_clockout_time,omitempty"`

	// If filtering results by their updated_at field, the end of the requested reporting period, in ISO 8601 format.
	EndUpdatedAt string `json:"end_updated_at,omitempty"`

	// The maximum integer number of employee entities to return in a single response. Default 100, maximum 200.
	Limit int64 `json:"limit,omitempty"`

	// The order in which timecards are listed in the response, based on their created_at field.
	// See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`
}

// Validate validates this v1 list timecards request
func (m *V1ListTimecardsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListTimecardsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListTimecardsRequest) UnmarshalBinary(b []byte) error {
	var res V1ListTimecardsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
