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

// CancelBookingResponse cancel booking response
// Example: {"booking":{"appointment_segments":[{"duration_minutes":60,"service_variation_id":"RU3PBTZTK7DXZDQFCJHOK2MC","service_variation_version":1599775456731,"team_member_id":"TMXUrsBWWcHTt79t"}],"created_at":"2020-10-28T15:47:41Z","customer_id":"EX2QSVGTZN4K1E5QE1CBFNVQ8M","customer_note":"","id":"zkras0xv0xwswx","location_id":"LEQHH0YY8B42M","seller_note":"","start_at":"2020-11-26T13:00:00Z","status":"CANCELLED_BY_CUSTOMER","updated_at":"2020-10-28T15:49:25Z","version":1},"errors":[]}
//
// swagger:model CancelBookingResponse
type CancelBookingResponse struct {

	// The booking that was cancelled.
	Booking *Booking `json:"booking,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this cancel booking response
func (m *CancelBookingResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBooking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelBookingResponse) validateBooking(formats strfmt.Registry) error {
	if swag.IsZero(m.Booking) { // not required
		return nil
	}

	if m.Booking != nil {
		if err := m.Booking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("booking")
			}
			return err
		}
	}

	return nil
}

func (m *CancelBookingResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this cancel booking response based on the context it is used
func (m *CancelBookingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBooking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelBookingResponse) contextValidateBooking(ctx context.Context, formats strfmt.Registry) error {

	if m.Booking != nil {
		if err := m.Booking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("booking")
			}
			return err
		}
	}

	return nil
}

func (m *CancelBookingResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CancelBookingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelBookingResponse) UnmarshalBinary(b []byte) error {
	var res CancelBookingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}