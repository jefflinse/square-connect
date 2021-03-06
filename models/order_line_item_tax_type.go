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

// OrderLineItemTaxType Indicates how the tax is applied to the associated line item or order.
//
// swagger:model OrderLineItemTaxType
type OrderLineItemTaxType string

const (

	// OrderLineItemTaxTypeUNKNOWNTAX captures enum value "UNKNOWN_TAX"
	OrderLineItemTaxTypeUNKNOWNTAX OrderLineItemTaxType = "UNKNOWN_TAX"

	// OrderLineItemTaxTypeADDITIVE captures enum value "ADDITIVE"
	OrderLineItemTaxTypeADDITIVE OrderLineItemTaxType = "ADDITIVE"

	// OrderLineItemTaxTypeINCLUSIVE captures enum value "INCLUSIVE"
	OrderLineItemTaxTypeINCLUSIVE OrderLineItemTaxType = "INCLUSIVE"
)

// for schema
var orderLineItemTaxTypeEnum []interface{}

func init() {
	var res []OrderLineItemTaxType
	if err := json.Unmarshal([]byte(`["UNKNOWN_TAX","ADDITIVE","INCLUSIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderLineItemTaxTypeEnum = append(orderLineItemTaxTypeEnum, v)
	}
}

func (m OrderLineItemTaxType) validateOrderLineItemTaxTypeEnum(path, location string, value OrderLineItemTaxType) error {
	if err := validate.EnumCase(path, location, value, orderLineItemTaxTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this order line item tax type
func (m OrderLineItemTaxType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderLineItemTaxTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this order line item tax type based on context it is used
func (m OrderLineItemTaxType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
