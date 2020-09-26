// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObtainTokenResponse obtain token response
//
// swagger:model ObtainTokenResponse
type ObtainTokenResponse struct {

	// A valid OAuth access token. OAuth access tokens are 64 bytes long.
	// Provide the access token in a header with every request to Connect API
	// endpoints. See the [Build with OAuth](https://developer.squareup.com/docs/authz/oauth/build-with-the-api) guide
	// for more information.
	AccessToken string `json:"access_token,omitempty"`

	// The date when access_token expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format.
	ExpiresAt string `json:"expires_at,omitempty"`

	// Then OpenID token belonging to this this person. Only present if the
	// OPENID scope is included in the authorize request.
	IDToken string `json:"id_token,omitempty"`

	// The ID of the authorizing merchant's business.
	MerchantID string `json:"merchant_id,omitempty"`

	// T__LEGACY FIELD__. The ID of the subscription plan the merchant signed
	// up for. Only present if the merchant signed up for a subscription during
	// authorization.
	PlanID string `json:"plan_id,omitempty"`

	// A refresh token. OAuth refresh tokens are 64 bytes long.
	// For more information, see [OAuth access token management](https://developer.squareup.com/docs/authz/oauth/how-it-works#oauth-access-token-management).
	RefreshToken string `json:"refresh_token,omitempty"`

	// __LEGACY FIELD__. The ID of a subscription plan the merchant signed up
	// for. Only present if the merchant signed up for a subscription during authorization.
	SubscriptionID string `json:"subscription_id,omitempty"`

	// This value is always _bearer_.
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this obtain token response
func (m *ObtainTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObtainTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObtainTokenResponse) UnmarshalBinary(b []byte) error {
	var res ObtainTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}