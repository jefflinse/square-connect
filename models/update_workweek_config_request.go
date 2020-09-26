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

// UpdateWorkweekConfigRequest A request to update a `WorkweekConfig` object
//
// swagger:model UpdateWorkweekConfigRequest
type UpdateWorkweekConfigRequest struct {

	// The updated `WorkweekConfig` object.
	// Required: true
	WorkweekConfig *WorkweekConfig `json:"workweek_config"`
}

// Validate validates this update workweek config request
func (m *UpdateWorkweekConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkweekConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateWorkweekConfigRequest) validateWorkweekConfig(formats strfmt.Registry) error {

	if err := validate.Required("workweek_config", "body", m.WorkweekConfig); err != nil {
		return err
	}

	if m.WorkweekConfig != nil {
		if err := m.WorkweekConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workweek_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateWorkweekConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateWorkweekConfigRequest) UnmarshalBinary(b []byte) error {
	var res UpdateWorkweekConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}