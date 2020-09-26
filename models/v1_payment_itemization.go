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

// V1PaymentItemization Payment include an` itemizations` field that lists the items purchased,
// along with associated fees, modifiers, and discounts. Each itemization has an
// `itemization_type` field that indicates which of the following the itemization
// represents:
//
// <ul>
// <li>An item variation from the merchant's item library</li>
// <li>A custom monetary amount</li>
// <li>
// An action performed on a Square gift card, such as activating or
// reloading it.
// </li>
// </ul>
//
// *Note**: itemization information included in a `Payment` object reflects
// details collected **at the time of the payment**. Details such as the name or
// price of items might have changed since the payment was processed.
//
// swagger:model V1PaymentItemization
type V1PaymentItemization struct {

	// The total of all discounts applied to the itemization. This value is always negative or zero.
	DiscountMoney *V1Money `json:"discount_money,omitempty"`

	// All discounts applied to this itemization.
	Discounts []*V1PaymentDiscount `json:"discounts"`

	// The total cost of the itemization and its modifiers, not including taxes or discounts.
	GrossSalesMoney *V1Money `json:"gross_sales_money,omitempty"`

	// Details of the item, including its unique identifier and the identifier of the item variation purchased.
	ItemDetail *V1PaymentItemDetail `json:"item_detail,omitempty"`

	// The name of the item variation purchased, if any.
	ItemVariationName string `json:"item_variation_name,omitempty"`

	// The type of purchase that the itemization represents, such as an ITEM or CUSTOM_AMOUNT
	// See [V1PaymentItemizationItemizationType](#type-v1paymentitemizationitemizationtype) for possible values
	ItemizationType string `json:"itemization_type,omitempty"`

	// All modifier options applied to this itemization.
	Modifiers []*V1PaymentModifier `json:"modifiers"`

	// The item's name.
	Name string `json:"name,omitempty"`

	// The sum of gross_sales_money and discount_money.
	NetSalesMoney *V1Money `json:"net_sales_money,omitempty"`

	// Notes entered by the merchant about the item at the time of payment, if any.
	Notes string `json:"notes,omitempty"`

	// The quantity of the item purchased. This can be a decimal value.
	Quantity float64 `json:"quantity,omitempty"`

	// The cost of a single unit of this item.
	SingleQuantityMoney *V1Money `json:"single_quantity_money,omitempty"`

	// All taxes applied to this itemization.
	Taxes []*V1PaymentTax `json:"taxes"`

	// The total cost of the item, including all taxes and discounts.
	TotalMoney *V1Money `json:"total_money,omitempty"`
}

// Validate validates this v1 payment itemization
func (m *V1PaymentItemization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrossSalesMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetSalesMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleQuantityMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PaymentItemization) validateDiscountMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountMoney) { // not required
		return nil
	}

	if m.DiscountMoney != nil {
		if err := m.DiscountMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discount_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaymentItemization) validateDiscounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Discounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Discounts); i++ {
		if swag.IsZero(m.Discounts[i]) { // not required
			continue
		}

		if m.Discounts[i] != nil {
			if err := m.Discounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PaymentItemization) validateGrossSalesMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.GrossSalesMoney) { // not required
		return nil
	}

	if m.GrossSalesMoney != nil {
		if err := m.GrossSalesMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gross_sales_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaymentItemization) validateItemDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemDetail) { // not required
		return nil
	}

	if m.ItemDetail != nil {
		if err := m.ItemDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item_detail")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaymentItemization) validateModifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Modifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Modifiers); i++ {
		if swag.IsZero(m.Modifiers[i]) { // not required
			continue
		}

		if m.Modifiers[i] != nil {
			if err := m.Modifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PaymentItemization) validateNetSalesMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.NetSalesMoney) { // not required
		return nil
	}

	if m.NetSalesMoney != nil {
		if err := m.NetSalesMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net_sales_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaymentItemization) validateSingleQuantityMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.SingleQuantityMoney) { // not required
		return nil
	}

	if m.SingleQuantityMoney != nil {
		if err := m.SingleQuantityMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("single_quantity_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1PaymentItemization) validateTaxes(formats strfmt.Registry) error {

	if swag.IsZero(m.Taxes) { // not required
		return nil
	}

	for i := 0; i < len(m.Taxes); i++ {
		if swag.IsZero(m.Taxes[i]) { // not required
			continue
		}

		if m.Taxes[i] != nil {
			if err := m.Taxes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PaymentItemization) validateTotalMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalMoney) { // not required
		return nil
	}

	if m.TotalMoney != nil {
		if err := m.TotalMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PaymentItemization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PaymentItemization) UnmarshalBinary(b []byte) error {
	var res V1PaymentItemization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}