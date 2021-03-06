// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerSort Specifies how searched customers profiles are sorted, including the sort key and sort order.
//
// swagger:model CustomerSort
type CustomerSort struct {

	//  Use one or more customer attributes as the sort key to sort searched customer profiles.
	// For example, use creation date (`created_at`) of customers or default attributes as the sort key.
	//
	//
	// Default: `DEFAULT`.
	// See [CustomerSortField](#type-customersortfield) for possible values
	Field string `json:"field,omitempty"`

	// Indicates the order in which results should be sorted based on the
	// sort field value. Strings use standard alphabetic comparison
	// to determine order. Strings representing numbers are sorted as strings.
	//
	// Default: `ASC`.
	// See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`
}

// Validate validates this customer sort
func (m *CustomerSort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer sort based on context it is used
func (m *CustomerSort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerSort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerSort) UnmarshalBinary(b []byte) error {
	var res CustomerSort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
