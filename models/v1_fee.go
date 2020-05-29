// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Fee V1Fee
//
// swagger:model V1Fee
type V1Fee struct {

	// The type of adjustment the fee applies to a payment. Currently, this value is TAX for all fees.
	// See [V1FeeAdjustmentType](#type-v1feeadjustmenttype) for possible values
	AdjustmentType string `json:"adjustment_type,omitempty"`

	// If true, the fee applies to custom amounts entered into Square Point of Sale that are not associated with a particular item.
	AppliesToCustomAmounts bool `json:"applies_to_custom_amounts,omitempty"`

	// Forthcoming
	// See [V1FeeCalculationPhase](#type-v1feecalculationphase) for possible values
	CalculationPhase string `json:"calculation_phase,omitempty"`

	// If true, the fee is applied to all appropriate items. If false, the fee is not applied at all.
	Enabled bool `json:"enabled,omitempty"`

	// The fee's unique ID.
	ID string `json:"id,omitempty"`

	// Whether the fee is ADDITIVE or INCLUSIVE.
	// See [V1FeeInclusionType](#type-v1feeinclusiontype) for possible values
	InclusionType string `json:"inclusion_type,omitempty"`

	// The fee's name.
	Name string `json:"name,omitempty"`

	// The rate of the fee, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%.
	Rate string `json:"rate,omitempty"`

	// In countries with multiple classifications for sales taxes, indicates which classification the fee falls under. Currently relevant only to Canadian merchants.
	// See [V1FeeType](#type-v1feetype) for possible values
	Type string `json:"type,omitempty"`

	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2ID string `json:"v2_id,omitempty"`
}

// Validate validates this v1 fee
func (m *V1Fee) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1Fee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Fee) UnmarshalBinary(b []byte) error {
	var res V1Fee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
