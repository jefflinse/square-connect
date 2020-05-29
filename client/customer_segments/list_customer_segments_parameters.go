// Code generated by go-swagger; DO NOT EDIT.

package customer_segments

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

// NewListCustomerSegmentsParams creates a new ListCustomerSegmentsParams object
// with the default values initialized.
func NewListCustomerSegmentsParams() *ListCustomerSegmentsParams {
	var ()
	return &ListCustomerSegmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCustomerSegmentsParamsWithTimeout creates a new ListCustomerSegmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCustomerSegmentsParamsWithTimeout(timeout time.Duration) *ListCustomerSegmentsParams {
	var ()
	return &ListCustomerSegmentsParams{

		timeout: timeout,
	}
}

// NewListCustomerSegmentsParamsWithContext creates a new ListCustomerSegmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCustomerSegmentsParamsWithContext(ctx context.Context) *ListCustomerSegmentsParams {
	var ()
	return &ListCustomerSegmentsParams{

		Context: ctx,
	}
}

// NewListCustomerSegmentsParamsWithHTTPClient creates a new ListCustomerSegmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCustomerSegmentsParamsWithHTTPClient(client *http.Client) *ListCustomerSegmentsParams {
	var ()
	return &ListCustomerSegmentsParams{
		HTTPClient: client,
	}
}

/*ListCustomerSegmentsParams contains all the parameters to send to the API endpoint
for the list customer segments operation typically these are written to a http.Request
*/
type ListCustomerSegmentsParams struct {

	/*Cursor
	  A pagination cursor returned by previous calls to __ListCustomerSegments__.
	Used to retrieve the next set of query results.

	See the [Pagination guide](https://developer.squareup.com/docs/docs/working-with-apis/pagination) for more information.

	*/
	Cursor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list customer segments params
func (o *ListCustomerSegmentsParams) WithTimeout(timeout time.Duration) *ListCustomerSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list customer segments params
func (o *ListCustomerSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list customer segments params
func (o *ListCustomerSegmentsParams) WithContext(ctx context.Context) *ListCustomerSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list customer segments params
func (o *ListCustomerSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list customer segments params
func (o *ListCustomerSegmentsParams) WithHTTPClient(client *http.Client) *ListCustomerSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list customer segments params
func (o *ListCustomerSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list customer segments params
func (o *ListCustomerSegmentsParams) WithCursor(cursor *string) *ListCustomerSegmentsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list customer segments params
func (o *ListCustomerSegmentsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WriteToRequest writes these params to a swagger request
func (o *ListCustomerSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
