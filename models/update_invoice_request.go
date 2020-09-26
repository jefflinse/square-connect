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

// UpdateInvoiceRequest Describes a `UpdateInvoice` request.
//
// swagger:model UpdateInvoiceRequest
type UpdateInvoiceRequest struct {

	// List of fields to clear.
	// For examples, see [Update an invoice](https://developer.squareup.com/docs/docs/invoices-api/overview#update-an-invoice).
	FieldsToClear []string `json:"fields_to_clear"`

	// A unique string that identifies the `UpdateInvoice` request. If you do not
	// provide `idempotency_key` (or provide an empty string as the value), the endpoint
	// treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/docs/working-with-apis/idempotency).
	// Max Length: 128
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The invoice fields to update. You need to only specify the fields you want to change.
	// The current invoice version must be specified in the version field. For more information,
	// see [Update an invoice](invoices-api/overview#update-an-invoice).
	// Required: true
	Invoice *Invoice `json:"invoice"`
}

// Validate validates this update invoice request
func (m *UpdateInvoiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdempotencyKey(formats); err != nil {
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

func (m *UpdateInvoiceRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if swag.IsZero(m.IdempotencyKey) { // not required
		return nil
	}

	if err := validate.MaxLength("idempotency_key", "body", string(m.IdempotencyKey), 128); err != nil {
		return err
	}

	return nil
}

func (m *UpdateInvoiceRequest) validateInvoice(formats strfmt.Registry) error {

	if err := validate.Required("invoice", "body", m.Invoice); err != nil {
		return err
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
func (m *UpdateInvoiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateInvoiceRequest) UnmarshalBinary(b []byte) error {
	var res UpdateInvoiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
