// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogCustomAttributeValue An instance of a custom attribute. Custom attributes can be defined and
// added to `ITEM` and `ITEM_VARIATION` type catalog objects.
// [Read more about custom attributes](/catalog-api/add-custom-attributes).
//
// swagger:model CatalogCustomAttributeValue
type CatalogCustomAttributeValue struct {

	// A `true` or `false` value. Populated if `type` = `BOOLEAN`.
	BooleanValue bool `json:"boolean_value,omitempty"`

	// __Read-only.__ The id of the `CatalogCustomAttributeDefinition` this value belongs to.
	CustomAttributeDefinitionID string `json:"custom_attribute_definition_id,omitempty"`

	// __Read-only.__ A copy of key from the associated `CatalogCustomAttributeDefinition`.
	Key string `json:"key,omitempty"`

	// The name of the custom attribute.
	Name string `json:"name,omitempty"`

	// Populated if `type` = `NUMBER`. Contains a string
	// representation of a decimal number, using a `.` as the decimal separator.
	NumberValue string `json:"number_value,omitempty"`

	// One or more choices from `allowed_selections`. Populated if `type` = `SELECTION`.
	SelectionUIDValues []string `json:"selection_uid_values"`

	// The string value of the custom attribute.  Populated if `type` = `STRING`.
	StringValue string `json:"string_value,omitempty"`

	// __Read-only.__ A copy of type from the associated `CatalogCustomAttributeDefinition`.
	// See [CatalogCustomAttributeDefinitionType](#type-catalogcustomattributedefinitiontype) for possible values
	Type string `json:"type,omitempty"`
}

// Validate validates this catalog custom attribute value
func (m *CatalogCustomAttributeValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog custom attribute value based on context it is used
func (m *CatalogCustomAttributeValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogCustomAttributeValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogCustomAttributeValue) UnmarshalBinary(b []byte) error {
	var res CatalogCustomAttributeValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
