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

// TenderType Indicates a tender's type.
//
// swagger:model TenderType
type TenderType string

const (

	// TenderTypeCARD captures enum value "CARD"
	TenderTypeCARD TenderType = "CARD"

	// TenderTypeCASH captures enum value "CASH"
	TenderTypeCASH TenderType = "CASH"

	// TenderTypeTHIRDPARTYCARD captures enum value "THIRD_PARTY_CARD"
	TenderTypeTHIRDPARTYCARD TenderType = "THIRD_PARTY_CARD"

	// TenderTypeSQUAREGIFTCARD captures enum value "SQUARE_GIFT_CARD"
	TenderTypeSQUAREGIFTCARD TenderType = "SQUARE_GIFT_CARD"

	// TenderTypeNOSALE captures enum value "NO_SALE"
	TenderTypeNOSALE TenderType = "NO_SALE"

	// TenderTypeBANKTRANSFER captures enum value "BANK_TRANSFER"
	TenderTypeBANKTRANSFER TenderType = "BANK_TRANSFER"

	// TenderTypeOTHER captures enum value "OTHER"
	TenderTypeOTHER TenderType = "OTHER"
)

// for schema
var tenderTypeEnum []interface{}

func init() {
	var res []TenderType
	if err := json.Unmarshal([]byte(`["CARD","CASH","THIRD_PARTY_CARD","SQUARE_GIFT_CARD","NO_SALE","BANK_TRANSFER","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tenderTypeEnum = append(tenderTypeEnum, v)
	}
}

func (m TenderType) validateTenderTypeEnum(path, location string, value TenderType) error {
	if err := validate.Enum(path, location, value, tenderTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tender type
func (m TenderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTenderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
