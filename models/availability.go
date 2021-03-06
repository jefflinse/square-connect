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

// Availability Describes a slot available for booking, encapsulating appointment segments, the location and starting time.
//
// swagger:model Availability
type Availability struct {

	// The list of appointment segments available for booking
	AppointmentSegments []*AppointmentSegment `json:"appointment_segments"`

	// The ID of the location available for booking.
	LocationID string `json:"location_id,omitempty"`

	// The RFC-3339 timestamp specifying the beginning time of the slot available for booking.
	StartAt string `json:"start_at,omitempty"`
}

// Validate validates this availability
func (m *Availability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointmentSegments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Availability) validateAppointmentSegments(formats strfmt.Registry) error {
	if swag.IsZero(m.AppointmentSegments) { // not required
		return nil
	}

	for i := 0; i < len(m.AppointmentSegments); i++ {
		if swag.IsZero(m.AppointmentSegments[i]) { // not required
			continue
		}

		if m.AppointmentSegments[i] != nil {
			if err := m.AppointmentSegments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appointment_segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this availability based on the context it is used
func (m *Availability) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppointmentSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Availability) contextValidateAppointmentSegments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppointmentSegments); i++ {

		if m.AppointmentSegments[i] != nil {
			if err := m.AppointmentSegments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appointment_segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Availability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Availability) UnmarshalBinary(b []byte) error {
	var res Availability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
