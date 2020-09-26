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

// InvoiceSortField Field to use for sorting.
//
// swagger:model InvoiceSortField
type InvoiceSortField string

const (

	// InvoiceSortFieldINVOICESORTDATE captures enum value "INVOICE_SORT_DATE"
	InvoiceSortFieldINVOICESORTDATE InvoiceSortField = "INVOICE_SORT_DATE"
)

// for schema
var invoiceSortFieldEnum []interface{}

func init() {
	var res []InvoiceSortField
	if err := json.Unmarshal([]byte(`["INVOICE_SORT_DATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceSortFieldEnum = append(invoiceSortFieldEnum, v)
	}
}

func (m InvoiceSortField) validateInvoiceSortFieldEnum(path, location string, value InvoiceSortField) error {
	if err := validate.Enum(path, location, value, invoiceSortFieldEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this invoice sort field
func (m InvoiceSortField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInvoiceSortFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
