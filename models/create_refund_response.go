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

// CreateRefundResponse Defines the fields that are included in the response body of
// a request to the [CreateRefund](#endpoint-createrefund) endpoint.
//
// One of `errors` or `refund` is present in a given response (never both).
//
// swagger:model CreateRefundResponse
type CreateRefundResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The created refund.
	Refund *Refund `json:"refund,omitempty"`
}

// Validate validates this create refund response
func (m *CreateRefundResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefund(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRefundResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *CreateRefundResponse) validateRefund(formats strfmt.Registry) error {

	if swag.IsZero(m.Refund) { // not required
		return nil
	}

	if m.Refund != nil {
		if err := m.Refund.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refund")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateRefundResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRefundResponse) UnmarshalBinary(b []byte) error {
	var res CreateRefundResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
