// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RevokeTokenResponse revoke token response
// Example: {"success":true}
//
// swagger:model RevokeTokenResponse
type RevokeTokenResponse struct {

	// If the request is successful, this is true.
	Success bool `json:"success,omitempty"`
}

// Validate validates this revoke token response
func (m *RevokeTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this revoke token response based on context it is used
func (m *RevokeTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RevokeTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevokeTokenResponse) UnmarshalBinary(b []byte) error {
	var res RevokeTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
