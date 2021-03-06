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

// BusinessBookingProfileBookingPolicy Policies for accepting bookings.
//
// swagger:model BusinessBookingProfileBookingPolicy
type BusinessBookingProfileBookingPolicy string

const (

	// BusinessBookingProfileBookingPolicyACCEPTALL captures enum value "ACCEPT_ALL"
	BusinessBookingProfileBookingPolicyACCEPTALL BusinessBookingProfileBookingPolicy = "ACCEPT_ALL"

	// BusinessBookingProfileBookingPolicyREQUIRESACCEPTANCE captures enum value "REQUIRES_ACCEPTANCE"
	BusinessBookingProfileBookingPolicyREQUIRESACCEPTANCE BusinessBookingProfileBookingPolicy = "REQUIRES_ACCEPTANCE"
)

// for schema
var businessBookingProfileBookingPolicyEnum []interface{}

func init() {
	var res []BusinessBookingProfileBookingPolicy
	if err := json.Unmarshal([]byte(`["ACCEPT_ALL","REQUIRES_ACCEPTANCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		businessBookingProfileBookingPolicyEnum = append(businessBookingProfileBookingPolicyEnum, v)
	}
}

func (m BusinessBookingProfileBookingPolicy) validateBusinessBookingProfileBookingPolicyEnum(path, location string, value BusinessBookingProfileBookingPolicy) error {
	if err := validate.EnumCase(path, location, value, businessBookingProfileBookingPolicyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this business booking profile booking policy
func (m BusinessBookingProfileBookingPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBusinessBookingProfileBookingPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this business booking profile booking policy based on context it is used
func (m BusinessBookingProfileBookingPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
