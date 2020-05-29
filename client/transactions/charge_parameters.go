// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewChargeParams creates a new ChargeParams object
// with the default values initialized.
func NewChargeParams() *ChargeParams {
	var ()
	return &ChargeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChargeParamsWithTimeout creates a new ChargeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChargeParamsWithTimeout(timeout time.Duration) *ChargeParams {
	var ()
	return &ChargeParams{

		timeout: timeout,
	}
}

// NewChargeParamsWithContext creates a new ChargeParams object
// with the default values initialized, and the ability to set a context for a request
func NewChargeParamsWithContext(ctx context.Context) *ChargeParams {
	var ()
	return &ChargeParams{

		Context: ctx,
	}
}

// NewChargeParamsWithHTTPClient creates a new ChargeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChargeParamsWithHTTPClient(client *http.Client) *ChargeParams {
	var ()
	return &ChargeParams{
		HTTPClient: client,
	}
}

/*ChargeParams contains all the parameters to send to the API endpoint
for the charge operation typically these are written to a http.Request
*/
type ChargeParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.ChargeRequest
	/*LocationID
	  The ID of the location to associate the created transaction with.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the charge params
func (o *ChargeParams) WithTimeout(timeout time.Duration) *ChargeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the charge params
func (o *ChargeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the charge params
func (o *ChargeParams) WithContext(ctx context.Context) *ChargeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the charge params
func (o *ChargeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the charge params
func (o *ChargeParams) WithHTTPClient(client *http.Client) *ChargeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the charge params
func (o *ChargeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the charge params
func (o *ChargeParams) WithBody(body *models.ChargeRequest) *ChargeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the charge params
func (o *ChargeParams) SetBody(body *models.ChargeRequest) {
	o.Body = body
}

// WithLocationID adds the locationID to the charge params
func (o *ChargeParams) WithLocationID(locationID string) *ChargeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the charge params
func (o *ChargeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ChargeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
