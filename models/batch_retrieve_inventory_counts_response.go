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

// BatchRetrieveInventoryCountsResponse batch retrieve inventory counts response
//
// swagger:model BatchRetrieveInventoryCountsResponse
type BatchRetrieveInventoryCountsResponse struct {

	// The current calculated inventory counts for the requested objects
	// and locations.
	Counts []*InventoryCount `json:"counts"`

	// The pagination cursor to be used in a subsequent request. If unset,
	// this is the final response.
	//
	// See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this batch retrieve inventory counts response
func (m *BatchRetrieveInventoryCountsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRetrieveInventoryCountsResponse) validateCounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Counts) { // not required
		return nil
	}

	for i := 0; i < len(m.Counts); i++ {
		if swag.IsZero(m.Counts[i]) { // not required
			continue
		}

		if m.Counts[i] != nil {
			if err := m.Counts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BatchRetrieveInventoryCountsResponse) validateErrors(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BatchRetrieveInventoryCountsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchRetrieveInventoryCountsResponse) UnmarshalBinary(b []byte) error {
	var res BatchRetrieveInventoryCountsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
