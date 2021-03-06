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

// ListBankAccountsResponse Response object returned by ListBankAccounts.
// Example: {"bank_accounts":[{"account_number_suffix":"971","account_type":"CHECKING","bank_name":"Bank Name","country":"US","creditable":false,"currency":"USD","debitable":false,"holder_name":"Jane Doe","id":"ao6iaQ9vhDiaQD7n3GB","location_id":"S8GWD5example","primary_bank_identification_number":"112200303","status":"VERIFICATION_IN_PROGRESS","version":5},{"account_number_suffix":"972","account_type":"CHECKING","bank_name":"Bank Name","country":"US","creditable":false,"currency":"USD","debitable":false,"holder_name":"Jane Doe","id":"4x7WXuaxrkQkVlka3GB","location_id":"S8GWD5example","primary_bank_identification_number":"112200303","status":"VERIFICATION_IN_PROGRESS","version":5}]}
//
// swagger:model ListBankAccountsResponse
type ListBankAccountsResponse struct {

	// List of BankAccounts associated with this account.
	BankAccounts []*BankAccount `json:"bank_accounts"`

	// When a response is truncated, it includes a cursor that you can
	// use in a subsequent request to fetch next set of bank accounts.
	// If empty, this is the final response.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this list bank accounts response
func (m *ListBankAccountsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankAccounts(formats); err != nil {
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

func (m *ListBankAccountsResponse) validateBankAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.BankAccounts) { // not required
		return nil
	}

	for i := 0; i < len(m.BankAccounts); i++ {
		if swag.IsZero(m.BankAccounts[i]) { // not required
			continue
		}

		if m.BankAccounts[i] != nil {
			if err := m.BankAccounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bank_accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListBankAccountsResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this list bank accounts response based on the context it is used
func (m *ListBankAccountsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBankAccounts(ctx, formats); err != nil {
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

func (m *ListBankAccountsResponse) contextValidateBankAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BankAccounts); i++ {

		if m.BankAccounts[i] != nil {
			if err := m.BankAccounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bank_accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListBankAccountsResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ListBankAccountsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListBankAccountsResponse) UnmarshalBinary(b []byte) error {
	var res ListBankAccountsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
