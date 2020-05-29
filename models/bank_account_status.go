// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BankAccountStatus Indicates the current verification status of a `BankAccount` object.
//
// swagger:model BankAccountStatus
type BankAccountStatus string

const (

	// BankAccountStatusVERIFICATIONINPROGRESS captures enum value "VERIFICATION_IN_PROGRESS"
	BankAccountStatusVERIFICATIONINPROGRESS BankAccountStatus = "VERIFICATION_IN_PROGRESS"

	// BankAccountStatusVERIFIED captures enum value "VERIFIED"
	BankAccountStatusVERIFIED BankAccountStatus = "VERIFIED"

	// BankAccountStatusDISABLED captures enum value "DISABLED"
	BankAccountStatusDISABLED BankAccountStatus = "DISABLED"
)

// for schema
var bankAccountStatusEnum []interface{}

func init() {
	var res []BankAccountStatus
	if err := json.Unmarshal([]byte(`["VERIFICATION_IN_PROGRESS","VERIFIED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bankAccountStatusEnum = append(bankAccountStatusEnum, v)
	}
}

func (m BankAccountStatus) validateBankAccountStatusEnum(path, location string, value BankAccountStatus) error {
	if err := validate.Enum(path, location, value, bankAccountStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this bank account status
func (m BankAccountStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBankAccountStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
