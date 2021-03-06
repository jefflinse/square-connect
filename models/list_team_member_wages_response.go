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

// ListTeamMemberWagesResponse The response to a request for a set of `TeamMemberWage` objects. Contains
// a set of `TeamMemberWage`.
// Example: {"cursor":"2fofTniCgT0yIPAq26kmk0YyFQJZfbWkh73OOnlTHmTAx13NgED","team_member_wages":[{"hourly_rate":{"amount":3250,"currency":"USD"},"id":"pXS3qCv7BERPnEGedM4S8mhm","team_member_id":"33fJchumvVdJwxV0H6L9","title":"Manager"},{"hourly_rate":{"amount":2600,"currency":"USD"},"id":"rZduCkzYDUVL3ovh1sQgbue6","team_member_id":"33fJchumvVdJwxV0H6L9","title":"Cook"},{"hourly_rate":{"amount":1600,"currency":"USD"},"id":"FxLbs5KpPUHa8wyt5ctjubDX","team_member_id":"33fJchumvVdJwxV0H6L9","title":"Barista"},{"hourly_rate":{"amount":1700,"currency":"USD"},"id":"vD1wCgijMDR3cX5TPnu7VXto","team_member_id":"33fJchumvVdJwxV0H6L9","title":"Cashier"}]}
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

// ContextValidate validate this list team member wages response based on the context it is used
func (m *ListTeamMemberWagesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamMemberWages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTeamMemberWagesResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ListTeamMemberWagesResponse) contextValidateTeamMemberWages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TeamMemberWages); i++ {

		if m.TeamMemberWages[i] != nil {
			if err := m.TeamMemberWages[i].ContextValidate(ctx, formats); err != nil {
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
