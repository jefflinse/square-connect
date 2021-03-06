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

// CatalogPricingType Indicates whether the price of a CatalogItemVariation should be entered manually at the time of sale.
//
// swagger:model CatalogPricingType
type CatalogPricingType string

const (

	// CatalogPricingTypeFIXEDPRICING captures enum value "FIXED_PRICING"
	CatalogPricingTypeFIXEDPRICING CatalogPricingType = "FIXED_PRICING"

	// CatalogPricingTypeVARIABLEPRICING captures enum value "VARIABLE_PRICING"
	CatalogPricingTypeVARIABLEPRICING CatalogPricingType = "VARIABLE_PRICING"
)

// for schema
var catalogPricingTypeEnum []interface{}

func init() {
	var res []CatalogPricingType
	if err := json.Unmarshal([]byte(`["FIXED_PRICING","VARIABLE_PRICING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogPricingTypeEnum = append(catalogPricingTypeEnum, v)
	}
}

func (m CatalogPricingType) validateCatalogPricingTypeEnum(path, location string, value CatalogPricingType) error {
	if err := validate.EnumCase(path, location, value, catalogPricingTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this catalog pricing type
func (m CatalogPricingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCatalogPricingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this catalog pricing type based on context it is used
func (m CatalogPricingType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
