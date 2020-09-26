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

// BatchChangeInventoryRequest batch change inventory request
//
// swagger:model BatchChangeInventoryRequest
type BatchChangeInventoryRequest struct {

	// The set of physical counts and inventory adjustments to be made.
	// Changes are applied based on the client-supplied timestamp and may be sent
	// out of order. Max size is 100 changes.
	Changes []*InventoryChange `json:"changes"`

	// A client-supplied, universally unique identifier (UUID) for the
	// request.
	//
	// See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) in the
	// [API Development 101](https://developer.squareup.com/docs/basics/api101/overview) section for more
	// information.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// Indicates whether the current physical count should be ignored if
	// the quantity is unchanged since the last physical count. Default: `true`.
	IgnoreUnchangedCounts bool `json:"ignore_unchanged_counts,omitempty"`
}

// Validate validates this batch change inventory request
func (m *BatchChangeInventoryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchChangeInventoryRequest) validateChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchChangeInventoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchChangeInventoryRequest) UnmarshalBinary(b []byte) error {
	var res BatchChangeInventoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}