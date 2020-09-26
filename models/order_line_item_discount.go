// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderLineItemDiscount Represents a discount that applies to one or more line items in an
// order.
//
// Fixed-amount, order-scoped discounts are distributed across all non-zero line item totals.
// The amount distributed to each line item is relative to the
// amount contributed by the item to the order subtotal.
//
// swagger:model OrderLineItemDiscount
type OrderLineItemDiscount struct {

	// The total declared monetary amount of the discount.
	//
	// `amount_money` is not set for percentage-based discounts.
	AmountMoney *Money `json:"amount_money,omitempty"`

	// The amount of discount actually applied to the line item.
	//
	// Represents the amount of money applied as a line item-scoped discount.
	// When an amount-based discount is scoped to the entire order, the value
	// of `applied_money` is different from `amount_money` because the total
	// amount of the discount is distributed across all line items.
	AppliedMoney *Money `json:"applied_money,omitempty"`

	// The catalog object id referencing `CatalogDiscount`.
	// Max Length: 192
	CatalogObjectID string `json:"catalog_object_id,omitempty"`

	// Application-defined data attached to this discount. Metadata fields are intended
	// to store descriptive references or associations with an entity in another system or store brief
	// information about the object. Square does not process this field; it only stores and returns it
	// in relevant API calls. Do not use metadata to store any sensitive information (personally
	// identifiable information, card details, etc.).
	//
	// Keys written by applications must be 60 characters or less and must be in the character set
	// `[a-zA-Z0-9_-]`. Entries may also include metadata generated by Square. These keys are prefixed
	// with a namespace, separated from the key with a ':' character.
	//
	// Values have a max length of 255 characters.
	//
	// An application may have up to 10 entries per metadata field.
	//
	// Entries written by applications are private and can only be read or modified by the same
	// application.
	//
	// See [Metadata](https://developer.squareup.com/docs/build-basics/metadata) for more information.
	Metadata map[string]string `json:"metadata,omitempty"`

	// The discount's name.
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// The percentage of the discount, as a string representation of a decimal number.
	// A value of `7.25` corresponds to a percentage of 7.25%.
	//
	// `percentage` is not set for amount-based discounts.
	// Max Length: 10
	Percentage string `json:"percentage,omitempty"`

	// The object identifier of a `pricing rule` to be applied automatically
	// to this discount. The specification and application of the discounts, to which a `pricing_rule_id` is
	// assigned, are completely controlled by the corresponding pricing rule.
	PricingRuleID string `json:"pricing_rule_id,omitempty"`

	// The reward identifiers corresponding to this discount. The application and
	// specification of discounts that have `reward_ids` are completely controlled by the backing
	// criteria corresponding to the reward tiers of the rewards that are added to the order
	// through the Loyalty API. To manually unapply discounts that are the result of added rewards,
	// the rewards must be removed from the order through the Loyalty API.
	RewardIds []string `json:"reward_ids"`

	// Indicates the level at which the discount applies. For `ORDER` scoped discounts,
	// Square generates references in `applied_discounts` on all order line items that do
	// not have them. For `LINE_ITEM` scoped discounts, the discount only applies to line items
	// with a discount reference in their `applied_discounts` field.
	//
	// This field is immutable. To change the scope of a discount you must delete
	// the discount and re-add it as a new discount.
	// See [OrderLineItemDiscountScope](#type-orderlineitemdiscountscope) for possible values
	Scope string `json:"scope,omitempty"`

	// The type of the discount.
	//
	// Discounts that don't reference a catalog object ID must have a type of
	// `FIXED_PERCENTAGE` or `FIXED_AMOUNT`.
	// See [OrderLineItemDiscountType](#type-orderlineitemdiscounttype) for possible values
	Type string `json:"type,omitempty"`

	// Unique ID that identifies the discount only within this order.
	// Max Length: 60
	UID string `json:"uid,omitempty"`
}

// Validate validates this order line item discount
func (m *OrderLineItemDiscount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppliedMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalogObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderLineItemDiscount) validateAmountMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.AmountMoney) { // not required
		return nil
	}

	if m.AmountMoney != nil {
		if err := m.AmountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLineItemDiscount) validateAppliedMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.AppliedMoney) { // not required
		return nil
	}

	if m.AppliedMoney != nil {
		if err := m.AppliedMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applied_money")
			}
			return err
		}
	}

	return nil
}

func (m *OrderLineItemDiscount) validateCatalogObjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.CatalogObjectID) { // not required
		return nil
	}

	if err := validate.MaxLength("catalog_object_id", "body", string(m.CatalogObjectID), 192); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemDiscount) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 255); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemDiscount) validatePercentage(formats strfmt.Registry) error {

	if swag.IsZero(m.Percentage) { // not required
		return nil
	}

	if err := validate.MaxLength("percentage", "body", string(m.Percentage), 10); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemDiscount) validateUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.MaxLength("uid", "body", string(m.UID), 60); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderLineItemDiscount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderLineItemDiscount) UnmarshalBinary(b []byte) error {
	var res OrderLineItemDiscount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
