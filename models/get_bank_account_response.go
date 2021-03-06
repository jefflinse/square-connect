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

// GetBankAccountResponse Response object returned by `GetBankAccount`.
// Example: {"bank_account":{"account_number_suffix":"971","account_type":"CHECKING","bank_name":"Bank Name","country":"US","creditable":false,"currency":"USD","debitable":false,"holder_name":"Jane Doe","id":"w3yRgCGYQnwmdl0R3GB","location_id":"S8GWD5example","primary_bank_identification_number":"112200303","status":"VERIFICATION_IN_PROGRESS","version":5}}
//
// swagger:model GetBankAccountResponse
type GetBankAccountResponse struct {

	// The requested `BankAccount` object.
	BankAccount *BankAccount `json:"bank_account,omitempty"`

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this get bank account response
func (m *GetBankAccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBankAccountResponse) validateBankAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BankAccount) { // not required
		return nil
	}

	if m.BankAccount != nil {
		if err := m.BankAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank_account")
			}
			return err
		}
	}

	return nil
}

func (m *GetBankAccountResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this get bank account response based on the context it is used
func (m *GetBankAccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBankAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBankAccountResponse) contextValidateBankAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BankAccount != nil {
		if err := m.BankAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank_account")
			}
			return err
		}
	}

	return nil
}

func (m *GetBankAccountResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *GetBankAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBankAccountResponse) UnmarshalBinary(b []byte) error {
	var res GetBankAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
