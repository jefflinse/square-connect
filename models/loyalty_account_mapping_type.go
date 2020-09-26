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

// LoyaltyAccountMappingType The type of mapping.
//
// swagger:model LoyaltyAccountMappingType
type LoyaltyAccountMappingType string

const (

	// LoyaltyAccountMappingTypePHONE captures enum value "PHONE"
	LoyaltyAccountMappingTypePHONE LoyaltyAccountMappingType = "PHONE"
)

// for schema
var loyaltyAccountMappingTypeEnum []interface{}

func init() {
	var res []LoyaltyAccountMappingType
	if err := json.Unmarshal([]byte(`["PHONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loyaltyAccountMappingTypeEnum = append(loyaltyAccountMappingTypeEnum, v)
	}
}

func (m LoyaltyAccountMappingType) validateLoyaltyAccountMappingTypeEnum(path, location string, value LoyaltyAccountMappingType) error {
	if err := validate.Enum(path, location, value, loyaltyAccountMappingTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this loyalty account mapping type
func (m LoyaltyAccountMappingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLoyaltyAccountMappingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}