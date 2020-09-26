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

// GetBankAccountByV1IDResponse Response object returned by GetBankAccountByV1Id.
//
// swagger:model GetBankAccountByV1IdResponse
type GetBankAccountByV1IDResponse struct {

	// The requested `BankAccount` object.
	BankAccount *BankAccount `json:"bank_account,omitempty"`

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this get bank account by v1 Id response
func (m *GetBankAccountByV1IDResponse) Validate(formats strfmt.Registry) error {
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

func (m *GetBankAccountByV1IDResponse) validateBankAccount(formats strfmt.Registry) error {

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

func (m *GetBankAccountByV1IDResponse) validateErrors(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *GetBankAccountByV1IDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBankAccountByV1IDResponse) UnmarshalBinary(b []byte) error {
	var res GetBankAccountByV1IDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}