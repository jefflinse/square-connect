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

// CancelPaymentResponse The return value from the [CancelPayment](#endpoint-payments-cancelpayment) endpoint.
// Example: {"payment":{"amount_money":{"amount":200,"currency":"USD"},"app_fee_money":{"amount":10,"currency":"USD"},"card_details":{"auth_result_code":"eWZBDh","avs_status":"AVS_ACCEPTED","card":{"bin":"411111","card_brand":"VISA","card_type":"DEBIT","exp_month":2,"exp_year":2024,"fingerprint":"sq-1-9PP0tWfcM6vIsYmfsesdjfhduHSDFNdJFNDfDNFjdfjpseirDErsaP","last_4":"1234","prepaid_type":"PREPAID"},"card_payment_timeline":{"authorized_at":"2018-10-17T20:38:46.753Z","voided_at":"2018-10-17T20:38:57.793Z"},"cvv_status":"CVV_ACCEPTED","entry_method":"KEYED","statement_description":"SQ *MY MERCHANT","status":"VOIDED"},"created_at":"2018-10-17T20:38:46.743Z","customer_id":"RDX9Z4XTIZR7MRZJUXNY9HUK6I","id":"GQTFp1ZlXdpoW4o6eGiZhbjosiDFf","location_id":"XTI0H92143A39","note":"Brief description","order_id":"m2Hr8Hk8A3CTyQQ1k4ynExg92tO3","reference_id":"123456","source_type":"CARD","status":"CANCELED","total_money":{"amount":200,"currency":"USD"},"updated_at":"2018-10-17T20:38:57.693Z"}}
//
// swagger:model CancelPaymentResponse
type CancelPaymentResponse struct {

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The successfully canceled `Payment` object.
	Payment *Payment `json:"payment,omitempty"`
}

// Validate validates this cancel payment response
func (m *CancelPaymentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelPaymentResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *CancelPaymentResponse) validatePayment(formats strfmt.Registry) error {
	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cancel payment response based on the context it is used
func (m *CancelPaymentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelPaymentResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CancelPaymentResponse) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if m.Payment != nil {
		if err := m.Payment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelPaymentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelPaymentResponse) UnmarshalBinary(b []byte) error {
	var res CancelPaymentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
