// Code generated by go-swagger; DO NOT EDIT.

package v1_employees

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

// NewCreateTimecardParams creates a new CreateTimecardParams object
// with the default values initialized.
func NewCreateTimecardParams() *CreateTimecardParams {
	var ()
	return &CreateTimecardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTimecardParamsWithTimeout creates a new CreateTimecardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTimecardParamsWithTimeout(timeout time.Duration) *CreateTimecardParams {
	var ()
	return &CreateTimecardParams{

		timeout: timeout,
	}
}

// NewCreateTimecardParamsWithContext creates a new CreateTimecardParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTimecardParamsWithContext(ctx context.Context) *CreateTimecardParams {
	var ()
	return &CreateTimecardParams{

		Context: ctx,
	}
}

// NewCreateTimecardParamsWithHTTPClient creates a new CreateTimecardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTimecardParamsWithHTTPClient(client *http.Client) *CreateTimecardParams {
	var ()
	return &CreateTimecardParams{
		HTTPClient: client,
	}
}

/*CreateTimecardParams contains all the parameters to send to the API endpoint
for the create timecard operation typically these are written to a http.Request
*/
type CreateTimecardParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.V1Timecard

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create timecard params
func (o *CreateTimecardParams) WithTimeout(timeout time.Duration) *CreateTimecardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create timecard params
func (o *CreateTimecardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create timecard params
func (o *CreateTimecardParams) WithContext(ctx context.Context) *CreateTimecardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create timecard params
func (o *CreateTimecardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create timecard params
func (o *CreateTimecardParams) WithHTTPClient(client *http.Client) *CreateTimecardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create timecard params
func (o *CreateTimecardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create timecard params
func (o *CreateTimecardParams) WithBody(body *models.V1Timecard) *CreateTimecardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create timecard params
func (o *CreateTimecardParams) SetBody(body *models.V1Timecard) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTimecardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
