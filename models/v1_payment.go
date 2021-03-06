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

// V1Payment A payment represents a paid transaction between a Square merchant and a
// customer. Payment details are usually available from Connect API endpoints
// within a few minutes after the transaction completes.
//
// Each Payment object includes several fields that end in `_money`. These fields
// describe the various amounts of money that contribute to the payment total:
//
// <ul>
// <li>
// Monetary values are <b>positive</b> if they represent an
// <em>increase</em> in the amount of money the merchant receives (e.g.,
// <code>tax_money</code>, <code>tip_money</code>).
// </li>
// <li>
// Monetary values are <b>negative</b> if they represent an
// <em>decrease</em> in the amount of money the merchant receives (e.g.,
// <code>discount_money</code>, <code>refunded_money</code>).
// </li>
// </ul>
//
// swagger:model V1Payment
type V1Payment struct {

	// All of the additive taxes associated with the payment.
	AdditiveTax []*V1PaymentTax `json:"additive_tax"`

	// The sum of all additive taxes associated with the payment.
	AdditiveTaxMoney *V1Money `json:"additive_tax_money,omitempty"`

	// The time when the payment was created, in ISO 8601 format. Reflects the time of the first payment if the object represents an incomplete partial payment, and the time of the last or complete payment otherwise.
	CreatedAt string `json:"created_at,omitempty"`

	// The unique identifier of the Square account that took the payment.
	CreatorID string `json:"creator_id,omitempty"`

	// The device that took the payment.
	Device *Device `json:"device,omitempty"`

	// The total of all discounts applied to the payment.
	DiscountMoney *V1Money `json:"discount_money,omitempty"`

	// The total of all sales, including any applicable taxes.
	GrossSalesMoney *V1Money `json:"gross_sales_money,omitempty"`

	// The payment's unique identifier.
	ID string `json:"id,omitempty"`

	// All of the inclusive taxes associated with the payment.
	InclusiveTax []*V1PaymentTax `json:"inclusive_tax"`

	// The sum of all inclusive taxes associated with the payment.
	InclusiveTaxMoney *V1Money `json:"inclusive_tax_money,omitempty"`

	// Indicates whether or not the payment is only partially paid for.
	// If true, this payment will have the tenders collected so far, but the
	// itemizations will be empty until the payment is completed.
	IsPartial bool `json:"is_partial,omitempty"`

	// The items purchased in the payment.
	Itemizations []*V1PaymentItemization `json:"itemizations"`

	// The unique identifier of the merchant that took the payment.
	MerchantID string `json:"merchant_id,omitempty"`

	// The total of all sales, minus any applicable taxes.
	NetSalesMoney *V1Money `json:"net_sales_money,omitempty"`

	// The amount to be deposited into the merchant's bank account for the payment.
	NetTotalMoney *V1Money `json:"net_total_money,omitempty"`

	// The URL of the payment's detail page in the merchant dashboard. The merchant must be signed in to the merchant dashboard to view this page.
	PaymentURL string `json:"payment_url,omitempty"`

	// The total of all processing fees collected by Square for the payment.
	ProcessingFeeMoney *V1Money `json:"processing_fee_money,omitempty"`

	// The URL of the receipt for the payment. Note that for split tender
	// payments, this URL corresponds to the receipt for the first tender
	// listed in the payment's tender field. Each Tender object has its own
	// receipt_url field you can use to get the other receipts associated with
	// a split tender payment.
	ReceiptURL string `json:"receipt_url,omitempty"`

	// The total of all refunds applied to the payment.
	RefundedMoney *V1Money `json:"refunded_money,omitempty"`

	// All of the refunds applied to the payment. Note that the value of all refunds on a payment can exceed the value of all tenders if a merchant chooses to refund money to a tender after previously accepting returned goods as part of an exchange.
	Refunds []*V1Refund `json:"refunds"`

	// The total of all surcharges applied to the payment.
	SurchargeMoney *V1Money `json:"surcharge_money,omitempty"`

	// A list of all surcharges associated with the payment.
	Surcharges []*V1PaymentSurcharge `json:"surcharges"`

	// The total of all sales, including any applicable taxes, rounded to the smallest legal unit of currency (e.g., the nearest penny in USD, the nearest nickel in CAD)
	SwedishRoundingMoney *V1Money `json:"swedish_rounding_money,omitempty"`

	// The total of all taxes applied to the payment. This is always the sum of inclusive_tax_money and additive_tax_money.
	TaxMoney *V1Money `json:"tax_money,omitempty"`

	// All of the tenders associated with the payment.
	Tender []*V1Tender `json:"tender"`

	// The total of all tips applied to the payment.
	TipMoney *V1Money `json:"tip_money,omitempty"`

	// The total of all discounts applied to the payment.
	TotalCollectedMoney *V1Money `json:"total_collected_money,omitempty"`
}

