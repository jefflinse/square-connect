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

// MerchantStatus merchant status
//
// swagger:model MerchantStatus
type MerchantStatus string

const (

	// MerchantStatusACTIVE captures enum value "ACTIVE"
	MerchantStatusACTIVE MerchantStatus = "ACTIVE"

	// MerchantStatusINACTIVE captures enum value "INACTIVE"
	MerchantStatusINACTIVE MerchantStatus = "INACTIVE"
)

// for schema
var merchantStatusEnum []interface{}

func init() {
	var res []MerchantStatus
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		merchantStatusEnum = append(merchantStatusEnum, v)
	}
}

func (m MerchantStatus) validateMerchantStatusEnum(path, location string, value MerchantStatus) error {
	if err := validate.EnumCase(path, location, value, merchantStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this merchant status
func (m MerchantStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMerchantStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this merchant status based on context it is used
func (m MerchantStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
