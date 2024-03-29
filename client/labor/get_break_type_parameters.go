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
)

// NewGetBreakTypeParams creates a new GetBreakTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBreakTypeParams() *GetBreakTypeParams {
	return &GetBreakTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBreakTypeParamsWithTimeout creates a new GetBreakTypeParams object
// with the ability to set a timeout on a request.
func NewGetBreakTypeParamsWithTimeout(timeout time.Duration) *GetBreakTypeParams {
	return &GetBreakTypeParams{
		timeout: timeout,
	}
}

// NewGetBreakTypeParamsWithContext creates a new GetBreakTypeParams object
// with the ability to set a context for a request.
func NewGetBreakTypeParamsWithContext(ctx context.Context) *GetBreakTypeParams {
	return &GetBreakTypeParams{
		Context: ctx,
	}
}

// NewGetBreakTypeParamsWithHTTPClient creates a new GetBreakTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBreakTypeParamsWithHTTPClient(client *http.Client) *GetBreakTypeParams {
	return &GetBreakTypeParams{
		HTTPClient: client,
	}
}

/* GetBreakTypeParams contains all the parameters to send to the API endpoint
   for the get break type operation.

   Typically these are written to a http.Request.
*/
type GetBreakTypeParams struct {

	/* ID.

	   UUID for the `BreakType` being retrieved.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get break type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBreakTypeParams) WithDefaults() *GetBreakTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get break type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBreakTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get break type params
func (o *GetBreakTypeParams) WithTimeout(timeout time.Duration) *GetBreakTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get break type params
func (o *GetBreakTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get break type params
func (o *GetBreakTypeParams) WithContext(ctx context.Context) *GetBreakTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get break type params
func (o *GetBreakTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get break type params
func (o *GetBreakTypeParams) WithHTTPClient(client *http.Client) *GetBreakTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get break type params
func (o *GetBreakTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get break type params
func (o *GetBreakTypeParams) WithID(id string) *GetBreakTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get break type params
func (o *GetBreakTypeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBreakTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
