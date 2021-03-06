// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TeamMemberAssignedLocations An object that represents a team member's assignment to locations.
//
// swagger:model TeamMemberAssignedLocations
type TeamMemberAssignedLocations struct {

	// The current assignment type of the team member.
	// See [TeamMemberAssignedLocationsAssignmentType](#type-teammemberassignedlocationsassignmenttype) for possible values
	AssignmentType string `json:"assignment_type,omitempty"`

	// The locations that the team member is assigned to.
	LocationIds []string `json:"location_ids"`
}

// Validate validates this team member assigned locations
func (m *TeamMemberAssignedLocations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this team member assigned locations based on context it is used
func (m *TeamMemberAssignedLocations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamMemberAssignedLocations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamMemberAssignedLocations) UnmarshalBinary(b []byte) error {
	var res TeamMemberAssignedLocations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
