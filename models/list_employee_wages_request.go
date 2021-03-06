// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListEmployeeWagesRequest A request for a set of `EmployeeWage` objects
// Example: {"request_params":"?employee_id=33fJchumvVdJwxV0H6L9\u0026limit=4\u0026cursor=s4R0Z6ecFTzTC4jz8sUDBQTudX3KE313OT9fCt3VUgsXM4sMgED"}
//
// swagger:model ListEmployeeWagesRequest
type ListEmployeeWagesRequest struct {

	// Pointer to the next page of Employee Wage results to fetch.
	Cursor string `json:"cursor,omitempty"`

	// Filter wages returned to only those that are associated with the specified employee.
	EmployeeID string `json:"employee_id,omitempty"`

	// Maximum number of Employee Wages to return per page. Can range between
	// 1 and 200. The default is the maximum at 200.
	// Maximum: 200
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`
}

// Validate validates this list employee wages request
func (m *ListEmployeeWagesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListEmployeeWagesRequest) validateLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("limit", "body", m.Limit, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "body", m.Limit, 200, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list employee wages request based on context it is used
func (m *ListEmployeeWagesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListEmployeeWagesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListEmployeeWagesRequest) UnmarshalBinary(b []byte) error {
	var res ListEmployeeWagesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
