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

// NewDeleteFeeParams creates a new DeleteFeeParams object
// with the default values initialized.
func NewDeleteFeeParams() *DeleteFeeParams {
	var ()
	return &DeleteFeeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFeeParamsWithTimeout creates a new DeleteFeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFeeParamsWithTimeout(timeout time.Duration) *DeleteFeeParams {
	var ()
	return &DeleteFeeParams{

		timeout: timeout,
	}
}

// NewDeleteFeeParamsWithContext creates a new DeleteFeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFeeParamsWithContext(ctx context.Context) *DeleteFeeParams {
	var ()
	return &DeleteFeeParams{

		Context: ctx,
	}
}

// NewDeleteFeeParamsWithHTTPClient creates a new DeleteFeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFeeParamsWithHTTPClient(client *http.Client) *DeleteFeeParams {
	var ()
	return &DeleteFeeParams{
		HTTPClient: client,
	}
}

/*DeleteFeeParams contains all the parameters to send to the API endpoint
for the delete fee operation typically these are written to a http.Request
*/
type DeleteFeeParams struct {

	/*FeeID
	  The ID of the fee to delete.

	*/
	FeeID string
	/*LocationID
	  The ID of the fee's associated location.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete fee params
func (o *DeleteFeeParams) WithTimeout(timeout time.Duration) *DeleteFeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fee params
func (o *DeleteFeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fee params
func (o *DeleteFeeParams) WithContext(ctx context.Context) *DeleteFeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fee params
func (o *DeleteFeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fee params
func (o *DeleteFeeParams) WithHTTPClient(client *http.Client) *DeleteFeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fee params
func (o *DeleteFeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeeID adds the feeID to the delete fee params
func (o *DeleteFeeParams) WithFeeID(feeID string) *DeleteFeeParams {
	o.SetFeeID(feeID)
	return o
}

// SetFeeID adds the feeId to the delete fee params
func (o *DeleteFeeParams) SetFeeID(feeID string) {
	o.FeeID = feeID
}

// WithLocationID adds the locationID to the delete fee params
func (o *DeleteFeeParams) WithLocationID(locationID string) *DeleteFeeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the delete fee params
func (o *DeleteFeeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fee_id
	if err := r.SetPathParam("fee_id", o.FeeID); err != nil {
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
