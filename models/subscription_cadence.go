// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SubscriptionCadence Determines the billing cadence of a `Subscription`
//
// swagger:model SubscriptionCadence
type SubscriptionCadence string

const (

	// SubscriptionCadenceDAILY captures enum value "DAILY"
	SubscriptionCadenceDAILY SubscriptionCadence = "DAILY"

	// SubscriptionCadenceWEEKLY captures enum value "WEEKLY"
	SubscriptionCadenceWEEKLY SubscriptionCadence = "WEEKLY"

	// SubscriptionCadenceEVERYTWOWEEKS captures enum value "EVERY_TWO_WEEKS"
	SubscriptionCadenceEVERYTWOWEEKS SubscriptionCadence = "EVERY_TWO_WEEKS"

	// SubscriptionCadenceTHIRTYDAYS captures enum value "THIRTY_DAYS"
	SubscriptionCadenceTHIRTYDAYS SubscriptionCadence = "THIRTY_DAYS"

	// SubscriptionCadenceSIXTYDAYS captures enum value "SIXTY_DAYS"
	SubscriptionCadenceSIXTYDAYS SubscriptionCadence = "SIXTY_DAYS"

	// SubscriptionCadenceNINETYDAYS captures enum value "NINETY_DAYS"
	SubscriptionCadenceNINETYDAYS SubscriptionCadence = "NINETY_DAYS"

	// SubscriptionCadenceMONTHLY captures enum value "MONTHLY"
	SubscriptionCadenceMONTHLY SubscriptionCadence = "MONTHLY"

	// SubscriptionCadenceEVERYTWOMONTHS captures enum value "EVERY_TWO_MONTHS"
	SubscriptionCadenceEVERYTWOMONTHS SubscriptionCadence = "EVERY_TWO_MONTHS"

	// SubscriptionCadenceQUARTERLY captures enum value "QUARTERLY"
	SubscriptionCadenceQUARTERLY SubscriptionCadence = "QUARTERLY"

	// SubscriptionCadenceEVERYFOURMONTHS captures enum value "EVERY_FOUR_MONTHS"
	SubscriptionCadenceEVERYFOURMONTHS SubscriptionCadence = "EVERY_FOUR_MONTHS"

	// SubscriptionCadenceEVERYSIXMONTHS captures enum value "EVERY_SIX_MONTHS"
	SubscriptionCadenceEVERYSIXMONTHS SubscriptionCadence = "EVERY_SIX_MONTHS"

	// SubscriptionCadenceANNUAL captures enum value "ANNUAL"
	SubscriptionCadenceANNUAL SubscriptionCadence = "ANNUAL"

	// SubscriptionCadenceEVERYTWOYEARS captures enum value "EVERY_TWO_YEARS"
	SubscriptionCadenceEVERYTWOYEARS SubscriptionCadence = "EVERY_TWO_YEARS"
)

// for schema
var subscriptionCadenceEnum []interface{}

func init() {
	var res []SubscriptionCadence
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","EVERY_TWO_WEEKS","THIRTY_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","EVERY_TWO_MONTHS","QUARTERLY","EVERY_FOUR_MONTHS","EVERY_SIX_MONTHS","ANNUAL","EVERY_TWO_YEARS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionCadenceEnum = append(subscriptionCadenceEnum, v)
	}
}

func (m SubscriptionCadence) validateSubscriptionCadenceEnum(path, location string, value SubscriptionCadence) error {
	if err := validate.EnumCase(path, location, value, subscriptionCadenceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this subscription cadence
func (m SubscriptionCadence) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSubscriptionCadenceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this subscription cadence based on context it is used
func (m SubscriptionCadence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
