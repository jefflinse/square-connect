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

// TransactionType Transaction type used in the disputed payment.
//
// swagger:model TransactionType
type TransactionType string

const (

	// TransactionTypeDEBIT captures enum value "DEBIT"
	TransactionTypeDEBIT TransactionType = "DEBIT"

	// TransactionTypeCREDIT captures enum value "CREDIT"
	TransactionTypeCREDIT TransactionType = "CREDIT"
)

// for schema
var transactionTypeEnum []interface{}

func init() {
	var res []TransactionType
	if err := json.Unmarshal([]byte(`["DEBIT","CREDIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeEnum = append(transactionTypeEnum, v)
	}
}

func (m TransactionType) validateTransactionTypeEnum(path, location string, value TransactionType) error {
	if err := validate.Enum(path, location, value, transactionTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this transaction type
func (m TransactionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTransactionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
