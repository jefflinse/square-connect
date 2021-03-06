// Code generated by go-swagger; DO NOT EDIT.

package bookings

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

// NewRetrieveBusinessBookingProfileParams creates a new RetrieveBusinessBookingProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveBusinessBookingProfileParams() *RetrieveBusinessBookingProfileParams {
	return &RetrieveBusinessBookingProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveBusinessBookingProfileParamsWithTimeout creates a new RetrieveBusinessBookingProfileParams object
// with the ability to set a timeout on a request.
func NewRetrieveBusinessBookingProfileParamsWithTimeout(timeout time.Duration) *RetrieveBusinessBookingProfileParams {
	return &RetrieveBusinessBookingProfileParams{
		timeout: timeout,
	}
}

// NewRetrieveBusinessBookingProfileParamsWithContext creates a new RetrieveBusinessBookingProfileParams object
// with the ability to set a context for a request.
func NewRetrieveBusinessBookingProfileParamsWithContext(ctx context.Context) *RetrieveBusinessBookingProfileParams {
	return &RetrieveBusinessBookingProfileParams{
		Context: ctx,
	}
}

// NewRetrieveBusinessBookingProfileParamsWithHTTPClient creates a new RetrieveBusinessBookingProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveBusinessBookingProfileParamsWithHTTPClient(client *http.Client) *RetrieveBusinessBookingProfileParams {
	return &RetrieveBusinessBookingProfileParams{
		HTTPClient: client,
	}
}

/* RetrieveBusinessBookingProfileParams contains all the parameters to send to the API endpoint
   for the retrieve business booking profile operation.

   Typically these are written to a http.Request.
*/
type RetrieveBusinessBookingProfileParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve business booking profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveBusinessBookingProfileParams) WithDefaults() *RetrieveBusinessBookingProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve business booking profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveBusinessBookingProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve business booking profile params
func (o *RetrieveBusinessBookingProfileParams) WithTimeout(timeout time.Duration) *RetrieveBusinessBookingProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve business booking profile params
func (o *RetrieveBusinessBookingProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve business booking profile params
func (o *RetrieveBusinessBookingProfileParams) WithContext(ctx context.Context) *RetrieveBusinessBookingProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve business booking profile params
func (o *RetrieveBusinessBookingProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve business booking profile params
func (o *RetrieveBusinessBookingProfileParams) WithHTTPClient(client *http.Client) *RetrieveBusinessBookingProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve business booking profile params
func (o *RetrieveBusinessBookingProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveBusinessBookingProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
