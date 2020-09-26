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

// PublishInvoiceRequest Describes a `PublishInvoice` request.
//
// swagger:model PublishInvoiceRequest
type PublishInvoiceRequest struct {

	// A unique string that identifies the `PublishInvoice` request. If you do not
	// provide `idempotency_key` (or provide an empty string as the value), the endpoint
	// treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/docs/working-with-apis/idempotency).
	// Max Length: 128
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The version of the `Invoice` to publish.
	// This must match the current version of the invoice,
	// otherwise the request is rejected.
	// Required: true
	Version *int64 `json:"version"`
}

// Validate validates this publish invoice request
func (m *PublishInvoiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublishInvoiceRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if swag.IsZero(m.IdempotencyKey) { // not required
		return nil
	}

	if err := validate.MaxLength("idempotency_key", "body", string(m.IdempotencyKey), 128); err != nil {
		return err
	}

	return nil
}

func (m *PublishInvoiceRequest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublishInvoiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublishInvoiceRequest) UnmarshalBinary(b []byte) error {
	var res PublishInvoiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}