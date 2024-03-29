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

// SearchOrdersSort Sorting criteria for a SearchOrders request. Results can only be sorted
// by a timestamp field.
//
// swagger:model SearchOrdersSort
type SearchOrdersSort struct {

	// The field to sort by.
	//
	// __Important:__ When using a `DateTimeFilter`,
	// `sort_field` must match the timestamp field that the DateTimeFilter uses to
	// filter. For example, If you set your `sort_field` to `CLOSED_AT` and you use a
	// DateTimeFilter, your DateTimeFilter must filter for orders by their `CLOSED_AT` date.
	// If this field does not match the timestamp field in `DateTimeFilter`,
	// SearchOrders will return an error.
	//
	// Default: `CREATED_AT`.
	// See [SearchOrdersSortField](#type-searchorderssortfield) for possible values
	// Required: true
	SortField *string `json:"sort_field"`

	// The chronological order in which results are returned. Defaults to `DESC`.
	// See [SortOrder](#type-sortorder) for possible values
	SortOrder string `json:"sort_order,omitempty"`
}

// Validate validates this search orders sort
func (m *SearchOrdersSort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSortField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchOrdersSort) validateSortField(formats strfmt.Registry) error {

	if err := validate.Required("sort_field", "body", m.SortField); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this search orders sort based on context it is used
func (m *SearchOrdersSort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchOrdersSort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchOrdersSort) UnmarshalBinary(b []byte) error {
	var res SearchOrdersSort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
