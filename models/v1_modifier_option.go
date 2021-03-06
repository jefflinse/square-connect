// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ModifierOption V1ModifierOption
//
// swagger:model V1ModifierOption
type V1ModifierOption struct {

	// The modifier option's unique ID.
	ID string `json:"id,omitempty"`

	// The ID of the modifier list the option belongs to.
	ModifierListID string `json:"modifier_list_id,omitempty"`

	// The modifier option's name.
	Name string `json:"name,omitempty"`

	// If true, the modifier option is the default option in a modifier list for which selection_type is SINGLE.
	OnByDefault bool `json:"on_by_default,omitempty"`

	// Indicates the modifier option's list position when displayed in Square Point of Sale and the merchant dashboard. If more than one modifier option in the same modifier list has the same ordinal value, those options are displayed in alphabetical order.
	Ordinal int64 `json:"ordinal,omitempty"`

	// The modifier option's price.
	PriceMoney *V1Money `json:"price_money,omitempty"`

	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2ID string `json:"v2_id,omitempty"`
}

// Validate validates this v1 modifier option
func (m *V1ModifierOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ModifierOption) validatePriceMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceMoney) { // not required
		return nil
	}

	if m.PriceMoney != nil {
		if err := m.PriceMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ModifierOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ModifierOption) UnmarshalBinary(b []byte) error {
	var res V1ModifierOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
