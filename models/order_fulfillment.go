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

// OrderFulfillment Contains details on how to fulfill this order.
//
// swagger:model OrderFulfillment
type OrderFulfillment struct {

	// Application-defined data attached to this fulfillment. Metadata fields are intended
	// to store descriptive references or associations with an entity in another system or store brief
	// information about the object. Square does not process this field; it only stores and returns it
	// in relevant API calls. Do not use metadata to store any sensitive information (personally
	// identifiable information, card details, etc.).
	//
	// Keys written by applications must be 60 characters or less and must be in the character set
	// `[a-zA-Z0-9_-]`. Entries may also include metadata generated by Square. These keys are prefixed
	// with a namespace, separated from the key with a ':' character.
	//
	// Values have a max length of 255 characters.
	//
	// An application may have up to 10 entries per metadata field.
	//
	// Entries written by applications are private and can only be read or modified by the same
	// application.
	//
	// See [Metadata](https://developer.squareup.com/docs/build-basics/metadata) for more information.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Contains details for a pickup fulfillment. Required when fulfillment
	// type is `PICKUP`.
	PickupDetails *OrderFulfillmentPickupDetails `json:"pickup_details,omitempty"`

	// Contains details for a shipment fulfillment. Required when fulfillment type
	// is `SHIPMENT`.
	//
	// A shipment fulfillment's relationship to fulfillment `state`:
	// `PROPOSED`: A shipment is requested.
	// `RESERVED`: Fulfillment accepted. Shipment processing.
	// `PREPARED`: Shipment packaged. Shipping label created.
	// `COMPLETED`: Package has been shipped.
	// `CANCELED`: Shipment has been canceled.
	// `FAILED`: Shipment has failed.
	ShipmentDetails *OrderFulfillmentShipmentDetails `json:"shipment_details,omitempty"`

	// The state of the fulfillment.
	// See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values
	State string `json:"state,omitempty"`

	// The type of the fulfillment.
	// See [OrderFulfillmentType](#type-orderfulfillmenttype) for possible values
	Type string `json:"type,omitempty"`

	// Unique ID that identifies the fulfillment only within this order.
	// Max Length: 60
	UID string `json:"uid,omitempty"`
}

// Validate validates this order fulfillment
func (m *OrderFulfillment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePickupDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderFulfillment) validatePickupDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.PickupDetails) { // not required
		return nil
	}

	if m.PickupDetails != nil {
		if err := m.PickupDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickup_details")
			}
			return err
		}
	}

	return nil
}

func (m *OrderFulfillment) validateShipmentDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentDetails) { // not required
		return nil
	}

	if m.ShipmentDetails != nil {
		if err := m.ShipmentDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipment_details")
			}
			return err
		}
	}

	return nil
}

func (m *OrderFulfillment) validateUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.MaxLength("uid", "body", m.UID, 60); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order fulfillment based on the context it is used
func (m *OrderFulfillment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePickupDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderFulfillment) contextValidatePickupDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.PickupDetails != nil {
		if err := m.PickupDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickup_details")
			}
			return err
		}
	}

	return nil
}

func (m *OrderFulfillment) contextValidateShipmentDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentDetails != nil {
		if err := m.ShipmentDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipment_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderFulfillment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderFulfillment) UnmarshalBinary(b []byte) error {
	var res OrderFulfillment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
