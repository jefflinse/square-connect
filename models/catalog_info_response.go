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

// CatalogInfoResponse catalog info response
//
// swagger:model CatalogInfoResponse
type CatalogInfoResponse struct {

	// The set of errors encountered.
	Errors []*Error `json:"errors"`

	// limits
	Limits *CatalogInfoResponseLimits `json:"limits,omitempty"`

	// Names and abbreviations for standard units.
	StandardUnitDescriptionGroup *StandardUnitDescriptionGroup `json:"standard_unit_description_group,omitempty"`
}

// Validate validates this catalog info response
func (m *CatalogInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardUnitDescriptionGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogInfoResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *CatalogInfoResponse) validateLimits(formats strfmt.Registry) error {

	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	if m.Limits != nil {
		if err := m.Limits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limits")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogInfoResponse) validateStandardUnitDescriptionGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.StandardUnitDescriptionGroup) { // not required
		return nil
	}

	if m.StandardUnitDescriptionGroup != nil {
		if err := m.StandardUnitDescriptionGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standard_unit_description_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogInfoResponse) UnmarshalBinary(b []byte) error {
	var res CatalogInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
