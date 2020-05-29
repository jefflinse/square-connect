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

// CatalogCustomAttributeDefinitionSellerVisibility Defines the visibility of a custom attribute to sellers in Square
// client applications, Square APIs or in Square UIs (including Square Point
// of Sale applications and Square Dashboard).
//
// swagger:model CatalogCustomAttributeDefinitionSellerVisibility
type CatalogCustomAttributeDefinitionSellerVisibility string

const (

	// CatalogCustomAttributeDefinitionSellerVisibilitySELLERVISIBILITYHIDDEN captures enum value "SELLER_VISIBILITY_HIDDEN"
	CatalogCustomAttributeDefinitionSellerVisibilitySELLERVISIBILITYHIDDEN CatalogCustomAttributeDefinitionSellerVisibility = "SELLER_VISIBILITY_HIDDEN"

	// CatalogCustomAttributeDefinitionSellerVisibilitySELLERVISIBILITYREADWRITEVALUES captures enum value "SELLER_VISIBILITY_READ_WRITE_VALUES"
	CatalogCustomAttributeDefinitionSellerVisibilitySELLERVISIBILITYREADWRITEVALUES CatalogCustomAttributeDefinitionSellerVisibility = "SELLER_VISIBILITY_READ_WRITE_VALUES"
)

// for schema
var catalogCustomAttributeDefinitionSellerVisibilityEnum []interface{}

func init() {
	var res []CatalogCustomAttributeDefinitionSellerVisibility
	if err := json.Unmarshal([]byte(`["SELLER_VISIBILITY_HIDDEN","SELLER_VISIBILITY_READ_WRITE_VALUES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogCustomAttributeDefinitionSellerVisibilityEnum = append(catalogCustomAttributeDefinitionSellerVisibilityEnum, v)
	}
}

func (m CatalogCustomAttributeDefinitionSellerVisibility) validateCatalogCustomAttributeDefinitionSellerVisibilityEnum(path, location string, value CatalogCustomAttributeDefinitionSellerVisibility) error {
	if err := validate.Enum(path, location, value, catalogCustomAttributeDefinitionSellerVisibilityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this catalog custom attribute definition seller visibility
func (m CatalogCustomAttributeDefinitionSellerVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCatalogCustomAttributeDefinitionSellerVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
