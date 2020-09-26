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

// CustomerCreationSource Indicates the method used to create the customer profile.
//
// swagger:model CustomerCreationSource
type CustomerCreationSource string

const (

	// CustomerCreationSourceOTHER captures enum value "OTHER"
	CustomerCreationSourceOTHER CustomerCreationSource = "OTHER"

	// CustomerCreationSourceAPPOINTMENTS captures enum value "APPOINTMENTS"
	CustomerCreationSourceAPPOINTMENTS CustomerCreationSource = "APPOINTMENTS"

	// CustomerCreationSourceCOUPON captures enum value "COUPON"
	CustomerCreationSourceCOUPON CustomerCreationSource = "COUPON"

	// CustomerCreationSourceDELETIONRECOVERY captures enum value "DELETION_RECOVERY"
	CustomerCreationSourceDELETIONRECOVERY CustomerCreationSource = "DELETION_RECOVERY"

	// CustomerCreationSourceDIRECTORY captures enum value "DIRECTORY"
	CustomerCreationSourceDIRECTORY CustomerCreationSource = "DIRECTORY"

	// CustomerCreationSourceEGIFTING captures enum value "EGIFTING"
	CustomerCreationSourceEGIFTING CustomerCreationSource = "EGIFTING"

	// CustomerCreationSourceEMAILCOLLECTION captures enum value "EMAIL_COLLECTION"
	CustomerCreationSourceEMAILCOLLECTION CustomerCreationSource = "EMAIL_COLLECTION"

	// CustomerCreationSourceFEEDBACK captures enum value "FEEDBACK"
	CustomerCreationSourceFEEDBACK CustomerCreationSource = "FEEDBACK"

	// CustomerCreationSourceIMPORT captures enum value "IMPORT"
	CustomerCreationSourceIMPORT CustomerCreationSource = "IMPORT"

	// CustomerCreationSourceINVOICES captures enum value "INVOICES"
	CustomerCreationSourceINVOICES CustomerCreationSource = "INVOICES"

	// CustomerCreationSourceLOYALTY captures enum value "LOYALTY"
	CustomerCreationSourceLOYALTY CustomerCreationSource = "LOYALTY"

	// CustomerCreationSourceMARKETING captures enum value "MARKETING"
	CustomerCreationSourceMARKETING CustomerCreationSource = "MARKETING"

	// CustomerCreationSourceMERGE captures enum value "MERGE"
	CustomerCreationSourceMERGE CustomerCreationSource = "MERGE"

	// CustomerCreationSourceONLINESTORE captures enum value "ONLINE_STORE"
	CustomerCreationSourceONLINESTORE CustomerCreationSource = "ONLINE_STORE"

	// CustomerCreationSourceINSTANTPROFILE captures enum value "INSTANT_PROFILE"
	CustomerCreationSourceINSTANTPROFILE CustomerCreationSource = "INSTANT_PROFILE"

	// CustomerCreationSourceTERMINAL captures enum value "TERMINAL"
	CustomerCreationSourceTERMINAL CustomerCreationSource = "TERMINAL"

	// CustomerCreationSourceTHIRDPARTY captures enum value "THIRD_PARTY"
	CustomerCreationSourceTHIRDPARTY CustomerCreationSource = "THIRD_PARTY"

	// CustomerCreationSourceTHIRDPARTYIMPORT captures enum value "THIRD_PARTY_IMPORT"
	CustomerCreationSourceTHIRDPARTYIMPORT CustomerCreationSource = "THIRD_PARTY_IMPORT"

	// CustomerCreationSourceUNMERGERECOVERY captures enum value "UNMERGE_RECOVERY"
	CustomerCreationSourceUNMERGERECOVERY CustomerCreationSource = "UNMERGE_RECOVERY"
)

// for schema
var customerCreationSourceEnum []interface{}

func init() {
	var res []CustomerCreationSource
	if err := json.Unmarshal([]byte(`["OTHER","APPOINTMENTS","COUPON","DELETION_RECOVERY","DIRECTORY","EGIFTING","EMAIL_COLLECTION","FEEDBACK","IMPORT","INVOICES","LOYALTY","MARKETING","MERGE","ONLINE_STORE","INSTANT_PROFILE","TERMINAL","THIRD_PARTY","THIRD_PARTY_IMPORT","UNMERGE_RECOVERY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerCreationSourceEnum = append(customerCreationSourceEnum, v)
	}
}

func (m CustomerCreationSource) validateCustomerCreationSourceEnum(path, location string, value CustomerCreationSource) error {
	if err := validate.Enum(path, location, value, customerCreationSourceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this customer creation source
func (m CustomerCreationSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustomerCreationSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}