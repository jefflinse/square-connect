// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteInvoiceRequest Describes a `DeleteInvoice` request.
//
// swagger:model DeleteInvoiceRequest
type DeleteInvoiceRequest struct {

	// The version of the `invoice` to delete.
	// If you do not know the version, you can call `GetInvoice` or
	// `ListInvoices`.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this delete invoice request
func (m *DeleteInvoiceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteInvoiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteInvoiceRequest) UnmarshalBinary(b []byte) error {
	var res DeleteInvoiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
