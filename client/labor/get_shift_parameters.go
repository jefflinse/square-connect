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

// NewGetShiftParams creates a new GetShiftParams object
// with the default values initialized.
func NewGetShiftParams() *GetShiftParams {
	var ()
	return &GetShiftParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetShiftParamsWithTimeout creates a new GetShiftParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetShiftParamsWithTimeout(timeout time.Duration) *GetShiftParams {
	var ()
	return &GetShiftParams{

		timeout: timeout,
	}
}

// NewGetShiftParamsWithContext creates a new GetShiftParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetShiftParamsWithContext(ctx context.Context) *GetShiftParams {
	var ()
	return &GetShiftParams{

		Context: ctx,
	}
}

// NewGetShiftParamsWithHTTPClient creates a new GetShiftParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetShiftParamsWithHTTPClient(client *http.Client) *GetShiftParams {
	var ()
	return &GetShiftParams{
		HTTPClient: client,
	}
}

/*GetShiftParams contains all the parameters to send to the API endpoint
for the get shift operation typically these are written to a http.Request
*/
type GetShiftParams struct {

	/*ID
	  UUID for the `Shift` being retrieved.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get shift params
func (o *GetShiftParams) WithTimeout(timeout time.Duration) *GetShiftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shift params
func (o *GetShiftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shift params
func (o *GetShiftParams) WithContext(ctx context.Context) *GetShiftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shift params
func (o *GetShiftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shift params
func (o *GetShiftParams) WithHTTPClient(client *http.Client) *GetShiftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shift params
func (o *GetShiftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get shift params
func (o *GetShiftParams) WithID(id string) *GetShiftParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get shift params
func (o *GetShiftParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetShiftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
