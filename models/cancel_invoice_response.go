// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CancelInvoiceResponse The response returned by the `CancelInvoice` request.
// Example: {"invoice":{"created_at":"2020-06-18T17:45:13Z","custom_fields":[{"label":"Event Reference Number","placement":"ABOVE_LINE_ITEMS","value":"Ref. #1234"},{"label":"Terms of Service","placement":"BELOW_LINE_ITEMS","value":"The terms of service are..."}],"delivery_method":"EMAIL","description":"We appreciate your business!","id":"gt2zv1z6mnUm1V7KMxxf3w","invoice_number":"inv-100","location_id":"ES0RJRZYEC39A","order_id":"CAISENgvlJ6jLWAzERDzjyHVybY","payment_requests":[{"automatic_payment_source":"NONE","computed_amount_money":{"amount":10000,"currency":"USD"},"due_date":"2030-01-24","reminders":[{"message":"Your invoice is due tomorrow","relative_scheduled_days":-1,"status":"PENDING","uid":"beebd363-e47f-4075-8785-c235aaa7df11"}],"request_type":"BALANCE","tipping_enabled":true,"total_completed_amount_money":{"amount":0,"currency":"USD"},"uid":"2da7964f-f3d2-4f43-81e8-5aa220bf3355"}],"primary_recipient":{"customer_id":"JDKYHBWT1D4F8MFH63DBMEN8Y4","email_address":"Amelia.Earhart@example.com","family_name":"Earhart","given_name":"Amelia","phone_number":"1-212-555-4240"},"scheduled_at":"2030-01-13T10:00:00Z","status":"CANCELED","timezone":"America/Los_Angeles","title":"Event Planning Services","updated_at":"2020-06-18T18:23:11Z","version":1}}
//
// swagger:model CancelInvoiceResponse
type CancelInvoiceResponse struct {

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The canceled invoice.
	Invoice *Invoice `json:"invoice,omitempty"`
}

// Validate validates this cancel invoice response
func (m *CancelInvoiceResponse) Validate(formats strfmt.Registry) error {
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

func (m *CancelInvoiceResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *CancelInvoiceResponse) validateInvoice(formats strfmt.Registry) error {
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

// ContextValidate validate this cancel invoice response based on the context it is used
func (m *CancelInvoiceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvoice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelInvoiceResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CancelInvoiceResponse) contextValidateInvoice(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CancelInvoiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelInvoiceResponse) UnmarshalBinary(b []byte) error {
	var res CancelInvoiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
