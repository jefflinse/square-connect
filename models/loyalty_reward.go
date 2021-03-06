// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoyaltyReward loyalty reward
//
// swagger:model LoyaltyReward
type LoyaltyReward struct {

	// The timestamp when the reward was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The Square-assigned ID of the loyalty reward.
	// Max Length: 36
	ID string `json:"id,omitempty"`

	// The Square-assigned ID of the `loyalty account` to which the reward belongs.
	// Required: true
	// Max Length: 36
	// Min Length: 1
	LoyaltyAccountID *string `json:"loyalty_account_id"`

	// The Square-assigned ID of the `order` to which the reward is attached.
	OrderID string `json:"order_id,omitempty"`

	// The number of loyalty points used for the reward.
	// Minimum: 1
	Points int64 `json:"points,omitempty"`

	// The timestamp when the reward was redeemed, in RFC 3339 format.
	RedeemedAt string `json:"redeemed_at,omitempty"`

	// The Square-assigned ID of the `reward tier` used to create the reward.
	// Required: true
	// Max Length: 36
	// Min Length: 1
	RewardTierID *string `json:"reward_tier_id"`

	// The status of a loyalty reward.
	// See [LoyaltyRewardStatus](#type-loyaltyrewardstatus) for possible values
	Status string `json:"status,omitempty"`

	// The timestamp when the reward was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this loyalty reward
func (m *LoyaltyReward) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoyaltyAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
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

func (m *LoyaltyReward) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID, 36); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyReward) validateLoyaltyAccountID(formats strfmt.Registry) error {

	if err := validate.Required("loyalty_account_id", "body", m.LoyaltyAccountID); err != nil {
		return err
	}

	if err := validate.MinLength("loyalty_account_id", "body", *m.LoyaltyAccountID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("loyalty_account_id", "body", *m.LoyaltyAccountID, 36); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyReward) validatePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if err := validate.MinimumInt("points", "body", m.Points, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyReward) validateRewardTierID(formats strfmt.Registry) error {

	if err := validate.Required("reward_tier_id", "body", m.RewardTierID); err != nil {
		return err
	}

	if err := validate.MinLength("reward_tier_id", "body", *m.RewardTierID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("reward_tier_id", "body", *m.RewardTierID, 36); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this loyalty reward based on context it is used
func (m *LoyaltyReward) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyReward) UnmarshalBinary(b []byte) error {
	var res LoyaltyReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
