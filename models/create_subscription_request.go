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

// CreateSubscriptionRequest Defines parameters in a
// [CreateSubscription](#endpoint-subscriptions-createsubscription) endpoint request.
//
// swagger:model CreateSubscriptionRequest
type CreateSubscriptionRequest struct {

	// The date when the subscription should be canceled, in
	// YYYY-MM-DD format (for example, 2025-02-29). This overrides the plan configuration
	// if it comes before the date the subscription would otherwise end.
	CanceledDate string `json:"canceled_date,omitempty"`

	// The ID of the `customer](#type-customer) [card` to charge.
	// If not specified, Square sends an invoice via email. For an example to
	// create a customer and add a card on file, see [Subscriptions Walkthrough](https://developer.squareup.com/docs/docs/subscriptions-api/walkthrough).
	CardID string `json:"card_id,omitempty"`

	// The ID of the `customer` profile.
	// Required: true
	// Min Length: 1
	CustomerID *string `json:"customer_id"`

	// A unique string that identifies this `CreateSubscription` request.
	// If you do not provide a unique string (or provide an empty string as the value),
	// the endpoint treats each request as independent.
	//
	// For more information, see [Idempotency keys](https://developer.squareup.com/docs/docs/working-with-apis/idempotency).
	// Required: true
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`

	// The ID of the location the subscription is associated with.
	// Required: true
	// Min Length: 1
	LocationID *string `json:"location_id"`

	// The ID of the subscription plan. For more information, see
	// [Subscription Plan Overview](https://developer.squareup.com/docs/docs/subscriptions/overview).
	// Required: true
	// Min Length: 1
	PlanID *string `json:"plan_id"`

	// A custom price to apply for the subscription. If specified,
	// it overrides the price configured by the subscription plan.
	PriceOverrideMoney *Money `json:"price_override_money,omitempty"`

	// The start date of the subscription, in YYYY-MM-DD format. For example,
	// 2013-01-15. If the start date is left empty, the subscription begins
	// immediately.
	StartDate string `json:"start_date,omitempty"`

	// The tax to add when billing the subscription.
	// The percentage is expressed in decimal form, using a `'.'` as the decimal
	// separator and without a `'%'` sign. For example, a value of 7.5
	// corresponds to 7.5%.
	// Max Length: 10
	TaxPercentage string `json:"tax_percentage,omitempty"`

	// The timezone that is used in date calculations for the subscription. If unset, defaults to
	// the location timezone. If a timezone is not configured for the location, defaults to "America/New_York".
	// Format: the IANA Timezone Database identifier for the location timezone. For
	// a list of time zones, see [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this create subscription request
func (m *CreateSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceOverrideMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxPercentage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSubscriptionRequest) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	if err := validate.MinLength("customer_id", "body", string(*m.CustomerID), 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", string(*m.IdempotencyKey), 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	if err := validate.MinLength("location_id", "body", string(*m.LocationID), 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("plan_id", "body", m.PlanID); err != nil {
		return err
	}

	if err := validate.MinLength("plan_id", "body", string(*m.PlanID), 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validatePriceOverrideMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceOverrideMoney) { // not required
		return nil
	}

	if m.PriceOverrideMoney != nil {
		if err := m.PriceOverrideMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price_override_money")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSubscriptionRequest) validateTaxPercentage(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxPercentage) { // not required
		return nil
	}

	if err := validate.MaxLength("tax_percentage", "body", string(m.TaxPercentage), 10); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSubscriptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSubscriptionRequest) UnmarshalBinary(b []byte) error {
	var res CreateSubscriptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
