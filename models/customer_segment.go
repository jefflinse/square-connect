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

// CustomerSegment Represents a group of customer profiles that match one or more predefined filter criteria.
//
// Segments (also known as Smart Groups) are defined and created within Customer Directory in the Square Dashboard or Point of Sale.
//
// swagger:model CustomerSegment
type CustomerSegment struct {

	// The timestamp when the segment was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// Unique Square-generated ID for the segment.
	// Max Length: 255
	ID string `json:"id,omitempty"`

	// Name of the segment.
	// Required: true
	Name *string `json:"name"`

	// The timestamp when the segment was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this customer segment
func (m *CustomerSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerSegment) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerSegment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this customer segment based on context it is used
func (m *CustomerSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerSegment) UnmarshalBinary(b []byte) error {
	var res CustomerSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
