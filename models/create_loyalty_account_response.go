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

// CreateLoyaltyAccountResponse A response that includes loyalty account created.
//
// swagger:model CreateLoyaltyAccountResponse
type CreateLoyaltyAccountResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The newly created loyalty account.
	LoyaltyAccount *LoyaltyAccount `json:"loyalty_account,omitempty"`
}

// Validate validates this create loyalty account response
func (m *CreateLoyaltyAccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoyaltyAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLoyaltyAccountResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *CreateLoyaltyAccountResponse) validateLoyaltyAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.LoyaltyAccount) { // not required
		return nil
	}

	if m.LoyaltyAccount != nil {
		if err := m.LoyaltyAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loyalty_account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateLoyaltyAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateLoyaltyAccountResponse) UnmarshalBinary(b []byte) error {
	var res CreateLoyaltyAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}