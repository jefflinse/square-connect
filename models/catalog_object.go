// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogObject The wrapper object for object types in the Catalog data model. The type
// of a particular `CatalogObject` is determined by the value of
// `type` and only the corresponding data field may be set.
//
// - if type = `ITEM`, only `item_data` will be populated and it will contain a valid `CatalogItem` object.
// - if type = `ITEM_VARIATION`, only `item_variation_data` will be populated and it will contain a valid `CatalogItemVariation` object.
// - if type = `MODIFIER`, only `modifier_data` will be populated and it will contain a valid `CatalogModifier` object.
// - if type = `MODIFIER_LIST`, only `modifier_list_data` will be populated and it will contain a valid `CatalogModifierList` object.
// - if type = `CATEGORY`, only `category_data` will be populated and it will contain a valid `CatalogCategory` object.
// - if type = `DISCOUNT`, only `discount_data` will be populated and it will contain a valid `CatalogDiscount` object.
// - if type = `TAX`, only `tax_data` will be populated and it will contain a valid `CatalogTax` object.
// - if type = `IMAGE`, only `image_data` will be populated and it will contain a valid `CatalogImage` object.
// - if type = `QUICK_AMOUNTS_SETTINGS`, only `quick_amounts_settings_data` will be populated and will contain a valid `CatalogQuickAmountsSettings` object.
//
// For a more detailed discussion of the Catalog data model, please see the
// [Design a Catalog](/catalog-api/design-a-catalog) guide.
//
// swagger:model CatalogObject
type CatalogObject struct {

	// A list of locations where the object is not present, even if `present_at_all_locations` is `true`.
	AbsentAtLocationIds []string `json:"absent_at_location_ids"`

	// The Connect v1 IDs for this object at each location where it is present, where they
	// differ from the object's Connect V2 ID. The field will only be present for objects that
	// have been created or modified by legacy APIs.
	CatalogV1Ids []*CatalogV1ID `json:"catalog_v1_ids"`

	// Structured data for a `CatalogCategory`, set for CatalogObjects of type `CATEGORY`.
	CategoryData *CatalogCategory `json:"category_data,omitempty"`

	// Structured data for a `CatalogCustomAttributeDefinition`, set for CatalogObjects of type `CUSTOM_ATTRIBUTE_DEFINITION`.
	CustomAttributeDefinitionData *CatalogCustomAttributeDefinition `json:"custom_attribute_definition_data,omitempty"`

	// Application-defined key/value attributes that are set at a global (location-independent) level.
	// Custom Attribute Values are intended to store additional information about a Catalog Object
	// or associations with an entity in another system. Do not use custom attributes
	// to store any sensitive information (personally identifiable information, card details, etc.).
	//
	// For CustomAttributesDefinitions defined by the app making the request, the map key is the key defined in the
	// `CatalogCustomAttributeDefinition` (e.g. “reference_id”). For custom attributes created by other apps, the map key is
	// the key defined in `CatalogCustomAttributeDefinition` prefixed with the application ID and a colon
	// (eg. “abcd1234:reference_id”).
	CustomAttributeValues map[string]CatalogCustomAttributeValue `json:"custom_attribute_values,omitempty"`

	// Structured data for a `CatalogDiscount`, set for CatalogObjects of type `DISCOUNT`.
	DiscountData *CatalogDiscount `json:"discount_data,omitempty"`

	// An identifier to reference this object in the catalog. When a new `CatalogObject`
	// is inserted, the client should set the id to a temporary identifier starting with
	// a "`#`" character. Other objects being inserted or updated within the same request
	// may use this identifier to refer to the new object.
	//
	// When the server receives the new object, it will supply a unique identifier that
	// replaces the temporary identifier for all future references.
	// Required: true
	// Min Length: 1
	ID *string `json:"id"`

	// Structured data for a `CatalogImage`, set for CatalogObjects of type `IMAGE`.
	ImageData *CatalogImage `json:"image_data,omitempty"`

	// Identifies the `CatalogImage` attached to this `CatalogObject`.
	ImageID string `json:"image_id,omitempty"`

	// If `true`, the object has been deleted from the database. Must be `false` for new objects
	// being inserted. When deleted, the `updated_at` field will equal the deletion time.
	IsDeleted bool `json:"is_deleted,omitempty"`

	// Structured data for a `CatalogItem`, set for CatalogObjects of type `ITEM`.
	ItemData *CatalogItem `json:"item_data,omitempty"`

	// Structured data for a `CatalogItemOption`, set for CatalogObjects of type `ITEM_OPTION`.
	ItemOptionData *CatalogItemOption `json:"item_option_data,omitempty"`

	// Structured data for a `CatalogItemOptionValue`, set for CatalogObjects of type `ITEM_OPTION_VAL`.
	ItemOptionValueData *CatalogItemOptionValue `json:"item_option_value_data,omitempty"`

	// Structured data for a `CatalogItemVariation`, set for CatalogObjects of type `ITEM_VARIATION`.
	ItemVariationData *CatalogItemVariation `json:"item_variation_data,omitempty"`

	// Structured data for a `CatalogMeasurementUnit`, set for CatalogObjects of type `MEASUREMENT_UNIT`.
	MeasurementUnitData *CatalogMeasurementUnit `json:"measurement_unit_data,omitempty"`

	// Structured data for a `CatalogModifier`, set for CatalogObjects of type `MODIFIER`.
	ModifierData *CatalogModifier `json:"modifier_data,omitempty"`

	// Structured data for a `CatalogModifierList`, set for CatalogObjects of type `MODIFIER_LIST`.
	ModifierListData *CatalogModifierList `json:"modifier_list_data,omitempty"`

	// If `true`, this object is present at all locations (including future locations), except where specified in
	// the `absent_at_location_ids` field. If `false`, this object is not present at any locations (including future locations),
	// except where specified in the `present_at_location_ids` field. If not specified, defaults to `true`.
	PresentAtAllLocations bool `json:"present_at_all_locations,omitempty"`

	// A list of locations where the object is present, even if `present_at_all_locations` is `false`.
	PresentAtLocationIds []string `json:"present_at_location_ids"`

	// Structured data for a `CatalogPricingRule`, set for CatalogObjects of type `PRICING_RULE`.
	PricingRuleData *CatalogPricingRule `json:"pricing_rule_data,omitempty"`

	// Structured data for a `CatalogProductSet`, set for CatalogObjects of type `PRODUCT_SET`.
	ProductSetData *CatalogProductSet `json:"product_set_data,omitempty"`

	// Structured data for a `CatalogQuickAmountsSettings`, set for CatalogObjects of type `QUICK_AMOUNTS_SETTINGS`.
	QuickAmountsSettingsData *CatalogQuickAmountsSettings `json:"quick_amounts_settings_data,omitempty"`

	// Structured data for a `CatalogTax`, set for CatalogObjects of type `TAX`.
	TaxData *CatalogTax `json:"tax_data,omitempty"`

	// Structured data for a `CatalogTimePeriod`, set for CatalogObjects of type `TIME_PERIOD`.
	TimePeriodData *CatalogTimePeriod `json:"time_period_data,omitempty"`

	// The type of this object. Each object type has expected
	// properties expressed in a structured format within its corresponding `*_data` field below.
	// See [CatalogObjectType](#type-catalogobjecttype) for possible values
	// Required: true
	Type *string `json:"type"`

	// Last modification [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) in RFC 3339 format, e.g., `"2016-08-15T23:59:33.123Z"`
	// would indicate the UTC time (denoted by `Z`) of August 15, 2016 at 23:59:33 and 123 milliseconds.
	UpdatedAt string `json:"updated_at,omitempty"`

	// The version of the object. When updating an object, the version supplied
	// must match the version in the database, otherwise the write will be rejected as conflicting.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this catalog object
func (m *CatalogObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogV1Ids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomAttributeDefinitionData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomAttributeValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscountData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemOptionData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemOptionValueData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemVariationData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeasurementUnitData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifierData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifierListData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingRuleData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductSetData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuickAmountsSettingsData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimePeriodData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogObject) validateCatalogV1Ids(formats strfmt.Registry) error {

	if swag.IsZero(m.CatalogV1Ids) { // not required
		return nil
	}

	for i := 0; i < len(m.CatalogV1Ids); i++ {
		if swag.IsZero(m.CatalogV1Ids[i]) { // not required
			continue
		}

		if m.CatalogV1Ids[i] != nil {
			if err := m.CatalogV1Ids[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalog_v1_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogObject) validateCategoryData(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryData) { // not required
		return nil
	}

	if m.CategoryData != nil {
		if err := m.CategoryData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateCustomAttributeDefinitionData(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomAttributeDefinitionData) { // not required
		return nil
	}

	if m.CustomAttributeDefinitionData != nil {
		if err := m.CustomAttributeDefinitionData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom_attribute_definition_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateCustomAttributeValues(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomAttributeValues) { // not required
		return nil
	}

	for k := range m.CustomAttributeValues {

		if err := validate.Required("custom_attribute_values"+"."+k, "body", m.CustomAttributeValues[k]); err != nil {
			return err
		}
		if val, ok := m.CustomAttributeValues[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CatalogObject) validateDiscountData(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountData) { // not required
		return nil
	}

	if m.DiscountData != nil {
		if err := m.DiscountData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discount_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *CatalogObject) validateImageData(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageData) { // not required
		return nil
	}

	if m.ImageData != nil {
		if err := m.ImageData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateItemData(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemData) { // not required
		return nil
	}

	if m.ItemData != nil {
		if err := m.ItemData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateItemOptionData(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemOptionData) { // not required
		return nil
	}

	if m.ItemOptionData != nil {
		if err := m.ItemOptionData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item_option_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateItemOptionValueData(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemOptionValueData) { // not required
		return nil
	}

	if m.ItemOptionValueData != nil {
		if err := m.ItemOptionValueData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item_option_value_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateItemVariationData(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemVariationData) { // not required
		return nil
	}

	if m.ItemVariationData != nil {
		if err := m.ItemVariationData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item_variation_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateMeasurementUnitData(formats strfmt.Registry) error {

	if swag.IsZero(m.MeasurementUnitData) { // not required
		return nil
	}

	if m.MeasurementUnitData != nil {
		if err := m.MeasurementUnitData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("measurement_unit_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateModifierData(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifierData) { // not required
		return nil
	}

	if m.ModifierData != nil {
		if err := m.ModifierData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifier_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateModifierListData(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifierListData) { // not required
		return nil
	}

	if m.ModifierListData != nil {
		if err := m.ModifierListData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifier_list_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validatePricingRuleData(formats strfmt.Registry) error {

	if swag.IsZero(m.PricingRuleData) { // not required
		return nil
	}

	if m.PricingRuleData != nil {
		if err := m.PricingRuleData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing_rule_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateProductSetData(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductSetData) { // not required
		return nil
	}

	if m.ProductSetData != nil {
		if err := m.ProductSetData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product_set_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateQuickAmountsSettingsData(formats strfmt.Registry) error {

	if swag.IsZero(m.QuickAmountsSettingsData) { // not required
		return nil
	}

	if m.QuickAmountsSettingsData != nil {
		if err := m.QuickAmountsSettingsData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quick_amounts_settings_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateTaxData(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxData) { // not required
		return nil
	}

	if m.TaxData != nil {
		if err := m.TaxData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateTimePeriodData(formats strfmt.Registry) error {

	if swag.IsZero(m.TimePeriodData) { // not required
		return nil
	}

	if m.TimePeriodData != nil {
		if err := m.TimePeriodData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_period_data")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogObject) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogObject) UnmarshalBinary(b []byte) error {
	var res CatalogObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}