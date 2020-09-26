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

// RegisterDomainResponseStatus The status of domain registration.
//
// swagger:model RegisterDomainResponseStatus
type RegisterDomainResponseStatus string

const (

	// RegisterDomainResponseStatusPENDING captures enum value "PENDING"
	RegisterDomainResponseStatusPENDING RegisterDomainResponseStatus = "PENDING"

	// RegisterDomainResponseStatusVERIFIED captures enum value "VERIFIED"
	RegisterDomainResponseStatusVERIFIED RegisterDomainResponseStatus = "VERIFIED"
)

// for schema
var registerDomainResponseStatusEnum []interface{}

func init() {
	var res []RegisterDomainResponseStatus
	if err := json.Unmarshal([]byte(`["PENDING","VERIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registerDomainResponseStatusEnum = append(registerDomainResponseStatusEnum, v)
	}
}

func (m RegisterDomainResponseStatus) validateRegisterDomainResponseStatusEnum(path, location string, value RegisterDomainResponseStatus) error {
	if err := validate.Enum(path, location, value, registerDomainResponseStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this register domain response status
func (m RegisterDomainResponseStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRegisterDomainResponseStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}