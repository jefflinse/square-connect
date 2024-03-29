// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateCustomerCardResponse Defines the fields that are included in the response body of
// a request to the CreateCustomerCard endpoint.
//
// One of `errors` or `card` is present in a given response (never both).
// Example: {"card":{"billing_address":{"address_line_1":"500 Electric Ave","address_line_2":"Suite 600","administrative_district_level_1":"NY","country":"US","locality":"New York","postal_code":"10003"},"card_brand":"VISA","cardholder_name":"Amelia Earhart","exp_month":11,"exp_year":2018,"id":"icard-card_id","last_4":"1111"}}
//
// swagger:model CreateCustomerCardResponse
type CreateCustomerCardResponse struct {

	// The created card on file.
	Card *Card `json:"card,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this create customer card response
func (m *CreateCustomerCardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCustomerCardResponse) validateCard(formats strfmt.Registry) error {
	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {
		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

func (m *CreateCustomerCardResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this create customer card response based on the context it is used
func (m *CreateCustomerCardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCustomerCardResponse) contextValidateCard(ctx context.Context, formats strfmt.Registry) error {

	if m.Card != nil {
		if err := m.Card.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("card")
			}
			return err
		}
	}

	return nil
}

func (m *CreateCustomerCardResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCustomerCardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCustomerCardResponse) UnmarshalBinary(b []byte) error {
	var res CreateCustomerCardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
