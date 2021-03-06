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

// TenderBankTransferDetailsStatus Indicates the bank transfer's current status.
//
// swagger:model TenderBankTransferDetailsStatus
type TenderBankTransferDetailsStatus string

const (

	// TenderBankTransferDetailsStatusPENDING captures enum value "PENDING"
	TenderBankTransferDetailsStatusPENDING TenderBankTransferDetailsStatus = "PENDING"

	// TenderBankTransferDetailsStatusCOMPLETED captures enum value "COMPLETED"
	TenderBankTransferDetailsStatusCOMPLETED TenderBankTransferDetailsStatus = "COMPLETED"

	// TenderBankTransferDetailsStatusFAILED captures enum value "FAILED"
	TenderBankTransferDetailsStatusFAILED TenderBankTransferDetailsStatus = "FAILED"
)

// for schema
var tenderBankTransferDetailsStatusEnum []interface{}

func init() {
	var res []TenderBankTransferDetailsStatus
	if err := json.Unmarshal([]byte(`["PENDING","COMPLETED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tenderBankTransferDetailsStatusEnum = append(tenderBankTransferDetailsStatusEnum, v)
	}
}

func (m TenderBankTransferDetailsStatus) validateTenderBankTransferDetailsStatusEnum(path, location string, value TenderBankTransferDetailsStatus) error {
	if err := validate.Enum(path, location, value, tenderBankTransferDetailsStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tender bank transfer details status
func (m TenderBankTransferDetailsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTenderBankTransferDetailsStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
