// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderFulfillmentUpdatedUpdate Information about fulfillment updates.
//
// swagger:model OrderFulfillmentUpdatedUpdate
type OrderFulfillmentUpdatedUpdate struct {

	// Unique ID that identifies the fulfillment only within this order.
	FulfillmentUID string `json:"fulfillment_uid,omitempty"`

	// The state of the fulfillment after the change. May be equal to old_state if a non-state
	// field was changed on the fulfillment (e.g. tracking number).
	// See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values
	NewState string `json:"new_state,omitempty"`

	// The state of the fulfillment before the change.
	// Will not be populated if the fulfillment is created with this new Order version.
	// See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values
	OldState string `json:"old_state,omitempty"`
}

// Validate validates this order fulfillment updated update
func (m *OrderFulfillmentUpdatedUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this order fulfillment updated update based on context it is used
func (m *OrderFulfillmentUpdatedUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderFulfillmentUpdatedUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderFulfillmentUpdatedUpdate) UnmarshalBinary(b []byte) error {
	var res OrderFulfillmentUpdatedUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
