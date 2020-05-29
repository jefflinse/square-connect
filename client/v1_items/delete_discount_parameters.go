// Code generated by go-swagger; DO NOT EDIT.

package v1_items

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

// NewDeleteDiscountParams creates a new DeleteDiscountParams object
// with the default values initialized.
func NewDeleteDiscountParams() *DeleteDiscountParams {
	var ()
	return &DeleteDiscountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDiscountParamsWithTimeout creates a new DeleteDiscountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDiscountParamsWithTimeout(timeout time.Duration) *DeleteDiscountParams {
	var ()
	return &DeleteDiscountParams{

		timeout: timeout,
	}
}

// NewDeleteDiscountParamsWithContext creates a new DeleteDiscountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDiscountParamsWithContext(ctx context.Context) *DeleteDiscountParams {
	var ()
	return &DeleteDiscountParams{

		Context: ctx,
	}
}

// NewDeleteDiscountParamsWithHTTPClient creates a new DeleteDiscountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDiscountParamsWithHTTPClient(client *http.Client) *DeleteDiscountParams {
	var ()
	return &DeleteDiscountParams{
		HTTPClient: client,
	}
}

/*DeleteDiscountParams contains all the parameters to send to the API endpoint
for the delete discount operation typically these are written to a http.Request
*/
type DeleteDiscountParams struct {

	/*DiscountID
	  The ID of the discount to delete.

	*/
	DiscountID string
	/*LocationID
	  The ID of the item's associated location.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete discount params
func (o *DeleteDiscountParams) WithTimeout(timeout time.Duration) *DeleteDiscountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete discount params
func (o *DeleteDiscountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete discount params
func (o *DeleteDiscountParams) WithContext(ctx context.Context) *DeleteDiscountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete discount params
func (o *DeleteDiscountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete discount params
func (o *DeleteDiscountParams) WithHTTPClient(client *http.Client) *DeleteDiscountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete discount params
func (o *DeleteDiscountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiscountID adds the discountID to the delete discount params
func (o *DeleteDiscountParams) WithDiscountID(discountID string) *DeleteDiscountParams {
	o.SetDiscountID(discountID)
	return o
}

// SetDiscountID adds the discountId to the delete discount params
func (o *DeleteDiscountParams) SetDiscountID(discountID string) {
	o.DiscountID = discountID
}

// WithLocationID adds the locationID to the delete discount params
func (o *DeleteDiscountParams) WithLocationID(locationID string) *DeleteDiscountParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the delete discount params
func (o *DeleteDiscountParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDiscountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param discount_id
	if err := r.SetPathParam("discount_id", o.DiscountID); err != nil {
		return err
	}

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
