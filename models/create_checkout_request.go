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

// CreateCheckoutRequest Defines the parameters that can be included in the body of
// a request to the __CreateCheckout__ endpoint.
// Example: {"request_body":{"additional_recipients":[{"amount_money":{"amount":60,"currency":"USD"},"description":"Application fees","location_id":"057P5VYJ4A5X1"}],"ask_for_shipping_address":true,"idempotency_key":"86ae1696-b1e3-4328-af6d-f1e04d947ad6","merchant_support_email":"merchant+support@website.com","order":{"idempotency_key":"12ae1696-z1e3-4328-af6d-f1e04d947gd4","order":{"customer_id":"customer_id","discounts":[{"amount_money":{"amount":100,"currency":"USD"},"scope":"LINE_ITEM","type":"FIXED_AMOUNT","uid":"56ae1696-z1e3-9328-af6d-f1e04d947gd4"}],"line_items":[{"applied_discounts":[{"discount_uid":"56ae1696-z1e3-9328-af6d-f1e04d947gd4"}],"applied_taxes":[{"tax_uid":"38ze1696-z1e3-5628-af6d-f1e04d947fg3"}],"base_price_money":{"amount":1500,"currency":"USD"},"name":"Printed T Shirt","quantity":"2"},{"base_price_money":{"amount":2500,"currency":"USD"},"name":"Slim Jeans","quantity":"1"},{"base_price_money":{"amount":3500,"currency":"USD"},"name":"Woven Sweater","quantity":"3"}],"location_id":"location_id","reference_id":"reference_id","taxes":[{"percentage":"7.75","scope":"LINE_ITEM","type":"INCLUSIVE","uid":"38ze1696-z1e3-5628-af6d-f1e04d947fg3"}]}},"pre_populate_buyer_email":"example@email.com","pre_populate_shipping_address":{"address_line_1":"1455 Market St.","address_line_2":"Suite 600","administrative_district_level_1":"CA","country":"US","first_name":"Jane","last_name":"Doe","locality":"San Francisco","postal_code":"94103"},"redirect_url":"https://merchant.website.com/order-confirm"}}
//
// swagger:model CreateCheckoutRequest
type CreateCheckoutRequest struct {

	// The basic primitive of multi-party transaction. The value is optional.
	// The transaction facilitated by you can be split from here.
	//
	// If you provide this value, the `amount_money` value in your additional_recipients
	// must not be more than 90% of the `total_money` calculated by Square for your order.
	// The `location_id` must be the valid location of the app owner merchant.
	//
	// This field requires `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission.
	//
	// This field is currently not supported in sandbox.
	AdditionalRecipients []*ChargeRequestAdditionalRecipient `json:"additional_recipients"`

	// If `true`, Square Checkout will collect shipping information on your
	// behalf and store that information with the transaction information in your
	// Square Dashboard.
	//
	// Default: `false`.
	AskForShippingAddress bool `json:"ask_for_shipping_address,omitempty"`

	// A unique string that identifies this checkout among others
	// you've created. It can be any valid string but must be unique for every
	// order sent to Square Checkout for a given location ID.
	//
	// The idempotency key is used to avoid processing the same order more than
	// once. If you're unsure whether a particular checkout was created
	// successfully, you can reattempt it with the same idempotency key and all the
	// same other parameters without worrying about creating duplicates.
	//
	// We recommend using a random number/string generator native to the language
	// you are working in to generate strings for your idempotency keys.
	//
	// See the [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency) guide for more information.
	// Required: true
	// Max Length: 192
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`

	// The email address to display on the Square Checkout confirmation page
	// and confirmation email that the buyer can use to contact the merchant.
	//
	// If this value is not set, the confirmation page and email will display the
	// primary email address associated with the merchant's Square account.
	//
	// Default: none; only exists if explicitly set.
	// Max Length: 254
	MerchantSupportEmail string `json:"merchant_support_email,omitempty"`

	// An optional note to associate with the checkout object.
	//
	// This value cannot exceed 60 characters.
	// Max Length: 60
	Note string `json:"note,omitempty"`

	// The order including line items to be checked out.
	// Required: true
	Order *CreateOrderRequest `json:"order"`

	// If provided, the buyer's email is pre-populated on the checkout page
	// as an editable text field.
	//
	// Default: none; only exists if explicitly set.
	// Max Length: 254
	PrePopulateBuyerEmail string `json:"pre_populate_buyer_email,omitempty"`

	// If provided, the buyer's shipping info is pre-populated on the
	// checkout page as editable text fields.
	//
	// Default: none; only exists if explicitly set.
	PrePopulateShippingAddress *Address `json:"pre_populate_shipping_address,omitempty"`

	// The URL to redirect to after checkout is completed with `checkoutId`,
	// Square's `orderId`, `transactionId`, and `referenceId` appended as URL
	// parameters. For example, if the provided redirect_url is
	// `http://www.example.com/order-complete`, a successful transaction redirects
	// the customer to:
	//
	// <pre><code>http://www.example.com/order-complete?checkoutId=xxxxxx&amp;orderId=xxxxxx&amp;referenceId=xxxxxx&amp;transactionId=xxxxxx</code></pre>
	//
	// If you do not provide a redirect URL, Square Checkout will display an order
	// confirmation page on your behalf; however Square strongly recommends that
	// you provide a redirect URL so you can verify the transaction results and
	// finalize the order through your existing/normal confirmation workflow.
	//
	// Default: none; only exists if explicitly set.
	// Max Length: 800
	RedirectURL string `json:"redirect_url,omitempty"`
}

