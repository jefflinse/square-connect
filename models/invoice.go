// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Invoice Stores information about an invoice. You use the Invoices API to create and process
// invoices. For more information, see [Manage Invoices Using the Invoices API](/docs/invoices-api/overview).
//
// swagger:model Invoice
type Invoice struct {

	// The timestamp when the invoice was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The description of the invoice. This is visible the customer receiving the invoice.
	// Max Length: 65536
	// Min Length: 1
	Description string `json:"description,omitempty"`

	// The Square-assigned ID of the invoice.
	ID string `json:"id,omitempty"`

	// A user-friendly invoice number. The value is unique within a location.
	// If not provided when creating an invoice, Square assigns a value.
	// It increments from 1 and padded with zeros making it 7 characters long
	// for example, 0000001, 0000002.
	// Max Length: 191
	// Min Length: 1
	InvoiceNumber string `json:"invoice_number,omitempty"`

	// The ID of the location that this invoice is associated with.
	// This field is required when creating an invoice.
	// Max Length: 255
	// Min Length: 1
	LocationID string `json:"location_id,omitempty"`

	// The current amount due for the invoice. In addition to the
	// amount due on the next payment request, this also includes any overdue payment amounts.
	NextPaymentAmountMoney *Money `json:"next_payment_amount_money,omitempty"`

	// The ID of the `order` for which the invoice is created.
	//
	// This order must be in the `OPEN` state and must belong to the `location_id`
	// specified for this invoice. This field is required when creating an invoice.
	// Max Length: 255
	// Min Length: 1
	OrderID string `json:"order_id,omitempty"`

	// An array of `InvoicePaymentRequest` objects. Each object defines
	// a payment request in an invoice payment schedule. It provides information
	// such as when and how Square processes payments. You can specify maximum of
	// nine payment requests. All all the payment requests must specify the
	// same `request_method`.
	//
	// This field is required when creating an invoice.
	PaymentRequests []*InvoicePaymentRequest `json:"payment_requests"`

	// The customer who gets the invoice. Square uses the contact information
	// to deliver the invoice. This field is required when creating an invoice.
	PrimaryRecipient *InvoiceRecipient `json:"primary_recipient,omitempty"`

	// The URL of the Square-hosted invoice page.
	// After you publish the invoice using the `PublishInvoice` endpoint, Square hosts the invoice
	// page and returns the page URL in the response.
	PublicURL string `json:"public_url,omitempty"`

	// The timestamp when the invoice is scheduled for processing, in RFC 3339 format.
	// At the specified time, depending on the `request_method`, Square sends the
	// invoice to the customer's email address or charge the customer's card on file.
	//
	// If the field is not set, Square processes the invoice immediately after publication.
	ScheduledAt string `json:"scheduled_at,omitempty"`

	// The status of the invoice.
	// See [InvoiceStatus](#type-invoicestatus) for possible values
	Status string `json:"status,omitempty"`

	// The time zone of the date values (for example, `due_date`) specified in the invoice.
	Timezone string `json:"timezone,omitempty"`

	// The title of the invoice.
	// Max Length: 255
	// Min Length: 1
	Title string `json:"title,omitempty"`

	// The timestamp when the invoice was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	// The version number, which is incremented each time an update is committed to the invoice.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextPaymentAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 65536); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceNumber) { // not required
		return nil
	}

	if err := validate.MinLength("invoice_number", "body", string(m.InvoiceNumber), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("invoice_number", "body", string(m.InvoiceNumber), 191); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateLocationID(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.MinLength("location_id", "body", string(m.LocationID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("location_id", "body", string(m.LocationID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateNextPaymentAmountMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.NextPaymentAmountMoney) { // not required
		return nil
	}

	if m.NextPaymentAmountMoney != nil {
		if err := m.NextPaymentAmountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_payment_amount_money")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) validateOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderID) { // not required
		return nil
	}

	if err := validate.MinLength("order_id", "body", string(m.OrderID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("order_id", "body", string(m.OrderID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validatePaymentRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentRequests); i++ {
		if swag.IsZero(m.PaymentRequests[i]) { // not required
			continue
		}

		if m.PaymentRequests[i] != nil {
			if err := m.PaymentRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) validatePrimaryRecipient(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryRecipient) { // not required
		return nil
	}

	if m.PrimaryRecipient != nil {
		if err := m.PrimaryRecipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_recipient")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) validateTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", string(m.Title), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", string(m.Title), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoice) UnmarshalBinary(b []byte) error {
	var res Invoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
