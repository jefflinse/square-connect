// Code generated by go-swagger; DO NOT EDIT.

package v1_transactions

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

// NewRetrievePaymentParams creates a new RetrievePaymentParams object
// with the default values initialized.
func NewRetrievePaymentParams() *RetrievePaymentParams {
	var ()
	return &RetrievePaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrievePaymentParamsWithTimeout creates a new RetrievePaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrievePaymentParamsWithTimeout(timeout time.Duration) *RetrievePaymentParams {
	var ()
	return &RetrievePaymentParams{

		timeout: timeout,
	}
}

// NewRetrievePaymentParamsWithContext creates a new RetrievePaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrievePaymentParamsWithContext(ctx context.Context) *RetrievePaymentParams {
	var ()
	return &RetrievePaymentParams{

		Context: ctx,
	}
}

// NewRetrievePaymentParamsWithHTTPClient creates a new RetrievePaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrievePaymentParamsWithHTTPClient(client *http.Client) *RetrievePaymentParams {
	var ()
	return &RetrievePaymentParams{
		HTTPClient: client,
	}
}

/*RetrievePaymentParams contains all the parameters to send to the API endpoint
for the retrieve payment operation typically these are written to a http.Request
*/
type RetrievePaymentParams struct {

	/*LocationID
	  The ID of the payment's associated location.

	*/
	LocationID string
	/*PaymentID
	  The Square-issued payment ID. payment_id comes from Payment objects returned by the List Payments endpoint, Settlement objects returned by the List Settlements endpoint, or Refund objects returned by the List Refunds endpoint.

	*/
	PaymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve payment params
func (o *RetrievePaymentParams) WithTimeout(timeout time.Duration) *RetrievePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve payment params
func (o *RetrievePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve payment params
func (o *RetrievePaymentParams) WithContext(ctx context.Context) *RetrievePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve payment params
func (o *RetrievePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve payment params
func (o *RetrievePaymentParams) WithHTTPClient(client *http.Client) *RetrievePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve payment params
func (o *RetrievePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationID adds the locationID to the retrieve payment params
func (o *RetrievePaymentParams) WithLocationID(locationID string) *RetrievePaymentParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the retrieve payment params
func (o *RetrievePaymentParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithPaymentID adds the paymentID to the retrieve payment params
func (o *RetrievePaymentParams) WithPaymentID(paymentID string) *RetrievePaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the retrieve payment params
func (o *RetrievePaymentParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrievePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	// path param payment_id
	if err := r.SetPathParam("payment_id", o.PaymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
