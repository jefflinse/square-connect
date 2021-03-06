// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TipSettings tip settings
//
// swagger:model TipSettings
type TipSettings struct {

	// Indicates whether tipping is enabled for this checkout. Defaults to false.
	AllowTipping bool `json:"allow_tipping,omitempty"`

	// Indicates whether custom tip amounts are allowed during the checkout flow. Defaults to false.
	CustomTipField bool `json:"custom_tip_field,omitempty"`

	// Indicates whether tip options should be presented on their own screen before presenting
	// the signature screen during card payment. Defaults to false.
	SeparateTipScreen bool `json:"separate_tip_screen,omitempty"`
}

// Validate validates this tip settings
func (m *TipSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TipSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TipSettings) UnmarshalBinary(b []byte) error {
	var res TipSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
