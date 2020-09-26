// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListMerchantsRequest Request object for the [ListMerchant](#endpoint-listmerchant) endpoint.
//
// swagger:model ListMerchantsRequest
type ListMerchantsRequest struct {

	// The cursor generated by the previous response.
	Cursor int64 `json:"cursor,omitempty"`
}

// Validate validates this list merchants request
func (m *ListMerchantsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListMerchantsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListMerchantsRequest) UnmarshalBinary(b []byte) error {
	var res ListMerchantsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}