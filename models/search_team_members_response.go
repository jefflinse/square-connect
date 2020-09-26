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

// SearchTeamMembersResponse Represents a response from a search request, containing a filtered list of `TeamMember` objects.
//
// swagger:model SearchTeamMembersResponse
type SearchTeamMembersResponse struct {

	// The opaque cursor for fetching the next page. Read about
	// [pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination) with Square APIs for more information.
	Cursor string `json:"cursor,omitempty"`

	// The errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The filtered list of `TeamMember` objects.
	TeamMembers []*TeamMember `json:"team_members"`
}

// Validate validates this search team members response
func (m *SearchTeamMembersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchTeamMembersResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *SearchTeamMembersResponse) validateTeamMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamMembers) { // not required
		return nil
	}

	for i := 0; i < len(m.TeamMembers); i++ {
		if swag.IsZero(m.TeamMembers[i]) { // not required
			continue
		}

		if m.TeamMembers[i] != nil {
			if err := m.TeamMembers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("team_members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchTeamMembersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchTeamMembersResponse) UnmarshalBinary(b []byte) error {
	var res SearchTeamMembersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}