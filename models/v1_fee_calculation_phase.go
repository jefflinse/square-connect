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

// V1FeeCalculationPhase v1 fee calculation phase
//
// swagger:model V1FeeCalculationPhase
type V1FeeCalculationPhase string

const (

	// V1FeeCalculationPhaseFEESUBTOTALPHASE captures enum value "FEE_SUBTOTAL_PHASE"
	V1FeeCalculationPhaseFEESUBTOTALPHASE V1FeeCalculationPhase = "FEE_SUBTOTAL_PHASE"

	// V1FeeCalculationPhaseOTHER captures enum value "OTHER"
	V1FeeCalculationPhaseOTHER V1FeeCalculationPhase = "OTHER"

	// V1FeeCalculationPhaseFEETOTALPHASE captures enum value "FEE_TOTAL_PHASE"
	V1FeeCalculationPhaseFEETOTALPHASE V1FeeCalculationPhase = "FEE_TOTAL_PHASE"
)

// for schema
var v1FeeCalculationPhaseEnum []interface{}

func init() {
	var res []V1FeeCalculationPhase
	if err := json.Unmarshal([]byte(`["FEE_SUBTOTAL_PHASE","OTHER","FEE_TOTAL_PHASE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1FeeCalculationPhaseEnum = append(v1FeeCalculationPhaseEnum, v)
	}
}

func (m V1FeeCalculationPhase) validateV1FeeCalculationPhaseEnum(path, location string, value V1FeeCalculationPhase) error {
	if err := validate.Enum(path, location, value, v1FeeCalculationPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 fee calculation phase
func (m V1FeeCalculationPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1FeeCalculationPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}