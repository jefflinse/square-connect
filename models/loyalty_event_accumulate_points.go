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

// LoyaltyEventAccumulatePoints Provides metadata when the event `type` is `ACCUMULATE_POINTS`.
//
// swagger:model LoyaltyEventAccumulatePoints
type LoyaltyEventAccumulatePoints struct {

	// The ID of the `loyalty program`.
	// Max Length: 36
	LoyaltyProgramID string `json:"loyalty_program_id,omitempty"`

	// The ID of the `order` for which the buyer accumulated the points.
	// This field is returned only if the Orders API is used to process orders.
	OrderID string `json:"order_id,omitempty"`

	// The number of points accumulated by the event.
	// Minimum: 1
	Points int64 `json:"points,omitempty"`
}

// Validate validates this loyalty event accumulate points
func (m *LoyaltyEventAccumulatePoints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoyaltyProgramID(formats); err != nil {
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

func (m *LoyaltyEventAccumulatePoints) validateLoyaltyProgramID(formats strfmt.Registry) error {

	if swag.IsZero(m.LoyaltyProgramID) { // not required
		return nil
	}

	if err := validate.MaxLength("loyalty_program_id", "body", string(m.LoyaltyProgramID), 36); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyEventAccumulatePoints) validatePoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if err := validate.MinimumInt("points", "body", int64(m.Points), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyEventAccumulatePoints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyEventAccumulatePoints) UnmarshalBinary(b []byte) error {
	var res LoyaltyEventAccumulatePoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
