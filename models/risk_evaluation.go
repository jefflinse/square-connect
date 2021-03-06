// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RiskEvaluation Represents fraud risk information for the associated payment.
//
// When you take a payment through Square's Payments API (using the `CreatePayment`
// endpoint), Square evaluates it and assigns a risk level to the payment. Sellers
// can use this information to determine the course of action (for example,
// provide the goods/services or refund the payment).
//
// swagger:model RiskEvaluation
type RiskEvaluation struct {

	// The timestamp when payment risk was evaluated, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The risk level associated with the payment
	// See [RiskEvaluationRiskLevel](#type-riskevaluationrisklevel) for possible values
	RiskLevel string `json:"risk_level,omitempty"`
}

// Validate validates this risk evaluation
func (m *RiskEvaluation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this risk evaluation based on context it is used
func (m *RiskEvaluation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RiskEvaluation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RiskEvaluation) UnmarshalBinary(b []byte) error {
	var res RiskEvaluation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}