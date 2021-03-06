// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewListLocationsParams creates a new ListLocationsParams object
// with the default values initialized.
func NewListLocationsParams() *ListLocationsParams {

	return &ListLocationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListLocationsParamsWithTimeout creates a new ListLocationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListLocationsParamsWithTimeout(timeout time.Duration) *ListLocationsParams {

	return &ListLocationsParams{

		timeout: timeout,
	}
}

// NewListLocationsParamsWithContext creates a new ListLocationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListLocationsParamsWithContext(ctx context.Context) *ListLocationsParams {

	return &ListLocationsParams{

		Context: ctx,
	}
}

// NewListLocationsParamsWithHTTPClient creates a new ListLocationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListLocationsParamsWithHTTPClient(client *http.Client) *ListLocationsParams {

	return &ListLocationsParams{
		HTTPClient: client,
	}
}

/*ListLocationsParams contains all the parameters to send to the API endpoint
for the list locations operation typically these are written to a http.Request
*/
type ListLocationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list locations params
func (o *ListLocationsParams) WithTimeout(timeout time.Duration) *ListLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list locations params
func (o *ListLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list locations params
func (o *ListLocationsParams) WithContext(ctx context.Context) *ListLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list locations params
func (o *ListLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list locations params
func (o *ListLocationsParams) WithHTTPClient(client *http.Client) *ListLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list locations params
func (o *ListLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
