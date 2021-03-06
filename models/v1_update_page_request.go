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

// V1UpdatePageRequest v1 update page request
//
// swagger:model V1UpdatePageRequest
type V1UpdatePageRequest struct {

	// An object containing the fields to POST for the request.
	//
	// See the corresponding object definition for field details.
	// Required: true
	Body *V1Page `json:"body"`
}

// Validate validates this v1 update page request
func (m *V1UpdatePageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdatePageRequest) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
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
func (m *V1UpdatePageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdatePageRequest) UnmarshalBinary(b []byte) error {
	var res V1UpdatePageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
