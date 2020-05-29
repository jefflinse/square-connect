// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewUpdateCustomerParams creates a new UpdateCustomerParams object
// with the default values initialized.
func NewUpdateCustomerParams() *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomerParamsWithTimeout creates a new UpdateCustomerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCustomerParamsWithTimeout(timeout time.Duration) *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{

		timeout: timeout,
	}
}

// NewUpdateCustomerParamsWithContext creates a new UpdateCustomerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCustomerParamsWithContext(ctx context.Context) *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{

		Context: ctx,
	}
}

// NewUpdateCustomerParamsWithHTTPClient creates a new UpdateCustomerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCustomerParamsWithHTTPClient(client *http.Client) *UpdateCustomerParams {
	var ()
	return &UpdateCustomerParams{
		HTTPClient: client,
	}
}

/*UpdateCustomerParams contains all the parameters to send to the API endpoint
for the update customer operation typically these are written to a http.Request
*/
type UpdateCustomerParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.UpdateCustomerRequest
	/*CustomerID
	  The ID of the customer to update.

	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update customer params
func (o *UpdateCustomerParams) WithTimeout(timeout time.Duration) *UpdateCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update customer params
func (o *UpdateCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update customer params
func (o *UpdateCustomerParams) WithContext(ctx context.Context) *UpdateCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update customer params
func (o *UpdateCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update customer params
func (o *UpdateCustomerParams) WithHTTPClient(client *http.Client) *UpdateCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update customer params
func (o *UpdateCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update customer params
func (o *UpdateCustomerParams) WithBody(body *models.UpdateCustomerRequest) *UpdateCustomerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update customer params
func (o *UpdateCustomerParams) SetBody(body *models.UpdateCustomerRequest) {
	o.Body = body
}

// WithCustomerID adds the customerID to the update customer params
func (o *UpdateCustomerParams) WithCustomerID(customerID string) *UpdateCustomerParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the update customer params
func (o *UpdateCustomerParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
