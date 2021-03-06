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

// BreakType A defined break template that sets an expectation for possible `Break`
// instances on a `Shift`.
//
// swagger:model BreakType
type BreakType struct {

	// A human-readable name for this type of break. Will be displayed to
	// employees in Square products.
	// Required: true
	// Min Length: 1
	BreakName *string `json:"break_name"`

	// A read-only timestamp in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`

	// Format: RFC-3339 P[n]Y[n]M[n]DT[n]H[n]M[n]S. The expected length of
	// this break. Precision below minutes is truncated.
	// Required: true
	// Min Length: 1
	ExpectedDuration *string `json:"expected_duration"`

	// UUID for this object.
	// Max Length: 255
	ID string `json:"id,omitempty"`

	// Whether this break counts towards time worked for compensation
	// purposes.
	// Required: true
	IsPaid *bool `json:"is_paid"`

	// The ID of the business location this type of break applies to.
	// Required: true
	// Min Length: 1
	LocationID *string `json:"location_id"`

	// A read-only timestamp in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	// Used for resolving concurrency issues; request will fail if version
	// provided does not match server version at time of request. If a value is not
	// provided, Square's servers execute a "blind" write; potentially
	// overwriting another writer's data.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this break type
func (m *BreakType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBreakName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPaid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BreakType) validateBreakName(formats strfmt.Registry) error {

	if err := validate.Required("break_name", "body", m.BreakName); err != nil {
		return err
	}

	if err := validate.MinLength("break_name", "body", *m.BreakName, 1); err != nil {
		return err
	}

	return nil
}

func (m *BreakType) validateExpectedDuration(formats strfmt.Registry) error {

	if err := validate.Required("expected_duration", "body", m.ExpectedDuration); err != nil {
		return err
	}

	if err := validate.MinLength("expected_duration", "body", *m.ExpectedDuration, 1); err != nil {
		return err
	}

	return nil
}

func (m *BreakType) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("id", "body", m.ID, 255); err != nil {
		return err
	}

	return nil
}

func (m *BreakType) validateIsPaid(formats strfmt.Registry) error {

	if err := validate.Required("is_paid", "body", m.IsPaid); err != nil {
		return err
	}

	return nil
}

func (m *BreakType) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	if err := validate.MinLength("location_id", "body", *m.LocationID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this break type based on context it is used
func (m *BreakType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BreakType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BreakType) UnmarshalBinary(b []byte) error {
	var res BreakType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
