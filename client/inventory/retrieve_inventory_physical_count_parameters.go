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
)

// NewRetrieveInventoryPhysicalCountParams creates a new RetrieveInventoryPhysicalCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveInventoryPhysicalCountParams() *RetrieveInventoryPhysicalCountParams {
	return &RetrieveInventoryPhysicalCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveInventoryPhysicalCountParamsWithTimeout creates a new RetrieveInventoryPhysicalCountParams object
// with the ability to set a timeout on a request.
func NewRetrieveInventoryPhysicalCountParamsWithTimeout(timeout time.Duration) *RetrieveInventoryPhysicalCountParams {
	return &RetrieveInventoryPhysicalCountParams{
		timeout: timeout,
	}
}

// NewRetrieveInventoryPhysicalCountParamsWithContext creates a new RetrieveInventoryPhysicalCountParams object
// with the ability to set a context for a request.
func NewRetrieveInventoryPhysicalCountParamsWithContext(ctx context.Context) *RetrieveInventoryPhysicalCountParams {
	return &RetrieveInventoryPhysicalCountParams{
		Context: ctx,
	}
}

// NewRetrieveInventoryPhysicalCountParamsWithHTTPClient creates a new RetrieveInventoryPhysicalCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveInventoryPhysicalCountParamsWithHTTPClient(client *http.Client) *RetrieveInventoryPhysicalCountParams {
	return &RetrieveInventoryPhysicalCountParams{
		HTTPClient: client,
	}
}

/* RetrieveInventoryPhysicalCountParams contains all the parameters to send to the API endpoint
   for the retrieve inventory physical count operation.

   Typically these are written to a http.Request.
*/
type RetrieveInventoryPhysicalCountParams struct {

	/* PhysicalCountID.

	     ID of the
	`InventoryPhysicalCount` to retrieve.
	*/
	PhysicalCountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve inventory physical count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveInventoryPhysicalCountParams) WithDefaults() *RetrieveInventoryPhysicalCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve inventory physical count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveInventoryPhysicalCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) WithTimeout(timeout time.Duration) *RetrieveInventoryPhysicalCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) WithContext(ctx context.Context) *RetrieveInventoryPhysicalCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) WithHTTPClient(client *http.Client) *RetrieveInventoryPhysicalCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPhysicalCountID adds the physicalCountID to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) WithPhysicalCountID(physicalCountID string) *RetrieveInventoryPhysicalCountParams {
	o.SetPhysicalCountID(physicalCountID)
	return o
}

// SetPhysicalCountID adds the physicalCountId to the retrieve inventory physical count params
func (o *RetrieveInventoryPhysicalCountParams) SetPhysicalCountID(physicalCountID string) {
	o.PhysicalCountID = physicalCountID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveInventoryPhysicalCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param physical_count_id
	if err := r.SetPathParam("physical_count_id", o.PhysicalCountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
