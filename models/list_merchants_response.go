// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListMerchantsResponse The response object returned by the [ListMerchant](#endpoint-listmerchant) endpoint.
//
// swagger:model ListMerchantsResponse
type ListMerchantsResponse struct {

	// If the  response is truncated, the cursor to use in next  request to fetch next set of objects.
	Cursor int64 `json:"cursor,omitempty"`

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The requested `Merchant` entities.
	Merchant []*Merchant `json:"merchant"`
}

// Validate validates this list merchants response
func (m *ListMerchantsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMerchant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListMerchantsResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListMerchantsResponse) validateMerchant(formats strfmt.Registry) error {

	if swag.IsZero(m.Merchant) { // not required
		return nil
	}

	for i := 0; i < len(m.Merchant); i++ {
		if swag.IsZero(m.Merchant[i]) { // not required
			continue
		}

		if m.Merchant[i] != nil {
			if err := m.Merchant[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("merchant" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListMerchantsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListMerchantsResponse) UnmarshalBinary(b []byte) error {
	var res ListMerchantsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
