// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListTeamMemberWagesResponse The response to a request for a set of `TeamMemberWage` objects. Contains
// a set of `TeamMemberWage`.
//
// swagger:model ListTeamMemberWagesResponse
type ListTeamMemberWagesResponse struct {

	// Value supplied in the subsequent request to fetch the next next page
	// of Team Member Wage results.
	Cursor string `json:"cursor,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// A page of Team Member Wage results.
	TeamMemberWages []*TeamMemberWage `json:"team_member_wages"`
}

// Validate validates this list team member wages response
func (m *ListTeamMemberWagesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamMemberWages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTeamMemberWagesResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListTeamMemberWagesResponse) validateTeamMemberWages(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamMemberWages) { // not required
		return nil
	}

	for i := 0; i < len(m.TeamMemberWages); i++ {
		if swag.IsZero(m.TeamMemberWages[i]) { // not required
			continue
		}

		if m.TeamMemberWages[i] != nil {
			if err := m.TeamMemberWages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("team_member_wages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListTeamMemberWagesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListTeamMemberWagesResponse) UnmarshalBinary(b []byte) error {
	var res ListTeamMemberWagesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
