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

// V1EmployeeRole V1EmployeeRole
//
// swagger:model V1EmployeeRole
type V1EmployeeRole struct {

	// The time when the employee entity was created, in ISO 8601 format. Is set by Square when the Role is created.
	CreatedAt string `json:"created_at,omitempty"`

	// The role's unique ID, Can only be set by Square.
	ID string `json:"id,omitempty"`

	// If true, employees with this role have all permissions, regardless of the values indicated in permissions.
	IsOwner bool `json:"is_owner,omitempty"`

	// The role's merchant-defined name.
	// Required: true
	Name *string `json:"name"`

	// The role's permissions.
	// See [V1EmployeeRolePermissions](#type-v1employeerolepermissions) for possible values
	// Required: true
	Permissions []string `json:"permissions"`

	// The time when the employee entity was most recently updated, in ISO 8601 format. Is set by Square when the Role updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this v1 employee role
func (m *V1EmployeeRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EmployeeRole) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1EmployeeRole) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EmployeeRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EmployeeRole) UnmarshalBinary(b []byte) error {
	var res V1EmployeeRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
