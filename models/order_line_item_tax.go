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

// OrderLineItemTax Represents a tax that applies to one or more line item in the order.
//
// Fixed-amount, order-scoped taxes are distributed across all non-zero line item totals.
// The amount distributed to each line item is relative to the amount the item
// contributes to the order subtotal.
//
// swagger:model OrderLineItemTax
type OrderLineItemTax struct {

	// The amount of the money applied by the tax in the order.
	AppliedMoney *Money `json:"applied_money,omitempty"`

	// Determines whether the tax was automatically applied to the order based on
	// the catalog configuration. For an example, see
	// [Automatically Apply Taxes to an Order](https://developer.squareup.com/docs/docs/orders-api/apply-taxes-and-discounts/auto-apply-taxes).
	AutoApplied bool `json:"auto_applied,omitempty"`

	// The catalog object id referencing `CatalogTax`.
	// Max Length: 192
	CatalogObjectID string `json:"catalog_object_id,omitempty"`

	// Application-defined data attached to this tax. Metadata fields are intended
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

	// The tax's name.
	// Max Length: 255
	Name string `json:"name,omitempty"`

	// The percentage of the tax, as a string representation of a decimal
	// number. For example, a value of `"7.25"` corresponds to a percentage of
	// 7.25%.
	// Max Length: 10
	Percentage string `json:"percentage,omitempty"`

	// Indicates the level at which the tax applies. For `ORDER` scoped taxes,
	// Square generates references in `applied_taxes` on all order line items that do
	// not have them. For `LINE_ITEM` scoped taxes, the tax will only apply to line items
	// with references in their `applied_taxes` field.
	//
	// This field is immutable. To change the scope, you must delete the tax and
	// re-add it as a new tax.
	// See [OrderLineItemTaxScope](#type-orderlineitemtaxscope) for possible values
	Scope string `json:"scope,omitempty"`

	// Indicates the calculation method used to apply the tax.
	// See [OrderLineItemTaxType](#type-orderlineitemtaxtype) for possible values
	Type string `json:"type,omitempty"`

	// Unique ID that identifies the tax only within this order.
	// Max Length: 60
	UID string `json:"uid,omitempty"`
}

// Validate validates this order line item tax
func (m *OrderLineItemTax) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *OrderLineItemTax) validateAppliedMoney(formats strfmt.Registry) error {
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

func (m *OrderLineItemTax) validateCatalogObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogObjectID) { // not required
		return nil
	}

	if err := validate.MaxLength("catalog_object_id", "body", m.CatalogObjectID, 192); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemTax) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemTax) validatePercentage(formats strfmt.Registry) error {
	if swag.IsZero(m.Percentage) { // not required
		return nil
	}

	if err := validate.MaxLength("percentage", "body", m.Percentage, 10); err != nil {
		return err
	}

	return nil
}

func (m *OrderLineItemTax) validateUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if err := validate.MaxLength("uid", "body", m.UID, 60); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order line item tax based on the context it is used
func (m *OrderLineItemTax) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppliedMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderLineItemTax) contextValidateAppliedMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.AppliedMoney != nil {
		if err := m.AppliedMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applied_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderLineItemTax) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderLineItemTax) UnmarshalBinary(b []byte) error {
	var res OrderLineItemTax
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
