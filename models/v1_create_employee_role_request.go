// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateEmployeeRoleRequest v1 create employee role request
//
// swagger:model V1CreateEmployeeRoleRequest
type V1CreateEmployeeRoleRequest struct {

	// An EmployeeRole object with a name and permissions, and an optional owner flag.
	EmployeeRole *V1EmployeeRole `json:"employee_role,omitempty"`
}

// Validate validates this v1 create employee role request
func (m *V1CreateEmployeeRoleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmployeeRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateEmployeeRoleRequest) validateEmployeeRole(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployeeRole) { // not required
		return nil
	}

	if m.EmployeeRole != nil {
		if err := m.EmployeeRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employee_role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateEmployeeRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateEmployeeRoleRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateEmployeeRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
