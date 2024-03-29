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

// SearchLoyaltyRewardsResponse A response that includes the loyalty rewards satisfying the search criteria.
// Example: {"rewards":[{"created_at":"2020-05-08T22:00:44Z","id":"d03f79f4-815f-3500-b851-cc1e68a457f9","loyalty_account_id":"5adcb100-07f1-4ee7-b8c6-6bb9ebc474bd","order_id":"PyATxhYLfsMqpVkcKJITPydgEYfZY","points":10,"redeemed_at":"2020-05-08T22:01:17Z","reward_tier_id":"e1b39225-9da5-43d1-a5db-782cdd8ad94f","status":"REDEEMED","updated_at":"2020-05-08T22:01:17Z"},{"created_at":"2020-05-08T21:55:42Z","id":"9f18ac21-233a-31c3-be77-b45840f5a810","loyalty_account_id":"5adcb100-07f1-4ee7-b8c6-6bb9ebc474bd","points":10,"redeemed_at":"2020-05-08T21:56:00Z","reward_tier_id":"e1b39225-9da5-43d1-a5db-782cdd8ad94f","status":"REDEEMED","updated_at":"2020-05-08T21:56:00Z"},{"created_at":"2020-05-01T21:49:54Z","id":"a8f43ebe-2ad6-3001-bdd5-7d7c2da08943","loyalty_account_id":"5adcb100-07f1-4ee7-b8c6-6bb9ebc474bd","order_id":"5NB69ZNh3FbsOs1ox43bh1xrli6YY","points":10,"reward_tier_id":"e1b39225-9da5-43d1-a5db-782cdd8ad94f","status":"DELETED","updated_at":"2020-05-08T21:55:10Z"},{"created_at":"2020-05-01T20:20:37Z","id":"a051254c-f840-3b45-8cf1-50bcd38ff92a","loyalty_account_id":"5adcb100-07f1-4ee7-b8c6-6bb9ebc474bd","order_id":"LQQ16znvi2VIUKPVhUfJefzr1eEZY","points":10,"reward_tier_id":"e1b39225-9da5-43d1-a5db-782cdd8ad94f","status":"ISSUED","updated_at":"2020-05-01T20:20:40Z"}]}
//
// swagger:model SearchLoyaltyRewardsResponse
type SearchLoyaltyRewardsResponse struct {

	// The pagination cursor to be used in a subsequent
	// request. If empty, this is the final response.
	Cursor string `json:"cursor,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The loyalty rewards that satisfy the search criteria.
	// These are returned in descending order by `updated_at`.
	Rewards []*LoyaltyReward `json:"rewards"`
}

// Validate validates this search loyalty rewards response
func (m *SearchLoyaltyRewardsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchLoyaltyRewardsResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *SearchLoyaltyRewardsResponse) validateRewards(formats strfmt.Registry) error {
	if swag.IsZero(m.Rewards) { // not required
		return nil
	}

	for i := 0; i < len(m.Rewards); i++ {
		if swag.IsZero(m.Rewards[i]) { // not required
			continue
		}

		if m.Rewards[i] != nil {
			if err := m.Rewards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search loyalty rewards response based on the context it is used
func (m *SearchLoyaltyRewardsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRewards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchLoyaltyRewardsResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SearchLoyaltyRewardsResponse) contextValidateRewards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rewards); i++ {

		if m.Rewards[i] != nil {
			if err := m.Rewards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchLoyaltyRewardsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchLoyaltyRewardsResponse) UnmarshalBinary(b []byte) error {
	var res SearchLoyaltyRewardsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
