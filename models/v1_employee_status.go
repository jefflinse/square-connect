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

// V1EmployeeStatus v1 employee status
//
// swagger:model V1EmployeeStatus
type V1EmployeeStatus string

const (

	// V1EmployeeStatusACTIVE captures enum value "ACTIVE"
	V1EmployeeStatusACTIVE V1EmployeeStatus = "ACTIVE"

	// V1EmployeeStatusINACTIVE captures enum value "INACTIVE"
	V1EmployeeStatusINACTIVE V1EmployeeStatus = "INACTIVE"
)

// for schema
var v1EmployeeStatusEnum []interface{}

func init() {
	var res []V1EmployeeStatus
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EmployeeStatusEnum = append(v1EmployeeStatusEnum, v)
	}
}

func (m V1EmployeeStatus) validateV1EmployeeStatusEnum(path, location string, value V1EmployeeStatus) error {
	if err := validate.Enum(path, location, value, v1EmployeeStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 employee status
func (m V1EmployeeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1EmployeeStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}