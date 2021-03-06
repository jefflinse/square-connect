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

// InventoryCount Represents Square's estimated quantity of items in a particular state at a
// particular location based on the known history of physical counts and
// inventory adjustments.
//
// swagger:model InventoryCount
type InventoryCount struct {

	// A read-only timestamp in RFC 3339 format that indicates when Square
	// received the most recent physical count or adjustment that had an affect
	// on the estimated count.
	// Max Length: 34
	CalculatedAt string `json:"calculated_at,omitempty"`

	// The Square generated ID of the
	// `CatalogObject` being tracked.
	// Max Length: 100
	CatalogObjectID string `json:"catalog_object_id,omitempty"`

	// The `CatalogObjectType` of the
	// `CatalogObject` being tracked. Tracking is only
	// supported for the `ITEM_VARIATION` type.
	// Max Length: 14
	CatalogObjectType string `json:"catalog_object_type,omitempty"`

	// The Square ID of the `Location` where the related
	// quantity of items are being tracked.
	// Max Length: 100
	LocationID string `json:"location_id,omitempty"`

	// The number of items affected by the estimated count as a decimal string.
	// Can support up to 5 digits after the decimal point.
	// Max Length: 26
	Quantity string `json:"quantity,omitempty"`

	// The current `InventoryState` for the related
	// quantity of items.
	// See [InventoryState](#type-inventorystate) for possible values
	State string `json:"state,omitempty"`
}

// Validate validates this inventory count
func (m *InventoryCount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalogObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalogObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryCount) validateCalculatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CalculatedAt) { // not required
		return nil
	}

	if err := validate.MaxLength("calculated_at", "body", m.CalculatedAt, 34); err != nil {
		return err
	}

	return nil
}

func (m *InventoryCount) validateCatalogObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogObjectID) { // not required
		return nil
	}

	if err := validate.MaxLength("catalog_object_id", "body", m.CatalogObjectID, 100); err != nil {
		return err
	}

	return nil
}

func (m *InventoryCount) validateCatalogObjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogObjectType) { // not required
		return nil
	}

	if err := validate.MaxLength("catalog_object_type", "body", m.CatalogObjectType, 14); err != nil {
		return err
	}

	return nil
}

func (m *InventoryCount) validateLocationID(formats strfmt.Registry) error {
	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.MaxLength("location_id", "body", m.LocationID, 100); err != nil {
		return err
	}

	return nil
}

func (m *InventoryCount) validateQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.Quantity) { // not required
		return nil
	}

	if err := validate.MaxLength("quantity", "body", m.Quantity, 26); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inventory count based on context it is used
func (m *InventoryCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryCount) UnmarshalBinary(b []byte) error {
	var res InventoryCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
