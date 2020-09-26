// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ItemImage V1ItemImage
//
// swagger:model V1ItemImage
type V1ItemImage struct {

	// The image's unique ID.
	ID string `json:"id,omitempty"`

	// The image's publicly accessible URL.
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 item image
func (m *V1ItemImage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ItemImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ItemImage) UnmarshalBinary(b []byte) error {
	var res V1ItemImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}