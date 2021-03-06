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

// V1SettlementEntryType v1 settlement entry type
//
// swagger:model V1SettlementEntryType
type V1SettlementEntryType string

const (

	// V1SettlementEntryTypeADJUSTMENT captures enum value "ADJUSTMENT"
	V1SettlementEntryTypeADJUSTMENT V1SettlementEntryType = "ADJUSTMENT"

	// V1SettlementEntryTypeBALANCECHARGE captures enum value "BALANCE_CHARGE"
	V1SettlementEntryTypeBALANCECHARGE V1SettlementEntryType = "BALANCE_CHARGE"

	// V1SettlementEntryTypeCHARGE captures enum value "CHARGE"
	V1SettlementEntryTypeCHARGE V1SettlementEntryType = "CHARGE"

	// V1SettlementEntryTypeFREEPROCESSING captures enum value "FREE_PROCESSING"
	V1SettlementEntryTypeFREEPROCESSING V1SettlementEntryType = "FREE_PROCESSING"

	// V1SettlementEntryTypeHOLDADJUSTMENT captures enum value "HOLD_ADJUSTMENT"
	V1SettlementEntryTypeHOLDADJUSTMENT V1SettlementEntryType = "HOLD_ADJUSTMENT"

	// V1SettlementEntryTypePAIDSERVICEFEE captures enum value "PAID_SERVICE_FEE"
	V1SettlementEntryTypePAIDSERVICEFEE V1SettlementEntryType = "PAID_SERVICE_FEE"

	// V1SettlementEntryTypePAIDSERVICEFEEREFUND captures enum value "PAID_SERVICE_FEE_REFUND"
	V1SettlementEntryTypePAIDSERVICEFEEREFUND V1SettlementEntryType = "PAID_SERVICE_FEE_REFUND"

	// V1SettlementEntryTypeREDEMPTIONCODE captures enum value "REDEMPTION_CODE"
	V1SettlementEntryTypeREDEMPTIONCODE V1SettlementEntryType = "REDEMPTION_CODE"

	// V1SettlementEntryTypeREFUND captures enum value "REFUND"
	V1SettlementEntryTypeREFUND V1SettlementEntryType = "REFUND"

	// V1SettlementEntryTypeRETURNEDPAYOUT captures enum value "RETURNED_PAYOUT"
	V1SettlementEntryTypeRETURNEDPAYOUT V1SettlementEntryType = "RETURNED_PAYOUT"

	// V1SettlementEntryTypeSQUARECAPITALADVANCE captures enum value "SQUARE_CAPITAL_ADVANCE"
	V1SettlementEntryTypeSQUARECAPITALADVANCE V1SettlementEntryType = "SQUARE_CAPITAL_ADVANCE"

	// V1SettlementEntryTypeSQUARECAPITALPAYMENT captures enum value "SQUARE_CAPITAL_PAYMENT"
	V1SettlementEntryTypeSQUARECAPITALPAYMENT V1SettlementEntryType = "SQUARE_CAPITAL_PAYMENT"

	// V1SettlementEntryTypeSQUARECAPITALREVERSEDPAYMENT captures enum value "SQUARE_CAPITAL_REVERSED_PAYMENT"
	V1SettlementEntryTypeSQUARECAPITALREVERSEDPAYMENT V1SettlementEntryType = "SQUARE_CAPITAL_REVERSED_PAYMENT"

	// V1SettlementEntryTypeSUBSCRIPTIONFEE captures enum value "SUBSCRIPTION_FEE"
	V1SettlementEntryTypeSUBSCRIPTIONFEE V1SettlementEntryType = "SUBSCRIPTION_FEE"

	// V1SettlementEntryTypeSUBSCRIPTIONFEEREFUND captures enum value "SUBSCRIPTION_FEE_REFUND"
	V1SettlementEntryTypeSUBSCRIPTIONFEEREFUND V1SettlementEntryType = "SUBSCRIPTION_FEE_REFUND"

	// V1SettlementEntryTypeOTHER captures enum value "OTHER"
	V1SettlementEntryTypeOTHER V1SettlementEntryType = "OTHER"

	// V1SettlementEntryTypeINCENTEDPAYMENT captures enum value "INCENTED_PAYMENT"
	V1SettlementEntryTypeINCENTEDPAYMENT V1SettlementEntryType = "INCENTED_PAYMENT"

	// V1SettlementEntryTypeRETURNEDACHENTRY captures enum value "RETURNED_ACH_ENTRY"
	V1SettlementEntryTypeRETURNEDACHENTRY V1SettlementEntryType = "RETURNED_ACH_ENTRY"

	// V1SettlementEntryTypeRETURNEDSQUARE275 captures enum value "RETURNED_SQUARE_275"
	V1SettlementEntryTypeRETURNEDSQUARE275 V1SettlementEntryType = "RETURNED_SQUARE_275"

	// V1SettlementEntryTypeSQUARE275 captures enum value "SQUARE_275"
	V1SettlementEntryTypeSQUARE275 V1SettlementEntryType = "SQUARE_275"

	// V1SettlementEntryTypeSQUARECARD captures enum value "SQUARE_CARD"
	V1SettlementEntryTypeSQUARECARD V1SettlementEntryType = "SQUARE_CARD"
)

// for schema
var v1SettlementEntryTypeEnum []interface{}

func init() {
	var res []V1SettlementEntryType
	if err := json.Unmarshal([]byte(`["ADJUSTMENT","BALANCE_CHARGE","CHARGE","FREE_PROCESSING","HOLD_ADJUSTMENT","PAID_SERVICE_FEE","PAID_SERVICE_FEE_REFUND","REDEMPTION_CODE","REFUND","RETURNED_PAYOUT","SQUARE_CAPITAL_ADVANCE","SQUARE_CAPITAL_PAYMENT","SQUARE_CAPITAL_REVERSED_PAYMENT","SUBSCRIPTION_FEE","SUBSCRIPTION_FEE_REFUND","OTHER","INCENTED_PAYMENT","RETURNED_ACH_ENTRY","RETURNED_SQUARE_275","SQUARE_275","SQUARE_CARD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SettlementEntryTypeEnum = append(v1SettlementEntryTypeEnum, v)
	}
}

func (m V1SettlementEntryType) validateV1SettlementEntryTypeEnum(path, location string, value V1SettlementEntryType) error {
	if err := validate.EnumCase(path, location, value, v1SettlementEntryTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 settlement entry type
func (m V1SettlementEntryType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SettlementEntryTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 settlement entry type based on context it is used
func (m V1SettlementEntryType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
