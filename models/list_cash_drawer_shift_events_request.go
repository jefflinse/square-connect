// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListCashDrawerShiftEventsRequest list cash drawer shift events request
//
// swagger:model ListCashDrawerShiftEventsRequest
type ListCashDrawerShiftEventsRequest struct {

	// Opaque cursor for fetching the next page of results.
	Cursor string `json:"cursor,omitempty"`

	// Number of resources to be returned in a page of results (200 by
	// default, 1000 max).
	// Maximum: 1000
	Limit int64 `json:"limit,omitempty"`

	// The ID of the location to list cash drawer shifts for.
	// Required: true
	// Min Length: 1
	LocationID *string `json:"location_id"`
}

// Validate validates this list cash drawer shift events request
func (m *ListCashDrawerShiftEventsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListCashDrawerShiftEventsRequest) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *ListCashDrawerShiftEventsRequest) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	if err := validate.MinLength("location_id", "body", string(*m.LocationID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListCashDrawerShiftEventsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCashDrawerShiftEventsRequest) UnmarshalBinary(b []byte) error {
	var res ListCashDrawerShiftEventsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
