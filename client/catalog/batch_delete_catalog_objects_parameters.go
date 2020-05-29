// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewBatchDeleteCatalogObjectsParams creates a new BatchDeleteCatalogObjectsParams object
// with the default values initialized.
func NewBatchDeleteCatalogObjectsParams() *BatchDeleteCatalogObjectsParams {
	var ()
	return &BatchDeleteCatalogObjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchDeleteCatalogObjectsParamsWithTimeout creates a new BatchDeleteCatalogObjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchDeleteCatalogObjectsParamsWithTimeout(timeout time.Duration) *BatchDeleteCatalogObjectsParams {
	var ()
	return &BatchDeleteCatalogObjectsParams{

		timeout: timeout,
	}
}

// NewBatchDeleteCatalogObjectsParamsWithContext creates a new BatchDeleteCatalogObjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchDeleteCatalogObjectsParamsWithContext(ctx context.Context) *BatchDeleteCatalogObjectsParams {
	var ()
	return &BatchDeleteCatalogObjectsParams{

		Context: ctx,
	}
}

// NewBatchDeleteCatalogObjectsParamsWithHTTPClient creates a new BatchDeleteCatalogObjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchDeleteCatalogObjectsParamsWithHTTPClient(client *http.Client) *BatchDeleteCatalogObjectsParams {
	var ()
	return &BatchDeleteCatalogObjectsParams{
		HTTPClient: client,
	}
}

/*BatchDeleteCatalogObjectsParams contains all the parameters to send to the API endpoint
for the batch delete catalog objects operation typically these are written to a http.Request
*/
type BatchDeleteCatalogObjectsParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.BatchDeleteCatalogObjectsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) WithTimeout(timeout time.Duration) *BatchDeleteCatalogObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) WithContext(ctx context.Context) *BatchDeleteCatalogObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) WithHTTPClient(client *http.Client) *BatchDeleteCatalogObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) WithBody(body *models.BatchDeleteCatalogObjectsRequest) *BatchDeleteCatalogObjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch delete catalog objects params
func (o *BatchDeleteCatalogObjectsParams) SetBody(body *models.BatchDeleteCatalogObjectsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchDeleteCatalogObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
