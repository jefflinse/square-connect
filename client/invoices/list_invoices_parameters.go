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
	"github.com/go-openapi/swag"
)

// NewListInvoicesParams creates a new ListInvoicesParams object
// with the default values initialized.
func NewListInvoicesParams() *ListInvoicesParams {
	var ()
	return &ListInvoicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListInvoicesParamsWithTimeout creates a new ListInvoicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListInvoicesParamsWithTimeout(timeout time.Duration) *ListInvoicesParams {
	var ()
	return &ListInvoicesParams{

		timeout: timeout,
	}
}

// NewListInvoicesParamsWithContext creates a new ListInvoicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListInvoicesParamsWithContext(ctx context.Context) *ListInvoicesParams {
	var ()
	return &ListInvoicesParams{

		Context: ctx,
	}
}

// NewListInvoicesParamsWithHTTPClient creates a new ListInvoicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListInvoicesParamsWithHTTPClient(client *http.Client) *ListInvoicesParams {
	var ()
	return &ListInvoicesParams{
		HTTPClient: client,
	}
}

/*ListInvoicesParams contains all the parameters to send to the API endpoint
for the list invoices operation typically these are written to a http.Request
*/
type ListInvoicesParams struct {

	/*Cursor
	  A pagination cursor returned by a previous call to this endpoint.
	Provide this cursor to retrieve the next set of results for your original query.

	For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).

	*/
	Cursor *string
	/*Limit
	  The maximum number of invoices to return (200 is the maximum `limit`).
	If not provided, the server
	uses a default limit of 100 invoices.

	*/
	Limit *int64
	/*LocationID
	  The ID of the location for which to list invoices.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list invoices params
func (o *ListInvoicesParams) WithTimeout(timeout time.Duration) *ListInvoicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list invoices params
func (o *ListInvoicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list invoices params
func (o *ListInvoicesParams) WithContext(ctx context.Context) *ListInvoicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list invoices params
func (o *ListInvoicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list invoices params
func (o *ListInvoicesParams) WithHTTPClient(client *http.Client) *ListInvoicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list invoices params
func (o *ListInvoicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list invoices params
func (o *ListInvoicesParams) WithCursor(cursor *string) *ListInvoicesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list invoices params
func (o *ListInvoicesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the list invoices params
func (o *ListInvoicesParams) WithLimit(limit *int64) *ListInvoicesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list invoices params
func (o *ListInvoicesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the list invoices params
func (o *ListInvoicesParams) WithLocationID(locationID string) *ListInvoicesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list invoices params
func (o *ListInvoicesParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ListInvoicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// query param location_id
	qrLocationID := o.LocationID
	qLocationID := qrLocationID
	if qLocationID != "" {
		if err := r.SetQueryParam("location_id", qLocationID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
