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

// V1PaymentItemizationItemizationType v1 payment itemization itemization type
//
// swagger:model V1PaymentItemizationItemizationType
type V1PaymentItemizationItemizationType string

const (

	// V1PaymentItemizationItemizationTypeITEM captures enum value "ITEM"
	V1PaymentItemizationItemizationTypeITEM V1PaymentItemizationItemizationType = "ITEM"

	// V1PaymentItemizationItemizationTypeCUSTOMAMOUNT captures enum value "CUSTOM_AMOUNT"
	V1PaymentItemizationItemizationTypeCUSTOMAMOUNT V1PaymentItemizationItemizationType = "CUSTOM_AMOUNT"

	// V1PaymentItemizationItemizationTypeGIFTCARDACTIVATION captures enum value "GIFT_CARD_ACTIVATION"
	V1PaymentItemizationItemizationTypeGIFTCARDACTIVATION V1PaymentItemizationItemizationType = "GIFT_CARD_ACTIVATION"

	// V1PaymentItemizationItemizationTypeGIFTCARDRELOAD captures enum value "GIFT_CARD_RELOAD"
	V1PaymentItemizationItemizationTypeGIFTCARDRELOAD V1PaymentItemizationItemizationType = "GIFT_CARD_RELOAD"

	// V1PaymentItemizationItemizationTypeGIFTCARDUNKNOWN captures enum value "GIFT_CARD_UNKNOWN"
	V1PaymentItemizationItemizationTypeGIFTCARDUNKNOWN V1PaymentItemizationItemizationType = "GIFT_CARD_UNKNOWN"

	// V1PaymentItemizationItemizationTypeOTHER captures enum value "OTHER"
	V1PaymentItemizationItemizationTypeOTHER V1PaymentItemizationItemizationType = "OTHER"
)

// for schema
var v1PaymentItemizationItemizationTypeEnum []interface{}

func init() {
	var res []V1PaymentItemizationItemizationType
	if err := json.Unmarshal([]byte(`["ITEM","CUSTOM_AMOUNT","GIFT_CARD_ACTIVATION","GIFT_CARD_RELOAD","GIFT_CARD_UNKNOWN","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PaymentItemizationItemizationTypeEnum = append(v1PaymentItemizationItemizationTypeEnum, v)
	}
}

func (m V1PaymentItemizationItemizationType) validateV1PaymentItemizationItemizationTypeEnum(path, location string, value V1PaymentItemizationItemizationType) error {
	if err := validate.Enum(path, location, value, v1PaymentItemizationItemizationTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 payment itemization itemization type
func (m V1PaymentItemizationItemizationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1PaymentItemizationItemizationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
