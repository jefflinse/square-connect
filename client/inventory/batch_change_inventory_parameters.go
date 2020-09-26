// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewBatchChangeInventoryParams creates a new BatchChangeInventoryParams object
// with the default values initialized.
func NewBatchChangeInventoryParams() *BatchChangeInventoryParams {
	var ()
	return &BatchChangeInventoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchChangeInventoryParamsWithTimeout creates a new BatchChangeInventoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchChangeInventoryParamsWithTimeout(timeout time.Duration) *BatchChangeInventoryParams {
	var ()
	return &BatchChangeInventoryParams{

		timeout: timeout,
	}
}

// NewBatchChangeInventoryParamsWithContext creates a new BatchChangeInventoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchChangeInventoryParamsWithContext(ctx context.Context) *BatchChangeInventoryParams {
	var ()
	return &BatchChangeInventoryParams{

		Context: ctx,
	}
}

// NewBatchChangeInventoryParamsWithHTTPClient creates a new BatchChangeInventoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchChangeInventoryParamsWithHTTPClient(client *http.Client) *BatchChangeInventoryParams {
	var ()
	return &BatchChangeInventoryParams{
		HTTPClient: client,
	}
}

/*BatchChangeInventoryParams contains all the parameters to send to the API endpoint
for the batch change inventory operation typically these are written to a http.Request
*/
type BatchChangeInventoryParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.BatchChangeInventoryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batch change inventory params
func (o *BatchChangeInventoryParams) WithTimeout(timeout time.Duration) *BatchChangeInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch change inventory params
func (o *BatchChangeInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch change inventory params
func (o *BatchChangeInventoryParams) WithContext(ctx context.Context) *BatchChangeInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch change inventory params
func (o *BatchChangeInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch change inventory params
func (o *BatchChangeInventoryParams) WithHTTPClient(client *http.Client) *BatchChangeInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch change inventory params
func (o *BatchChangeInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch change inventory params
func (o *BatchChangeInventoryParams) WithBody(body *models.BatchChangeInventoryRequest) *BatchChangeInventoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch change inventory params
func (o *BatchChangeInventoryParams) SetBody(body *models.BatchChangeInventoryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchChangeInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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