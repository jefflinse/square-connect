// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Employee An employee object that is used by the external API.
//
// swagger:model Employee
type Employee struct {

	// A read-only timestamp in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// The employee's email address
	Email string `json:"email,omitempty"`

	// The employee's first name.
	FirstName string `json:"first_name,omitempty"`

	// UUID for this object.
	ID string `json:"id,omitempty"`

	// Whether this employee is the owner of the merchant. Each merchant
	// has one owner employee, and that employee has full authority over
	// the account.
	IsOwner bool `json:"is_owner,omitempty"`

	// The employee's last name.
	LastName string `json:"last_name,omitempty"`

	// A list of location IDs where this employee has access to.
	LocationIds []string `json:"location_ids"`

	// The employee's phone number in E.164 format, i.e. "+12125554250"
	PhoneNumber string `json:"phone_number,omitempty"`

	// Specifies the status of the employees being fetched.
	// See [EmployeeStatus](#type-employeestatus) for possible values
	Status string `json:"status,omitempty"`

	// A read-only timestamp in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this employee
func (m *Employee) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Employee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Employee) UnmarshalBinary(b []byte) error {
	var res Employee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}