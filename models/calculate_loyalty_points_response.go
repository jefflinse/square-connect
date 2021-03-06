// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CalculateLoyaltyPointsResponse A response that includes the points that the buyer can earn from
// a specified purchase.
//
// swagger:model CalculateLoyaltyPointsResponse
type CalculateLoyaltyPointsResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The points that the buyer can earn from a specified purchase.
	// Minimum: 0
	Points *int64 `json:"points,omitempty"`
}

// Validate validates this calculate loyalty points response
func (m *CalculateLoyaltyPointsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateLoyaltyPointsResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *CalculateLoyaltyPointsResponse) validatePoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if err := validate.MinimumInt("points", "body", int64(*m.Points), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CalculateLoyaltyPointsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculateLoyaltyPointsResponse) UnmarshalBinary(b []byte) error {
	var res CalculateLoyaltyPointsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
