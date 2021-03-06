// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemVariationLocationOverrides Price and inventory alerting overrides for a `CatalogItemVariation` at a specific `Location`.
//
// swagger:model ItemVariationLocationOverrides
type ItemVariationLocationOverrides struct {

	// If the inventory quantity for the variation is less than or equal to this value and `inventory_alert_type`
	// is `LOW_QUANTITY`, the variation displays an alert in the merchant dashboard.
	//
	// This value is always an integer.
	InventoryAlertThreshold int64 `json:"inventory_alert_threshold,omitempty"`

	// Indicates whether the `CatalogItemVariation` displays an alert when its inventory
	// quantity is less than or equal to its `inventory_alert_threshold`.
	// See [InventoryAlertType](#type-inventoryalerttype) for possible values
	InventoryAlertType string `json:"inventory_alert_type,omitempty"`

	// The ID of the `Location`.
	LocationID string `json:"location_id,omitempty"`

	// The price of the `CatalogItemVariation` at the given `Location`, or blank for variable pricing.
	PriceMoney *Money `json:"price_money,omitempty"`

	// The pricing type (fixed or variable) for the `CatalogItemVariation` at the given `Location`.
	// See [CatalogPricingType](#type-catalogpricingtype) for possible values
	PricingType string `json:"pricing_type,omitempty"`

	// If `true`, inventory tracking is active for the `CatalogItemVariation` at this `Location`.
	TrackInventory bool `json:"track_inventory,omitempty"`
}

// Validate validates this item variation location overrides
func (m *ItemVariationLocationOverrides) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePriceMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemVariationLocationOverrides) validatePriceMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.PriceMoney) { // not required
		return nil
	}

	if m.PriceMoney != nil {
		if err := m.PriceMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price_money")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this item variation location overrides based on the context it is used
func (m *ItemVariationLocationOverrides) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePriceMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemVariationLocationOverrides) contextValidatePriceMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.PriceMoney != nil {
		if err := m.PriceMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemVariationLocationOverrides) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemVariationLocationOverrides) UnmarshalBinary(b []byte) error {
	var res ItemVariationLocationOverrides
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
