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

// RetrieveCashDrawerShiftResponse retrieve cash drawer shift response
//
// swagger:model RetrieveCashDrawerShiftResponse
type RetrieveCashDrawerShiftResponse struct {

	// The cash drawer shift queried for.
	CashDrawerShift *CashDrawerShift `json:"cash_drawer_shift,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this retrieve cash drawer shift response
func (m *RetrieveCashDrawerShiftResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCashDrawerShift(formats); err != nil {
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

func (m *RetrieveCashDrawerShiftResponse) validateCashDrawerShift(formats strfmt.Registry) error {

	if swag.IsZero(m.CashDrawerShift) { // not required
		return nil
	}

	if m.CashDrawerShift != nil {
		if err := m.CashDrawerShift.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cash_drawer_shift")
			}
			return err
		}
	}

	return nil
}

func (m *RetrieveCashDrawerShiftResponse) validateErrors(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RetrieveCashDrawerShiftResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveCashDrawerShiftResponse) UnmarshalBinary(b []byte) error {
	var res RetrieveCashDrawerShiftResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}