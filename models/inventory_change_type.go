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

// InventoryChangeType Indicates how the inventory change was applied to a tracked quantity of items.
//
// swagger:model InventoryChangeType
type InventoryChangeType string

const (

	// InventoryChangeTypePHYSICALCOUNT captures enum value "PHYSICAL_COUNT"
	InventoryChangeTypePHYSICALCOUNT InventoryChangeType = "PHYSICAL_COUNT"

	// InventoryChangeTypeADJUSTMENT captures enum value "ADJUSTMENT"
	InventoryChangeTypeADJUSTMENT InventoryChangeType = "ADJUSTMENT"

	// InventoryChangeTypeTRANSFER captures enum value "TRANSFER"
	InventoryChangeTypeTRANSFER InventoryChangeType = "TRANSFER"
)

// for schema
var inventoryChangeTypeEnum []interface{}

func init() {
	var res []InventoryChangeType
	if err := json.Unmarshal([]byte(`["PHYSICAL_COUNT","ADJUSTMENT","TRANSFER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryChangeTypeEnum = append(inventoryChangeTypeEnum, v)
	}
}

func (m InventoryChangeType) validateInventoryChangeTypeEnum(path, location string, value InventoryChangeType) error {
	if err := validate.EnumCase(path, location, value, inventoryChangeTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this inventory change type
func (m InventoryChangeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInventoryChangeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this inventory change type based on context it is used
func (m InventoryChangeType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
