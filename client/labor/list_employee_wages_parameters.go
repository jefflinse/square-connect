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

// NewListEmployeeWagesParams creates a new ListEmployeeWagesParams object
// with the default values initialized.
func NewListEmployeeWagesParams() *ListEmployeeWagesParams {
	var ()
	return &ListEmployeeWagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEmployeeWagesParamsWithTimeout creates a new ListEmployeeWagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEmployeeWagesParamsWithTimeout(timeout time.Duration) *ListEmployeeWagesParams {
	var ()
	return &ListEmployeeWagesParams{

		timeout: timeout,
	}
}

// NewListEmployeeWagesParamsWithContext creates a new ListEmployeeWagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEmployeeWagesParamsWithContext(ctx context.Context) *ListEmployeeWagesParams {
	var ()
	return &ListEmployeeWagesParams{

		Context: ctx,
	}
}

// NewListEmployeeWagesParamsWithHTTPClient creates a new ListEmployeeWagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEmployeeWagesParamsWithHTTPClient(client *http.Client) *ListEmployeeWagesParams {
	var ()
	return &ListEmployeeWagesParams{
		HTTPClient: client,
	}
}

/*ListEmployeeWagesParams contains all the parameters to send to the API endpoint
for the list employee wages operation typically these are written to a http.Request
*/
type ListEmployeeWagesParams struct {

	/*Cursor
	  Pointer to the next page of Employee Wage results to fetch.

	*/
	Cursor *string
	/*EmployeeID
	  Filter wages returned to only those that are associated with the
	specified employee.

	*/
	EmployeeID *string
	/*Limit
	  Maximum number of Employee Wages to return per page. Can range between
	1 and 200. The default is the maximum at 200.

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list employee wages params
func (o *ListEmployeeWagesParams) WithTimeout(timeout time.Duration) *ListEmployeeWagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list employee wages params
func (o *ListEmployeeWagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list employee wages params
func (o *ListEmployeeWagesParams) WithContext(ctx context.Context) *ListEmployeeWagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list employee wages params
func (o *ListEmployeeWagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list employee wages params
func (o *ListEmployeeWagesParams) WithHTTPClient(client *http.Client) *ListEmployeeWagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list employee wages params
func (o *ListEmployeeWagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list employee wages params
func (o *ListEmployeeWagesParams) WithCursor(cursor *string) *ListEmployeeWagesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list employee wages params
func (o *ListEmployeeWagesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithEmployeeID adds the employeeID to the list employee wages params
func (o *ListEmployeeWagesParams) WithEmployeeID(employeeID *string) *ListEmployeeWagesParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the list employee wages params
func (o *ListEmployeeWagesParams) SetEmployeeID(employeeID *string) {
	o.EmployeeID = employeeID
}

// WithLimit adds the limit to the list employee wages params
func (o *ListEmployeeWagesParams) WithLimit(limit *int64) *ListEmployeeWagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list employee wages params
func (o *ListEmployeeWagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *ListEmployeeWagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EmployeeID != nil {

		// query param employee_id
		var qrEmployeeID string
		if o.EmployeeID != nil {
			qrEmployeeID = *o.EmployeeID
		}
		qEmployeeID := qrEmployeeID
		if qEmployeeID != "" {
			if err := r.SetQueryParam("employee_id", qEmployeeID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
