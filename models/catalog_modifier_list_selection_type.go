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

// CatalogModifierListSelectionType Indicates whether a CatalogModifierList supports multiple selections.
//
// swagger:model CatalogModifierListSelectionType
type CatalogModifierListSelectionType string

const (

	// CatalogModifierListSelectionTypeSINGLE captures enum value "SINGLE"
	CatalogModifierListSelectionTypeSINGLE CatalogModifierListSelectionType = "SINGLE"

	// CatalogModifierListSelectionTypeMULTIPLE captures enum value "MULTIPLE"
	CatalogModifierListSelectionTypeMULTIPLE CatalogModifierListSelectionType = "MULTIPLE"
)

// for schema
var catalogModifierListSelectionTypeEnum []interface{}

func init() {
	var res []CatalogModifierListSelectionType
	if err := json.Unmarshal([]byte(`["SINGLE","MULTIPLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogModifierListSelectionTypeEnum = append(catalogModifierListSelectionTypeEnum, v)
	}
}

func (m CatalogModifierListSelectionType) validateCatalogModifierListSelectionTypeEnum(path, location string, value CatalogModifierListSelectionType) error {
	if err := validate.EnumCase(path, location, value, catalogModifierListSelectionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this catalog modifier list selection type
func (m CatalogModifierListSelectionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCatalogModifierListSelectionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this catalog modifier list selection type based on context it is used
func (m CatalogModifierListSelectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
