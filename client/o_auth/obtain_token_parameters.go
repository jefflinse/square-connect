// Code generated by go-swagger; DO NOT EDIT.

package o_auth

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

// NewObtainTokenParams creates a new ObtainTokenParams object
// with the default values initialized.
func NewObtainTokenParams() *ObtainTokenParams {
	var ()
	return &ObtainTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObtainTokenParamsWithTimeout creates a new ObtainTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObtainTokenParamsWithTimeout(timeout time.Duration) *ObtainTokenParams {
	var ()
	return &ObtainTokenParams{

		timeout: timeout,
	}
}

// NewObtainTokenParamsWithContext creates a new ObtainTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewObtainTokenParamsWithContext(ctx context.Context) *ObtainTokenParams {
	var ()
	return &ObtainTokenParams{

		Context: ctx,
	}
}

// NewObtainTokenParamsWithHTTPClient creates a new ObtainTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObtainTokenParamsWithHTTPClient(client *http.Client) *ObtainTokenParams {
	var ()
	return &ObtainTokenParams{
		HTTPClient: client,
	}
}

/*ObtainTokenParams contains all the parameters to send to the API endpoint
for the obtain token operation typically these are written to a http.Request
*/
type ObtainTokenParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.ObtainTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the obtain token params
func (o *ObtainTokenParams) WithTimeout(timeout time.Duration) *ObtainTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the obtain token params
func (o *ObtainTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the obtain token params
func (o *ObtainTokenParams) WithContext(ctx context.Context) *ObtainTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the obtain token params
func (o *ObtainTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the obtain token params
func (o *ObtainTokenParams) WithHTTPClient(client *http.Client) *ObtainTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the obtain token params
func (o *ObtainTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the obtain token params
func (o *ObtainTokenParams) WithBody(body *models.ObtainTokenRequest) *ObtainTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the obtain token params
func (o *ObtainTokenParams) SetBody(body *models.ObtainTokenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ObtainTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
