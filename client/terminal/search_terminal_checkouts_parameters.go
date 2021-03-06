// Code generated by go-swagger; DO NOT EDIT.

package terminal

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

// NewSearchTerminalCheckoutsParams creates a new SearchTerminalCheckoutsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchTerminalCheckoutsParams() *SearchTerminalCheckoutsParams {
	return &SearchTerminalCheckoutsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchTerminalCheckoutsParamsWithTimeout creates a new SearchTerminalCheckoutsParams object
// with the ability to set a timeout on a request.
func NewSearchTerminalCheckoutsParamsWithTimeout(timeout time.Duration) *SearchTerminalCheckoutsParams {
	return &SearchTerminalCheckoutsParams{
		timeout: timeout,
	}
}

// NewSearchTerminalCheckoutsParamsWithContext creates a new SearchTerminalCheckoutsParams object
// with the ability to set a context for a request.
func NewSearchTerminalCheckoutsParamsWithContext(ctx context.Context) *SearchTerminalCheckoutsParams {
	return &SearchTerminalCheckoutsParams{
		Context: ctx,
	}
}

// NewSearchTerminalCheckoutsParamsWithHTTPClient creates a new SearchTerminalCheckoutsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchTerminalCheckoutsParamsWithHTTPClient(client *http.Client) *SearchTerminalCheckoutsParams {
	return &SearchTerminalCheckoutsParams{
		HTTPClient: client,
	}
}

/* SearchTerminalCheckoutsParams contains all the parameters to send to the API endpoint
   for the search terminal checkouts operation.

   Typically these are written to a http.Request.
*/
type SearchTerminalCheckoutsParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.SearchTerminalCheckoutsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search terminal checkouts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTerminalCheckoutsParams) WithDefaults() *SearchTerminalCheckoutsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search terminal checkouts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTerminalCheckoutsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) WithTimeout(timeout time.Duration) *SearchTerminalCheckoutsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) WithContext(ctx context.Context) *SearchTerminalCheckoutsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) WithHTTPClient(client *http.Client) *SearchTerminalCheckoutsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) WithBody(body *models.SearchTerminalCheckoutsRequest) *SearchTerminalCheckoutsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search terminal checkouts params
func (o *SearchTerminalCheckoutsParams) SetBody(body *models.SearchTerminalCheckoutsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchTerminalCheckoutsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
