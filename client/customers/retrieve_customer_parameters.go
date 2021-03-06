// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewRetrieveCustomerParams creates a new RetrieveCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveCustomerParams() *RetrieveCustomerParams {
	return &RetrieveCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveCustomerParamsWithTimeout creates a new RetrieveCustomerParams object
// with the ability to set a timeout on a request.
func NewRetrieveCustomerParamsWithTimeout(timeout time.Duration) *RetrieveCustomerParams {
	return &RetrieveCustomerParams{
		timeout: timeout,
	}
}

// NewRetrieveCustomerParamsWithContext creates a new RetrieveCustomerParams object
// with the ability to set a context for a request.
func NewRetrieveCustomerParamsWithContext(ctx context.Context) *RetrieveCustomerParams {
	return &RetrieveCustomerParams{
		Context: ctx,
	}
}

// NewRetrieveCustomerParamsWithHTTPClient creates a new RetrieveCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveCustomerParamsWithHTTPClient(client *http.Client) *RetrieveCustomerParams {
	return &RetrieveCustomerParams{
		HTTPClient: client,
	}
}

/* RetrieveCustomerParams contains all the parameters to send to the API endpoint
   for the retrieve customer operation.

   Typically these are written to a http.Request.
*/
type RetrieveCustomerParams struct {

	/* CustomerID.

	   The ID of the customer to retrieve.
	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveCustomerParams) WithDefaults() *RetrieveCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve customer params
func (o *RetrieveCustomerParams) WithTimeout(timeout time.Duration) *RetrieveCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve customer params
func (o *RetrieveCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve customer params
func (o *RetrieveCustomerParams) WithContext(ctx context.Context) *RetrieveCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve customer params
func (o *RetrieveCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve customer params
func (o *RetrieveCustomerParams) WithHTTPClient(client *http.Client) *RetrieveCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve customer params
func (o *RetrieveCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the retrieve customer params
func (o *RetrieveCustomerParams) WithCustomerID(customerID string) *RetrieveCustomerParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the retrieve customer params
func (o *RetrieveCustomerParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
