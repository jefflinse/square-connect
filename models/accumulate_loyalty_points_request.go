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

// AccumulateLoyaltyPointsRequest A request to accumulate points for a purchase.
//
// swagger:model AccumulateLoyaltyPointsRequest
type AccumulateLoyaltyPointsRequest struct {

	// The points to add to the account.
	// If you are using the Orders API to manage orders, you
	// specify the order ID. Otherwise, specify the
	// points to add.
	// Required: true
	AccumulatePoints *LoyaltyEventAccumulatePoints `json:"accumulate_points"`

	// A unique string that identifies the `AccumulateLoyaltyPoints` request.
	// Keys can be any valid string but must be unique for every request.
	// Required: true
	// Max Length: 128
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`

	// The `location` where the purchase was made.
	// Required: true
	LocationID *string `json:"location_id"`
}

// Validate validates this accumulate loyalty points request
func (m *AccumulateLoyaltyPointsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccumulatePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
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

func (m *AccumulateLoyaltyPointsRequest) validateAccumulatePoints(formats strfmt.Registry) error {

	if err := validate.Required("accumulate_points", "body", m.AccumulatePoints); err != nil {
		return err
	}

	if m.AccumulatePoints != nil {
		if err := m.AccumulatePoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accumulate_points")
			}
			return err
		}
	}

	return nil
}

func (m *AccumulateLoyaltyPointsRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", string(*m.IdempotencyKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", string(*m.IdempotencyKey), 128); err != nil {
		return err
	}

	return nil
}

func (m *AccumulateLoyaltyPointsRequest) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccumulateLoyaltyPointsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccumulateLoyaltyPointsRequest) UnmarshalBinary(b []byte) error {
	var res AccumulateLoyaltyPointsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}