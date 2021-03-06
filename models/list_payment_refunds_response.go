// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListPaymentRefundsResponse Defines the fields that are included in the response body of
// a request to the [ListPaymentRefunds](#endpoint-refunds-listpaymentrefunds) endpoint.
//
// One of `errors` or `refunds` is present in a given response (never both).
//
// swagger:model ListPaymentRefundsResponse
type ListPaymentRefundsResponse struct {

	// The pagination cursor to be used in a subsequent request. If empty,
	// this is the final response.
	//
	// See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The list of requested refunds.
	Refunds []*PaymentRefund `json:"refunds"`
}

// Validate validates this list payment refunds response
func (m *ListPaymentRefundsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefunds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListPaymentRefundsResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListPaymentRefundsResponse) validateRefunds(formats strfmt.Registry) error {

	if swag.IsZero(m.Refunds) { // not required
		return nil
	}

	for i := 0; i < len(m.Refunds); i++ {
		if swag.IsZero(m.Refunds[i]) { // not required
			continue
		}

		if m.Refunds[i] != nil {
			if err := m.Refunds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("refunds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListPaymentRefundsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListPaymentRefundsResponse) UnmarshalBinary(b []byte) error {
	var res ListPaymentRefundsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
