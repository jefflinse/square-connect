// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateLoyaltyRewardResponse A response that includes the loyalty reward created.
//
// swagger:model CreateLoyaltyRewardResponse
type CreateLoyaltyRewardResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The loyalty reward created.
	Reward *LoyaltyReward `json:"reward,omitempty"`
}

// Validate validates this create loyalty reward response
func (m *CreateLoyaltyRewardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReward(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLoyaltyRewardResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *CreateLoyaltyRewardResponse) validateReward(formats strfmt.Registry) error {

	if swag.IsZero(m.Reward) { // not required
		return nil
	}

	if m.Reward != nil {
		if err := m.Reward.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reward")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateLoyaltyRewardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateLoyaltyRewardResponse) UnmarshalBinary(b []byte) error {
	var res CreateLoyaltyRewardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}