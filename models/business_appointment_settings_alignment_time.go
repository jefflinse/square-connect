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

// BusinessAppointmentSettingsAlignmentTime Time units of a service duration for bookings.
//
// swagger:model BusinessAppointmentSettingsAlignmentTime
type BusinessAppointmentSettingsAlignmentTime string

const (

	// BusinessAppointmentSettingsAlignmentTimeSERVICEDURATION captures enum value "SERVICE_DURATION"
	BusinessAppointmentSettingsAlignmentTimeSERVICEDURATION BusinessAppointmentSettingsAlignmentTime = "SERVICE_DURATION"

	// BusinessAppointmentSettingsAlignmentTimeQUARTERHOURLY captures enum value "QUARTER_HOURLY"
	BusinessAppointmentSettingsAlignmentTimeQUARTERHOURLY BusinessAppointmentSettingsAlignmentTime = "QUARTER_HOURLY"

	// BusinessAppointmentSettingsAlignmentTimeHALFHOURLY captures enum value "HALF_HOURLY"
	BusinessAppointmentSettingsAlignmentTimeHALFHOURLY BusinessAppointmentSettingsAlignmentTime = "HALF_HOURLY"

	// BusinessAppointmentSettingsAlignmentTimeHOURLY captures enum value "HOURLY"
	BusinessAppointmentSettingsAlignmentTimeHOURLY BusinessAppointmentSettingsAlignmentTime = "HOURLY"
)

// for schema
var businessAppointmentSettingsAlignmentTimeEnum []interface{}

func init() {
	var res []BusinessAppointmentSettingsAlignmentTime
	if err := json.Unmarshal([]byte(`["SERVICE_DURATION","QUARTER_HOURLY","HALF_HOURLY","HOURLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		businessAppointmentSettingsAlignmentTimeEnum = append(businessAppointmentSettingsAlignmentTimeEnum, v)
	}
}

func (m BusinessAppointmentSettingsAlignmentTime) validateBusinessAppointmentSettingsAlignmentTimeEnum(path, location string, value BusinessAppointmentSettingsAlignmentTime) error {
	if err := validate.EnumCase(path, location, value, businessAppointmentSettingsAlignmentTimeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this business appointment settings alignment time
func (m BusinessAppointmentSettingsAlignmentTime) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBusinessAppointmentSettingsAlignmentTimeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this business appointment settings alignment time based on context it is used
func (m BusinessAppointmentSettingsAlignmentTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
