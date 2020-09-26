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

// NewRetrieveInventoryCountParams creates a new RetrieveInventoryCountParams object
// with the default values initialized.
func NewRetrieveInventoryCountParams() *RetrieveInventoryCountParams {
	var ()
	return &RetrieveInventoryCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveInventoryCountParamsWithTimeout creates a new RetrieveInventoryCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveInventoryCountParamsWithTimeout(timeout time.Duration) *RetrieveInventoryCountParams {
	var ()
	return &RetrieveInventoryCountParams{

		timeout: timeout,
	}
}

// NewRetrieveInventoryCountParamsWithContext creates a new RetrieveInventoryCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveInventoryCountParamsWithContext(ctx context.Context) *RetrieveInventoryCountParams {
	var ()
	return &RetrieveInventoryCountParams{

		Context: ctx,
	}
}

// NewRetrieveInventoryCountParamsWithHTTPClient creates a new RetrieveInventoryCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveInventoryCountParamsWithHTTPClient(client *http.Client) *RetrieveInventoryCountParams {
	var ()
	return &RetrieveInventoryCountParams{
		HTTPClient: client,
	}
}

/*RetrieveInventoryCountParams contains all the parameters to send to the API endpoint
for the retrieve inventory count operation typically these are written to a http.Request
*/
type RetrieveInventoryCountParams struct {

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

// WithTimeout adds the timeout to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) WithTimeout(timeout time.Duration) *RetrieveInventoryCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) WithContext(ctx context.Context) *RetrieveInventoryCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) WithHTTPClient(client *http.Client) *RetrieveInventoryCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogObjectID adds the catalogObjectID to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) WithCatalogObjectID(catalogObjectID string) *RetrieveInventoryCountParams {
	o.SetCatalogObjectID(catalogObjectID)
	return o
}

// SetCatalogObjectID adds the catalogObjectId to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) SetCatalogObjectID(catalogObjectID string) {
	o.CatalogObjectID = catalogObjectID
}

// WithCursor adds the cursor to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) WithCursor(cursor *string) *RetrieveInventoryCountParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLocationIds adds the locationIds to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) WithLocationIds(locationIds *string) *RetrieveInventoryCountParams {
	o.SetLocationIds(locationIds)
	return o
}

// SetLocationIds adds the locationIds to the retrieve inventory count params
func (o *RetrieveInventoryCountParams) SetLocationIds(locationIds *string) {
	o.LocationIds = locationIds
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveInventoryCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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