// Code generated by go-swagger; DO NOT EDIT.

package checkout

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

// NewCreateCheckoutParams creates a new CreateCheckoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCheckoutParams() *CreateCheckoutParams {
	return &CreateCheckoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCheckoutParamsWithTimeout creates a new CreateCheckoutParams object
// with the ability to set a timeout on a request.
func NewCreateCheckoutParamsWithTimeout(timeout time.Duration) *CreateCheckoutParams {
	return &CreateCheckoutParams{
		timeout: timeout,
	}
}

// NewCreateCheckoutParamsWithContext creates a new CreateCheckoutParams object
// with the ability to set a context for a request.
func NewCreateCheckoutParamsWithContext(ctx context.Context) *CreateCheckoutParams {
	return &CreateCheckoutParams{
		Context: ctx,
	}
}

// NewCreateCheckoutParamsWithHTTPClient creates a new CreateCheckoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCheckoutParamsWithHTTPClient(client *http.Client) *CreateCheckoutParams {
	return &CreateCheckoutParams{
		HTTPClient: client,
	}
}

/* CreateCheckoutParams contains all the parameters to send to the API endpoint
   for the create checkout operation.

   Typically these are written to a http.Request.
*/
type CreateCheckoutParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.CreateCheckoutRequest

	/* LocationID.

	   The ID of the business location to associate the checkout with.
	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create checkout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCheckoutParams) WithDefaults() *CreateCheckoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create checkout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCheckoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create checkout params
func (o *CreateCheckoutParams) WithTimeout(timeout time.Duration) *CreateCheckoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create checkout params
func (o *CreateCheckoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create checkout params
func (o *CreateCheckoutParams) WithContext(ctx context.Context) *CreateCheckoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create checkout params
func (o *CreateCheckoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create checkout params
func (o *CreateCheckoutParams) WithHTTPClient(client *http.Client) *CreateCheckoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create checkout params
func (o *CreateCheckoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create checkout params
func (o *CreateCheckoutParams) WithBody(body *models.CreateCheckoutRequest) *CreateCheckoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create checkout params
func (o *CreateCheckoutParams) SetBody(body *models.CreateCheckoutRequest) {
	o.Body = body
}

// WithLocationID adds the locationID to the create checkout params
func (o *CreateCheckoutParams) WithLocationID(locationID string) *CreateCheckoutParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the create checkout params
func (o *CreateCheckoutParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCheckoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
