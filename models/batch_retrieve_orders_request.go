// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BatchRetrieveOrdersRequest Defines the fields that are included in requests to the
// BatchRetrieveOrders endpoint.
//
// swagger:model BatchRetrieveOrdersRequest
type BatchRetrieveOrdersRequest struct {

	// The IDs of the orders to retrieve. A maximum of 100 orders can be retrieved per request.
	// Required: true
	OrderIds []string `json:"order_ids"`
}

// Validate validates this batch retrieve orders request
func (m *BatchRetrieveOrdersRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRetrieveOrdersRequest) validateOrderIds(formats strfmt.Registry) error {

	if err := validate.Required("order_ids", "body", m.OrderIds); err != nil {
		return err
	}

	for i := 0; i < len(m.OrderIds); i++ {

		if err := validate.MinLength("order_ids"+"."+strconv.Itoa(i), "body", string(m.OrderIds[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchRetrieveOrdersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchRetrieveOrdersRequest) UnmarshalBinary(b []byte) error {
	var res BatchRetrieveOrdersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}