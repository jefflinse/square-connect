// Code generated by go-swagger; DO NOT EDIT.

package v1_items

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

// NewRetrieveItemParams creates a new RetrieveItemParams object
// with the default values initialized.
func NewRetrieveItemParams() *RetrieveItemParams {
	var ()
	return &RetrieveItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveItemParamsWithTimeout creates a new RetrieveItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveItemParamsWithTimeout(timeout time.Duration) *RetrieveItemParams {
	var ()
	return &RetrieveItemParams{

		timeout: timeout,
	}
}

// NewRetrieveItemParamsWithContext creates a new RetrieveItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveItemParamsWithContext(ctx context.Context) *RetrieveItemParams {
	var ()
	return &RetrieveItemParams{

		Context: ctx,
	}
}

// NewRetrieveItemParamsWithHTTPClient creates a new RetrieveItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveItemParamsWithHTTPClient(client *http.Client) *RetrieveItemParams {
	var ()
	return &RetrieveItemParams{
		HTTPClient: client,
	}
}

/*RetrieveItemParams contains all the parameters to send to the API endpoint
for the retrieve item operation typically these are written to a http.Request
*/
type RetrieveItemParams struct {

	/*ItemID
	  The item's ID.

	*/
	ItemID string
	/*LocationID
	  The ID of the item's associated location.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve item params
func (o *RetrieveItemParams) WithTimeout(timeout time.Duration) *RetrieveItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve item params
func (o *RetrieveItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve item params
func (o *RetrieveItemParams) WithContext(ctx context.Context) *RetrieveItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve item params
func (o *RetrieveItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve item params
func (o *RetrieveItemParams) WithHTTPClient(client *http.Client) *RetrieveItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve item params
func (o *RetrieveItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemID adds the itemID to the retrieve item params
func (o *RetrieveItemParams) WithItemID(itemID string) *RetrieveItemParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the retrieve item params
func (o *RetrieveItemParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithLocationID adds the locationID to the retrieve item params
func (o *RetrieveItemParams) WithLocationID(locationID string) *RetrieveItemParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the retrieve item params
func (o *RetrieveItemParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param item_id
	if err := r.SetPathParam("item_id", o.ItemID); err != nil {
		return err
	}

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
