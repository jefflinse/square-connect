// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerSort Indicates the field to use for sorting customer profiles.
//
// swagger:model CustomerSort
type CustomerSort struct {

	// Indicates the information used to sort the results. For example,
	// by creation date.
	//
	// Default: `DEFAULT`.
	// See [CustomerSortField](#type-customersortfield) for possible values
	Field string `json:"field,omitempty"`

	// Indicates the order in which results should be displayed based on the
	// value of the sort field. String comparisons use standard alphabetic comparison
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