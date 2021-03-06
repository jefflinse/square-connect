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

// RetrieveLoyaltyAccountResponse A response that includes the loyalty account.
// Example: {"loyalty_account":{"balance":10,"created_at":"2020-05-08T21:44:32Z","customer_id":"Q8002FAM9V1EZ0ADB2T5609X6NET1H0","id":"79b807d2-d786-46a9-933b-918028d7a8c5","lifetime_points":20,"mappings":[{"created_at":"2020-05-08T21:44:32Z","id":"66aaab3f-da99-49ed-8b19-b87f851c844f","type":"PHONE","value":"+14155551234"}],"program_id":"d619f755-2d17-41f3-990d-c04ecedd64dd","updated_at":"2020-05-08T21:44:32Z"}}
//
// swagger:model RetrieveLoyaltyAccountResponse
type RetrieveLoyaltyAccountResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The loyalty account.
	LoyaltyAccount *LoyaltyAccount `json:"loyalty_account,omitempty"`
}

// Validate validates this retrieve loyalty account response
func (m *RetrieveLoyaltyAccountResponse) Validate(formats strfmt.Registry) error {
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

func (m *RetrieveLoyaltyAccountResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *RetrieveLoyaltyAccountResponse) validateLoyaltyAccount(formats strfmt.Registry) error {
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

// ContextValidate validate this retrieve loyalty account response based on the context it is used
func (m *RetrieveLoyaltyAccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoyaltyAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveLoyaltyAccountResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RetrieveLoyaltyAccountResponse) contextValidateLoyaltyAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.LoyaltyAccount != nil {
		if err := m.LoyaltyAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loyalty_account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetrieveLoyaltyAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveLoyaltyAccountResponse) UnmarshalBinary(b []byte) error {
	var res RetrieveLoyaltyAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
