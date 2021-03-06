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

// TaxCalculationPhase When to calculate the taxes due on a cart.
//
// swagger:model TaxCalculationPhase
type TaxCalculationPhase string

const (

	// TaxCalculationPhaseTAXSUBTOTALPHASE captures enum value "TAX_SUBTOTAL_PHASE"
	TaxCalculationPhaseTAXSUBTOTALPHASE TaxCalculationPhase = "TAX_SUBTOTAL_PHASE"

	// TaxCalculationPhaseTAXTOTALPHASE captures enum value "TAX_TOTAL_PHASE"
	TaxCalculationPhaseTAXTOTALPHASE TaxCalculationPhase = "TAX_TOTAL_PHASE"
)

// for schema
var taxCalculationPhaseEnum []interface{}

func init() {
	var res []TaxCalculationPhase
	if err := json.Unmarshal([]byte(`["TAX_SUBTOTAL_PHASE","TAX_TOTAL_PHASE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taxCalculationPhaseEnum = append(taxCalculationPhaseEnum, v)
	}
}

func (m TaxCalculationPhase) validateTaxCalculationPhaseEnum(path, location string, value TaxCalculationPhase) error {
	if err := validate.Enum(path, location, value, taxCalculationPhaseEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tax calculation phase
func (m TaxCalculationPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTaxCalculationPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
