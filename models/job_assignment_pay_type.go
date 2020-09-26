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

// JobAssignmentPayType Enumerates the possible pay types that a job can be assigned.
//
// swagger:model JobAssignmentPayType
type JobAssignmentPayType string

const (

	// JobAssignmentPayTypeNONE captures enum value "NONE"
	JobAssignmentPayTypeNONE JobAssignmentPayType = "NONE"

	// JobAssignmentPayTypeHOURLY captures enum value "HOURLY"
	JobAssignmentPayTypeHOURLY JobAssignmentPayType = "HOURLY"

	// JobAssignmentPayTypeSALARY captures enum value "SALARY"
	JobAssignmentPayTypeSALARY JobAssignmentPayType = "SALARY"
)

// for schema
var jobAssignmentPayTypeEnum []interface{}

func init() {
	var res []JobAssignmentPayType
	if err := json.Unmarshal([]byte(`["NONE","HOURLY","SALARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobAssignmentPayTypeEnum = append(jobAssignmentPayTypeEnum, v)
	}
}

func (m JobAssignmentPayType) validateJobAssignmentPayTypeEnum(path, location string, value JobAssignmentPayType) error {
	if err := validate.Enum(path, location, value, jobAssignmentPayTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this job assignment pay type
func (m JobAssignmentPayType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobAssignmentPayTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
