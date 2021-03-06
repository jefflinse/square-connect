// Code generated by go-swagger; DO NOT EDIT.

package v1_employees

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

// NewV1RetrieveEmployeeParams creates a new V1RetrieveEmployeeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1RetrieveEmployeeParams() *V1RetrieveEmployeeParams {
	return &V1RetrieveEmployeeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1RetrieveEmployeeParamsWithTimeout creates a new V1RetrieveEmployeeParams object
// with the ability to set a timeout on a request.
func NewV1RetrieveEmployeeParamsWithTimeout(timeout time.Duration) *V1RetrieveEmployeeParams {
	return &V1RetrieveEmployeeParams{
		timeout: timeout,
	}
}

// NewV1RetrieveEmployeeParamsWithContext creates a new V1RetrieveEmployeeParams object
// with the ability to set a context for a request.
func NewV1RetrieveEmployeeParamsWithContext(ctx context.Context) *V1RetrieveEmployeeParams {
	return &V1RetrieveEmployeeParams{
		Context: ctx,
	}
}

// NewV1RetrieveEmployeeParamsWithHTTPClient creates a new V1RetrieveEmployeeParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1RetrieveEmployeeParamsWithHTTPClient(client *http.Client) *V1RetrieveEmployeeParams {
	return &V1RetrieveEmployeeParams{
		HTTPClient: client,
	}
}

/* V1RetrieveEmployeeParams contains all the parameters to send to the API endpoint
   for the v1 retrieve employee operation.

   Typically these are written to a http.Request.
*/
type V1RetrieveEmployeeParams struct {

	/* EmployeeID.

	   The employee's ID.
	*/
	EmployeeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 retrieve employee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RetrieveEmployeeParams) WithDefaults() *V1RetrieveEmployeeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 retrieve employee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RetrieveEmployeeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) WithTimeout(timeout time.Duration) *V1RetrieveEmployeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) WithContext(ctx context.Context) *V1RetrieveEmployeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) WithHTTPClient(client *http.Client) *V1RetrieveEmployeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmployeeID adds the employeeID to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) WithEmployeeID(employeeID string) *V1RetrieveEmployeeParams {
	o.SetEmployeeID(employeeID)
	return o
}

// SetEmployeeID adds the employeeId to the v1 retrieve employee params
func (o *V1RetrieveEmployeeParams) SetEmployeeID(employeeID string) {
	o.EmployeeID = employeeID
}

// WriteToRequest writes these params to a swagger request
func (o *V1RetrieveEmployeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param employee_id
	if err := r.SetPathParam("employee_id", o.EmployeeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
