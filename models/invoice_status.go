// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// InvoiceStatus Indicates the status of an invoice.
//
// swagger:model InvoiceStatus
type InvoiceStatus string

const (

	// InvoiceStatusDRAFT captures enum value "DRAFT"
	InvoiceStatusDRAFT InvoiceStatus = "DRAFT"

	// InvoiceStatusUNPAID captures enum value "UNPAID"
	InvoiceStatusUNPAID InvoiceStatus = "UNPAID"

	// InvoiceStatusSCHEDULED captures enum value "SCHEDULED"
	InvoiceStatusSCHEDULED InvoiceStatus = "SCHEDULED"

	// InvoiceStatusPARTIALLYPAID captures enum value "PARTIALLY_PAID"
	InvoiceStatusPARTIALLYPAID InvoiceStatus = "PARTIALLY_PAID"

	// InvoiceStatusPAID captures enum value "PAID"
	InvoiceStatusPAID InvoiceStatus = "PAID"

	// InvoiceStatusPARTIALLYREFUNDED captures enum value "PARTIALLY_REFUNDED"
	InvoiceStatusPARTIALLYREFUNDED InvoiceStatus = "PARTIALLY_REFUNDED"

	// InvoiceStatusREFUNDED captures enum value "REFUNDED"
	InvoiceStatusREFUNDED InvoiceStatus = "REFUNDED"

	// InvoiceStatusCANCELED captures enum value "CANCELED"
	InvoiceStatusCANCELED InvoiceStatus = "CANCELED"

	// InvoiceStatusFAILED captures enum value "FAILED"
	InvoiceStatusFAILED InvoiceStatus = "FAILED"

	// InvoiceStatusPAYMENTPENDING captures enum value "PAYMENT_PENDING"
	InvoiceStatusPAYMENTPENDING InvoiceStatus = "PAYMENT_PENDING"
)

// for schema
var invoiceStatusEnum []interface{}

func init() {
	var res []InvoiceStatus
	if err := json.Unmarshal([]byte(`["DRAFT","UNPAID","SCHEDULED","PARTIALLY_PAID","PAID","PARTIALLY_REFUNDED","REFUNDED","CANCELED","FAILED","PAYMENT_PENDING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceStatusEnum = append(invoiceStatusEnum, v)
	}
}

func (m InvoiceStatus) validateInvoiceStatusEnum(path, location string, value InvoiceStatus) error {
	if err := validate.EnumCase(path, location, value, invoiceStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this invoice status
func (m InvoiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvoiceStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this invoice status based on context it is used
func (m InvoiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
