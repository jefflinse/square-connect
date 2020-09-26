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

// BankAccount Represents a bank account. For more information about
// linking a bank account to a Square account, see
// [Bank Accounts API](/docs/bank-accounts-api).
//
// swagger:model BankAccount
type BankAccount struct {

	// The last few digits of the account number.
	// Required: true
	// Min Length: 1
	AccountNumberSuffix *string `json:"account_number_suffix"`

	// The financial purpose of the associated bank account.
	// See [BankAccountType](#type-bankaccounttype) for possible values
	// Required: true
	AccountType *string `json:"account_type"`

	// Read only. Name of actual financial institution.
	// For example "Bank of America".
	// Max Length: 100
	BankName string `json:"bank_name,omitempty"`

	// The ISO 3166 Alpha-2 country code where the bank account is based.
	// See [Country](#type-country) for possible values
	// Required: true
	Country *string `json:"country"`

	// Indicates whether it is possible for Square to send money to this bank account.
	// Required: true
	Creditable *bool `json:"creditable"`

	// The 3-character ISO 4217 currency code indicating the operating
	// currency of the bank account. For example, the currency code for US dollars
	// is `USD`.
	// See [Currency](#type-currency) for possible values
	// Required: true
	Currency *string `json:"currency"`

	// Reference identifier that will be displayed to UK bank account owners
	// when collecting direct debit authorization. Only required for UK bank accounts.
	DebitMandateReferenceID string `json:"debit_mandate_reference_id,omitempty"`

	// Indicates whether it is possible for Square to take money from this
	// bank account.
	// Required: true
	Debitable *bool `json:"debitable"`

	// A Square-assigned, unique identifier for the bank account based on the
	// account information. The account fingerprint can be used to compare account
	// entries and determine if the they represent the same real-world bank account.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name of the account holder. This name must match the name
	// on the targeted bank account record.
	// Required: true
	// Min Length: 1
	HolderName *string `json:"holder_name"`

	// The unique, Square-issued identifier for the bank account.
	// Required: true
	// Max Length: 30
	// Min Length: 1
	ID *string `json:"id"`

	// The location to which the bank account belongs.
	LocationID string `json:"location_id,omitempty"`

	// Primary identifier for the bank. For more information, see
	// [Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api).
	// Required: true
	// Max Length: 40
	PrimaryBankIdentificationNumber *string `json:"primary_bank_identification_number"`

	// Client-provided identifier for linking the banking account to an entity
	// in a third-party system (for example, a bank account number or a user identifier).
	ReferenceID string `json:"reference_id,omitempty"`

	// Secondary identifier for the bank. For more information, see
	// [Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api).
	// Max Length: 40
	SecondaryBankIdentificationNumber string `json:"secondary_bank_identification_number,omitempty"`

	// Read-only. The current verification status of this BankAccount object.
	// See [BankAccountStatus](#type-bankaccountstatus) for possible values
	// Required: true
	Status *string `json:"status"`

	// The current version of the `BankAccount`.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this bank account
func (m *BankAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberSuffix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebitable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHolderName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryBankIdentificationNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryBankIdentificationNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BankAccount) validateAccountNumberSuffix(formats strfmt.Registry) error {

	if err := validate.Required("account_number_suffix", "body", m.AccountNumberSuffix); err != nil {
		return err
	}

	if err := validate.MinLength("account_number_suffix", "body", string(*m.AccountNumberSuffix), 1); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateAccountType(formats strfmt.Registry) error {

	if err := validate.Required("account_type", "body", m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateBankName(formats strfmt.Registry) error {

	if swag.IsZero(m.BankName) { // not required
		return nil
	}

	if err := validate.MaxLength("bank_name", "body", string(m.BankName), 100); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateCreditable(formats strfmt.Registry) error {

	if err := validate.Required("creditable", "body", m.Creditable); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateDebitable(formats strfmt.Registry) error {

	if err := validate.Required("debitable", "body", m.Debitable); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateHolderName(formats strfmt.Registry) error {

	if err := validate.Required("holder_name", "body", m.HolderName); err != nil {
		return err
	}

	if err := validate.MinLength("holder_name", "body", string(*m.HolderName), 1); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 30); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validatePrimaryBankIdentificationNumber(formats strfmt.Registry) error {

	if err := validate.Required("primary_bank_identification_number", "body", m.PrimaryBankIdentificationNumber); err != nil {
		return err
	}

	if err := validate.MaxLength("primary_bank_identification_number", "body", string(*m.PrimaryBankIdentificationNumber), 40); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateSecondaryBankIdentificationNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryBankIdentificationNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("secondary_bank_identification_number", "body", string(m.SecondaryBankIdentificationNumber), 40); err != nil {
		return err
	}

	return nil
}

func (m *BankAccount) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BankAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BankAccount) UnmarshalBinary(b []byte) error {
	var res BankAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}