// Code generated by go-swagger; DO NOT EDIT.

package terminal

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

// NewCancelTerminalCheckoutParams creates a new CancelTerminalCheckoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelTerminalCheckoutParams() *CancelTerminalCheckoutParams {
	return &CancelTerminalCheckoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelTerminalCheckoutParamsWithTimeout creates a new CancelTerminalCheckoutParams object
// with the ability to set a timeout on a request.
func NewCancelTerminalCheckoutParamsWithTimeout(timeout time.Duration) *CancelTerminalCheckoutParams {
	return &CancelTerminalCheckoutParams{
		timeout: timeout,
	}
}

// NewCancelTerminalCheckoutParamsWithContext creates a new CancelTerminalCheckoutParams object
// with the ability to set a context for a request.
func NewCancelTerminalCheckoutParamsWithContext(ctx context.Context) *CancelTerminalCheckoutParams {
	return &CancelTerminalCheckoutParams{
		Context: ctx,
	}
}

// NewCancelTerminalCheckoutParamsWithHTTPClient creates a new CancelTerminalCheckoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelTerminalCheckoutParamsWithHTTPClient(client *http.Client) *CancelTerminalCheckoutParams {
	return &CancelTerminalCheckoutParams{
		HTTPClient: client,
	}
}

/* CancelTerminalCheckoutParams contains all the parameters to send to the API endpoint
   for the cancel terminal checkout operation.

   Typically these are written to a http.Request.
*/
type CancelTerminalCheckoutParams struct {

	/* CheckoutID.

	   Unique ID for the desired `TerminalCheckout`
	*/
	CheckoutID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel terminal checkout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelTerminalCheckoutParams) WithDefaults() *CancelTerminalCheckoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel terminal checkout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelTerminalCheckoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) WithTimeout(timeout time.Duration) *CancelTerminalCheckoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) WithContext(ctx context.Context) *CancelTerminalCheckoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) WithHTTPClient(client *http.Client) *CancelTerminalCheckoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckoutID adds the checkoutID to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) WithCheckoutID(checkoutID string) *CancelTerminalCheckoutParams {
	o.SetCheckoutID(checkoutID)
	return o
}

// SetCheckoutID adds the checkoutId to the cancel terminal checkout params
func (o *CancelTerminalCheckoutParams) SetCheckoutID(checkoutID string) {
	o.CheckoutID = checkoutID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelTerminalCheckoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param checkout_id
	if err := r.SetPathParam("checkout_id", o.CheckoutID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
