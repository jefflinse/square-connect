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

// DisputeState List of possible dispute states.
//
// swagger:model DisputeState
type DisputeState string

const (

	// DisputeStateUNKNOWNSTATE captures enum value "UNKNOWN_STATE"
	DisputeStateUNKNOWNSTATE DisputeState = "UNKNOWN_STATE"

	// DisputeStateINQUIRYEVIDENCEREQUIRED captures enum value "INQUIRY_EVIDENCE_REQUIRED"
	DisputeStateINQUIRYEVIDENCEREQUIRED DisputeState = "INQUIRY_EVIDENCE_REQUIRED"

	// DisputeStateINQUIRYPROCESSING captures enum value "INQUIRY_PROCESSING"
	DisputeStateINQUIRYPROCESSING DisputeState = "INQUIRY_PROCESSING"

	// DisputeStateINQUIRYCLOSED captures enum value "INQUIRY_CLOSED"
	DisputeStateINQUIRYCLOSED DisputeState = "INQUIRY_CLOSED"

	// DisputeStateEVIDENCEREQUIRED captures enum value "EVIDENCE_REQUIRED"
	DisputeStateEVIDENCEREQUIRED DisputeState = "EVIDENCE_REQUIRED"

	// DisputeStatePROCESSING captures enum value "PROCESSING"
	DisputeStatePROCESSING DisputeState = "PROCESSING"

	// DisputeStateWON captures enum value "WON"
	DisputeStateWON DisputeState = "WON"

	// DisputeStateLOST captures enum value "LOST"
	DisputeStateLOST DisputeState = "LOST"

	// DisputeStateACCEPTED captures enum value "ACCEPTED"
	DisputeStateACCEPTED DisputeState = "ACCEPTED"

	// DisputeStateWAITINGTHIRDPARTY captures enum value "WAITING_THIRD_PARTY"
	DisputeStateWAITINGTHIRDPARTY DisputeState = "WAITING_THIRD_PARTY"
)

// for schema
var disputeStateEnum []interface{}

func init() {
	var res []DisputeState
	if err := json.Unmarshal([]byte(`["UNKNOWN_STATE","INQUIRY_EVIDENCE_REQUIRED","INQUIRY_PROCESSING","INQUIRY_CLOSED","EVIDENCE_REQUIRED","PROCESSING","WON","LOST","ACCEPTED","WAITING_THIRD_PARTY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		disputeStateEnum = append(disputeStateEnum, v)
	}
}

func (m DisputeState) validateDisputeStateEnum(path, location string, value DisputeState) error {
	if err := validate.Enum(path, location, value, disputeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this dispute state
func (m DisputeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDisputeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}