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

// CatalogQueryItemsForTax The query filter to return the items containing the specified tax IDs.
//
// swagger:model CatalogQueryItemsForTax
type CatalogQueryItemsForTax struct {

	// A set of `CatalogTax` IDs to be used to find associated `CatalogItem`s.
	// Required: true
	TaxIds []string `json:"tax_ids"`
}

// Validate validates this catalog query items for tax
func (m *CatalogQueryItemsForTax) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogQueryItemsForTax) validateTaxIds(formats strfmt.Registry) error {

	if err := validate.Required("tax_ids", "body", m.TaxIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this catalog query items for tax based on context it is used
func (m *CatalogQueryItemsForTax) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogQueryItemsForTax) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogQueryItemsForTax) UnmarshalBinary(b []byte) error {
	var res CatalogQueryItemsForTax
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
