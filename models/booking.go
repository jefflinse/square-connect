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
	"github.com/go-openapi/validate"
)

// Booking Represents a booking as a time-bound service contract for a seller's staff member to provide a specified service
// at a given location to a requesting customer in one or more appointment segments.
//
// swagger:model Booking
type Booking struct {

	// A list of appointment segments for this booking.
	AppointmentSegments []*AppointmentSegment `json:"appointment_segments"`

	// The timestamp specifying the creation time of this booking.
	CreatedAt string `json:"created_at,omitempty"`

	// The ID of the `Customer` object representing the customer attending this booking
	CustomerID string `json:"customer_id,omitempty"`

	// The free-text field for the customer to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a relevant `CatalogObject` instance.
	// Max Length: 4096
	CustomerNote string `json:"customer_note,omitempty"`

	// A unique ID of this object representing a booking.
	ID string `json:"id,omitempty"`

	// The ID of the `Location` object representing the location where the booked service is provided.
	LocationID string `json:"location_id,omitempty"`

	// The free-text field for the seller to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a specific `CatalogObject` instance.
	// This field should not be visible to customers.
	// Max Length: 4096
	SellerNote string `json:"seller_note,omitempty"`

	// The timestamp specifying the starting time of this booking.
	StartAt string `json:"start_at,omitempty"`

	// The status of the booking, describing where the booking stands with respect to the booking state machine.
	// See [BookingStatus](#type-bookingstatus) for possible values
	Status string `json:"status,omitempty"`

	// The timestamp specifying the most recent update time of this booking.
	UpdatedAt string `json:"updated_at,omitempty"`

	// The revision number for the booking used for optimistic concurrency.
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this booking
func (m *Booking) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointmentSegments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Booking) validateAppointmentSegments(formats strfmt.Registry) error {
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

func (m *Booking) validateCustomerNote(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerNote) { // not required
		return nil
	}

	if err := validate.MaxLength("customer_note", "body", m.CustomerNote, 4096); err != nil {
		return err
	}

	return nil
}

func (m *Booking) validateSellerNote(formats strfmt.Registry) error {
	if swag.IsZero(m.SellerNote) { // not required
		return nil
	}

	if err := validate.MaxLength("seller_note", "body", m.SellerNote, 4096); err != nil {
		return err
	}

	return nil
}

func (m *Booking) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", *m.Version, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this booking based on the context it is used
func (m *Booking) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppointmentSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Booking) contextValidateAppointmentSegments(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Booking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Booking) UnmarshalBinary(b []byte) error {
	var res Booking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
