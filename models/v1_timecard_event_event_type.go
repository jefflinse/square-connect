// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1TimecardEventEventType Actions that resulted in a change to a timecard. All timecard
// events created with the Connect API have an event type that begins with
// `API`.
//
// swagger:model V1TimecardEventEventType
type V1TimecardEventEventType string

const (

	// V1TimecardEventEventTypeAPICREATE captures enum value "API_CREATE"
	V1TimecardEventEventTypeAPICREATE V1TimecardEventEventType = "API_CREATE"

	// V1TimecardEventEventTypeAPIEDIT captures enum value "API_EDIT"
	V1TimecardEventEventTypeAPIEDIT V1TimecardEventEventType = "API_EDIT"

	// V1TimecardEventEventTypeAPIDELETE captures enum value "API_DELETE"
	V1TimecardEventEventTypeAPIDELETE V1TimecardEventEventType = "API_DELETE"

	// V1TimecardEventEventTypeREGISTERCLOCKIN captures enum value "REGISTER_CLOCKIN"
	V1TimecardEventEventTypeREGISTERCLOCKIN V1TimecardEventEventType = "REGISTER_CLOCKIN"

	// V1TimecardEventEventTypeREGISTERCLOCKOUT captures enum value "REGISTER_CLOCKOUT"
	V1TimecardEventEventTypeREGISTERCLOCKOUT V1TimecardEventEventType = "REGISTER_CLOCKOUT"

	// V1TimecardEventEventTypeDASHBOARDSUPERVISORCLOSE captures enum value "DASHBOARD_SUPERVISOR_CLOSE"
	V1TimecardEventEventTypeDASHBOARDSUPERVISORCLOSE V1TimecardEventEventType = "DASHBOARD_SUPERVISOR_CLOSE"

	// V1TimecardEventEventTypeDASHBOARDEDIT captures enum value "DASHBOARD_EDIT"
	V1TimecardEventEventTypeDASHBOARDEDIT V1TimecardEventEventType = "DASHBOARD_EDIT"

	// V1TimecardEventEventTypeDASHBOARDDELETE captures enum value "DASHBOARD_DELETE"
	V1TimecardEventEventTypeDASHBOARDDELETE V1TimecardEventEventType = "DASHBOARD_DELETE"
)

// for schema
var v1TimecardEventEventTypeEnum []interface{}

func init() {
	var res []V1TimecardEventEventType
	if err := json.Unmarshal([]byte(`["API_CREATE","API_EDIT","API_DELETE","REGISTER_CLOCKIN","REGISTER_CLOCKOUT","DASHBOARD_SUPERVISOR_CLOSE","DASHBOARD_EDIT","DASHBOARD_DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1TimecardEventEventTypeEnum = append(v1TimecardEventEventTypeEnum, v)
	}
}

func (m V1TimecardEventEventType) validateV1TimecardEventEventTypeEnum(path, location string, value V1TimecardEventEventType) error {
	if err := validate.Enum(path, location, value, v1TimecardEventEventTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 timecard event event type
func (m V1TimecardEventEventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1TimecardEventEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}