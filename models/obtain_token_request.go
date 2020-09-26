// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObtainTokenRequest obtain token request
//
// swagger:model ObtainTokenRequest
type ObtainTokenRequest struct {

	// The Square-issued ID of your application, available from the
	// [application dashboard](https://connect.squareup.com/apps).
	// Required: true
	ClientID *string `json:"client_id"`

	// The Square-issued application secret for your application, available
	// from the [application dashboard](https://connect.squareup.com/apps).
	// Required: true
	ClientSecret *string `json:"client_secret"`

	// The authorization code to exchange.
	// This is required if `grant_type` is set to `authorization_code`, to indicate that
	// the application wants to exchange an authorization code for an OAuth access token.
	Code string `json:"code,omitempty"`

	// Specifies the method to request an OAuth access token.
	// Valid values are: `authorization_code`, `refresh_token`, and `migration_token`
	// Required: true
	GrantType *string `json:"grant_type"`

	// Legacy OAuth access token obtained using a Connect API version prior
	// to 2019-03-13. This parameter is required if `grant_type` is set to
	// `migration_token` to indicate that the application wants to get a replacement
	// OAuth access token. The response also returns a refresh token.
	// For more information, see [Migrate to Using Refresh Tokens](https://developer.squareup.com/docs/authz/oauth/migration).
	MigrationToken string `json:"migration_token,omitempty"`

	// The redirect URL assigned in the [application dashboard](https://connect.squareup.com/apps).
	RedirectURI string `json:"redirect_uri,omitempty"`

	// A valid refresh token for generating a new OAuth access token.
	// A valid refresh token is required if `grant_type` is set to `refresh_token` ,
	// to indicate the application wants a replacement for an expired OAuth access token.
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Validate validates this obtain token request
func (m *ObtainTokenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObtainTokenRequest) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *ObtainTokenRequest) validateClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("client_secret", "body", m.ClientSecret); err != nil {
		return err
	}

	return nil
}

func (m *ObtainTokenRequest) validateGrantType(formats strfmt.Registry) error {

	if err := validate.Required("grant_type", "body", m.GrantType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObtainTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObtainTokenRequest) UnmarshalBinary(b []byte) error {
	var res ObtainTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}