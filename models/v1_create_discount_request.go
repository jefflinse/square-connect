// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateDiscountRequest v1 create discount request
//
// swagger:model V1CreateDiscountRequest
type V1CreateDiscountRequest struct {

	// An object containing the fields to POST for the request.
	//
	// See the corresponding object definition for field details.
	Body *V1Discount `json:"body,omitempty"`
}

// Validate validates this v1 create discount request
func (m *V1CreateDiscountRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateDiscountRequest) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateDiscountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateDiscountRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateDiscountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}