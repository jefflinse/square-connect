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

// OrderLineItemTaxScope Indicates whether this is a line item or order level tax.
//
// swagger:model OrderLineItemTaxScope
type OrderLineItemTaxScope string

const (

	// OrderLineItemTaxScopeOTHERTAXSCOPE captures enum value "OTHER_TAX_SCOPE"
	OrderLineItemTaxScopeOTHERTAXSCOPE OrderLineItemTaxScope = "OTHER_TAX_SCOPE"

	// OrderLineItemTaxScopeLINEITEM captures enum value "LINE_ITEM"
	OrderLineItemTaxScopeLINEITEM OrderLineItemTaxScope = "LINE_ITEM"

	// OrderLineItemTaxScopeORDER captures enum value "ORDER"
	OrderLineItemTaxScopeORDER OrderLineItemTaxScope = "ORDER"
)

// for schema
var orderLineItemTaxScopeEnum []interface{}

func init() {
	var res []OrderLineItemTaxScope
	if err := json.Unmarshal([]byte(`["OTHER_TAX_SCOPE","LINE_ITEM","ORDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderLineItemTaxScopeEnum = append(orderLineItemTaxScopeEnum, v)
	}
}

func (m OrderLineItemTaxScope) validateOrderLineItemTaxScopeEnum(path, location string, value OrderLineItemTaxScope) error {
	if err := validate.EnumCase(path, location, value, orderLineItemTaxScopeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this order line item tax scope
func (m OrderLineItemTaxScope) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderLineItemTaxScopeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this order line item tax scope based on context it is used
func (m OrderLineItemTaxScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
