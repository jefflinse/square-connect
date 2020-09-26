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
)

// NewRetrieveInventoryChangesParams creates a new RetrieveInventoryChangesParams object
// with the default values initialized.
func NewRetrieveInventoryChangesParams() *RetrieveInventoryChangesParams {
	var ()
	return &RetrieveInventoryChangesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveInventoryChangesParamsWithTimeout creates a new RetrieveInventoryChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveInventoryChangesParamsWithTimeout(timeout time.Duration) *RetrieveInventoryChangesParams {
	var ()
	return &RetrieveInventoryChangesParams{

		timeout: timeout,
	}
}

// NewRetrieveInventoryChangesParamsWithContext creates a new RetrieveInventoryChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveInventoryChangesParamsWithContext(ctx context.Context) *RetrieveInventoryChangesParams {
	var ()
	return &RetrieveInventoryChangesParams{

		Context: ctx,
	}
}

// NewRetrieveInventoryChangesParamsWithHTTPClient creates a new RetrieveInventoryChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveInventoryChangesParamsWithHTTPClient(client *http.Client) *RetrieveInventoryChangesParams {
	var ()
	return &RetrieveInventoryChangesParams{
		HTTPClient: client,
	}
}

/*RetrieveInventoryChangesParams contains all the parameters to send to the API endpoint
for the retrieve inventory changes operation typically these are written to a http.Request
*/
type RetrieveInventoryChangesParams struct {

	/*CatalogObjectID
	  ID of the `CatalogObject` to retrieve.

	*/
	CatalogObjectID string
	/*Cursor
	  A pagination cursor returned by a previous call to this endpoint.
	Provide this to retrieve the next set of results for the original query.

	See the [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination) guide for more information.

	*/
	Cursor *string
	/*LocationIds
	  The `Location` IDs to look up as a comma-separated
	list. An empty list queries all locations.

	*/
	LocationIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) WithTimeout(timeout time.Duration) *RetrieveInventoryChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) WithContext(ctx context.Context) *RetrieveInventoryChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) WithHTTPClient(client *http.Client) *RetrieveInventoryChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogObjectID adds the catalogObjectID to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) WithCatalogObjectID(catalogObjectID string) *RetrieveInventoryChangesParams {
	o.SetCatalogObjectID(catalogObjectID)
	return o
}

// SetCatalogObjectID adds the catalogObjectId to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) SetCatalogObjectID(catalogObjectID string) {
	o.CatalogObjectID = catalogObjectID
}

// WithCursor adds the cursor to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) WithCursor(cursor *string) *RetrieveInventoryChangesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLocationIds adds the locationIds to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) WithLocationIds(locationIds *string) *RetrieveInventoryChangesParams {
	o.SetLocationIds(locationIds)
	return o
}

// SetLocationIds adds the locationIds to the retrieve inventory changes params
func (o *RetrieveInventoryChangesParams) SetLocationIds(locationIds *string) {
	o.LocationIds = locationIds
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveInventoryChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param catalog_object_id
	if err := r.SetPathParam("catalog_object_id", o.CatalogObjectID); err != nil {
		return err
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.LocationIds != nil {

		// query param location_ids
		var qrLocationIds string
		if o.LocationIds != nil {
			qrLocationIds = *o.LocationIds
		}
		qLocationIds := qrLocationIds
		if qLocationIds != "" {
			if err := r.SetQueryParam("location_ids", qLocationIds); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}