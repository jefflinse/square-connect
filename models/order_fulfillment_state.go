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

// OrderFulfillmentState The current state of this fulfillment.
//
// swagger:model OrderFulfillmentState
type OrderFulfillmentState string

const (

	// OrderFulfillmentStatePROPOSED captures enum value "PROPOSED"
	OrderFulfillmentStatePROPOSED OrderFulfillmentState = "PROPOSED"

	// OrderFulfillmentStateRESERVED captures enum value "RESERVED"
	OrderFulfillmentStateRESERVED OrderFulfillmentState = "RESERVED"

	// OrderFulfillmentStatePREPARED captures enum value "PREPARED"
	OrderFulfillmentStatePREPARED OrderFulfillmentState = "PREPARED"

	// OrderFulfillmentStateCOMPLETED captures enum value "COMPLETED"
	OrderFulfillmentStateCOMPLETED OrderFulfillmentState = "COMPLETED"

	// OrderFulfillmentStateCANCELED captures enum value "CANCELED"
	OrderFulfillmentStateCANCELED OrderFulfillmentState = "CANCELED"

	// OrderFulfillmentStateFAILED captures enum value "FAILED"
	OrderFulfillmentStateFAILED OrderFulfillmentState = "FAILED"
)

// for schema
var orderFulfillmentStateEnum []interface{}

func init() {
	var res []OrderFulfillmentState
	if err := json.Unmarshal([]byte(`["PROPOSED","RESERVED","PREPARED","COMPLETED","CANCELED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderFulfillmentStateEnum = append(orderFulfillmentStateEnum, v)
	}
}

func (m OrderFulfillmentState) validateOrderFulfillmentStateEnum(path, location string, value OrderFulfillmentState) error {
	if err := validate.Enum(path, location, value, orderFulfillmentStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this order fulfillment state
func (m OrderFulfillmentState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderFulfillmentStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}