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

	"github.com/jefflinse/square-connect/models"
)

// NewCreateBookingParams creates a new CreateBookingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBookingParams() *CreateBookingParams {
	return &CreateBookingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBookingParamsWithTimeout creates a new CreateBookingParams object
// with the ability to set a timeout on a request.
func NewCreateBookingParamsWithTimeout(timeout time.Duration) *CreateBookingParams {
	return &CreateBookingParams{
		timeout: timeout,
	}
}

// NewCreateBookingParamsWithContext creates a new CreateBookingParams object
// with the ability to set a context for a request.
func NewCreateBookingParamsWithContext(ctx context.Context) *CreateBookingParams {
	return &CreateBookingParams{
		Context: ctx,
	}
}

// NewCreateBookingParamsWithHTTPClient creates a new CreateBookingParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBookingParamsWithHTTPClient(client *http.Client) *CreateBookingParams {
	return &CreateBookingParams{
		HTTPClient: client,
	}
}

/* CreateBookingParams contains all the parameters to send to the API endpoint
   for the create booking operation.

   Typically these are written to a http.Request.
*/
type CreateBookingParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.CreateBookingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create booking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBookingParams) WithDefaults() *CreateBookingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create booking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBookingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create booking params
func (o *CreateBookingParams) WithTimeout(timeout time.Duration) *CreateBookingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create booking params
func (o *CreateBookingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create booking params
func (o *CreateBookingParams) WithContext(ctx context.Context) *CreateBookingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create booking params
func (o *CreateBookingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create booking params
func (o *CreateBookingParams) WithHTTPClient(client *http.Client) *CreateBookingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create booking params
func (o *CreateBookingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create booking params
func (o *CreateBookingParams) WithBody(body *models.CreateBookingRequest) *CreateBookingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create booking params
func (o *CreateBookingParams) SetBody(body *models.CreateBookingRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBookingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}