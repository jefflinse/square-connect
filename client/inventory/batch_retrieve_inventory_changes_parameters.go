// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewBatchRetrieveInventoryChangesParams creates a new BatchRetrieveInventoryChangesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchRetrieveInventoryChangesParams() *BatchRetrieveInventoryChangesParams {
	return &BatchRetrieveInventoryChangesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBatchRetrieveInventoryChangesParamsWithTimeout creates a new BatchRetrieveInventoryChangesParams object
// with the ability to set a timeout on a request.
func NewBatchRetrieveInventoryChangesParamsWithTimeout(timeout time.Duration) *BatchRetrieveInventoryChangesParams {
	return &BatchRetrieveInventoryChangesParams{
		timeout: timeout,
	}
}

// NewBatchRetrieveInventoryChangesParamsWithContext creates a new BatchRetrieveInventoryChangesParams object
// with the ability to set a context for a request.
func NewBatchRetrieveInventoryChangesParamsWithContext(ctx context.Context) *BatchRetrieveInventoryChangesParams {
	return &BatchRetrieveInventoryChangesParams{
		Context: ctx,
	}
}

// NewBatchRetrieveInventoryChangesParamsWithHTTPClient creates a new BatchRetrieveInventoryChangesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchRetrieveInventoryChangesParamsWithHTTPClient(client *http.Client) *BatchRetrieveInventoryChangesParams {
	return &BatchRetrieveInventoryChangesParams{
		HTTPClient: client,
	}
}

/* BatchRetrieveInventoryChangesParams contains all the parameters to send to the API endpoint
   for the batch retrieve inventory changes operation.

   Typically these are written to a http.Request.
*/
type BatchRetrieveInventoryChangesParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.BatchRetrieveInventoryChangesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the batch retrieve inventory changes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchRetrieveInventoryChangesParams) WithDefaults() *BatchRetrieveInventoryChangesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch retrieve inventory changes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchRetrieveInventoryChangesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) WithTimeout(timeout time.Duration) *BatchRetrieveInventoryChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) WithContext(ctx context.Context) *BatchRetrieveInventoryChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) WithHTTPClient(client *http.Client) *BatchRetrieveInventoryChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) WithBody(body *models.BatchRetrieveInventoryChangesRequest) *BatchRetrieveInventoryChangesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch retrieve inventory changes params
func (o *BatchRetrieveInventoryChangesParams) SetBody(body *models.BatchRetrieveInventoryChangesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchRetrieveInventoryChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
