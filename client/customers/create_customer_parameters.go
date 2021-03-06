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

	"github.com/jefflinse/square-connect/models"
)

// NewCreateCustomerParams creates a new CreateCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCustomerParams() *CreateCustomerParams {
	return &CreateCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomerParamsWithTimeout creates a new CreateCustomerParams object
// with the ability to set a timeout on a request.
func NewCreateCustomerParamsWithTimeout(timeout time.Duration) *CreateCustomerParams {
	return &CreateCustomerParams{
		timeout: timeout,
	}
}

// NewCreateCustomerParamsWithContext creates a new CreateCustomerParams object
// with the ability to set a context for a request.
func NewCreateCustomerParamsWithContext(ctx context.Context) *CreateCustomerParams {
	return &CreateCustomerParams{
		Context: ctx,
	}
}

// NewCreateCustomerParamsWithHTTPClient creates a new CreateCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCustomerParamsWithHTTPClient(client *http.Client) *CreateCustomerParams {
	return &CreateCustomerParams{
		HTTPClient: client,
	}
}

/* CreateCustomerParams contains all the parameters to send to the API endpoint
   for the create customer operation.

   Typically these are written to a http.Request.
*/
type CreateCustomerParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.CreateCustomerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCustomerParams) WithDefaults() *CreateCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create customer params
func (o *CreateCustomerParams) WithTimeout(timeout time.Duration) *CreateCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create customer params
func (o *CreateCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create customer params
func (o *CreateCustomerParams) WithContext(ctx context.Context) *CreateCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create customer params
func (o *CreateCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create customer params
func (o *CreateCustomerParams) WithHTTPClient(client *http.Client) *CreateCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create customer params
func (o *CreateCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create customer params
func (o *CreateCustomerParams) WithBody(body *models.CreateCustomerRequest) *CreateCustomerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create customer params
func (o *CreateCustomerParams) SetBody(body *models.CreateCustomerRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
