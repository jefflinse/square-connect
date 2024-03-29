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

// CreateSubscriptionResponse Defines the fields that are included in the response from the
// [CreateSubscription](#endpoint-subscriptions-createsubscription) endpoint.
// Example: {"subscription":{"card_id":"ccof:qy5x8hHGYsgLrp4Q4GB","created_at":"2020-08-03T21:53:10Z","customer_id":"CHFGVKYY8RSV93M5KCYTG4PN0G","id":"56214fb2-cc85-47a1-93bc-44f3766bb56f","location_id":"S8GWD5R9QB376","plan_id":"6JHXF3B2CW3YKHDV4XEM674H","price_override_money":{"amount":100,"currency":"USD"},"start_date":"2020-08-01","status":"PENDING","tax_percentage":"5","timezone":"America/Los_Angeles","version":1594155459464}}
//
// swagger:model CreateSubscriptionResponse
type CreateSubscriptionResponse struct {

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The newly created subscription.
	//
	// For more information, see
	// [Subscription object](https://developer.squareup.com/docs/docs/subscriptions-api/overview#subscription-object).
	Subscription *Subscription `json:"subscription,omitempty"`
}

// Validate validates this create subscription response
func (m *CreateSubscriptionResponse) Validate(formats strfmt.Registry) error {
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

func (m *CreateSubscriptionResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *CreateSubscriptionResponse) validateSubscription(formats strfmt.Registry) error {
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

// ContextValidate validate this create subscription response based on the context it is used
func (m *CreateSubscriptionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSubscriptionResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateSubscriptionResponse) contextValidateSubscription(ctx context.Context, formats strfmt.Registry) error {

	if m.Subscription != nil {
		if err := m.Subscription.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscription")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSubscriptionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSubscriptionResponse) UnmarshalBinary(b []byte) error {
	var res CreateSubscriptionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
