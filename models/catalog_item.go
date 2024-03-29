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
	"github.com/go-openapi/validate"
)

// CatalogItem A [CatalogObject](#type-CatalogObject) instance of the `ITEM` type, also referred to as an item, in the catalog.
// Example: {"object":{"id":"#Cocoa","item_data":{"abbreviation":"Ch","description":"Hot chocolate","name":"Cocoa","visibility":"PRIVATE"},"present_at_all_locations":true,"type":"ITEM"}}
//
// swagger:model CatalogItem
type CatalogItem struct {

	// The text of the item's display label in the Square Point of Sale app. Only up to the first five characters of the string are used.
	// This attribute is searchable, and its value length is of Unicode code points.
	// Max Length: 24
	Abbreviation string `json:"abbreviation,omitempty"`

	// If `true`, the item can be added to electronically fulfilled orders from the merchant's online store.
	AvailableElectronically bool `json:"available_electronically,omitempty"`

	// If `true`, the item can be added to pickup orders from the merchant's online store.
	AvailableForPickup bool `json:"available_for_pickup,omitempty"`

	// If `true`, the item can be added to shipping orders from the merchant's online store.
	AvailableOnline bool `json:"available_online,omitempty"`

	// The ID of the item's category, if any.
	CategoryID string `json:"category_id,omitempty"`

	// The item's description. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	// Max Length: 4096
	Description string `json:"description,omitempty"`

	// List of item options IDs for this item. Used to manage and group item
	// variations in a specified order.
	//
	// Maximum: 6 item options.
	ItemOptions []*CatalogItemOptionForItem `json:"item_options"`

	// The color of the item's display label in the Square Point of Sale app. This must be a valid hex color code.
	LabelColor string `json:"label_color,omitempty"`

	// A set of `CatalogItemModifierListInfo` objects
	// representing the modifier lists that apply to this item, along with the overrides and min
	// and max limits that are specific to this item. Modifier lists
	// may also be added to or deleted from an item using `UpdateItemModifierLists`.
	ModifierListInfo []*CatalogItemModifierListInfo `json:"modifier_list_info"`

	// The item's name. This is a searchable attribute for use in applicable query filters, its value must not be empty, and the length is of Unicode code points.
	// Max Length: 512
	Name string `json:"name,omitempty"`

	// The product type of the item. May not be changed once an item has been created.
	//
	// Only items of product type `REGULAR` or `APPOINTMENTS_SERVICE` may be created by this API; items with other product
	// types are read-only.
	// See [CatalogItemProductType](#type-catalogitemproducttype) for possible values
	ProductType string `json:"product_type,omitempty"`

	// If `false`, the Square Point of Sale app will present the `CatalogItem`'s
	// details screen immediately, allowing the merchant to choose `CatalogModifier`s
	// before adding the item to the cart.  This is the default behavior.
	//
	// If `true`, the Square Point of Sale app will immediately add the item to the cart with the pre-selected
	// modifiers, and merchants can edit modifiers by drilling down onto the item's details.
	//
	// Third-party clients are encouraged to implement similar behaviors.
	SkipModifierScreen bool `json:"skip_modifier_screen,omitempty"`

	// A set of IDs indicating the taxes enabled for
	// this item. When updating an item, any taxes listed here will be added to the item.
	// Taxes may also be added to or deleted from an item using `UpdateItemTaxes`.
	TaxIds []string `json:"tax_ids"`

	// A list of CatalogObjects containing the `CatalogItemVariation`s for this item.
	Variations []*CatalogObject `json:"variations"`
}

// Validate validates this catalog item
func (m *CatalogItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbbreviation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifierListInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItem) validateAbbreviation(formats strfmt.Registry) error {
	if swag.IsZero(m.Abbreviation) { // not required
		return nil
	}

	if err := validate.MaxLength("abbreviation", "body", m.Abbreviation, 24); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItem) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 4096); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItem) validateItemOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.ItemOptions); i++ {
		if swag.IsZero(m.ItemOptions[i]) { // not required
			continue
		}

		if m.ItemOptions[i] != nil {
			if err := m.ItemOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("item_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogItem) validateModifierListInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifierListInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifierListInfo); i++ {
		if swag.IsZero(m.ModifierListInfo[i]) { // not required
			continue
		}

		if m.ModifierListInfo[i] != nil {
			if err := m.ModifierListInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifier_list_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogItem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 512); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItem) validateVariations(formats strfmt.Registry) error {
	if swag.IsZero(m.Variations) { // not required
		return nil
	}

	for i := 0; i < len(m.Variations); i++ {
		if swag.IsZero(m.Variations[i]) { // not required
			continue
		}

		if m.Variations[i] != nil {
			if err := m.Variations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this catalog item based on the context it is used
func (m *CatalogItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItemOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifierListInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItem) contextValidateItemOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ItemOptions); i++ {

		if m.ItemOptions[i] != nil {
			if err := m.ItemOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("item_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogItem) contextValidateModifierListInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifierListInfo); i++ {

		if m.ModifierListInfo[i] != nil {
			if err := m.ModifierListInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifier_list_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogItem) contextValidateVariations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variations); i++ {

		if m.Variations[i] != nil {
			if err := m.Variations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItem) UnmarshalBinary(b []byte) error {
	var res CatalogItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
