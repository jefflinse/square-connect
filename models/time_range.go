// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeRange Represents a generic time range. The start and end values are
// represented in RFC 3339 format. Time ranges are customized to be
// inclusive or exclusive based on the needs of a particular endpoint.
// Refer to the relevant endpoint-specific documentation to determine
// how time ranges are handled.
//
// swagger:model TimeRange
type TimeRange struct {

	// A datetime value in RFC 3339 format indicating when the time range
	// ends.
	EndAt string `json:"end_at,omitempty"`

	// A datetime value in RFC 3339 format indicating when the time range
	// starts.
	StartAt string `json:"start_at,omitempty"`
}

// Validate validates this time range
func (m *TimeRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this time range based on context it is used
func (m *TimeRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeRange) UnmarshalBinary(b []byte) error {
	var res TimeRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
