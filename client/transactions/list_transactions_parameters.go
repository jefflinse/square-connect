// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewListTransactionsParams creates a new ListTransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTransactionsParams() *ListTransactionsParams {
	return &ListTransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTransactionsParamsWithTimeout creates a new ListTransactionsParams object
// with the ability to set a timeout on a request.
func NewListTransactionsParamsWithTimeout(timeout time.Duration) *ListTransactionsParams {
	return &ListTransactionsParams{
		timeout: timeout,
	}
}

// NewListTransactionsParamsWithContext creates a new ListTransactionsParams object
// with the ability to set a context for a request.
func NewListTransactionsParamsWithContext(ctx context.Context) *ListTransactionsParams {
	return &ListTransactionsParams{
		Context: ctx,
	}
}

// NewListTransactionsParamsWithHTTPClient creates a new ListTransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTransactionsParamsWithHTTPClient(client *http.Client) *ListTransactionsParams {
	return &ListTransactionsParams{
		HTTPClient: client,
	}
}

/* ListTransactionsParams contains all the parameters to send to the API endpoint
   for the list transactions operation.

   Typically these are written to a http.Request.
*/
type ListTransactionsParams struct {

	/* BeginTime.

	     The beginning of the requested reporting period, in RFC 3339 format.

	See [Date ranges](#dateranges) for details on date inclusivity/exclusivity.

	Default value: The current time minus one year.
	*/
	BeginTime *string

	/* Cursor.

	     A pagination cursor returned by a previous call to this endpoint.
	Provide this to retrieve the next set of results for your original query.

	See [Paginating results](#paginatingresults) for more information.
	*/
	Cursor *string

	/* EndTime.

	     The end of the requested reporting period, in RFC 3339 format.

	See [Date ranges](#dateranges) for details on date inclusivity/exclusivity.

	Default value: The current time.
	*/
	EndTime *string

	/* LocationID.

	   The ID of the location to list transactions for.
	*/
	LocationID string

	/* SortOrder.

	     The order in which results are listed in the response (`ASC` for
	oldest first, `DESC` for newest first).

	Default value: `DESC`
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTransactionsParams) WithDefaults() *ListTransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTransactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list transactions params
func (o *ListTransactionsParams) WithTimeout(timeout time.Duration) *ListTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list transactions params
func (o *ListTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list transactions params
func (o *ListTransactionsParams) WithContext(ctx context.Context) *ListTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list transactions params
func (o *ListTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list transactions params
func (o *ListTransactionsParams) WithHTTPClient(client *http.Client) *ListTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list transactions params
func (o *ListTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBeginTime adds the beginTime to the list transactions params
func (o *ListTransactionsParams) WithBeginTime(beginTime *string) *ListTransactionsParams {
	o.SetBeginTime(beginTime)
	return o
}

// SetBeginTime adds the beginTime to the list transactions params
func (o *ListTransactionsParams) SetBeginTime(beginTime *string) {
	o.BeginTime = beginTime
}

// WithCursor adds the cursor to the list transactions params
func (o *ListTransactionsParams) WithCursor(cursor *string) *ListTransactionsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list transactions params
func (o *ListTransactionsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithEndTime adds the endTime to the list transactions params
func (o *ListTransactionsParams) WithEndTime(endTime *string) *ListTransactionsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the list transactions params
func (o *ListTransactionsParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithLocationID adds the locationID to the list transactions params
func (o *ListTransactionsParams) WithLocationID(locationID string) *ListTransactionsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list transactions params
func (o *ListTransactionsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithSortOrder adds the sortOrder to the list transactions params
func (o *ListTransactionsParams) WithSortOrder(sortOrder *string) *ListTransactionsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the list transactions params
func (o *ListTransactionsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *ListTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BeginTime != nil {

		// query param begin_time
		var qrBeginTime string

		if o.BeginTime != nil {
			qrBeginTime = *o.BeginTime
		}
		qBeginTime := qrBeginTime
		if qBeginTime != "" {

			if err := r.SetQueryParam("begin_time", qBeginTime); err != nil {
				return err
			}
		}
	}

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

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime string

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}
	}

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	if o.SortOrder != nil {

		// query param sort_order
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sort_order", qSortOrder); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
