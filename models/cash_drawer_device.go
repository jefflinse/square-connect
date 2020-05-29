// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CashDrawerDevice cash drawer device
//
// swagger:model CashDrawerDevice
type CashDrawerDevice struct {

	// The device Square-issued ID
	ID string `json:"id,omitempty"`

	// The device merchant-specified name.
	Name string `json:"name,omitempty"`
}

// Validate validates this cash drawer device
func (m *CashDrawerDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CashDrawerDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CashDrawerDevice) UnmarshalBinary(b []byte) error {
	var res CashDrawerDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
