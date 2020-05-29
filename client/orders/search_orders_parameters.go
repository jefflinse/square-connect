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

// NewSearchOrdersParams creates a new SearchOrdersParams object
// with the default values initialized.
func NewSearchOrdersParams() *SearchOrdersParams {
	var ()
	return &SearchOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchOrdersParamsWithTimeout creates a new SearchOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchOrdersParamsWithTimeout(timeout time.Duration) *SearchOrdersParams {
	var ()
	return &SearchOrdersParams{

		timeout: timeout,
	}
}

// NewSearchOrdersParamsWithContext creates a new SearchOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchOrdersParamsWithContext(ctx context.Context) *SearchOrdersParams {
	var ()
	return &SearchOrdersParams{

		Context: ctx,
	}
}

// NewSearchOrdersParamsWithHTTPClient creates a new SearchOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchOrdersParamsWithHTTPClient(client *http.Client) *SearchOrdersParams {
	var ()
	return &SearchOrdersParams{
		HTTPClient: client,
	}
}

/*SearchOrdersParams contains all the parameters to send to the API endpoint
for the search orders operation typically these are written to a http.Request
*/
type SearchOrdersParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.SearchOrdersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search orders params
func (o *SearchOrdersParams) WithTimeout(timeout time.Duration) *SearchOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search orders params
func (o *SearchOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search orders params
func (o *SearchOrdersParams) WithContext(ctx context.Context) *SearchOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search orders params
func (o *SearchOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search orders params
func (o *SearchOrdersParams) WithHTTPClient(client *http.Client) *SearchOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search orders params
func (o *SearchOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search orders params
func (o *SearchOrdersParams) WithBody(body *models.SearchOrdersRequest) *SearchOrdersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search orders params
func (o *SearchOrdersParams) SetBody(body *models.SearchOrdersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
