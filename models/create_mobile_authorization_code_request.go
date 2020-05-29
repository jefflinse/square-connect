// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateMobileAuthorizationCodeRequest Defines the body parameters that can be provided in a request to the
// __CreateMobileAuthorizationCode__ endpoint.
//
// swagger:model CreateMobileAuthorizationCodeRequest
type CreateMobileAuthorizationCodeRequest struct {

	// The Square location ID the authorization code should be tied to.
	LocationID string `json:"location_id,omitempty"`
}

// Validate validates this create mobile authorization code request
func (m *CreateMobileAuthorizationCodeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateMobileAuthorizationCodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMobileAuthorizationCodeRequest) UnmarshalBinary(b []byte) error {
	var res CreateMobileAuthorizationCodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
