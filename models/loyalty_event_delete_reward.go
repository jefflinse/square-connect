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

// LoyaltyEventDeleteReward Provides metadata when the event `type` is `DELETE_REWARD`.
//
// swagger:model LoyaltyEventDeleteReward
type LoyaltyEventDeleteReward struct {

	// The ID of the `loyalty program`.
	// Required: true
	// Max Length: 36
	// Min Length: 1
	LoyaltyProgramID *string `json:"loyalty_program_id"`

	// The number of points returned to the loyalty account.
	// Required: true
	// Minimum: 0
	Points *int64 `json:"points"`

	// The ID of the deleted `loyalty reward`.
	// This field is returned only if the event source is `LOYALTY_API`.
	// Max Length: 36
	RewardID string `json:"reward_id,omitempty"`
}

// Validate validates this loyalty event delete reward
func (m *LoyaltyEventDeleteReward) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoyaltyProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyEventDeleteReward) validateLoyaltyProgramID(formats strfmt.Registry) error {

	if err := validate.Required("loyalty_program_id", "body", m.LoyaltyProgramID); err != nil {
		return err
	}

	if err := validate.MinLength("loyalty_program_id", "body", string(*m.LoyaltyProgramID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("loyalty_program_id", "body", string(*m.LoyaltyProgramID), 36); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyEventDeleteReward) validatePoints(formats strfmt.Registry) error {

	if err := validate.Required("points", "body", m.Points); err != nil {
		return err
	}

	if err := validate.MinimumInt("points", "body", int64(*m.Points), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyEventDeleteReward) validateRewardID(formats strfmt.Registry) error {

	if swag.IsZero(m.RewardID) { // not required
		return nil
	}

	if err := validate.MaxLength("reward_id", "body", string(m.RewardID), 36); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyEventDeleteReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyEventDeleteReward) UnmarshalBinary(b []byte) error {
	var res LoyaltyEventDeleteReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
