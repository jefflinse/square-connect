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

// BulkCreateTeamMembersRequest Represents a bulk create request for `TeamMember` objects.
//
// swagger:model BulkCreateTeamMembersRequest
type BulkCreateTeamMembersRequest struct {

	// The data which will be used to create the `TeamMember` objects. Each key is the `idempotency_key` that maps to the `CreateTeamMemberRequest`.
	// Required: true
	TeamMembers map[string]CreateTeamMemberRequest `json:"team_members"`
}

// Validate validates this bulk create team members request
func (m *BulkCreateTeamMembersRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCreateTeamMembersRequest) validateTeamMembers(formats strfmt.Registry) error {

	for k := range m.TeamMembers {

		if err := validate.Required("team_members"+"."+k, "body", m.TeamMembers[k]); err != nil {
			return err
		}
		if val, ok := m.TeamMembers[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkCreateTeamMembersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkCreateTeamMembersRequest) UnmarshalBinary(b []byte) error {
	var res BulkCreateTeamMembersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
