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

// CatalogItemVariation An item variation (i.e., product) in the Catalog object model. Each item
// may have a maximum of 250 item variations.
//
// swagger:model CatalogItemVariation
type CatalogItemVariation struct {

	// If the inventory quantity for the variation is less than or equal to this value and `inventory_alert_type`
	// is `LOW_QUANTITY`, the variation displays an alert in the merchant dashboard.
	//
	// This value is always an integer.
	InventoryAlertThreshold int64 `json:"inventory_alert_threshold,omitempty"`

	// Indicates whether the item variation displays an alert when its inventory quantity is less than or equal
	// to its `inventory_alert_threshold`.
	// See [InventoryAlertType](#type-inventoryalerttype) for possible values
	InventoryAlertType string `json:"inventory_alert_type,omitempty"`

	// The ID of the `CatalogItem` associated with this item variation. Searchable.
	ItemID string `json:"item_id,omitempty"`

	// List of item option values associated with this item variation. Listed
	// in the same order as the item options of the parent item.
	ItemOptionValues []*CatalogItemOptionValueForItemVariation `json:"item_option_values"`

	// Per-location price and inventory overrides.
	LocationOverrides []*ItemVariationLocationOverrides `json:"location_overrides"`

	// ID of the ‘CatalogMeasurementUnit’ that is used to measure the quantity
	// sold of this item variation. If left unset, the item will be sold in
	// whole quantities.
	MeasurementUnitID string `json:"measurement_unit_id,omitempty"`

	// The item variation's name. Searchable. This field has max length of 255 Unicode code points.
	Name string `json:"name,omitempty"`

	// The order in which this item variation should be displayed. This value is read-only. On writes, the ordinal
	// for each item variation within a parent `CatalogItem` is set according to the item variations's
	// position. On reads, the value is not guaranteed to be sequential or unique.
	Ordinal int64 `json:"ordinal,omitempty"`

	// The item variation's price, if fixed pricing is used.
	PriceMoney *Money `json:"price_money,omitempty"`

	// Indicates whether the item variation's price is fixed or determined at the time
	// of sale.
	// See [CatalogPricingType](#type-catalogpricingtype) for possible values
	PricingType string `json:"pricing_type,omitempty"`

	// If the `CatalogItem` that owns this item variation is of type
	// `APPOINTMENTS_SERVICE`, then this is the duration of the service in milliseconds. For
	// example, a 30 minute appointment would have the value `1800000`, which is equal to
	// 30 (minutes) * 60 (seconds per minute) * 1000 (milliseconds per second).
	ServiceDuration int64 `json:"service_duration,omitempty"`

	// The item variation's SKU, if any. Searchable.
	Sku string `json:"sku,omitempty"`

	// If `true`, inventory tracking is active for the variation.
	TrackInventory bool `json:"track_inventory,omitempty"`

	// The item variation's UPC, if any. Searchable in the Connect API.
	// This field is only exposed in the Connect API. It is not exposed in Square's Dashboard,
	// Square Point of Sale app or Retail Point of Sale app.
	Upc string `json:"upc,omitempty"`

	// Arbitrary user metadata to associate with the item variation. Searchable. This field has max length of 255 Unicode code points.
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this catalog item variation
func (m *CatalogItemVariation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemOptionValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItemVariation) validateItemOptionValues(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemOptionValues) { // not required
		return nil
	}

	for i := 0; i < len(m.ItemOptionValues); i++ {
		if swag.IsZero(m.ItemOptionValues[i]) { // not required
			continue
		}

		if m.ItemOptionValues[i] != nil {
			if err := m.ItemOptionValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("item_option_values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogItemVariation) validateLocationOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationOverrides) { // not required
		return nil
	}

	for i := 0; i < len(m.LocationOverrides); i++ {
		if swag.IsZero(m.LocationOverrides[i]) { // not required
			continue
		}

		if m.LocationOverrides[i] != nil {
			if err := m.LocationOverrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("location_overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogItemVariation) validatePriceMoney(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CatalogItemVariation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemVariation) UnmarshalBinary(b []byte) error {
	var res CatalogItemVariation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}