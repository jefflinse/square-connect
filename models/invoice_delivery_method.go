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

// InvoiceDeliveryMethod Indicates how Square delivers the `invoice` to the customer.
//
// swagger:model InvoiceDeliveryMethod
type InvoiceDeliveryMethod string

const (

	// InvoiceDeliveryMethodEMAIL captures enum value "EMAIL"
	InvoiceDeliveryMethodEMAIL InvoiceDeliveryMethod = "EMAIL"

	// InvoiceDeliveryMethodSHAREMANUALLY captures enum value "SHARE_MANUALLY"
	InvoiceDeliveryMethodSHAREMANUALLY InvoiceDeliveryMethod = "SHARE_MANUALLY"
)

// for schema
var invoiceDeliveryMethodEnum []interface{}

func init() {
	var res []InvoiceDeliveryMethod
	if err := json.Unmarshal([]byte(`["EMAIL","SHARE_MANUALLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDeliveryMethodEnum = append(invoiceDeliveryMethodEnum, v)
	}
}

func (m InvoiceDeliveryMethod) validateInvoiceDeliveryMethodEnum(path, location string, value InvoiceDeliveryMethod) error {
	if err := validate.EnumCase(path, location, value, invoiceDeliveryMethodEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this invoice delivery method
func (m InvoiceDeliveryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvoiceDeliveryMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this invoice delivery method based on context it is used
func (m InvoiceDeliveryMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
