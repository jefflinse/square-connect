// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TeamMemberWage The hourly wage rate that a team member will earn on a `Shift` for doing the job
// specified by the `title` property of this object.
//
// swagger:model TeamMemberWage
type TeamMemberWage struct {

	// Can be a custom-set hourly wage or the calculated effective hourly
	// wage based on annual wage and hours worked per week.
	HourlyRate *Money `json:"hourly_rate,omitempty"`

	// UUID for this object.
	ID string `json:"id,omitempty"`

	// The `Team Member` that this wage is assigned to.
	TeamMemberID string `json:"team_member_id,omitempty"`

	// The job title that this wage relates to.
	Title string `json:"title,omitempty"`
}

// Validate validates this team member wage
func (m *TeamMemberWage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHourlyRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamMemberWage) validateHourlyRate(formats strfmt.Registry) error {

	if swag.IsZero(m.HourlyRate) { // not required
		return nil
	}

	if m.HourlyRate != nil {
		if err := m.HourlyRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourly_rate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamMemberWage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamMemberWage) UnmarshalBinary(b []byte) error {
	var res TeamMemberWage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
