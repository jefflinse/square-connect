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

// MeasurementUnitWeight Unit of weight used to measure a quantity.
//
// swagger:model MeasurementUnitWeight
type MeasurementUnitWeight string

const (

	// MeasurementUnitWeightIMPERIALWEIGHTOUNCE captures enum value "IMPERIAL_WEIGHT_OUNCE"
	MeasurementUnitWeightIMPERIALWEIGHTOUNCE MeasurementUnitWeight = "IMPERIAL_WEIGHT_OUNCE"

	// MeasurementUnitWeightIMPERIALPOUND captures enum value "IMPERIAL_POUND"
	MeasurementUnitWeightIMPERIALPOUND MeasurementUnitWeight = "IMPERIAL_POUND"

	// MeasurementUnitWeightIMPERIALSTONE captures enum value "IMPERIAL_STONE"
	MeasurementUnitWeightIMPERIALSTONE MeasurementUnitWeight = "IMPERIAL_STONE"

	// MeasurementUnitWeightMETRICMILLIGRAM captures enum value "METRIC_MILLIGRAM"
	MeasurementUnitWeightMETRICMILLIGRAM MeasurementUnitWeight = "METRIC_MILLIGRAM"

	// MeasurementUnitWeightMETRICGRAM captures enum value "METRIC_GRAM"
	MeasurementUnitWeightMETRICGRAM MeasurementUnitWeight = "METRIC_GRAM"

	// MeasurementUnitWeightMETRICKILOGRAM captures enum value "METRIC_KILOGRAM"
	MeasurementUnitWeightMETRICKILOGRAM MeasurementUnitWeight = "METRIC_KILOGRAM"
)

// for schema
var measurementUnitWeightEnum []interface{}

func init() {
	var res []MeasurementUnitWeight
	if err := json.Unmarshal([]byte(`["IMPERIAL_WEIGHT_OUNCE","IMPERIAL_POUND","IMPERIAL_STONE","METRIC_MILLIGRAM","METRIC_GRAM","METRIC_KILOGRAM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		measurementUnitWeightEnum = append(measurementUnitWeightEnum, v)
	}
}

func (m MeasurementUnitWeight) validateMeasurementUnitWeightEnum(path, location string, value MeasurementUnitWeight) error {
	if err := validate.Enum(path, location, value, measurementUnitWeightEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this measurement unit weight
func (m MeasurementUnitWeight) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMeasurementUnitWeightEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}