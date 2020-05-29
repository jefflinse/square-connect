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
	"github.com/go-openapi/swag"
)

// NewListBreakTypesParams creates a new ListBreakTypesParams object
// with the default values initialized.
func NewListBreakTypesParams() *ListBreakTypesParams {
	var ()
	return &ListBreakTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBreakTypesParamsWithTimeout creates a new ListBreakTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBreakTypesParamsWithTimeout(timeout time.Duration) *ListBreakTypesParams {
	var ()
	return &ListBreakTypesParams{

		timeout: timeout,
	}
}

// NewListBreakTypesParamsWithContext creates a new ListBreakTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBreakTypesParamsWithContext(ctx context.Context) *ListBreakTypesParams {
	var ()
	return &ListBreakTypesParams{

		Context: ctx,
	}
}

// NewListBreakTypesParamsWithHTTPClient creates a new ListBreakTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBreakTypesParamsWithHTTPClient(client *http.Client) *ListBreakTypesParams {
	var ()
	return &ListBreakTypesParams{
		HTTPClient: client,
	}
}

/*ListBreakTypesParams contains all the parameters to send to the API endpoint
for the list break types operation typically these are written to a http.Request
*/
type ListBreakTypesParams struct {

	/*Cursor
	  Pointer to the next page of Break Type results to fetch.

	*/
	Cursor *string
	/*Limit
	  Maximum number of Break Types to return per page. Can range between 1
	and 200. The default is the maximum at 200.

	*/
	Limit *int64
	/*LocationID
	  Filter Break Types returned to only those that are associated with the
	specified location.

	*/
	LocationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list break types params
func (o *ListBreakTypesParams) WithTimeout(timeout time.Duration) *ListBreakTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list break types params
func (o *ListBreakTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list break types params
func (o *ListBreakTypesParams) WithContext(ctx context.Context) *ListBreakTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list break types params
func (o *ListBreakTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list break types params
func (o *ListBreakTypesParams) WithHTTPClient(client *http.Client) *ListBreakTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list break types params
func (o *ListBreakTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list break types params
func (o *ListBreakTypesParams) WithCursor(cursor *string) *ListBreakTypesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list break types params
func (o *ListBreakTypesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the list break types params
func (o *ListBreakTypesParams) WithLimit(limit *int64) *ListBreakTypesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list break types params
func (o *ListBreakTypesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the list break types params
func (o *ListBreakTypesParams) WithLocationID(locationID *string) *ListBreakTypesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list break types params
func (o *ListBreakTypesParams) SetLocationID(locationID *string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ListBreakTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.LocationID != nil {

		// query param location_id
		var qrLocationID string
		if o.LocationID != nil {
			qrLocationID = *o.LocationID
		}
		qLocationID := qrLocationID
		if qLocationID != "" {
			if err := r.SetQueryParam("location_id", qLocationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
