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

// ListWorkweekConfigsResponse The response to a request for a set of `WorkweekConfig` objects. Contains
// the requested `WorkweekConfig` objects. May contain a set of `Error` objects if
// the request resulted in errors.
// Example: {"cursor":"2fofTniCgT0yIPAq26kmk0YyFQJZfbWkh73OOnlTHmTAx13NgED","workweek_configs":[{"created_at":"2016-02-04T00:58:24Z","id":"FY4VCAQN700GM","start_of_day_local_time":"10:00","start_of_week":"MON","updated_at":"2019-02-28T01:04:35Z","version":11}]}
//
// swagger:model ListWorkweekConfigsResponse
type ListWorkweekConfigsResponse struct {

	// Value supplied in the subsequent request to fetch the next page of
	// Employee Wage results.
	Cursor string `json:"cursor,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// A page of Employee Wage results.
	WorkweekConfigs []*WorkweekConfig `json:"workweek_configs"`
}

// Validate validates this list workweek configs response
func (m *ListWorkweekConfigsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkweekConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListWorkweekConfigsResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *ListWorkweekConfigsResponse) validateWorkweekConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkweekConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkweekConfigs); i++ {
		if swag.IsZero(m.WorkweekConfigs[i]) { // not required
			continue
		}

		if m.WorkweekConfigs[i] != nil {
			if err := m.WorkweekConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workweek_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list workweek configs response based on the context it is used
func (m *ListWorkweekConfigsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkweekConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListWorkweekConfigsResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ListWorkweekConfigsResponse) contextValidateWorkweekConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WorkweekConfigs); i++ {

		if m.WorkweekConfigs[i] != nil {
			if err := m.WorkweekConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workweek_configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListWorkweekConfigsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListWorkweekConfigsResponse) UnmarshalBinary(b []byte) error {
	var res ListWorkweekConfigsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
