// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogQueryText The query filter to return the search result whose searchable attribute values contain all of the specified keywords or tokens, independent of the token order or case.
//
// swagger:model CatalogQueryText
type CatalogQueryText struct {

	// A list of 1, 2, or 3 search keywords. Keywords with fewer than 3 characters are ignored.
	// Required: true
	Keywords []string `json:"keywords"`
}

// Validate validates this catalog query text
func (m *CatalogQueryText) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogQueryText) validateKeywords(formats strfmt.Registry) error {

	if err := validate.Required("keywords", "body", m.Keywords); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this catalog query text based on context it is used
func (m *CatalogQueryText) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogQueryText) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogQueryText) UnmarshalBinary(b []byte) error {
	var res CatalogQueryText
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
