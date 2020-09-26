// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderPricingOptions Pricing options for an order. The options affect how the order's price is calculated.
// They can be used, for example, to apply automatic price adjustments that are based on pre-configured
// [pricing rules](/reference/square/objects/CatalogPricingRule).
//
// swagger:model OrderPricingOptions
type OrderPricingOptions struct {

	// The option to determine whether or not pricing rule-based discounts are automatically applied to an order.
	AutoApplyDiscounts bool `json:"auto_apply_discounts,omitempty"`
}

// Validate validates this order pricing options
func (m *OrderPricingOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderPricingOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderPricingOptions) UnmarshalBinary(b []byte) error {
	var res OrderPricingOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
