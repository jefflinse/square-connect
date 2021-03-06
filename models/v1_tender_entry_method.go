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

// V1TenderEntryMethod v1 tender entry method
//
// swagger:model V1TenderEntryMethod
type V1TenderEntryMethod string

const (

	// V1TenderEntryMethodMANUAL captures enum value "MANUAL"
	V1TenderEntryMethodMANUAL V1TenderEntryMethod = "MANUAL"

	// V1TenderEntryMethodSCANNED captures enum value "SCANNED"
	V1TenderEntryMethodSCANNED V1TenderEntryMethod = "SCANNED"

	// V1TenderEntryMethodSQUARECASH captures enum value "SQUARE_CASH"
	V1TenderEntryMethodSQUARECASH V1TenderEntryMethod = "SQUARE_CASH"

	// V1TenderEntryMethodSQUAREWALLET captures enum value "SQUARE_WALLET"
	V1TenderEntryMethodSQUAREWALLET V1TenderEntryMethod = "SQUARE_WALLET"

	// V1TenderEntryMethodSWIPED captures enum value "SWIPED"
	V1TenderEntryMethodSWIPED V1TenderEntryMethod = "SWIPED"

	// V1TenderEntryMethodWEBFORM captures enum value "WEB_FORM"
	V1TenderEntryMethodWEBFORM V1TenderEntryMethod = "WEB_FORM"

	// V1TenderEntryMethodOTHER captures enum value "OTHER"
	V1TenderEntryMethodOTHER V1TenderEntryMethod = "OTHER"
)

// for schema
var v1TenderEntryMethodEnum []interface{}

func init() {
	var res []V1TenderEntryMethod
	if err := json.Unmarshal([]byte(`["MANUAL","SCANNED","SQUARE_CASH","SQUARE_WALLET","SWIPED","WEB_FORM","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1TenderEntryMethodEnum = append(v1TenderEntryMethodEnum, v)
	}
}

func (m V1TenderEntryMethod) validateV1TenderEntryMethodEnum(path, location string, value V1TenderEntryMethod) error {
	if err := validate.Enum(path, location, value, v1TenderEntryMethodEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 tender entry method
func (m V1TenderEntryMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1TenderEntryMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
