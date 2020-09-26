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

// DeviceCodeStatus DeviceCode.Status enum.
//
// swagger:model DeviceCodeStatus
type DeviceCodeStatus string

const (

	// DeviceCodeStatusUNPAIRED captures enum value "UNPAIRED"
	DeviceCodeStatusUNPAIRED DeviceCodeStatus = "UNPAIRED"

	// DeviceCodeStatusPAIRED captures enum value "PAIRED"
	DeviceCodeStatusPAIRED DeviceCodeStatus = "PAIRED"
)

// for schema
var deviceCodeStatusEnum []interface{}

func init() {
	var res []DeviceCodeStatus
	if err := json.Unmarshal([]byte(`["UNPAIRED","PAIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceCodeStatusEnum = append(deviceCodeStatusEnum, v)
	}
}

func (m DeviceCodeStatus) validateDeviceCodeStatusEnum(path, location string, value DeviceCodeStatus) error {
	if err := validate.Enum(path, location, value, deviceCodeStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this device code status
func (m DeviceCodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceCodeStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}