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

// CatalogDiscountType How to apply a CatalogDiscount to a CatalogItem.
//
// swagger:model CatalogDiscountType
type CatalogDiscountType string

const (

	// CatalogDiscountTypeFIXEDPERCENTAGE captures enum value "FIXED_PERCENTAGE"
	CatalogDiscountTypeFIXEDPERCENTAGE CatalogDiscountType = "FIXED_PERCENTAGE"

	// CatalogDiscountTypeFIXEDAMOUNT captures enum value "FIXED_AMOUNT"
	CatalogDiscountTypeFIXEDAMOUNT CatalogDiscountType = "FIXED_AMOUNT"

	// CatalogDiscountTypeVARIABLEPERCENTAGE captures enum value "VARIABLE_PERCENTAGE"
	CatalogDiscountTypeVARIABLEPERCENTAGE CatalogDiscountType = "VARIABLE_PERCENTAGE"

	// CatalogDiscountTypeVARIABLEAMOUNT captures enum value "VARIABLE_AMOUNT"
	CatalogDiscountTypeVARIABLEAMOUNT CatalogDiscountType = "VARIABLE_AMOUNT"
)

// for schema
var catalogDiscountTypeEnum []interface{}

func init() {
	var res []CatalogDiscountType
	if err := json.Unmarshal([]byte(`["FIXED_PERCENTAGE","FIXED_AMOUNT","VARIABLE_PERCENTAGE","VARIABLE_AMOUNT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogDiscountTypeEnum = append(catalogDiscountTypeEnum, v)
	}
}

func (m CatalogDiscountType) validateCatalogDiscountTypeEnum(path, location string, value CatalogDiscountType) error {
	if err := validate.Enum(path, location, value, catalogDiscountTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this catalog discount type
func (m CatalogDiscountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCatalogDiscountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
