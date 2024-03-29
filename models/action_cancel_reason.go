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

// ActionCancelReason action cancel reason
//
// swagger:model ActionCancelReason
type ActionCancelReason string

const (

	// ActionCancelReasonBUYERCANCELED captures enum value "BUYER_CANCELED"
	ActionCancelReasonBUYERCANCELED ActionCancelReason = "BUYER_CANCELED"

	// ActionCancelReasonSELLERCANCELED captures enum value "SELLER_CANCELED"
	ActionCancelReasonSELLERCANCELED ActionCancelReason = "SELLER_CANCELED"

	// ActionCancelReasonTIMEDOUT captures enum value "TIMED_OUT"
	ActionCancelReasonTIMEDOUT ActionCancelReason = "TIMED_OUT"
)

// for schema
var actionCancelReasonEnum []interface{}

func init() {
	var res []ActionCancelReason
	if err := json.Unmarshal([]byte(`["BUYER_CANCELED","SELLER_CANCELED","TIMED_OUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionCancelReasonEnum = append(actionCancelReasonEnum, v)
	}
}

func (m ActionCancelReason) validateActionCancelReasonEnum(path, location string, value ActionCancelReason) error {
	if err := validate.EnumCase(path, location, value, actionCancelReasonEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this action cancel reason
func (m ActionCancelReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActionCancelReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this action cancel reason based on context it is used
func (m ActionCancelReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
