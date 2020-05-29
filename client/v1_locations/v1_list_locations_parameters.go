// Code generated by go-swagger; DO NOT EDIT.

package v1_locations

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

// NewV1ListLocationsParams creates a new V1ListLocationsParams object
// with the default values initialized.
func NewV1ListLocationsParams() *V1ListLocationsParams {

	return &V1ListLocationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ListLocationsParamsWithTimeout creates a new V1ListLocationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ListLocationsParamsWithTimeout(timeout time.Duration) *V1ListLocationsParams {

	return &V1ListLocationsParams{

		timeout: timeout,
	}
}

// NewV1ListLocationsParamsWithContext creates a new V1ListLocationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ListLocationsParamsWithContext(ctx context.Context) *V1ListLocationsParams {

	return &V1ListLocationsParams{

		Context: ctx,
	}
}

// NewV1ListLocationsParamsWithHTTPClient creates a new V1ListLocationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ListLocationsParamsWithHTTPClient(client *http.Client) *V1ListLocationsParams {

	return &V1ListLocationsParams{
		HTTPClient: client,
	}
}

/*V1ListLocationsParams contains all the parameters to send to the API endpoint
for the v1 list locations operation typically these are written to a http.Request
*/
type V1ListLocationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 list locations params
func (o *V1ListLocationsParams) WithTimeout(timeout time.Duration) *V1ListLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 list locations params
func (o *V1ListLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 list locations params
func (o *V1ListLocationsParams) WithContext(ctx context.Context) *V1ListLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 list locations params
func (o *V1ListLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 list locations params
func (o *V1ListLocationsParams) WithHTTPClient(client *http.Client) *V1ListLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 list locations params
func (o *V1ListLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1ListLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
