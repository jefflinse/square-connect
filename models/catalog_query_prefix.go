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

// CatalogQueryPrefix catalog query prefix
//
// swagger:model CatalogQueryPrefix
type CatalogQueryPrefix struct {

	// The name of the attribute to be searched.
	// Required: true
	// Min Length: 1
	AttributeName *string `json:"attribute_name"`

	// The desired prefix of the search attribute value.
	// Required: true
	// Min Length: 1
	AttributePrefix *string `json:"attribute_prefix"`
}

// Validate validates this catalog query prefix
func (m *CatalogQueryPrefix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogQueryPrefix) validateAttributeName(formats strfmt.Registry) error {

	if err := validate.Required("attribute_name", "body", m.AttributeName); err != nil {
		return err
	}

	if err := validate.MinLength("attribute_name", "body", string(*m.AttributeName), 1); err != nil {
		return err
	}

	return nil
}

func (m *CatalogQueryPrefix) validateAttributePrefix(formats strfmt.Registry) error {

	if err := validate.Required("attribute_prefix", "body", m.AttributePrefix); err != nil {
		return err
	}

	if err := validate.MinLength("attribute_prefix", "body", string(*m.AttributePrefix), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogQueryPrefix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogQueryPrefix) UnmarshalBinary(b []byte) error {
	var res CatalogQueryPrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
