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

// TenderCardDetailsStatus Indicates the card transaction's current status.
//
// swagger:model TenderCardDetailsStatus
type TenderCardDetailsStatus string

const (

	// TenderCardDetailsStatusAUTHORIZED captures enum value "AUTHORIZED"
	TenderCardDetailsStatusAUTHORIZED TenderCardDetailsStatus = "AUTHORIZED"

	// TenderCardDetailsStatusCAPTURED captures enum value "CAPTURED"
	TenderCardDetailsStatusCAPTURED TenderCardDetailsStatus = "CAPTURED"

	// TenderCardDetailsStatusVOIDED captures enum value "VOIDED"
	TenderCardDetailsStatusVOIDED TenderCardDetailsStatus = "VOIDED"

	// TenderCardDetailsStatusFAILED captures enum value "FAILED"
	TenderCardDetailsStatusFAILED TenderCardDetailsStatus = "FAILED"
)

// for schema
var tenderCardDetailsStatusEnum []interface{}

func init() {
	var res []TenderCardDetailsStatus
	if err := json.Unmarshal([]byte(`["AUTHORIZED","CAPTURED","VOIDED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tenderCardDetailsStatusEnum = append(tenderCardDetailsStatusEnum, v)
	}
}

func (m TenderCardDetailsStatus) validateTenderCardDetailsStatusEnum(path, location string, value TenderCardDetailsStatus) error {
	if err := validate.EnumCase(path, location, value, tenderCardDetailsStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this tender card details status
func (m TenderCardDetailsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTenderCardDetailsStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tender card details status based on context it is used
func (m TenderCardDetailsStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
