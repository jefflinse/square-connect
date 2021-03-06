// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewCreateDeviceCodeParams creates a new CreateDeviceCodeParams object
// with the default values initialized.
func NewCreateDeviceCodeParams() *CreateDeviceCodeParams {
	var ()
	return &CreateDeviceCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeviceCodeParamsWithTimeout creates a new CreateDeviceCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDeviceCodeParamsWithTimeout(timeout time.Duration) *CreateDeviceCodeParams {
	var ()
	return &CreateDeviceCodeParams{

		timeout: timeout,
	}
}

// NewCreateDeviceCodeParamsWithContext creates a new CreateDeviceCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDeviceCodeParamsWithContext(ctx context.Context) *CreateDeviceCodeParams {
	var ()
	return &CreateDeviceCodeParams{

		Context: ctx,
	}
}

// NewCreateDeviceCodeParamsWithHTTPClient creates a new CreateDeviceCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDeviceCodeParamsWithHTTPClient(client *http.Client) *CreateDeviceCodeParams {
	var ()
	return &CreateDeviceCodeParams{
		HTTPClient: client,
	}
}

/*CreateDeviceCodeParams contains all the parameters to send to the API endpoint
for the create device code operation typically these are written to a http.Request
*/
type CreateDeviceCodeParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.CreateDeviceCodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create device code params
func (o *CreateDeviceCodeParams) WithTimeout(timeout time.Duration) *CreateDeviceCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device code params
func (o *CreateDeviceCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device code params
func (o *CreateDeviceCodeParams) WithContext(ctx context.Context) *CreateDeviceCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device code params
func (o *CreateDeviceCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device code params
func (o *CreateDeviceCodeParams) WithHTTPClient(client *http.Client) *CreateDeviceCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device code params
func (o *CreateDeviceCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create device code params
func (o *CreateDeviceCodeParams) WithBody(body *models.CreateDeviceCodeRequest) *CreateDeviceCodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create device code params
func (o *CreateDeviceCodeParams) SetBody(body *models.CreateDeviceCodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeviceCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
