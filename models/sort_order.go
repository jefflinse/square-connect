// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SortOrder The order (e.g., chronological or alphabetical) in which results from a request are returned.
//
// swagger:model SortOrder
type SortOrder string

const (

	// SortOrderDESC captures enum value "DESC"
	SortOrderDESC SortOrder = "DESC"

	// SortOrderASC captures enum value "ASC"
	SortOrderASC SortOrder = "ASC"
)

// for schema
var sortOrderEnum []interface{}

func init() {
	var res []SortOrder
	if err := json.Unmarshal([]byte(`["DESC","ASC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sortOrderEnum = append(sortOrderEnum, v)
	}
}

func (m SortOrder) validateSortOrderEnum(path, location string, value SortOrder) error {
	if err := validate.EnumCase(path, location, value, sortOrderEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sort order
func (m SortOrder) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSortOrderEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sort order based on context it is used
func (m SortOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
