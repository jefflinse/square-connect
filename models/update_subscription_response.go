// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateSubscriptionResponse Defines the fields that are included in the response from the
// [UpdateSubscription](#endpoint-subscriptions-updatesubscription) endpoint.
//
// swagger:model UpdateSubscriptionResponse
type UpdateSubscriptionResponse struct {

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The modified `Subscription` object.
	Subscription *Subscription `json:"subscription,omitempty"`
}

// Validate validates this update subscription response
func (m *UpdateSubscriptionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSubscriptionResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *UpdateSubscriptionResponse) validateSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscription) { // not required
		return nil
	}

	if m.Subscription != nil {
		if err := m.Subscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSubscriptionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSubscriptionResponse) UnmarshalBinary(b []byte) error {
	var res UpdateSubscriptionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}