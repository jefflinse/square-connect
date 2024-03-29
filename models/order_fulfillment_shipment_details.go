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

// OrderFulfillmentShipmentDetails Contains details necessary to fulfill a shipment order.
//
// swagger:model OrderFulfillmentShipmentDetails
type OrderFulfillmentShipmentDetails struct {

	// A description of why the shipment was canceled.
	// Max Length: 100
	CancelReason string `json:"cancel_reason,omitempty"`

	// The [timestamp](#workingwithdates) indicating the shipment was canceled.
	// Must be in RFC 3339 timestamp format, e.g., "2016-09-04T23:59:33.123Z".
	CanceledAt string `json:"canceled_at,omitempty"`

	// The shipping carrier being used to ship this fulfillment
	// e.g. UPS, FedEx, USPS, etc.
	// Max Length: 50
	Carrier string `json:"carrier,omitempty"`

	// The [timestamp](#workingwithdates) indicating when the shipment is
	// expected to be delivered to the shipping carrier. Must be in RFC 3339 timestamp
	// format, e.g., "2016-09-04T23:59:33.123Z".
	ExpectedShippedAt string `json:"expected_shipped_at,omitempty"`

	// The [timestamp](#workingwithdates) indicating when the shipment
	// failed to be completed. Must be in RFC 3339 timestamp format, e.g.,
	// "2016-09-04T23:59:33.123Z".
	FailedAt string `json:"failed_at,omitempty"`

	// A description of why the shipment failed to be completed.
	// Max Length: 100
	FailureReason string `json:"failure_reason,omitempty"`

	// The [timestamp](#workingwithdates) indicating when this fulfillment was
	// moved to the `RESERVED` state. Indicates that preparation of this shipment has begun.
	// Must be in RFC 3339 timestamp format, e.g., "2016-09-04T23:59:33.123Z".
	InProgressAt string `json:"in_progress_at,omitempty"`

	// The [timestamp](#workingwithdates) indicating when this fulfillment
	// was moved to the `PREPARED` state. Indicates that the fulfillment is packaged.
	// Must be in RFC 3339 timestamp format, e.g., "2016-09-04T23:59:33.123Z".
	PackagedAt string `json:"packaged_at,omitempty"`

	// The [timestamp](#workingwithdates) indicating when the shipment was
	// requested. Must be in RFC 3339 timestamp format, e.g., "2016-09-04T23:59:33.123Z".
	PlacedAt string `json:"placed_at,omitempty"`

	// Information on the person meant to receive this shipment fulfillment.
	Recipient *OrderFulfillmentRecipient `json:"recipient,omitempty"`

	// The [timestamp](#workingwithdates) indicating when this fulfillment
	// was moved to the `COMPLETED`state. Indicates that the fulfillment has been given
	// to the shipping carrier. Must be in RFC 3339 timestamp format, e.g., "2016-09-04T23:59:33.123Z".
	ShippedAt string `json:"shipped_at,omitempty"`

	// A note with additional information for the shipping carrier.
	// Max Length: 500
	ShippingNote string `json:"shipping_note,omitempty"`

	// A description of the type of shipping product purchased from the carrier.
	// e.g. First Class, Priority, Express
	// Max Length: 50
	ShippingType string `json:"shipping_type,omitempty"`

	// The reference number provided by the carrier to track the shipment's progress.
	// Max Length: 100
	TrackingNumber string `json:"tracking_number,omitempty"`

	// A link to the tracking webpage on the carrier's website.
	// Max Length: 2000
	TrackingURL string `json:"tracking_url,omitempty"`
}

// Validate validates this order fulfillment shipment details
func (m *OrderFulfillmentShipmentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancelReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCarrier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailureReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateCancelReason(formats strfmt.Registry) error {
	if swag.IsZero(m.CancelReason) { // not required
		return nil
	}

	if err := validate.MaxLength("cancel_reason", "body", m.CancelReason, 100); err != nil {
		return err
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateCarrier(formats strfmt.Registry) error {
	if swag.IsZero(m.Carrier) { // not required
		return nil
	}

	if err := validate.MaxLength("carrier", "body", m.Carrier, 50); err != nil {
		return err
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateFailureReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FailureReason) { // not required
		return nil
	}

	if err := validate.MaxLength("failure_reason", "body", m.FailureReason, 100); err != nil {
		return err
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateRecipient(formats strfmt.Registry) error {
	if swag.IsZero(m.Recipient) { // not required
		return nil
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateShippingNote(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingNote) { // not required
		return nil
	}

	if err := validate.MaxLength("shipping_note", "body", m.ShippingNote, 500); err != nil {
		return err
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateShippingType(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingType) { // not required
		return nil
	}

	if err := validate.MaxLength("shipping_type", "body", m.ShippingType, 50); err != nil {
		return err
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateTrackingNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackingNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("tracking_number", "body", m.TrackingNumber, 100); err != nil {
		return err
	}

	return nil
}

func (m *OrderFulfillmentShipmentDetails) validateTrackingURL(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackingURL) { // not required
		return nil
	}

	if err := validate.MaxLength("tracking_url", "body", m.TrackingURL, 2000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order fulfillment shipment details based on the context it is used
func (m *OrderFulfillmentShipmentDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecipient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderFulfillmentShipmentDetails) contextValidateRecipient(ctx context.Context, formats strfmt.Registry) error {

	if m.Recipient != nil {
		if err := m.Recipient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderFulfillmentShipmentDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderFulfillmentShipmentDetails) UnmarshalBinary(b []byte) error {
	var res OrderFulfillmentShipmentDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
