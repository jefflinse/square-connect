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

// NewDeleteBreakTypeParams creates a new DeleteBreakTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBreakTypeParams() *DeleteBreakTypeParams {
	return &DeleteBreakTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBreakTypeParamsWithTimeout creates a new DeleteBreakTypeParams object
// with the ability to set a timeout on a request.
func NewDeleteBreakTypeParamsWithTimeout(timeout time.Duration) *DeleteBreakTypeParams {
	return &DeleteBreakTypeParams{
		timeout: timeout,
	}
}

// NewDeleteBreakTypeParamsWithContext creates a new DeleteBreakTypeParams object
// with the ability to set a context for a request.
func NewDeleteBreakTypeParamsWithContext(ctx context.Context) *DeleteBreakTypeParams {
	return &DeleteBreakTypeParams{
		Context: ctx,
	}
}

// NewDeleteBreakTypeParamsWithHTTPClient creates a new DeleteBreakTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBreakTypeParamsWithHTTPClient(client *http.Client) *DeleteBreakTypeParams {
	return &DeleteBreakTypeParams{
		HTTPClient: client,
	}
}

/* DeleteBreakTypeParams contains all the parameters to send to the API endpoint
   for the delete break type operation.

   Typically these are written to a http.Request.
*/
type DeleteBreakTypeParams struct {

	/* ID.

	   UUID for the `BreakType` being deleted.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete break type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBreakTypeParams) WithDefaults() *DeleteBreakTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete break type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBreakTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete break type params
func (o *DeleteBreakTypeParams) WithTimeout(timeout time.Duration) *DeleteBreakTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete break type params
func (o *DeleteBreakTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete break type params
func (o *DeleteBreakTypeParams) WithContext(ctx context.Context) *DeleteBreakTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete break type params
func (o *DeleteBreakTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete break type params
func (o *DeleteBreakTypeParams) WithHTTPClient(client *http.Client) *DeleteBreakTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete break type params
func (o *DeleteBreakTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete break type params
func (o *DeleteBreakTypeParams) WithID(id string) *DeleteBreakTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete break type params
func (o *DeleteBreakTypeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBreakTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
