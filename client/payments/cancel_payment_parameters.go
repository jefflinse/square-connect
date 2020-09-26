// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCancelPaymentParams creates a new CancelPaymentParams object
// with the default values initialized.
func NewCancelPaymentParams() *CancelPaymentParams {
	var ()
	return &CancelPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelPaymentParamsWithTimeout creates a new CancelPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelPaymentParamsWithTimeout(timeout time.Duration) *CancelPaymentParams {
	var ()
	return &CancelPaymentParams{

		timeout: timeout,
	}
}

// NewCancelPaymentParamsWithContext creates a new CancelPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelPaymentParamsWithContext(ctx context.Context) *CancelPaymentParams {
	var ()
	return &CancelPaymentParams{

		Context: ctx,
	}
}

// NewCancelPaymentParamsWithHTTPClient creates a new CancelPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelPaymentParamsWithHTTPClient(client *http.Client) *CancelPaymentParams {
	var ()
	return &CancelPaymentParams{
		HTTPClient: client,
	}
}

/*CancelPaymentParams contains all the parameters to send to the API endpoint
for the cancel payment operation typically these are written to a http.Request
*/
type CancelPaymentParams struct {

	/*PaymentID
	  `payment_id` identifying the payment to be canceled.

	*/
	PaymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel payment params
func (o *CancelPaymentParams) WithTimeout(timeout time.Duration) *CancelPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel payment params
func (o *CancelPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel payment params
func (o *CancelPaymentParams) WithContext(ctx context.Context) *CancelPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel payment params
func (o *CancelPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel payment params
func (o *CancelPaymentParams) WithHTTPClient(client *http.Client) *CancelPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel payment params
func (o *CancelPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentID adds the paymentID to the cancel payment params
func (o *CancelPaymentParams) WithPaymentID(paymentID string) *CancelPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the cancel payment params
func (o *CancelPaymentParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param payment_id
	if err := r.SetPathParam("payment_id", o.PaymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}