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

// CreateInvoiceRequest Describes a `CreateInvoice` request.
// Example: {"request_body":{"custom_fields":[{"label":"Event Reference Number","placement":"ABOVE_LINE_ITEMS","value":"Ref. #1234"},{"label":"Terms of Service","placement":"BELOW_LINE_ITEMS","value":"The terms of service are..."}],"idempotency_key":"ce3748f9-5fc1-4762-aa12-aae5e843f1f4","invoice":{"delivery_method":"EMAIL","description":"We appreciate your business!","invoice_number":"inv-100","location_id":"ES0RJRZYEC39A","order_id":"CAISENgvlJ6jLWAzERDzjyHVybY","payment_requests":[{"automatic_payment_source":"NONE","due_date":"2030-01-24","reminders":[{"message":"Your invoice is due tomorrow","relative_scheduled_days":-1}],"request_type":"BALANCE","tipping_enabled":true}],"primary_recipient":{"customer_id":"JDKYHBWT1D4F8MFH63DBMEN8Y4"},"scheduled_at":"2030-01-13T10:00:00Z","title":"Event Planning Services"}}}
//
// swagger:model CreateInvoiceRequest
type CreateInvoiceRequest struct {

	// A unique string that identifies the `CreateInvoice` request. If you do not
	// provide `idempotency_key` (or provide an empty string as the value), the endpoint
	// treats each request as independent.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/docs/working-with-apis/idempotency).
	// Max Length: 128
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// The invoice to create.
	// Required: true
	Invoice *Invoice `json:"invoice"`
}

// Validate validates this create invoice request
func (m *CreateInvoiceRequest) Validate(formats strfmt.Registry) error {
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

func (m *CreateInvoiceRequest) validateIdempotencyKey(formats strfmt.Registry) error {
	if swag.IsZero(m.IdempotencyKey) { // not required
		return nil
	}

	if err := validate.MaxLength("idempotency_key", "body", m.IdempotencyKey, 128); err != nil {
		return err
	}

	return nil
}

func (m *CreateInvoiceRequest) validateInvoice(formats strfmt.Registry) error {

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

// ContextValidate validate this create invoice request based on the context it is used
func (m *CreateInvoiceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvoice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateInvoiceRequest) contextValidateInvoice(ctx context.Context, formats strfmt.Registry) error {

	if m.Invoice != nil {
		if err := m.Invoice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoice")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateInvoiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateInvoiceRequest) UnmarshalBinary(b []byte) error {
	var res CreateInvoiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
