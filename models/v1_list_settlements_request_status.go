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

// V1ListSettlementsRequestStatus v1 list settlements request status
//
// swagger:model V1ListSettlementsRequestStatus
type V1ListSettlementsRequestStatus string

const (

	// V1ListSettlementsRequestStatusSENT captures enum value "SENT"
	V1ListSettlementsRequestStatusSENT V1ListSettlementsRequestStatus = "SENT"

	// V1ListSettlementsRequestStatusFAILED captures enum value "FAILED"
	V1ListSettlementsRequestStatusFAILED V1ListSettlementsRequestStatus = "FAILED"
)

// for schema
var v1ListSettlementsRequestStatusEnum []interface{}

func init() {
	var res []V1ListSettlementsRequestStatus
	if err := json.Unmarshal([]byte(`["SENT","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ListSettlementsRequestStatusEnum = append(v1ListSettlementsRequestStatusEnum, v)
	}
}

func (m V1ListSettlementsRequestStatus) validateV1ListSettlementsRequestStatusEnum(path, location string, value V1ListSettlementsRequestStatus) error {
	if err := validate.Enum(path, location, value, v1ListSettlementsRequestStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 list settlements request status
func (m V1ListSettlementsRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ListSettlementsRequestStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}