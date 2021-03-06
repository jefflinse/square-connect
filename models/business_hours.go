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

// BusinessHours Represents the hours of operation for a business location.
//
// swagger:model BusinessHours
type BusinessHours struct {

	// The list of time periods during which the business is open. There may be at most 10
	// periods per day.
	Periods []*BusinessHoursPeriod `json:"periods"`
}

// Validate validates this business hours
func (m *BusinessHours) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeriods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BusinessHours) validatePeriods(formats strfmt.Registry) error {
	if swag.IsZero(m.Periods) { // not required
		return nil
	}

	for i := 0; i < len(m.Periods); i++ {
		if swag.IsZero(m.Periods[i]) { // not required
			continue
		}

		if m.Periods[i] != nil {
			if err := m.Periods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this business hours based on the context it is used
func (m *BusinessHours) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeriods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BusinessHours) contextValidatePeriods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Periods); i++ {

		if m.Periods[i] != nil {
			if err := m.Periods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BusinessHours) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessHours) UnmarshalBinary(b []byte) error {
	var res BusinessHours
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
