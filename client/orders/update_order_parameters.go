// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// NewUpdateOrderParams creates a new UpdateOrderParams object
// with the default values initialized.
func NewUpdateOrderParams() *UpdateOrderParams {
	var ()
	return &UpdateOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrderParamsWithTimeout creates a new UpdateOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOrderParamsWithTimeout(timeout time.Duration) *UpdateOrderParams {
	var ()
	return &UpdateOrderParams{

		timeout: timeout,
	}
}

// NewUpdateOrderParamsWithContext creates a new UpdateOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOrderParamsWithContext(ctx context.Context) *UpdateOrderParams {
	var ()
	return &UpdateOrderParams{

		Context: ctx,
	}
}

// NewUpdateOrderParamsWithHTTPClient creates a new UpdateOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOrderParamsWithHTTPClient(client *http.Client) *UpdateOrderParams {
	var ()
	return &UpdateOrderParams{
		HTTPClient: client,
	}
}

/*UpdateOrderParams contains all the parameters to send to the API endpoint
for the update order operation typically these are written to a http.Request
*/
type UpdateOrderParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.UpdateOrderRequest
	/*LocationID
	  The ID of the order's associated location.

	*/
	LocationID string
	/*OrderID
	  The ID of the order to update.

	*/
	OrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update order params
func (o *UpdateOrderParams) WithTimeout(timeout time.Duration) *UpdateOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update order params
func (o *UpdateOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update order params
func (o *UpdateOrderParams) WithContext(ctx context.Context) *UpdateOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update order params
func (o *UpdateOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update order params
func (o *UpdateOrderParams) WithHTTPClient(client *http.Client) *UpdateOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update order params
func (o *UpdateOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update order params
func (o *UpdateOrderParams) WithBody(body *models.UpdateOrderRequest) *UpdateOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update order params
func (o *UpdateOrderParams) SetBody(body *models.UpdateOrderRequest) {
	o.Body = body
}

// WithLocationID adds the locationID to the update order params
func (o *UpdateOrderParams) WithLocationID(locationID string) *UpdateOrderParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the update order params
func (o *UpdateOrderParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithOrderID adds the orderID to the update order params
func (o *UpdateOrderParams) WithOrderID(orderID string) *UpdateOrderParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the update order params
func (o *UpdateOrderParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
