// Code generated by go-swagger; DO NOT EDIT.

package v1_transactions

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

// NewListOrdersParams creates a new ListOrdersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOrdersParams() *ListOrdersParams {
	return &ListOrdersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOrdersParamsWithTimeout creates a new ListOrdersParams object
// with the ability to set a timeout on a request.
func NewListOrdersParamsWithTimeout(timeout time.Duration) *ListOrdersParams {
	return &ListOrdersParams{
		timeout: timeout,
	}
}

// NewListOrdersParamsWithContext creates a new ListOrdersParams object
// with the ability to set a context for a request.
func NewListOrdersParamsWithContext(ctx context.Context) *ListOrdersParams {
	return &ListOrdersParams{
		Context: ctx,
	}
}

// NewListOrdersParamsWithHTTPClient creates a new ListOrdersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOrdersParamsWithHTTPClient(client *http.Client) *ListOrdersParams {
	return &ListOrdersParams{
		HTTPClient: client,
	}
}

/* ListOrdersParams contains all the parameters to send to the API endpoint
   for the list orders operation.

   Typically these are written to a http.Request.
*/
type ListOrdersParams struct {

	/* BatchToken.

	     A pagination cursor to retrieve the next set of results for your
	original query to the endpoint.
	*/
	BatchToken *string

	/* Limit.

	   The maximum number of payments to return in a single response. This value cannot exceed 200.
	*/
	Limit *int64

	/* LocationID.

	   The ID of the location to list online store orders for.
	*/
	LocationID string

	/* Order.

	   The order in which payments are listed in the response.
	*/
	Order *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrdersParams) WithDefaults() *ListOrdersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrdersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list orders params
func (o *ListOrdersParams) WithTimeout(timeout time.Duration) *ListOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list orders params
func (o *ListOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list orders params
func (o *ListOrdersParams) WithContext(ctx context.Context) *ListOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list orders params
func (o *ListOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list orders params
func (o *ListOrdersParams) WithHTTPClient(client *http.Client) *ListOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list orders params
func (o *ListOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatchToken adds the batchToken to the list orders params
func (o *ListOrdersParams) WithBatchToken(batchToken *string) *ListOrdersParams {
	o.SetBatchToken(batchToken)
	return o
}

// SetBatchToken adds the batchToken to the list orders params
func (o *ListOrdersParams) SetBatchToken(batchToken *string) {
	o.BatchToken = batchToken
}

// WithLimit adds the limit to the list orders params
func (o *ListOrdersParams) WithLimit(limit *int64) *ListOrdersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list orders params
func (o *ListOrdersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the list orders params
func (o *ListOrdersParams) WithLocationID(locationID string) *ListOrdersParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list orders params
func (o *ListOrdersParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithOrder adds the order to the list orders params
func (o *ListOrdersParams) WithOrder(order *string) *ListOrdersParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the list orders params
func (o *ListOrdersParams) SetOrder(order *string) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BatchToken != nil {

		// query param batch_token
		var qrBatchToken string

		if o.BatchToken != nil {
			qrBatchToken = *o.BatchToken
		}
		qBatchToken := qrBatchToken
		if qBatchToken != "" {

			if err := r.SetQueryParam("batch_token", qBatchToken); err != nil {
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

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	if o.Order != nil {

		// query param order
		var qrOrder string

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {

			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
