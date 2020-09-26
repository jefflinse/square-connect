// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PaymentItemDetail V1PaymentItemDetail
//
// swagger:model V1PaymentItemDetail
type V1PaymentItemDetail struct {

	// The name of the item's merchant-defined category, if any.
	CategoryName string `json:"category_name,omitempty"`

	// The unique ID of the item purchased, if any.
	ItemID string `json:"item_id,omitempty"`

	// The unique ID of the item variation purchased, if any.
	ItemVariationID string `json:"item_variation_id,omitempty"`

	//  The item's merchant-defined SKU, if any.
	Sku string `json:"sku,omitempty"`
}

// Validate validates this v1 payment item detail
func (m *V1PaymentItemDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PaymentItemDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PaymentItemDetail) UnmarshalBinary(b []byte) error {
	var res V1PaymentItemDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}