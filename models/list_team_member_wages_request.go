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

// ListTeamMemberWagesRequest A request for a set of `TeamMemberWage` objects
//
// swagger:model ListTeamMemberWagesRequest
type ListTeamMemberWagesRequest struct {

	// Pointer to the next page of Employee Wage results to fetch.
	Cursor string `json:"cursor,omitempty"`

	// Maximum number of Team Member Wages to return per page. Can range between
	// 1 and 200. The default is the maximum at 200.
	// Maximum: 200
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// Filter wages returned to only those that are associated with the
	// specified team member.
	TeamMemberID string `json:"team_member_id,omitempty"`
}

// Validate validates this list team member wages request
func (m *ListTeamMemberWagesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTeamMemberWagesRequest) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", int64(m.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 200, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListTeamMemberWagesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListTeamMemberWagesRequest) UnmarshalBinary(b []byte) error {
	var res ListTeamMemberWagesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
