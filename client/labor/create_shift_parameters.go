// Code generated by go-swagger; DO NOT EDIT.

package labor

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

// NewCreateShiftParams creates a new CreateShiftParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateShiftParams() *CreateShiftParams {
	return &CreateShiftParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateShiftParamsWithTimeout creates a new CreateShiftParams object
// with the ability to set a timeout on a request.
func NewCreateShiftParamsWithTimeout(timeout time.Duration) *CreateShiftParams {
	return &CreateShiftParams{
		timeout: timeout,
	}
}

// NewCreateShiftParamsWithContext creates a new CreateShiftParams object
// with the ability to set a context for a request.
func NewCreateShiftParamsWithContext(ctx context.Context) *CreateShiftParams {
	return &CreateShiftParams{
		Context: ctx,
	}
}

// NewCreateShiftParamsWithHTTPClient creates a new CreateShiftParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateShiftParamsWithHTTPClient(client *http.Client) *CreateShiftParams {
	return &CreateShiftParams{
		HTTPClient: client,
	}
}

/* CreateShiftParams contains all the parameters to send to the API endpoint
   for the create shift operation.

   Typically these are written to a http.Request.
*/
type CreateShiftParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.CreateShiftRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create shift params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateShiftParams) WithDefaults() *CreateShiftParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create shift params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateShiftParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create shift params
func (o *CreateShiftParams) WithTimeout(timeout time.Duration) *CreateShiftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create shift params
func (o *CreateShiftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create shift params
func (o *CreateShiftParams) WithContext(ctx context.Context) *CreateShiftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create shift params
func (o *CreateShiftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create shift params
func (o *CreateShiftParams) WithHTTPClient(client *http.Client) *CreateShiftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create shift params
func (o *CreateShiftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create shift params
func (o *CreateShiftParams) WithBody(body *models.CreateShiftRequest) *CreateShiftParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create shift params
func (o *CreateShiftParams) SetBody(body *models.CreateShiftRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateShiftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
