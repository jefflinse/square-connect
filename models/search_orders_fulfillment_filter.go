// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchOrdersFulfillmentFilter Filter based on [Order Fulfillment](#type-orderfulfillment) information.
//
// swagger:model SearchOrdersFulfillmentFilter
type SearchOrdersFulfillmentFilter struct {

	// List of `fulfillment states` to filter
	// for. Will return orders if any of its fulfillments match any of the
	// fulfillment states listed in this field.
	// See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values
	FulfillmentStates []string `json:"fulfillment_states"`

	// List of `fulfillment types` to filter
	// for. Will return orders if any of its fulfillments match any of the fulfillment types
	// listed in this field.
	// See [OrderFulfillmentType](#type-orderfulfillmenttype) for possible values
	// Required: true
	FulfillmentTypes []string `json:"fulfillment_types"`
}

// Validate validates this search orders fulfillment filter
func (m *SearchOrdersFulfillmentFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFulfillmentTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchOrdersFulfillmentFilter) validateFulfillmentTypes(formats strfmt.Registry) error {

	if err := validate.Required("fulfillment_types", "body", m.FulfillmentTypes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchOrdersFulfillmentFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchOrdersFulfillmentFilter) UnmarshalBinary(b []byte) error {
	var res SearchOrdersFulfillmentFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}