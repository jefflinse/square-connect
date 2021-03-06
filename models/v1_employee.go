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

// V1Employee Represents one of a business's employees.
//
// swagger:model V1Employee
type V1Employee struct {

	// The IDs of the locations the employee is allowed to clock in at.
	AuthorizedLocationIds []string `json:"authorized_location_ids"`

	// The time when the employee entity was created, in ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The employee's email address.
	Email string `json:"email,omitempty"`

	// An ID the merchant can set to associate the employee with an entity in another system.
	ExternalID string `json:"external_id,omitempty"`

	// The employee's first name.
	// Required: true
	FirstName *string `json:"first_name"`

	// The employee's unique ID.
	ID string `json:"id,omitempty"`

	// The employee's last name.
	// Required: true
	LastName *string `json:"last_name"`

	// The ids of the employee's associated roles. Currently, you can specify only one or zero roles per employee.
	RoleIds []string `json:"role_ids"`

	// CWhether the employee is ACTIVE or INACTIVE. Inactive employees cannot sign in to Square Register.Merchants update this field from the Square Dashboard.
	// See [V1EmployeeStatus](#type-v1employeestatus) for possible values
	Status string `json:"status,omitempty"`

	// The time when the employee entity was most recently updated, in ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this v1 employee
func (m *V1Employee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Employee) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *V1Employee) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Employee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Employee) UnmarshalBinary(b []byte) error {
	var res V1Employee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
