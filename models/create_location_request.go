// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateLocationRequest Request object for the [CreateLocation](#endpoint-createlocation) endpoint.
// Example: {"request_body":{"location":{"address":{"address_line_1":"1234 Peachtree St. NE","administrative_district_level_1":"GA","locality":"Atlanta","postal_code":"30309"},"description":"My new location.","facebook_url":null,"name":"New location name"}}}
//
// swagger:model CreateLocationRequest
type CreateLocationRequest struct {

	// The initial values of the location being created. The `name` field is required and must be unique within a seller account.
	// All other fields are optional. Unspecified fields will be set to default
	// values using existing location data.
	Location *Location `json:"location,omitempty"`
}

// Validate validates this create location request
func (m *CreateLocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLocationRequest) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create location request based on the context it is used
func (m *CreateLocationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateLocationRequest) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateLocationRequest) UnmarshalBinary(b []byte) error {
	var res CreateLocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
