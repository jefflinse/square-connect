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

// SearchTeamMembersQuery Represents the parameters in a search for `TeamMember` objects.
//
// swagger:model SearchTeamMembersQuery
type SearchTeamMembersQuery struct {

	// The options to filter by.
	Filter *SearchTeamMembersFilter `json:"filter,omitempty"`
}

// Validate validates this search team members query
func (m *SearchTeamMembersQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchTeamMembersQuery) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search team members query based on the context it is used
func (m *SearchTeamMembersQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchTeamMembersQuery) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.Filter != nil {
		if err := m.Filter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchTeamMembersQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchTeamMembersQuery) UnmarshalBinary(b []byte) error {
	var res SearchTeamMembersQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
