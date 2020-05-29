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
)

// NewListCatalogParams creates a new ListCatalogParams object
// with the default values initialized.
func NewListCatalogParams() *ListCatalogParams {
	var ()
	return &ListCatalogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCatalogParamsWithTimeout creates a new ListCatalogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCatalogParamsWithTimeout(timeout time.Duration) *ListCatalogParams {
	var ()
	return &ListCatalogParams{

		timeout: timeout,
	}
}

// NewListCatalogParamsWithContext creates a new ListCatalogParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCatalogParamsWithContext(ctx context.Context) *ListCatalogParams {
	var ()
	return &ListCatalogParams{

		Context: ctx,
	}
}

// NewListCatalogParamsWithHTTPClient creates a new ListCatalogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCatalogParamsWithHTTPClient(client *http.Client) *ListCatalogParams {
	var ()
	return &ListCatalogParams{
		HTTPClient: client,
	}
}

/*ListCatalogParams contains all the parameters to send to the API endpoint
for the list catalog operation typically these are written to a http.Request
*/
type ListCatalogParams struct {

	/*Cursor
	  The pagination cursor returned in the previous response. Leave unset for an initial request.
	See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information.

	*/
	Cursor *string
	/*Types
	  An optional case-insensitive, comma-separated list of object types to retrieve, for example
	`ITEM,ITEM_VARIATION,CATEGORY,IMAGE`.

	The legal values are taken from the CatalogObjectType enum:
	`ITEM`, `ITEM_VARIATION`, `CATEGORY`, `DISCOUNT`, `TAX`,
	`MODIFIER`, `MODIFIER_LIST`, or `IMAGE`.

	*/
	Types *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list catalog params
func (o *ListCatalogParams) WithTimeout(timeout time.Duration) *ListCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list catalog params
func (o *ListCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list catalog params
func (o *ListCatalogParams) WithContext(ctx context.Context) *ListCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list catalog params
func (o *ListCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list catalog params
func (o *ListCatalogParams) WithHTTPClient(client *http.Client) *ListCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list catalog params
func (o *ListCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list catalog params
func (o *ListCatalogParams) WithCursor(cursor *string) *ListCatalogParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list catalog params
func (o *ListCatalogParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithTypes adds the types to the list catalog params
func (o *ListCatalogParams) WithTypes(types *string) *ListCatalogParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the list catalog params
func (o *ListCatalogParams) SetTypes(types *string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *ListCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Types != nil {

		// query param types
		var qrTypes string
		if o.Types != nil {
			qrTypes = *o.Types
		}
		qTypes := qrTypes
		if qTypes != "" {
			if err := r.SetQueryParam("types", qTypes); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
