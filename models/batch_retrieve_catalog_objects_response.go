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

// BatchRetrieveCatalogObjectsResponse batch retrieve catalog objects response
//
// swagger:model BatchRetrieveCatalogObjectsResponse
type BatchRetrieveCatalogObjectsResponse struct {

	// The set of `Error`s encountered.
	Errors []*Error `json:"errors"`

	// A list of `CatalogObject`s returned.
	Objects []*CatalogObject `json:"objects"`

	// A list of `CatalogObject`s referenced by the object in the `objects` field.
	RelatedObjects []*CatalogObject `json:"related_objects"`
}

// Validate validates this batch retrieve catalog objects response
func (m *BatchRetrieveCatalogObjectsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRetrieveCatalogObjectsResponse) validateErrors(formats strfmt.Registry) error {

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

func (m *BatchRetrieveCatalogObjectsResponse) validateObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BatchRetrieveCatalogObjectsResponse) validateRelatedObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedObjects); i++ {
		if swag.IsZero(m.RelatedObjects[i]) { // not required
			continue
		}

		if m.RelatedObjects[i] != nil {
			if err := m.RelatedObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("related_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchRetrieveCatalogObjectsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchRetrieveCatalogObjectsResponse) UnmarshalBinary(b []byte) error {
	var res BatchRetrieveCatalogObjectsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
