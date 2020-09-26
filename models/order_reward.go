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

// OrderReward Represents a reward that may be applied to an order if the necessary
// reward tier criteria are met. Rewards are created through the Loyalty API.
//
// swagger:model OrderReward
type OrderReward struct {

	// The identifier of the reward.
	// Required: true
	// Min Length: 1
	ID *string `json:"id"`

	// The identifier of the reward tier corresponding to this reward.
	// Required: true
	// Min Length: 1
	RewardTierID *string `json:"reward_tier_id"`
}

// Validate validates this order reward
func (m *OrderReward) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardTierID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderReward) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *OrderReward) validateRewardTierID(formats strfmt.Registry) error {

	if err := validate.Required("reward_tier_id", "body", m.RewardTierID); err != nil {
		return err
	}

	if err := validate.MinLength("reward_tier_id", "body", string(*m.RewardTierID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderReward) UnmarshalBinary(b []byte) error {
	var res OrderReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}