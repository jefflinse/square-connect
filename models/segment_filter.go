// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SegmentFilter A query filter to search for appointment segments by.
//
// swagger:model SegmentFilter
type SegmentFilter struct {

	// The ID of the `CatalogItemVariation` representing the service booked in this segment.
	// Required: true
	// Min Length: 1
	ServiceVariationID *string `json:"service_variation_id"`

	// A query expression specifying which team members satisfy the condition. Supported expressions are
	// - `ANY`: include team members whose IDs match any member of the specified list.
	// - `NONE`: exclude team members whose IDs match members of the specified list.
	//
	// The `ALL` expression is not supported in the Bookings API.
	// When no expression is specified, any service-providing team member is eligible to fulfill the Booking.
	TeamMemberIDFilter *FilterValue `json:"team_member_id_filter,omitempty"`
}

// Validate validates this segment filter
func (m *SegmentFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceVariationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamMemberIDFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentFilter) validateServiceVariationID(formats strfmt.Registry) error {

	if err := validate.Required("service_variation_id", "body", m.ServiceVariationID); err != nil {
		return err
	}

	if err := validate.MinLength("service_variation_id", "body", *m.ServiceVariationID, 1); err != nil {
		return err
	}

	return nil
}

func (m *SegmentFilter) validateTeamMemberIDFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamMemberIDFilter) { // not required
		return nil
	}

	if m.TeamMemberIDFilter != nil {
		if err := m.TeamMemberIDFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team_member_id_filter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this segment filter based on the context it is used
func (m *SegmentFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTeamMemberIDFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentFilter) contextValidateTeamMemberIDFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.TeamMemberIDFilter != nil {
		if err := m.TeamMemberIDFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team_member_id_filter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SegmentFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentFilter) UnmarshalBinary(b []byte) error {
	var res SegmentFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
