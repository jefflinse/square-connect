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

// InventoryAlertType Indicates whether Square should alert the merchant when the inventory quantity of a CatalogItemVariation is low.
//
// swagger:model InventoryAlertType
type InventoryAlertType string

const (

	// InventoryAlertTypeNONE captures enum value "NONE"
	InventoryAlertTypeNONE InventoryAlertType = "NONE"

	// InventoryAlertTypeLOWQUANTITY captures enum value "LOW_QUANTITY"
	InventoryAlertTypeLOWQUANTITY InventoryAlertType = "LOW_QUANTITY"
)

// for schema
var inventoryAlertTypeEnum []interface{}

func init() {
	var res []InventoryAlertType
	if err := json.Unmarshal([]byte(`["NONE","LOW_QUANTITY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryAlertTypeEnum = append(inventoryAlertTypeEnum, v)
	}
}

func (m InventoryAlertType) validateInventoryAlertTypeEnum(path, location string, value InventoryAlertType) error {
	if err := validate.EnumCase(path, location, value, inventoryAlertTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this inventory alert type
func (m InventoryAlertType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInventoryAlertTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this inventory alert type based on context it is used
func (m InventoryAlertType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
