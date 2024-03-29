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

// GetPaymentResponse Defines the fields that are included in the response body of
// a request to the [GetPayment](#endpoint-payments-getpayment) endpoint.
// Example: {"payment":{"amount_money":{"amount":200,"currency":"USD"},"app_fee_money":{"amount":10,"currency":"USD"},"card_details":{"auth_result_code":"nsAyY2","avs_status":"AVS_ACCEPTED","card":{"bin":"411111","card_brand":"VISA","card_type":"DEBIT","exp_month":7,"exp_year":2026,"fingerprint":"sq-1-TpmjbNBMFdibiIjpQI5LiRgNUBC7u1689i0TgHjnlyHEWYB7tnn-K4QbW4ttvtaqXw","last_4":"1111","prepaid_type":"PREPAID"},"card_payment_timeline":{"authorized_at":"2019-07-10T13:23:49.234Z","captured_at":"2019-07-10T13:23:49.446Z"},"cvv_status":"CVV_ACCEPTED","entry_method":"ON_FILE","statement_description":"SQ *MY MERCHANT","status":"CAPTURED"},"created_at":"2019-07-10T13:23:49.154Z","customer_id":"RDX9Z4XTIZR7MRZJUXNY9HUK6I","id":"GQTFp1ZlXdpoW4o6eGiZhbjosiDFf","location_id":"XTI0H92143A39","note":"Brief description","order_id":"m2Hr8Hk8A3CTyQQ1k4ynExg92tO3","processing_fee":[{"amount_money":{"amount":36,"currency":"USD"},"effective_at":"2019-07-10T15:23:49.000Z","type":"INITIAL"}],"receipt_number":"GQTF","receipt_url":"https://squareup.com/receipt/preview/GQTFp1ZlXdpoW4o6eGiZhbjosiDFf","reference_id":"123456","source_type":"CARD","status":"COMPLETED","total_money":{"amount":200,"currency":"USD"},"updated_at":"2019-07-10T13:23:49.446Z"}}
//
// swagger:model GetPaymentResponse
type GetPaymentResponse struct {

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The requested `Payment`.
	Payment *Payment `json:"payment,omitempty"`
}

// Validate validates this get payment response
func (m *GetPaymentResponse) Validate(formats strfmt.Registry) error {
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

func (m *GetPaymentResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *GetPaymentResponse) validatePayment(formats strfmt.Registry) error {
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

// ContextValidate validate this get payment response based on the context it is used
func (m *GetPaymentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetPaymentResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetPaymentResponse) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetPaymentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPaymentResponse) UnmarshalBinary(b []byte) error {
	var res GetPaymentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
