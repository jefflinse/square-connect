// Code generated by go-swagger; DO NOT EDIT.

package team

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

// NewBulkUpdateTeamMembersParams creates a new BulkUpdateTeamMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkUpdateTeamMembersParams() *BulkUpdateTeamMembersParams {
	return &BulkUpdateTeamMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkUpdateTeamMembersParamsWithTimeout creates a new BulkUpdateTeamMembersParams object
// with the ability to set a timeout on a request.
func NewBulkUpdateTeamMembersParamsWithTimeout(timeout time.Duration) *BulkUpdateTeamMembersParams {
	return &BulkUpdateTeamMembersParams{
		timeout: timeout,
	}
}

// NewBulkUpdateTeamMembersParamsWithContext creates a new BulkUpdateTeamMembersParams object
// with the ability to set a context for a request.
func NewBulkUpdateTeamMembersParamsWithContext(ctx context.Context) *BulkUpdateTeamMembersParams {
	return &BulkUpdateTeamMembersParams{
		Context: ctx,
	}
}

// NewBulkUpdateTeamMembersParamsWithHTTPClient creates a new BulkUpdateTeamMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkUpdateTeamMembersParamsWithHTTPClient(client *http.Client) *BulkUpdateTeamMembersParams {
	return &BulkUpdateTeamMembersParams{
		HTTPClient: client,
	}
}

/* BulkUpdateTeamMembersParams contains all the parameters to send to the API endpoint
   for the bulk update team members operation.

   Typically these are written to a http.Request.
*/
type BulkUpdateTeamMembersParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.BulkUpdateTeamMembersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk update team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkUpdateTeamMembersParams) WithDefaults() *BulkUpdateTeamMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk update team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkUpdateTeamMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) WithTimeout(timeout time.Duration) *BulkUpdateTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) WithContext(ctx context.Context) *BulkUpdateTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) WithHTTPClient(client *http.Client) *BulkUpdateTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) WithBody(body *models.BulkUpdateTeamMembersRequest) *BulkUpdateTeamMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk update team members params
func (o *BulkUpdateTeamMembersParams) SetBody(body *models.BulkUpdateTeamMembersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BulkUpdateTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
