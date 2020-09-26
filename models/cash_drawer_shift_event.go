// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CashDrawerShiftEvent cash drawer shift event
//
// swagger:model CashDrawerShiftEvent
type CashDrawerShiftEvent struct {

	// The event time in ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`

	// An optional description of the event, entered by the employee that
	// created the event.
	Description string `json:"description,omitempty"`

	// The ID of the employee that created the event.
	EmployeeID string `json:"employee_id,omitempty"`

	// The amount of money that was added to or removed from the cash drawer
	// in the event. The amount can be positive (for added money), negative
	// (for removed money), or zero (for other tender type payments).
	EventMoney *Money `json:"event_money,omitempty"`

	// The type of cash drawer shift event.
	// See [CashDrawerEventType](#type-cashdrawereventtype) for possible values
	EventType string `json:"event_type,omitempty"`

	// The unique ID of the event.
	ID string `json:"id,omitempty"`
}

// Validate validates this cash drawer shift event
func (m *CashDrawerShiftEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CashDrawerShiftEvent) validateEventMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.EventMoney) { // not required
		return nil
	}

	if m.EventMoney != nil {
		if err := m.EventMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CashDrawerShiftEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CashDrawerShiftEvent) UnmarshalBinary(b []byte) error {
	var res CashDrawerShiftEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}