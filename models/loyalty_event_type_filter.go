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

// LoyaltyEventTypeFilter Filter events by event type.
//
// swagger:model LoyaltyEventTypeFilter
type LoyaltyEventTypeFilter struct {

	// The loyalty event types used to filter the result.
	// If multiple values are specified, the endpoint uses a
	// logical OR to combine them.
	// See [LoyaltyEventType](#type-loyaltyeventtype) for possible values
	// Required: true
	Types []string `json:"types"`
}

// Validate validates this loyalty event type filter
func (m *LoyaltyEventTypeFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyEventTypeFilter) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("types", "body", m.Types); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyEventTypeFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyEventTypeFilter) UnmarshalBinary(b []byte) error {
	var res LoyaltyEventTypeFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}