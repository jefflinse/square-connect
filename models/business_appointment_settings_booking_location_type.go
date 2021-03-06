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

// BusinessAppointmentSettingsBookingLocationType Types of location where service is provided.
//
// swagger:model BusinessAppointmentSettingsBookingLocationType
type BusinessAppointmentSettingsBookingLocationType string

const (

	// BusinessAppointmentSettingsBookingLocationTypeBUSINESSLOCATION captures enum value "BUSINESS_LOCATION"
	BusinessAppointmentSettingsBookingLocationTypeBUSINESSLOCATION BusinessAppointmentSettingsBookingLocationType = "BUSINESS_LOCATION"

	// BusinessAppointmentSettingsBookingLocationTypeCUSTOMERLOCATION captures enum value "CUSTOMER_LOCATION"
	BusinessAppointmentSettingsBookingLocationTypeCUSTOMERLOCATION BusinessAppointmentSettingsBookingLocationType = "CUSTOMER_LOCATION"

	// BusinessAppointmentSettingsBookingLocationTypePHONE captures enum value "PHONE"
	BusinessAppointmentSettingsBookingLocationTypePHONE BusinessAppointmentSettingsBookingLocationType = "PHONE"
)

// for schema
var businessAppointmentSettingsBookingLocationTypeEnum []interface{}

func init() {
	var res []BusinessAppointmentSettingsBookingLocationType
	if err := json.Unmarshal([]byte(`["BUSINESS_LOCATION","CUSTOMER_LOCATION","PHONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		businessAppointmentSettingsBookingLocationTypeEnum = append(businessAppointmentSettingsBookingLocationTypeEnum, v)
	}
}

func (m BusinessAppointmentSettingsBookingLocationType) validateBusinessAppointmentSettingsBookingLocationTypeEnum(path, location string, value BusinessAppointmentSettingsBookingLocationType) error {
	if err := validate.EnumCase(path, location, value, businessAppointmentSettingsBookingLocationTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this business appointment settings booking location type
func (m BusinessAppointmentSettingsBookingLocationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBusinessAppointmentSettingsBookingLocationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this business appointment settings booking location type based on context it is used
func (m BusinessAppointmentSettingsBookingLocationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