// Validate validates this v1 payment
func (m *V1Payment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditiveTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdditiveTaxMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrossSalesMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclusiveTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclusiveTaxMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetSalesMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetTotalMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingFeeMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefundedMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurchargeMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurcharges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwedishRoundingMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTipMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCollectedMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Payment) validateAdditiveTax(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditiveTax) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditiveTax); i++ {
		if swag.IsZero(m.AdditiveTax[i]) { // not required
			continue
		}

		if m.AdditiveTax[i] != nil {
			if err := m.AdditiveTax[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additive_tax" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) validateAdditiveTaxMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditiveTaxMoney) { // not required
		return nil
	}

	if m.AdditiveTaxMoney != nil {
		if err := m.AdditiveTaxMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additive_tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateDiscountMoney(formats strfmt.Registry) error {
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

func (m *V1Payment) validateGrossSalesMoney(formats strfmt.Registry) error {
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

func (m *V1Payment) validateInclusiveTax(formats strfmt.Registry) error {
	if swag.IsZero(m.InclusiveTax) { // not required
		return nil
	}

	for i := 0; i < len(m.InclusiveTax); i++ {
		if swag.IsZero(m.InclusiveTax[i]) { // not required
			continue
		}

		if m.InclusiveTax[i] != nil {
			if err := m.InclusiveTax[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inclusive_tax" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) validateInclusiveTaxMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.InclusiveTaxMoney) { // not required
		return nil
	}

	if m.InclusiveTaxMoney != nil {
		if err := m.InclusiveTaxMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inclusive_tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateItemizations(formats strfmt.Registry) error {
	if swag.IsZero(m.Itemizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Itemizations); i++ {
		if swag.IsZero(m.Itemizations[i]) { // not required
			continue
		}

		if m.Itemizations[i] != nil {
			if err := m.Itemizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) validateNetSalesMoney(formats strfmt.Registry) error {
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

func (m *V1Payment) validateNetTotalMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.NetTotalMoney) { // not required
		return nil
	}

	if m.NetTotalMoney != nil {
		if err := m.NetTotalMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net_total_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateProcessingFeeMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingFeeMoney) { // not required
		return nil
	}

	if m.ProcessingFeeMoney != nil {
		if err := m.ProcessingFeeMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processing_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateRefundedMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.RefundedMoney) { // not required
		return nil
	}

	if m.RefundedMoney != nil {
		if err := m.RefundedMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refunded_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateRefunds(formats strfmt.Registry) error {
	if swag.IsZero(m.Refunds) { // not required
		return nil
	}

	for i := 0; i < len(m.Refunds); i++ {
		if swag.IsZero(m.Refunds[i]) { // not required
			continue
		}

		if m.Refunds[i] != nil {
			if err := m.Refunds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("refunds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) validateSurchargeMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.SurchargeMoney) { // not required
		return nil
	}

	if m.SurchargeMoney != nil {
		if err := m.SurchargeMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surcharge_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateSurcharges(formats strfmt.Registry) error {
	if swag.IsZero(m.Surcharges) { // not required
		return nil
	}

	for i := 0; i < len(m.Surcharges); i++ {
		if swag.IsZero(m.Surcharges[i]) { // not required
			continue
		}

		if m.Surcharges[i] != nil {
			if err := m.Surcharges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surcharges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) validateSwedishRoundingMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.SwedishRoundingMoney) { // not required
		return nil
	}

	if m.SwedishRoundingMoney != nil {
		if err := m.SwedishRoundingMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swedish_rounding_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateTaxMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxMoney) { // not required
		return nil
	}

	if m.TaxMoney != nil {
		if err := m.TaxMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateTender(formats strfmt.Registry) error {
	if swag.IsZero(m.Tender) { // not required
		return nil
	}

	for i := 0; i < len(m.Tender); i++ {
		if swag.IsZero(m.Tender[i]) { // not required
			continue
		}

		if m.Tender[i] != nil {
			if err := m.Tender[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tender" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) validateTipMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TipMoney) { // not required
		return nil
	}

	if m.TipMoney != nil {
		if err := m.TipMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tip_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) validateTotalCollectedMoney(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalCollectedMoney) { // not required
		return nil
	}

	if m.TotalCollectedMoney != nil {
		if err := m.TotalCollectedMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_collected_money")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 payment based on the context it is used
func (m *V1Payment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditiveTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdditiveTaxMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiscountMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrossSalesMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInclusiveTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInclusiveTaxMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetSalesMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetTotalMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessingFeeMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefundedMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefunds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSurchargeMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSurcharges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwedishRoundingMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTipMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalCollectedMoney(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Payment) contextValidateAdditiveTax(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditiveTax); i++ {

		if m.AdditiveTax[i] != nil {
			if err := m.AdditiveTax[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additive_tax" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) contextValidateAdditiveTaxMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.AdditiveTaxMoney != nil {
		if err := m.AdditiveTaxMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additive_tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateDiscountMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.DiscountMoney != nil {
		if err := m.DiscountMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discount_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateGrossSalesMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.GrossSalesMoney != nil {
		if err := m.GrossSalesMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gross_sales_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateInclusiveTax(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InclusiveTax); i++ {

		if m.InclusiveTax[i] != nil {
			if err := m.InclusiveTax[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inclusive_tax" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) contextValidateInclusiveTaxMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.InclusiveTaxMoney != nil {
		if err := m.InclusiveTaxMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inclusive_tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateItemizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Itemizations); i++ {

		if m.Itemizations[i] != nil {
			if err := m.Itemizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) contextValidateNetSalesMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.NetSalesMoney != nil {
		if err := m.NetSalesMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net_sales_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateNetTotalMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.NetTotalMoney != nil {
		if err := m.NetTotalMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net_total_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateProcessingFeeMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessingFeeMoney != nil {
		if err := m.ProcessingFeeMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processing_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateRefundedMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.RefundedMoney != nil {
		if err := m.RefundedMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refunded_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateRefunds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Refunds); i++ {

		if m.Refunds[i] != nil {
			if err := m.Refunds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("refunds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) contextValidateSurchargeMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.SurchargeMoney != nil {
		if err := m.SurchargeMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surcharge_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateSurcharges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Surcharges); i++ {

		if m.Surcharges[i] != nil {
			if err := m.Surcharges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surcharges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) contextValidateSwedishRoundingMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.SwedishRoundingMoney != nil {
		if err := m.SwedishRoundingMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swedish_rounding_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateTaxMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxMoney != nil {
		if err := m.TaxMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateTender(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tender); i++ {

		if m.Tender[i] != nil {
			if err := m.Tender[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tender" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Payment) contextValidateTipMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TipMoney != nil {
		if err := m.TipMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tip_money")
			}
			return err
		}
	}

	return nil
}

func (m *V1Payment) contextValidateTotalCollectedMoney(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalCollectedMoney != nil {
		if err := m.TotalCollectedMoney.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total_collected_money")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Payment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Payment) UnmarshalBinary(b []byte) error {
	var res V1Payment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
