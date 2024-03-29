// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

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

// NewRetrieveCustomerGroupParams creates a new RetrieveCustomerGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveCustomerGroupParams() *RetrieveCustomerGroupParams {
	return &RetrieveCustomerGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveCustomerGroupParamsWithTimeout creates a new RetrieveCustomerGroupParams object
// with the ability to set a timeout on a request.
func NewRetrieveCustomerGroupParamsWithTimeout(timeout time.Duration) *RetrieveCustomerGroupParams {
	return &RetrieveCustomerGroupParams{
		timeout: timeout,
	}
}

// NewRetrieveCustomerGroupParamsWithContext creates a new RetrieveCustomerGroupParams object
// with the ability to set a context for a request.
func NewRetrieveCustomerGroupParamsWithContext(ctx context.Context) *RetrieveCustomerGroupParams {
	return &RetrieveCustomerGroupParams{
		Context: ctx,
	}
}

// NewRetrieveCustomerGroupParamsWithHTTPClient creates a new RetrieveCustomerGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveCustomerGroupParamsWithHTTPClient(client *http.Client) *RetrieveCustomerGroupParams {
	return &RetrieveCustomerGroupParams{
		HTTPClient: client,
	}
}

/* RetrieveCustomerGroupParams contains all the parameters to send to the API endpoint
   for the retrieve customer group operation.

   Typically these are written to a http.Request.
*/
type RetrieveCustomerGroupParams struct {

	/* GroupID.

	   The ID of the customer group to retrieve.
	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve customer group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveCustomerGroupParams) WithDefaults() *RetrieveCustomerGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve customer group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveCustomerGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) WithTimeout(timeout time.Duration) *RetrieveCustomerGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) WithContext(ctx context.Context) *RetrieveCustomerGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) WithHTTPClient(client *http.Client) *RetrieveCustomerGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) WithGroupID(groupID string) *RetrieveCustomerGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the retrieve customer group params
func (o *RetrieveCustomerGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveCustomerGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group_id
	if err := r.SetPathParam("group_id", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
