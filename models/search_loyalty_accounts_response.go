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

// SearchLoyaltyAccountsResponse A response that includes loyalty accounts that satisfy the search criteria.
//
// swagger:model SearchLoyaltyAccountsResponse
type SearchLoyaltyAccountsResponse struct {

	// The pagination cursor to use in a subsequent
	// request. If empty, this is the final response.
	// For more information,
	// see [Pagination](https://developer.squareup.com/docs/docs/basics/api101/pagination).
	Cursor string `json:"cursor,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The loyalty accounts that met the search criteria,
	// in order of creation date.
	LoyaltyAccounts []*LoyaltyAccount `json:"loyalty_accounts"`
}

// Validate validates this search loyalty accounts response
func (m *SearchLoyaltyAccountsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoyaltyAccounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchLoyaltyAccountsResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *SearchLoyaltyAccountsResponse) validateLoyaltyAccounts(formats strfmt.Registry) error {

	if swag.IsZero(m.LoyaltyAccounts) { // not required
		return nil
	}

	for i := 0; i < len(m.LoyaltyAccounts); i++ {
		if swag.IsZero(m.LoyaltyAccounts[i]) { // not required
			continue
		}

		if m.LoyaltyAccounts[i] != nil {
			if err := m.LoyaltyAccounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("loyalty_accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchLoyaltyAccountsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchLoyaltyAccountsResponse) UnmarshalBinary(b []byte) error {
	var res SearchLoyaltyAccountsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
