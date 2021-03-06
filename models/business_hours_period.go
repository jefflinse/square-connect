// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BusinessHoursPeriod Represents a period of time during which a business location is open.
//
// swagger:model BusinessHoursPeriod
type BusinessHoursPeriod struct {

	// The day of week for this time period.
	// See [DayOfWeek](#type-dayofweek) for possible values
	DayOfWeek string `json:"day_of_week,omitempty"`

	// The end time of a business hours period, specified in local time using partial-time
	// RFC 3339 format.
	EndLocalTime string `json:"end_local_time,omitempty"`

	// The start time of a business hours period, specified in local time using partial-time
	// RFC 3339 format.
	StartLocalTime string `json:"start_local_time,omitempty"`
}

// Validate validates this business hours period
func (m *BusinessHoursPeriod) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this business hours period based on context it is used
func (m *BusinessHoursPeriod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BusinessHoursPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessHoursPeriod) UnmarshalBinary(b []byte) error {
	var res BusinessHoursPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
