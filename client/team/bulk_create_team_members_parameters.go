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

// NewBulkCreateTeamMembersParams creates a new BulkCreateTeamMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkCreateTeamMembersParams() *BulkCreateTeamMembersParams {
	return &BulkCreateTeamMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkCreateTeamMembersParamsWithTimeout creates a new BulkCreateTeamMembersParams object
// with the ability to set a timeout on a request.
func NewBulkCreateTeamMembersParamsWithTimeout(timeout time.Duration) *BulkCreateTeamMembersParams {
	return &BulkCreateTeamMembersParams{
		timeout: timeout,
	}
}

// NewBulkCreateTeamMembersParamsWithContext creates a new BulkCreateTeamMembersParams object
// with the ability to set a context for a request.
func NewBulkCreateTeamMembersParamsWithContext(ctx context.Context) *BulkCreateTeamMembersParams {
	return &BulkCreateTeamMembersParams{
		Context: ctx,
	}
}

// NewBulkCreateTeamMembersParamsWithHTTPClient creates a new BulkCreateTeamMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkCreateTeamMembersParamsWithHTTPClient(client *http.Client) *BulkCreateTeamMembersParams {
	return &BulkCreateTeamMembersParams{
		HTTPClient: client,
	}
}

/* BulkCreateTeamMembersParams contains all the parameters to send to the API endpoint
   for the bulk create team members operation.

   Typically these are written to a http.Request.
*/
type BulkCreateTeamMembersParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.BulkCreateTeamMembersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk create team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkCreateTeamMembersParams) WithDefaults() *BulkCreateTeamMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk create team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkCreateTeamMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk create team members params
func (o *BulkCreateTeamMembersParams) WithTimeout(timeout time.Duration) *BulkCreateTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk create team members params
func (o *BulkCreateTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk create team members params
func (o *BulkCreateTeamMembersParams) WithContext(ctx context.Context) *BulkCreateTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk create team members params
func (o *BulkCreateTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk create team members params
func (o *BulkCreateTeamMembersParams) WithHTTPClient(client *http.Client) *BulkCreateTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk create team members params
func (o *BulkCreateTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bulk create team members params
func (o *BulkCreateTeamMembersParams) WithBody(body *models.BulkCreateTeamMembersRequest) *BulkCreateTeamMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk create team members params
func (o *BulkCreateTeamMembersParams) SetBody(body *models.BulkCreateTeamMembersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BulkCreateTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
