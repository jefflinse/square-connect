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

// OrderLineItemPricingBlocklistsBlockedDiscount A discount to block from applying to a line item. The discount must be
// identified by either `discount_uid` or `discount_catalog_object_id`, but not both.
//
// swagger:model OrderLineItemPricingBlocklistsBlockedDiscount
type OrderLineItemPricingBlocklistsBlockedDiscount struct {

	// The `catalog_object_id` of the discount that should be blocked.
	// Use this field to block catalog discounts. For ad-hoc discounts use the
	// `discount_uid` field.
	// Max Length: 192
	DiscountCatalogObjectID string `json:"discount_catalog_object_id,omitempty"`

	// The `uid` of the discount that should be blocked. Use this field to block
	// ad-hoc discounts. For catalog discounts use the `discount_catalog_object_id` field.
	// Max Length: 60
	DiscountUID string `json:"discount_uid,omitempty"`

	// Unique ID of the `BlockedDiscount` within the order.
	// Max Length: 60
	UID string `json:"uid,omitempty"`
}

// Validate validates this order line item pricing blocklists blocked discount
func (m *OrderLineItemPricingBlocklistsBlockedDiscount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountCatalogObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscountUID(formats); err != nil {
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

func (m *OrderLineItemPricingBlocklistsBlockedDiscount) validateDiscountCatalogObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscountCatalogObjectID) { // not required
		return nil
	}

	if err := validate.MaxLength("discount_catalog_object_id", "body", m.DiscountCatalogObjectID, 192); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemPricingBlocklistsBlockedDiscount) validateDiscountUID(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscountUID) { // not required
		return nil
	}

	if err := validate.MaxLength("discount_uid", "body", m.DiscountUID, 60); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemPricingBlocklistsBlockedDiscount) validateUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.MaxLength("uid", "body", m.UID, 60); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this order line item pricing blocklists blocked discount based on context it is used
func (m *OrderLineItemPricingBlocklistsBlockedDiscount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderLineItemPricingBlocklistsBlockedDiscount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderLineItemPricingBlocklistsBlockedDiscount) UnmarshalBinary(b []byte) error {
	var res OrderLineItemPricingBlocklistsBlockedDiscount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}