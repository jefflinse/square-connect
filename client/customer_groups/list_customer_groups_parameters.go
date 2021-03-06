// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

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

// NewListCustomerGroupsParams creates a new ListCustomerGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCustomerGroupsParams() *ListCustomerGroupsParams {
	return &ListCustomerGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCustomerGroupsParamsWithTimeout creates a new ListCustomerGroupsParams object
// with the ability to set a timeout on a request.
func NewListCustomerGroupsParamsWithTimeout(timeout time.Duration) *ListCustomerGroupsParams {
	return &ListCustomerGroupsParams{
		timeout: timeout,
	}
}

// NewListCustomerGroupsParamsWithContext creates a new ListCustomerGroupsParams object
// with the ability to set a context for a request.
func NewListCustomerGroupsParamsWithContext(ctx context.Context) *ListCustomerGroupsParams {
	return &ListCustomerGroupsParams{
		Context: ctx,
	}
}

// NewListCustomerGroupsParamsWithHTTPClient creates a new ListCustomerGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCustomerGroupsParamsWithHTTPClient(client *http.Client) *ListCustomerGroupsParams {
	return &ListCustomerGroupsParams{
		HTTPClient: client,
	}
}

/* ListCustomerGroupsParams contains all the parameters to send to the API endpoint
   for the list customer groups operation.

   Typically these are written to a http.Request.
*/
type ListCustomerGroupsParams struct {

	/* Cursor.

	     A pagination cursor returned by a previous call to this endpoint.
	Provide this to retrieve the next set of results for your original query.

	See the [Pagination guide](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
	*/
	Cursor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list customer groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCustomerGroupsParams) WithDefaults() *ListCustomerGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list customer groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCustomerGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list customer groups params
func (o *ListCustomerGroupsParams) WithTimeout(timeout time.Duration) *ListCustomerGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list customer groups params
func (o *ListCustomerGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list customer groups params
func (o *ListCustomerGroupsParams) WithContext(ctx context.Context) *ListCustomerGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list customer groups params
func (o *ListCustomerGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list customer groups params
func (o *ListCustomerGroupsParams) WithHTTPClient(client *http.Client) *ListCustomerGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list customer groups params
func (o *ListCustomerGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list customer groups params
func (o *ListCustomerGroupsParams) WithCursor(cursor *string) *ListCustomerGroupsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list customer groups params
func (o *ListCustomerGroupsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WriteToRequest writes these params to a swagger request
func (o *ListCustomerGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
