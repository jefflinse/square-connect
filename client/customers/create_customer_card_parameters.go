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

// NewCreateCustomerCardParams creates a new CreateCustomerCardParams object
// with the default values initialized.
func NewCreateCustomerCardParams() *CreateCustomerCardParams {
	var ()
	return &CreateCustomerCardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomerCardParamsWithTimeout creates a new CreateCustomerCardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomerCardParamsWithTimeout(timeout time.Duration) *CreateCustomerCardParams {
	var ()
	return &CreateCustomerCardParams{

		timeout: timeout,
	}
}

// NewCreateCustomerCardParamsWithContext creates a new CreateCustomerCardParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomerCardParamsWithContext(ctx context.Context) *CreateCustomerCardParams {
	var ()
	return &CreateCustomerCardParams{

		Context: ctx,
	}
}

// NewCreateCustomerCardParamsWithHTTPClient creates a new CreateCustomerCardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomerCardParamsWithHTTPClient(client *http.Client) *CreateCustomerCardParams {
	var ()
	return &CreateCustomerCardParams{
		HTTPClient: client,
	}
}

/*CreateCustomerCardParams contains all the parameters to send to the API endpoint
for the create customer card operation typically these are written to a http.Request
*/
type CreateCustomerCardParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.CreateCustomerCardRequest
	/*CustomerID
	  The Square ID of the customer profile the card is linked to.

	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create customer card params
func (o *CreateCustomerCardParams) WithTimeout(timeout time.Duration) *CreateCustomerCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create customer card params
func (o *CreateCustomerCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create customer card params
func (o *CreateCustomerCardParams) WithContext(ctx context.Context) *CreateCustomerCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create customer card params
func (o *CreateCustomerCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create customer card params
func (o *CreateCustomerCardParams) WithHTTPClient(client *http.Client) *CreateCustomerCardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create customer card params
func (o *CreateCustomerCardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create customer card params
func (o *CreateCustomerCardParams) WithBody(body *models.CreateCustomerCardRequest) *CreateCustomerCardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create customer card params
func (o *CreateCustomerCardParams) SetBody(body *models.CreateCustomerCardRequest) {
	o.Body = body
}

// WithCustomerID adds the customerID to the create customer card params
func (o *CreateCustomerCardParams) WithCustomerID(customerID string) *CreateCustomerCardParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the create customer card params
func (o *CreateCustomerCardParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomerCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
