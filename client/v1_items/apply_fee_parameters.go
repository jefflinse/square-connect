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

// NewApplyFeeParams creates a new ApplyFeeParams object
// with the default values initialized.
func NewApplyFeeParams() *ApplyFeeParams {
	var ()
	return &ApplyFeeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplyFeeParamsWithTimeout creates a new ApplyFeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplyFeeParamsWithTimeout(timeout time.Duration) *ApplyFeeParams {
	var ()
	return &ApplyFeeParams{

		timeout: timeout,
	}
}

// NewApplyFeeParamsWithContext creates a new ApplyFeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewApplyFeeParamsWithContext(ctx context.Context) *ApplyFeeParams {
	var ()
	return &ApplyFeeParams{

		Context: ctx,
	}
}

// NewApplyFeeParamsWithHTTPClient creates a new ApplyFeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApplyFeeParamsWithHTTPClient(client *http.Client) *ApplyFeeParams {
	var ()
	return &ApplyFeeParams{
		HTTPClient: client,
	}
}

/*ApplyFeeParams contains all the parameters to send to the API endpoint
for the apply fee operation typically these are written to a http.Request
*/
type ApplyFeeParams struct {

	/*FeeID
	  The ID of the fee to apply.

	*/
	FeeID string
	/*ItemID
	  The ID of the item to add the fee to.

	*/
	ItemID string
	/*LocationID
	  The ID of the fee's associated location.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the apply fee params
func (o *ApplyFeeParams) WithTimeout(timeout time.Duration) *ApplyFeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apply fee params
func (o *ApplyFeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apply fee params
func (o *ApplyFeeParams) WithContext(ctx context.Context) *ApplyFeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apply fee params
func (o *ApplyFeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apply fee params
func (o *ApplyFeeParams) WithHTTPClient(client *http.Client) *ApplyFeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apply fee params
func (o *ApplyFeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeeID adds the feeID to the apply fee params
func (o *ApplyFeeParams) WithFeeID(feeID string) *ApplyFeeParams {
	o.SetFeeID(feeID)
	return o
}

// SetFeeID adds the feeId to the apply fee params
func (o *ApplyFeeParams) SetFeeID(feeID string) {
	o.FeeID = feeID
}

// WithItemID adds the itemID to the apply fee params
func (o *ApplyFeeParams) WithItemID(itemID string) *ApplyFeeParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the apply fee params
func (o *ApplyFeeParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithLocationID adds the locationID to the apply fee params
func (o *ApplyFeeParams) WithLocationID(locationID string) *ApplyFeeParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the apply fee params
func (o *ApplyFeeParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ApplyFeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fee_id
	if err := r.SetPathParam("fee_id", o.FeeID); err != nil {
		return err
	}

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