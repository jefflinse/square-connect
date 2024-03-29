// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateTeamMemberRequest Represents a create request for a `TeamMember` object.
// Example: {"request_body":{"idempotency_key":"idempotency-key-0","team_member":{"assigned_locations":{"assignment_type":"EXPLICIT_LOCATIONS","location_ids":["YSGH2WBKG94QZ","GA2Y9HSJ8KRYT"]},"email_address":"joe_doe@gmail.com","family_name":"Doe","given_name":"Joe","phone_number":"+14159283333","reference_id":"reference_id_1","status":"ACTIVE"}}}
//
// swagger:model CreateTeamMemberRequest
type CreateTeamMemberRequest struct {

	// A unique string that identifies this CreateTeamMember request.
	// Keys can be any valid string but must be unique for every request.
	// See [Idempotency keys](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	// <br>
	// <b>Min Length 1    Max Length 45</b>
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// <b>Required</b> The data which will be used to create the `TeamMember` object.
	TeamMember *TeamMember `json:"team_member,omitempty"`
}

// Validate validates this create team member request
func (m *CreateTeamMemberRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTeamMemberRequest) validateTeamMember(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamMember) { // not required
		return nil
	}

	if m.TeamMember != nil {
		if err := m.TeamMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team_member")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create team member request based on the context it is used
func (m *CreateTeamMemberRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTeamMember(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTeamMemberRequest) contextValidateTeamMember(ctx context.Context, formats strfmt.Registry) error {

	if m.TeamMember != nil {
		if err := m.TeamMember.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team_member")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTeamMemberRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTeamMemberRequest) UnmarshalBinary(b []byte) error {
	var res CreateTeamMemberRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
