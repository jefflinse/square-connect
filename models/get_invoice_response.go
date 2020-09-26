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

// GetInvoiceResponse Describes a `GetInvoice` response.
//
// swagger:model GetInvoiceResponse
type GetInvoiceResponse struct {

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The invoice requested.
	Invoice *Invoice `json:"invoice,omitempty"`
}

// Validate validates this get invoice response
func (m *GetInvoiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInvoiceResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *GetInvoiceResponse) validateInvoice(formats strfmt.Registry) error {

	if swag.IsZero(m.Invoice) { // not required
		return nil
	}

	if m.Invoice != nil {
		if err := m.Invoice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetInvoiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInvoiceResponse) UnmarshalBinary(b []byte) error {
	var res GetInvoiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}