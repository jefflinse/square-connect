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

// InvoiceRequestMethod Specifies the action for Square to take for processing the invoice. For example,
// email the invoice, charge a customer's card on file, or do nothing. DEPRECATED at version 2021-01-21. The corresponding `request_method` field is replaced by the `Invoice.delivery_method` and `InvoicePaymentRequest.automatic_payment_source` fields.
//
// swagger:model InvoiceRequestMethod
type InvoiceRequestMethod string

const (

	// InvoiceRequestMethodEMAIL captures enum value "EMAIL"
	InvoiceRequestMethodEMAIL InvoiceRequestMethod = "EMAIL"

	// InvoiceRequestMethodCHARGECARDONFILE captures enum value "CHARGE_CARD_ON_FILE"
	InvoiceRequestMethodCHARGECARDONFILE InvoiceRequestMethod = "CHARGE_CARD_ON_FILE"

	// InvoiceRequestMethodSHAREMANUALLY captures enum value "SHARE_MANUALLY"
	InvoiceRequestMethodSHAREMANUALLY InvoiceRequestMethod = "SHARE_MANUALLY"

	// InvoiceRequestMethodCHARGEBANKONFILE captures enum value "CHARGE_BANK_ON_FILE"
	InvoiceRequestMethodCHARGEBANKONFILE InvoiceRequestMethod = "CHARGE_BANK_ON_FILE"
)

// for schema
var invoiceRequestMethodEnum []interface{}

func init() {
	var res []InvoiceRequestMethod
	if err := json.Unmarshal([]byte(`["EMAIL","CHARGE_CARD_ON_FILE","SHARE_MANUALLY","CHARGE_BANK_ON_FILE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceRequestMethodEnum = append(invoiceRequestMethodEnum, v)
	}
}

func (m InvoiceRequestMethod) validateInvoiceRequestMethodEnum(path, location string, value InvoiceRequestMethod) error {
	if err := validate.EnumCase(path, location, value, invoiceRequestMethodEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this invoice request method
func (m InvoiceRequestMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvoiceRequestMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this invoice request method based on context it is used
func (m InvoiceRequestMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
