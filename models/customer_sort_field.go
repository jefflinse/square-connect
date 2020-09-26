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

// CustomerSortField Specifies customer attributes as the sort key to customer profiles returned from a search.
//
// swagger:model CustomerSortField
type CustomerSortField string

const (

	// CustomerSortFieldDEFAULT captures enum value "DEFAULT"
	CustomerSortFieldDEFAULT CustomerSortField = "DEFAULT"

	// CustomerSortFieldCREATEDAT captures enum value "CREATED_AT"
	CustomerSortFieldCREATEDAT CustomerSortField = "CREATED_AT"
)

// for schema
var customerSortFieldEnum []interface{}

func init() {
	var res []CustomerSortField
	if err := json.Unmarshal([]byte(`["DEFAULT","CREATED_AT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerSortFieldEnum = append(customerSortFieldEnum, v)
	}
}

func (m CustomerSortField) validateCustomerSortFieldEnum(path, location string, value CustomerSortField) error {
	if err := validate.Enum(path, location, value, customerSortFieldEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this customer sort field
func (m CustomerSortField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustomerSortFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
