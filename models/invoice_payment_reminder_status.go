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

// InvoicePaymentReminderStatus The status of a payment request reminder.
//
// swagger:model InvoicePaymentReminderStatus
type InvoicePaymentReminderStatus string

const (

	// InvoicePaymentReminderStatusPENDING captures enum value "PENDING"
	InvoicePaymentReminderStatusPENDING InvoicePaymentReminderStatus = "PENDING"

	// InvoicePaymentReminderStatusNOTAPPLICABLE captures enum value "NOT_APPLICABLE"
	InvoicePaymentReminderStatusNOTAPPLICABLE InvoicePaymentReminderStatus = "NOT_APPLICABLE"

	// InvoicePaymentReminderStatusSENT captures enum value "SENT"
	InvoicePaymentReminderStatusSENT InvoicePaymentReminderStatus = "SENT"
)

// for schema
var invoicePaymentReminderStatusEnum []interface{}

func init() {
	var res []InvoicePaymentReminderStatus
	if err := json.Unmarshal([]byte(`["PENDING","NOT_APPLICABLE","SENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoicePaymentReminderStatusEnum = append(invoicePaymentReminderStatusEnum, v)
	}
}

func (m InvoicePaymentReminderStatus) validateInvoicePaymentReminderStatusEnum(path, location string, value InvoicePaymentReminderStatus) error {
	if err := validate.Enum(path, location, value, invoicePaymentReminderStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this invoice payment reminder status
func (m InvoicePaymentReminderStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvoicePaymentReminderStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}