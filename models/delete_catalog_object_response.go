// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteCatalogObjectResponse delete catalog object response
// Example: {"deleted_at":"2016-11-16T22:25:24.878Z","deleted_object_ids":["7SB3ZQYJ5GDMVFL7JK46JCHT","KQLFFHA6K6J3YQAQAWDQAL57"]}
//
// swagger:model DeleteCatalogObjectResponse
type DeleteCatalogObjectResponse struct {

	// The database [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)
	// of this deletion in RFC 3339 format, e.g., `2016-09-04T23:59:33.123Z`.
	DeletedAt string `json:"deleted_at,omitempty"`

	// The IDs of all catalog objects deleted by this request.
	// Multiple IDs may be returned when associated objects are also deleted, for example
	// a catalog item variation will be deleted (and its ID included in this field)
	// when its parent catalog item is deleted.
	DeletedObjectIds []string `json:"deleted_object_ids"`

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this delete catalog object response
func (m *DeleteCatalogObjectResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteCatalogObjectResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this delete catalog object response based on the context it is used
func (m *DeleteCatalogObjectResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteCatalogObjectResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
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
func (m *DeleteCatalogObjectResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteCatalogObjectResponse) UnmarshalBinary(b []byte) error {
	var res DeleteCatalogObjectResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
