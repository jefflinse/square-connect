// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchOrdersStateFilter Filter by current Order `state`.
//
// swagger:model SearchOrdersStateFilter
type SearchOrdersStateFilter struct {

	// States to filter for.
	// See [OrderState](#type-orderstate) for possible values
	// Required: true
	States []string `json:"states"`
}

// Validate validates this search orders state filter
func (m *SearchOrdersStateFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchOrdersStateFilter) validateStates(formats strfmt.Registry) error {

	if err := validate.Required("states", "body", m.States); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this search orders state filter based on context it is used
func (m *SearchOrdersStateFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchOrdersStateFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchOrdersStateFilter) UnmarshalBinary(b []byte) error {
	var res SearchOrdersStateFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
