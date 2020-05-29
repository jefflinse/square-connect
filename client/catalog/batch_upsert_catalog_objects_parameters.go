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

// NewBatchUpsertCatalogObjectsParams creates a new BatchUpsertCatalogObjectsParams object
// with the default values initialized.
func NewBatchUpsertCatalogObjectsParams() *BatchUpsertCatalogObjectsParams {
	var ()
	return &BatchUpsertCatalogObjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchUpsertCatalogObjectsParamsWithTimeout creates a new BatchUpsertCatalogObjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchUpsertCatalogObjectsParamsWithTimeout(timeout time.Duration) *BatchUpsertCatalogObjectsParams {
	var ()
	return &BatchUpsertCatalogObjectsParams{

		timeout: timeout,
	}
}

// NewBatchUpsertCatalogObjectsParamsWithContext creates a new BatchUpsertCatalogObjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchUpsertCatalogObjectsParamsWithContext(ctx context.Context) *BatchUpsertCatalogObjectsParams {
	var ()
	return &BatchUpsertCatalogObjectsParams{

		Context: ctx,
	}
}

// NewBatchUpsertCatalogObjectsParamsWithHTTPClient creates a new BatchUpsertCatalogObjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchUpsertCatalogObjectsParamsWithHTTPClient(client *http.Client) *BatchUpsertCatalogObjectsParams {
	var ()
	return &BatchUpsertCatalogObjectsParams{
		HTTPClient: client,
	}
}

/*BatchUpsertCatalogObjectsParams contains all the parameters to send to the API endpoint
for the batch upsert catalog objects operation typically these are written to a http.Request
*/
type BatchUpsertCatalogObjectsParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.BatchUpsertCatalogObjectsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) WithTimeout(timeout time.Duration) *BatchUpsertCatalogObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) WithContext(ctx context.Context) *BatchUpsertCatalogObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) WithHTTPClient(client *http.Client) *BatchUpsertCatalogObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) WithBody(body *models.BatchUpsertCatalogObjectsRequest) *BatchUpsertCatalogObjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batch upsert catalog objects params
func (o *BatchUpsertCatalogObjectsParams) SetBody(body *models.BatchUpsertCatalogObjectsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchUpsertCatalogObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