// Validate validates this create checkout request
func (m *CreateCheckoutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMerchantSupportEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrePopulateBuyerEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrePopulateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCheckoutRequest) validateAdditionalRecipients(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalRecipients) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalRecipients); i++ {
		if swag.IsZero(m.AdditionalRecipients[i]) { // not required
			continue
		}

		if m.AdditionalRecipients[i] != nil {
			if err := m.AdditionalRecipients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateCheckoutRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", *m.IdempotencyKey, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("idempotency_key", "body", *m.IdempotencyKey, 192); err != nil {
		return err
	}

	return nil
}

func (m *CreateCheckoutRequest) validateMerchantSupportEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.MerchantSupportEmail) { // not required
		return nil
	}

	if err := validate.MaxLength("merchant_support_email", "body", m.MerchantSupportEmail, 254); err != nil {
		return err
	}

	return nil
}

func (m *CreateCheckoutRequest) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	if err := validate.MaxLength("note", "body", m.Note, 60); err != nil {
		return err
	}

	return nil
}

func (m *CreateCheckoutRequest) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
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

func (m *CreateCheckoutRequest) validatePrePopulateBuyerEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.PrePopulateBuyerEmail) { // not required
		return nil
	}

	if err := validate.MaxLength("pre_populate_buyer_email", "body", m.PrePopulateBuyerEmail, 254); err != nil {
		return err
	}

	return nil
}

func (m *CreateCheckoutRequest) validatePrePopulateShippingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PrePopulateShippingAddress) { // not required
		return nil
	}

	if m.PrePopulateShippingAddress != nil {
		if err := m.PrePopulateShippingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pre_populate_shipping_address")
			}
			return err
		}
	}

	return nil
}

func (m *CreateCheckoutRequest) validateRedirectURL(formats strfmt.Registry) error {
	if swag.IsZero(m.RedirectURL) { // not required
		return nil
	}

	if err := validate.MaxLength("redirect_url", "body", m.RedirectURL, 800); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create checkout request based on the context it is used
func (m *CreateCheckoutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalRecipients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrePopulateShippingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCheckoutRequest) contextValidateAdditionalRecipients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalRecipients); i++ {

		if m.AdditionalRecipients[i] != nil {
			if err := m.AdditionalRecipients[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateCheckoutRequest) contextValidateOrder(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateCheckoutRequest) contextValidatePrePopulateShippingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.PrePopulateShippingAddress != nil {
		if err := m.PrePopulateShippingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pre_populate_shipping_address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCheckoutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCheckoutRequest) UnmarshalBinary(b []byte) error {
	var res CreateCheckoutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
