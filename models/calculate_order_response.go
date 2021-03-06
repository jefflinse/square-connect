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
)

// CalculateOrderResponse calculate order response
// Example: {"order":{"created_at":"2020-05-18T16:30:49.614Z","discounts":[{"applied_money":{"amount":550,"currency":"USD"},"name":"50% Off","percentage":"50","scope":"ORDER","type":"FIXED_PERCENTAGE","uid":"zGsRZP69aqSSR9lq9euSPB"}],"line_items":[{"applied_discounts":[{"applied_money":{"amount":250,"currency":"USD"},"discount_uid":"zGsRZP69aqSSR9lq9euSPB","uid":"9zr9S4dxvPAixvn0lpa1VC"}],"base_price_money":{"amount":500,"currency":"USD"},"gross_sales_money":{"amount":500,"currency":"USD"},"name":"Item 1","quantity":"1","total_discount_money":{"amount":250,"currency":"USD"},"total_money":{"amount":250,"currency":"USD"},"total_tax_money":{"amount":0,"currency":"USD"},"uid":"ULkg0tQTRK2bkU9fNv3IJD","variation_total_price_money":{"amount":500,"currency":"USD"}},{"applied_discounts":[{"applied_money":{"amount":300,"currency":"USD"},"discount_uid":"zGsRZP69aqSSR9lq9euSPB","uid":"qa8LwwZK82FgSEkQc2HYVC"}],"base_price_money":{"amount":300,"currency":"USD"},"gross_sales_money":{"amount":600,"currency":"USD"},"name":"Item 2","quantity":"2","total_discount_money":{"amount":300,"currency":"USD"},"total_money":{"amount":300,"currency":"USD"},"total_tax_money":{"amount":0,"currency":"USD"},"uid":"mumY8Nun4BC5aKe2yyx5a","variation_total_price_money":{"amount":600,"currency":"USD"}}],"location_id":"D7AVYMEAPJ3A3","net_amounts":{"discount_money":{"amount":550,"currency":"USD"},"service_charge_money":{"amount":0,"currency":"USD"},"tax_money":{"amount":0,"currency":"USD"},"tip_money":{"amount":0,"currency":"USD"},"total_money":{"amount":550,"currency":"USD"}},"state":"OPEN","total_discount_money":{"amount":550,"currency":"USD"},"total_money":{"amount":550,"currency":"USD"},"total_service_charge_money":{"amount":0,"currency":"USD"},"total_tax_money":{"amount":0,"currency":"USD"},"total_tip_money":{"amount":0,"currency":"USD"},"updated_at":"2020-05-18T16:30:49.614Z","version":1}}
//
// swagger:model CalculateOrderResponse
type CalculateOrderResponse struct {

	// Any errors that occurred during the request.
	Errors []*Error `json:"errors"`

	// The calculated version of the order provided in the request.
	Order *Order `json:"order,omitempty"`
}

// Validate validates this calculate order response
func (m *CalculateOrderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateOrderResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CalculateOrderResponse) validateOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.Order) { // not required
		return nil
	}

	if m.Order != nil {
		if err := m.Order.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this calculate order response based on the context it is used
func (m *CalculateOrderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateOrderResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CalculateOrderResponse) contextValidateOrder(ctx context.Context, formats strfmt.Registry) error {

	if m.Order != nil {
		if err := m.Order.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("order")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CalculateOrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculateOrderResponse) UnmarshalBinary(b []byte) error {
	var res CalculateOrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
