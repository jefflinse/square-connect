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

// InvoiceCustomField An additional seller-defined and customer-facing field to include on the invoice. For more information,
// see [Custom fields](/docs/invoices-api/overview#custom-fields).
//
// swagger:model InvoiceCustomField
type InvoiceCustomField struct {

	// The label or title of the custom field. This field is required for a custom field.
	// Max Length: 30
	Label string `json:"label,omitempty"`

	// The location of the custom field on the invoice. This field is required for a custom field.
	// See [InvoiceCustomFieldPlacement](#type-invoicecustomfieldplacement) for possible values
	Placement string `json:"placement,omitempty"`

	// The text of the custom field. If omitted, only the label is rendered.
	// Max Length: 2000
	Value string `json:"value,omitempty"`
}

// Validate validates this invoice custom field
func (m *InvoiceCustomField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceCustomField) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 30); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceCustomField) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.MaxLength("value", "body", m.Value, 2000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this invoice custom field based on context it is used
func (m *InvoiceCustomField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceCustomField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceCustomField) UnmarshalBinary(b []byte) error {
	var res InvoiceCustomField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}