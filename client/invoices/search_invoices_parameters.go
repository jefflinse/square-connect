// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewSearchInvoicesParams creates a new SearchInvoicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchInvoicesParams() *SearchInvoicesParams {
	return &SearchInvoicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchInvoicesParamsWithTimeout creates a new SearchInvoicesParams object
// with the ability to set a timeout on a request.
func NewSearchInvoicesParamsWithTimeout(timeout time.Duration) *SearchInvoicesParams {
	return &SearchInvoicesParams{
		timeout: timeout,
	}
}

// NewSearchInvoicesParamsWithContext creates a new SearchInvoicesParams object
// with the ability to set a context for a request.
func NewSearchInvoicesParamsWithContext(ctx context.Context) *SearchInvoicesParams {
	return &SearchInvoicesParams{
		Context: ctx,
	}
}

// NewSearchInvoicesParamsWithHTTPClient creates a new SearchInvoicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchInvoicesParamsWithHTTPClient(client *http.Client) *SearchInvoicesParams {
	return &SearchInvoicesParams{
		HTTPClient: client,
	}
}

/* SearchInvoicesParams contains all the parameters to send to the API endpoint
   for the search invoices operation.

   Typically these are written to a http.Request.
*/
type SearchInvoicesParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.SearchInvoicesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search invoices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchInvoicesParams) WithDefaults() *SearchInvoicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search invoices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchInvoicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search invoices params
func (o *SearchInvoicesParams) WithTimeout(timeout time.Duration) *SearchInvoicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search invoices params
func (o *SearchInvoicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search invoices params
func (o *SearchInvoicesParams) WithContext(ctx context.Context) *SearchInvoicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search invoices params
func (o *SearchInvoicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search invoices params
func (o *SearchInvoicesParams) WithHTTPClient(client *http.Client) *SearchInvoicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search invoices params
func (o *SearchInvoicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search invoices params
func (o *SearchInvoicesParams) WithBody(body *models.SearchInvoicesRequest) *SearchInvoicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search invoices params
func (o *SearchInvoicesParams) SetBody(body *models.SearchInvoicesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchInvoicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
