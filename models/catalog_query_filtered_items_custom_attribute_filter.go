// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogQueryFilteredItemsCustomAttributeFilter catalog query filtered items custom attribute filter
//
// swagger:model CatalogQueryFilteredItemsCustomAttributeFilter
type CatalogQueryFilteredItemsCustomAttributeFilter struct {

	// custom attribute definition ids
	CustomAttributeDefinitionIds []string `json:"custom_attribute_definition_ids"`

	// custom attribute max value
	CustomAttributeMaxValue string `json:"custom_attribute_max_value,omitempty"`

	// custom attribute min value
	CustomAttributeMinValue string `json:"custom_attribute_min_value,omitempty"`

	// custom attribute value exact
	CustomAttributeValueExact string `json:"custom_attribute_value_exact,omitempty"`

	// custom attribute value prefix
	CustomAttributeValuePrefix string `json:"custom_attribute_value_prefix,omitempty"`

	//
	// See [CatalogQueryFilteredItemsCustomAttributeFilterFilterType](#type-catalogqueryfiltereditemscustomattributefilterfiltertype) for possible values
	FilterType string `json:"filter_type,omitempty"`
}

// Validate validates this catalog query filtered items custom attribute filter
func (m *CatalogQueryFilteredItemsCustomAttributeFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogQueryFilteredItemsCustomAttributeFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogQueryFilteredItemsCustomAttributeFilter) UnmarshalBinary(b []byte) error {
	var res CatalogQueryFilteredItemsCustomAttributeFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}