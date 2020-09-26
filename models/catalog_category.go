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

// CatalogCategory A category to which a `CatalogItem` instance belongs.
//
// swagger:model CatalogCategory
type CatalogCategory struct {

	// The category name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	// Max Length: 255
	Name string `json:"name,omitempty"`
}

// Validate validates this catalog category
func (m *CatalogCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogCategory) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogCategory) UnmarshalBinary(b []byte) error {
	var res CatalogCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
