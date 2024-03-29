// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrieveTeamMemberResponse Represents a response from a retrieve request, containing a `TeamMember` object or error messages.
// Example: {"team_member":{"assigned_locations":{"assignment_type":"EXPLICIT_LOCATIONS","location_ids":["GA2Y9HSJ8KRYT","YSGH2WBKG94QZ"]},"created_at":"2020-06-11T22:55:45.867Z","email_address":"joe_doe@gmail.com","family_name":"Doe","given_name":"Joe","id":"1yJlHapkseYnNPETIU1B","is_owner":false,"phone_number":"+14159283333","reference_id":"reference_id_1","status":"ACTIVE","updated_at":"2020-06-11T22:55:45.867Z"}}
//
// swagger:model RetrieveTeamMemberResponse
type RetrieveTeamMemberResponse struct {

	// The errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The successfully retrieved `TeamMember` object.
	TeamMember *TeamMember `json:"team_member,omitempty"`
}

// Validate validates this retrieve team member response
func (m *RetrieveTeamMemberResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveTeamMemberResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *RetrieveTeamMemberResponse) validateTeamMember(formats strfmt.Registry) error {
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

// ContextValidate validate this retrieve team member response based on the context it is used
func (m *RetrieveTeamMemberResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamMember(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveTeamMemberResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RetrieveTeamMemberResponse) contextValidateTeamMember(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RetrieveTeamMemberResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveTeamMemberResponse) UnmarshalBinary(b []byte) error {
	var res RetrieveTeamMemberResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
