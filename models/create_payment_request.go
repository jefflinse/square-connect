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

// CreatePaymentRequest Creates a payment from the source (nonce, card on file, etc.)
//
// The `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission is required
// to enable application fees.
//
// For more information, see [Payments and Refunds Overview](/payments-api/overview).
//
// For information about application fees in a payment, see
// [Collect Fees](/payments-api/take-payments-and-collect-fees).
//
// swagger:model CreatePaymentRequest
type CreatePaymentRequest struct {

	// If set to true and charging a Square Gift Card, a payment may be returned with
	// amount_money equal to less than what was requested.  Example, a request for $20 when charging
	// a Square Gift Card with balance of $5 wil result in an APPROVED payment of $5.  You may choose
	// to prompt the buyer for an additional payment to cover the remainder, or cancel the gift card
	// payment.  Cannot be `true` when `autocomplete = true`.
	//
	// For more information, see
	// [Partial amount with Square gift cards](https://developer.squareup.com/docs/payments-api/take-payments#partial-payment-gift-card).
	//
	// Default: false
	AcceptPartialAuthorization bool `json:"accept_partial_authorization,omitempty"`

	// The amount of money to accept for this payment, not including `tip_money`.
	//
	// Must be specified in the smallest denomination of the applicable currency.
	// For example, US dollar amounts are specified in cents. See
	// [Working with monetary amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts) for details.
	//
	// The currency code must match the currency associated with the business
	// that is accepting the payment.
	// Required: true
	AmountMoney *Money `json:"amount_money"`

	// The amount of money the developer is taking as a fee
	// for facilitating the payment on behalf of the seller.
	//
	// Cannot be more than 90% of the total amount of the Payment.
	//
	// Must be specified in the smallest denomination of the applicable currency.
	// For example, US dollar amounts are specified in cents. See
	// [Working with monetary amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts) for details.
	//
	// The fee currency code must match the currency associated with the merchant
	// that is accepting the payment. The application must be from a developer
	// account in the same country, and using the same currency code, as the merchant.
	//
	// For more information about the application fee scenario, see
	// [Collect Fees](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees).
	AppFeeMoney *Money `json:"app_fee_money,omitempty"`

	// If set to `true`, this payment will be completed when possible. If
	// set to `false`, this payment will be held in an approved state until either
	// explicitly completed (captured) or canceled (voided). For more information, see
	// [Delayed Payments](https://developer.squareup.com/docs/payments-api/take-payments#delayed-payments).
	//
	// Default: true
	Autocomplete bool `json:"autocomplete,omitempty"`

	// The buyer's billing address.
	BillingAddress *Address `json:"billing_address,omitempty"`

	// The buyer's e-mail address
	// Max Length: 255
	BuyerEmailAddress string `json:"buyer_email_address,omitempty"`

	// The `Customer` ID of the customer associated with the payment.
	// Required if the `source_id` refers to a card on file created using the Customers API.
	CustomerID string `json:"customer_id,omitempty"`

	// The duration of time after the payment's creation when Square automatically cancels the
	// payment. This automatic cancellation applies only to payments that don't reach a terminal state
	// (COMPLETED, CANCELED, or FAILED) before the `delay_duration` time period.
	//
	// This parameter should be specified as a time duration, in RFC 3339 format, with a minimum value
	// of 1 minute.
	//
	// Notes:
	// This feature is only supported for card payments. This parameter can only be set for a delayed
	// capture payment (`autocomplete=false`).
	//
	// Default:
	//
	// - Card Present payments: "PT36H" (36 hours) from the creation time.
	// - Card Not Present payments: "P7D" (7 days) from the creation time.
	DelayDuration string `json:"delay_duration,omitempty"`

	// A unique string that identifies this CreatePayment request. Keys can be any valid string
	// but must be unique for every CreatePayment request.
	//
	// Max: 45 characters
	//
	// Note: The number of allowed characters might be less than the stated maximum, if multi-byte
	// characters are used.
	//
	// See [Idempotency keys](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	// Required: true
	// Max Length: 45
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`

	// The location ID to associate with the payment. If not specified, the default location is
	// used.
	LocationID string `json:"location_id,omitempty"`

	// An optional note to be entered by the developer when creating a payment
	//
	// Limit 500 characters.
	// Max Length: 500
	Note string `json:"note,omitempty"`

	// Associate a previously created order with this payment
	OrderID string `json:"order_id,omitempty"`

	// A user-defined ID to associate with the payment.
	// You can use this field to associate the payment to an entity in an external system.
	// For example, you might specify an order ID that is generated by a third-party shopping cart.
	//
	// Limit 40 characters.
	// Max Length: 40
	ReferenceID string `json:"reference_id,omitempty"`

	// The buyer's shipping address.
	ShippingAddress *Address `json:"shipping_address,omitempty"`

	// The ID for the source of funds for this payment.  This can be a nonce
	// generated by the Payment Form or a card on file made with the Customers API.
	// Required: true
	// Min Length: 1
	SourceID *string `json:"source_id"`

	// Optional additional payment information to include on the customer's card statement
	// as part of statement description. This can be, for example, an invoice number, ticket number,
	// or short description that uniquely identifies the purchase.
	//
	// Note that the `statement_description_identifier` may get truncated on the statement description
	// to fit the required information including the Square identifier (SQ *) and name of the
	// merchant taking the payment.
	// Max Length: 20
	StatementDescriptionIdentifier string `json:"statement_description_identifier,omitempty"`

	// The amount designated as a tip, in addition to `amount_money`
	//
	// Must be specified in the smallest denomination of the applicable currency.
	// For example, US dollar amounts are specified in cents. See
	// [Working with monetary amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts) for details.
	//
	// The currency code must match the currency associated with the business
	// that is accepting the payment.
	TipMoney *Money `json:"tip_money,omitempty"`

	// An identifying token generated by `SqPaymentForm.verifyBuyer()`.
	// Verification tokens encapsulate customer device information and 3-D Secure
	// challenge results to indicate that Square has verified the buyer identity.
	//
	// See the [SCA Overview](https://developer.squareup.com/docs/sca-overview).
	VerificationToken string `json:"verification_token,omitempty"`
}

// Validate validates this create payment request
func (m *CreatePaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppFeeMoney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuyerEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementDescriptionIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTipMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePaymentRequest) validateAmountMoney(formats strfmt.Registry) error {

	if err := validate.Required("amount_money", "body", m.AmountMoney); err != nil {
		return err
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

func (m *CreatePaymentRequest) validateAppFeeMoney(formats strfmt.Registry) error {

	if swag.IsZero(m.AppFeeMoney) { // not required
		return nil
	}

	if m.AppFeeMoney != nil {
		if err := m.AppFeeMoney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_fee_money")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePaymentRequest) validateBillingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAddress) { // not required
		return nil
	}

	if m.BillingAddress != nil {
		if err := m.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePaymentRequest) validateBuyerEmailAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.BuyerEmailAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("buyer_email_address", "body", string(m.BuyerEmailAddress), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", string(*m.IdempotencyKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", string(*m.IdempotencyKey), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateNote(formats strfmt.Registry) error {

	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if err := validate.MaxLength("note", "body", string(m.Note), 500); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateReferenceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceID) { // not required
		return nil
	}

	if err := validate.MaxLength("reference_id", "body", string(m.ReferenceID), 40); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateShippingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if m.ShippingAddress != nil {
		if err := m.ShippingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipping_address")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePaymentRequest) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("source_id", "body", m.SourceID); err != nil {
		return err
	}

	if err := validate.MinLength("source_id", "body", string(*m.SourceID), 1); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateStatementDescriptionIdentifier(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementDescriptionIdentifier) { // not required
		return nil
	}

	if err := validate.MaxLength("statement_description_identifier", "body", string(m.StatementDescriptionIdentifier), 20); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateTipMoney(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CreatePaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePaymentRequest) UnmarshalBinary(b []byte) error {
	var res CreatePaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
