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

// NewRetrieveEmployeeRoleParams creates a new RetrieveEmployeeRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveEmployeeRoleParams() *RetrieveEmployeeRoleParams {
	return &RetrieveEmployeeRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveEmployeeRoleParamsWithTimeout creates a new RetrieveEmployeeRoleParams object
// with the ability to set a timeout on a request.
func NewRetrieveEmployeeRoleParamsWithTimeout(timeout time.Duration) *RetrieveEmployeeRoleParams {
	return &RetrieveEmployeeRoleParams{
		timeout: timeout,
	}
}

// NewRetrieveEmployeeRoleParamsWithContext creates a new RetrieveEmployeeRoleParams object
// with the ability to set a context for a request.
func NewRetrieveEmployeeRoleParamsWithContext(ctx context.Context) *RetrieveEmployeeRoleParams {
	return &RetrieveEmployeeRoleParams{
		Context: ctx,
	}
}

// NewRetrieveEmployeeRoleParamsWithHTTPClient creates a new RetrieveEmployeeRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveEmployeeRoleParamsWithHTTPClient(client *http.Client) *RetrieveEmployeeRoleParams {
	return &RetrieveEmployeeRoleParams{
		HTTPClient: client,
	}
}

/* RetrieveEmployeeRoleParams contains all the parameters to send to the API endpoint
   for the retrieve employee role operation.

   Typically these are written to a http.Request.
*/
type RetrieveEmployeeRoleParams struct {

	/* RoleID.

	   The role's ID.
	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve employee role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveEmployeeRoleParams) WithDefaults() *RetrieveEmployeeRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve employee role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveEmployeeRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) WithTimeout(timeout time.Duration) *RetrieveEmployeeRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) WithContext(ctx context.Context) *RetrieveEmployeeRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) WithHTTPClient(client *http.Client) *RetrieveEmployeeRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) WithRoleID(roleID string) *RetrieveEmployeeRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the retrieve employee role params
func (o *RetrieveEmployeeRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveEmployeeRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
