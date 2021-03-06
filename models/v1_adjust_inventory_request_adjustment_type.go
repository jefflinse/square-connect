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

// V1AdjustInventoryRequestAdjustmentType v1 adjust inventory request adjustment type
//
// swagger:model V1AdjustInventoryRequestAdjustmentType
type V1AdjustInventoryRequestAdjustmentType string

const (

	// V1AdjustInventoryRequestAdjustmentTypeSALE captures enum value "SALE"
	V1AdjustInventoryRequestAdjustmentTypeSALE V1AdjustInventoryRequestAdjustmentType = "SALE"

	// V1AdjustInventoryRequestAdjustmentTypeRECEIVESTOCK captures enum value "RECEIVE_STOCK"
	V1AdjustInventoryRequestAdjustmentTypeRECEIVESTOCK V1AdjustInventoryRequestAdjustmentType = "RECEIVE_STOCK"

	// V1AdjustInventoryRequestAdjustmentTypeMANUALADJUST captures enum value "MANUAL_ADJUST"
	V1AdjustInventoryRequestAdjustmentTypeMANUALADJUST V1AdjustInventoryRequestAdjustmentType = "MANUAL_ADJUST"
)

// for schema
var v1AdjustInventoryRequestAdjustmentTypeEnum []interface{}

func init() {
	var res []V1AdjustInventoryRequestAdjustmentType
	if err := json.Unmarshal([]byte(`["SALE","RECEIVE_STOCK","MANUAL_ADJUST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AdjustInventoryRequestAdjustmentTypeEnum = append(v1AdjustInventoryRequestAdjustmentTypeEnum, v)
	}
}

func (m V1AdjustInventoryRequestAdjustmentType) validateV1AdjustInventoryRequestAdjustmentTypeEnum(path, location string, value V1AdjustInventoryRequestAdjustmentType) error {
	if err := validate.Enum(path, location, value, v1AdjustInventoryRequestAdjustmentTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 adjust inventory request adjustment type
func (m V1AdjustInventoryRequestAdjustmentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1AdjustInventoryRequestAdjustmentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
