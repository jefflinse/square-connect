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

// SearchShiftsResponse The response to a request for `Shift` objects. Contains
// the requested `Shift` objects. May contain a set of `Error` objects if
// the request resulted in errors.
// Example: {"shifts":[{"breaks":[{"break_type_id":"REGS1EQR1TPZ5","end_at":"2019-01-21T06:11:00-05:00","expected_duration":"PT10M","id":"SJW7X6AKEJQ65","is_paid":true,"name":"Tea Break","start_at":"2019-01-21T06:11:00-05:00"}],"created_at":"2019-01-24T01:12:03Z","employee_id":"ormj0jJJZ5OZIzxrZYJI","end_at":"2019-01-21T13:11:00-05:00","id":"X714F3HA6D1PT","location_id":"PAA1RJZZKXBFG","start_at":"2019-01-21T03:11:00-05:00","status":"CLOSED","team_member_id":"ormj0jJJZ5OZIzxrZYJI","timezone":"America/New_York","updated_at":"2019-02-07T22:21:08Z","version":6,"wage":{"hourly_rate":{"amount":1100,"currency":"USD"},"title":"Barista"}},{"breaks":[{"break_type_id":"WQX00VR99F53J","end_at":"2019-01-23T14:40:00-05:00","expected_duration":"PT10M","id":"BKS6VR7WR748A","is_paid":true,"name":"Tea Break","start_at":"2019-01-23T14:30:00-05:00"},{"break_type_id":"P6Q468ZFRN1AC","end_at":"2019-01-22T12:44:00-05:00","expected_duration":"PT15M","id":"BQFEZSHFZSC51","is_paid":false,"name":"Coffee Break","start_at":"2019-01-22T12:30:00-05:00"}],"created_at":"2019-01-23T23:32:45Z","employee_id":"33fJchumvVdJwxV0H6L9","end_at":"2019-01-22T13:02:00-05:00","id":"GDHYBZYWK0P2V","location_id":"PAA1RJZZKXBFG","start_at":"2019-01-22T12:02:00-05:00","status":"CLOSED","team_member_id":"33fJchumvVdJwxV0H6L9","timezone":"America/New_York","updated_at":"2019-01-24T00:56:25Z","version":16,"wage":{"hourly_rate":{"amount":1600,"currency":"USD"},"title":"Cook"}}]}
//
// swagger:model SearchShiftsResponse
type SearchShiftsResponse struct {

	// Opaque cursor for fetching the next page.
	Cursor string `json:"cursor,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// Shifts
	Shifts []*Shift `json:"shifts"`
}

// Validate validates this search shifts response
func (m *SearchShiftsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShifts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchShiftsResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *SearchShiftsResponse) validateShifts(formats strfmt.Registry) error {
	if swag.IsZero(m.Shifts) { // not required
		return nil
	}

	for i := 0; i < len(m.Shifts); i++ {
		if swag.IsZero(m.Shifts[i]) { // not required
			continue
		}

		if m.Shifts[i] != nil {
			if err := m.Shifts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shifts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search shifts response based on the context it is used
func (m *SearchShiftsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShifts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchShiftsResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SearchShiftsResponse) contextValidateShifts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Shifts); i++ {

		if m.Shifts[i] != nil {
			if err := m.Shifts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shifts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchShiftsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchShiftsResponse) UnmarshalBinary(b []byte) error {
	var res SearchShiftsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
