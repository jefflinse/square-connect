// Code generated by go-swagger; DO NOT EDIT.

package employees

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

// NewListEmployeesParams creates a new ListEmployeesParams object
// with the default values initialized.
func NewListEmployeesParams() *ListEmployeesParams {
	var ()
	return &ListEmployeesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEmployeesParamsWithTimeout creates a new ListEmployeesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEmployeesParamsWithTimeout(timeout time.Duration) *ListEmployeesParams {
	var ()
	return &ListEmployeesParams{

		timeout: timeout,
	}
}

// NewListEmployeesParamsWithContext creates a new ListEmployeesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEmployeesParamsWithContext(ctx context.Context) *ListEmployeesParams {
	var ()
	return &ListEmployeesParams{

		Context: ctx,
	}
}

// NewListEmployeesParamsWithHTTPClient creates a new ListEmployeesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEmployeesParamsWithHTTPClient(client *http.Client) *ListEmployeesParams {
	var ()
	return &ListEmployeesParams{
		HTTPClient: client,
	}
}

/*ListEmployeesParams contains all the parameters to send to the API endpoint
for the list employees operation typically these are written to a http.Request
*/
type ListEmployeesParams struct {

	/*Cursor
	  The token required to retrieve the specified page of results.

	*/
	Cursor *string
	/*Limit
	  The number of employees to be returned on each page.

	*/
	Limit *int64
	/*LocationID
	  Filter employees returned to only those that are associated with the specified location.

	*/
	LocationID *string
	/*Status
	  Specifies the EmployeeStatus to filter the employee by.

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list employees params
func (o *ListEmployeesParams) WithTimeout(timeout time.Duration) *ListEmployeesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list employees params
func (o *ListEmployeesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list employees params
func (o *ListEmployeesParams) WithContext(ctx context.Context) *ListEmployeesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list employees params
func (o *ListEmployeesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list employees params
func (o *ListEmployeesParams) WithHTTPClient(client *http.Client) *ListEmployeesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list employees params
func (o *ListEmployeesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list employees params
func (o *ListEmployeesParams) WithCursor(cursor *string) *ListEmployeesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list employees params
func (o *ListEmployeesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the list employees params
func (o *ListEmployeesParams) WithLimit(limit *int64) *ListEmployeesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list employees params
func (o *ListEmployeesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the list employees params
func (o *ListEmployeesParams) WithLocationID(locationID *string) *ListEmployeesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list employees params
func (o *ListEmployeesParams) SetLocationID(locationID *string) {
	o.LocationID = locationID
}

// WithStatus adds the status to the list employees params
func (o *ListEmployeesParams) WithStatus(status *string) *ListEmployeesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list employees params
func (o *ListEmployeesParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListEmployeesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}