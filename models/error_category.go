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

// ErrorCategory Indicates which high-level category of error has occurred during a
// request to the Connect API.
//
// swagger:model ErrorCategory
type ErrorCategory string

const (

	// ErrorCategoryAPIERROR captures enum value "API_ERROR"
	ErrorCategoryAPIERROR ErrorCategory = "API_ERROR"

	// ErrorCategoryAUTHENTICATIONERROR captures enum value "AUTHENTICATION_ERROR"
	ErrorCategoryAUTHENTICATIONERROR ErrorCategory = "AUTHENTICATION_ERROR"

	// ErrorCategoryINVALIDREQUESTERROR captures enum value "INVALID_REQUEST_ERROR"
	ErrorCategoryINVALIDREQUESTERROR ErrorCategory = "INVALID_REQUEST_ERROR"

	// ErrorCategoryRATELIMITERROR captures enum value "RATE_LIMIT_ERROR"
	ErrorCategoryRATELIMITERROR ErrorCategory = "RATE_LIMIT_ERROR"

	// ErrorCategoryPAYMENTMETHODERROR captures enum value "PAYMENT_METHOD_ERROR"
	ErrorCategoryPAYMENTMETHODERROR ErrorCategory = "PAYMENT_METHOD_ERROR"

	// ErrorCategoryREFUNDERROR captures enum value "REFUND_ERROR"
	ErrorCategoryREFUNDERROR ErrorCategory = "REFUND_ERROR"
)

// for schema
var errorCategoryEnum []interface{}

func init() {
	var res []ErrorCategory
	if err := json.Unmarshal([]byte(`["API_ERROR","AUTHENTICATION_ERROR","INVALID_REQUEST_ERROR","RATE_LIMIT_ERROR","PAYMENT_METHOD_ERROR","REFUND_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorCategoryEnum = append(errorCategoryEnum, v)
	}
}

func (m ErrorCategory) validateErrorCategoryEnum(path, location string, value ErrorCategory) error {
	if err := validate.EnumCase(path, location, value, errorCategoryEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this error category
func (m ErrorCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateErrorCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this error category based on context it is used
func (m ErrorCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
