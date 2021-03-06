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
)

// NewRetrieveOrderParams creates a new RetrieveOrderParams object
// with the default values initialized.
func NewRetrieveOrderParams() *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveOrderParamsWithTimeout creates a new RetrieveOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveOrderParamsWithTimeout(timeout time.Duration) *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{

		timeout: timeout,
	}
}

// NewRetrieveOrderParamsWithContext creates a new RetrieveOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveOrderParamsWithContext(ctx context.Context) *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{

		Context: ctx,
	}
}

// NewRetrieveOrderParamsWithHTTPClient creates a new RetrieveOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveOrderParamsWithHTTPClient(client *http.Client) *RetrieveOrderParams {
	var ()
	return &RetrieveOrderParams{
		HTTPClient: client,
	}
}

/*RetrieveOrderParams contains all the parameters to send to the API endpoint
for the retrieve order operation typically these are written to a http.Request
*/
type RetrieveOrderParams struct {

	/*LocationID
	  The ID of the order's associated location.

	*/
	LocationID string
	/*OrderID
	  The order's Square-issued ID. You obtain this value from Order objects returned by the List Orders endpoint

	*/
	OrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve order params
func (o *RetrieveOrderParams) WithTimeout(timeout time.Duration) *RetrieveOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve order params
func (o *RetrieveOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve order params
func (o *RetrieveOrderParams) WithContext(ctx context.Context) *RetrieveOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve order params
func (o *RetrieveOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve order params
func (o *RetrieveOrderParams) WithHTTPClient(client *http.Client) *RetrieveOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve order params
func (o *RetrieveOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationID adds the locationID to the retrieve order params
func (o *RetrieveOrderParams) WithLocationID(locationID string) *RetrieveOrderParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the retrieve order params
func (o *RetrieveOrderParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithOrderID adds the orderID to the retrieve order params
func (o *RetrieveOrderParams) WithOrderID(orderID string) *RetrieveOrderParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the retrieve order params
func (o *RetrieveOrderParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	// path param order_id
	if err := r.SetPathParam("order_id", o.OrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
