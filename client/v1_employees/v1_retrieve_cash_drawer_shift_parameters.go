// Code generated by go-swagger; DO NOT EDIT.

package v1_employees

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

// NewV1RetrieveCashDrawerShiftParams creates a new V1RetrieveCashDrawerShiftParams object
// with the default values initialized.
func NewV1RetrieveCashDrawerShiftParams() *V1RetrieveCashDrawerShiftParams {
	var ()
	return &V1RetrieveCashDrawerShiftParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1RetrieveCashDrawerShiftParamsWithTimeout creates a new V1RetrieveCashDrawerShiftParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1RetrieveCashDrawerShiftParamsWithTimeout(timeout time.Duration) *V1RetrieveCashDrawerShiftParams {
	var ()
	return &V1RetrieveCashDrawerShiftParams{

		timeout: timeout,
	}
}

// NewV1RetrieveCashDrawerShiftParamsWithContext creates a new V1RetrieveCashDrawerShiftParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1RetrieveCashDrawerShiftParamsWithContext(ctx context.Context) *V1RetrieveCashDrawerShiftParams {
	var ()
	return &V1RetrieveCashDrawerShiftParams{

		Context: ctx,
	}
}

// NewV1RetrieveCashDrawerShiftParamsWithHTTPClient creates a new V1RetrieveCashDrawerShiftParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1RetrieveCashDrawerShiftParamsWithHTTPClient(client *http.Client) *V1RetrieveCashDrawerShiftParams {
	var ()
	return &V1RetrieveCashDrawerShiftParams{
		HTTPClient: client,
	}
}

/*V1RetrieveCashDrawerShiftParams contains all the parameters to send to the API endpoint
for the v1 retrieve cash drawer shift operation typically these are written to a http.Request
*/
type V1RetrieveCashDrawerShiftParams struct {

	/*LocationID
	  The ID of the location to list cash drawer shifts for.

	*/
	LocationID string
	/*ShiftID
	  The shift's ID.

	*/
	ShiftID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) WithTimeout(timeout time.Duration) *V1RetrieveCashDrawerShiftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) WithContext(ctx context.Context) *V1RetrieveCashDrawerShiftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) WithHTTPClient(client *http.Client) *V1RetrieveCashDrawerShiftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationID adds the locationID to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) WithLocationID(locationID string) *V1RetrieveCashDrawerShiftParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithShiftID adds the shiftID to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) WithShiftID(shiftID string) *V1RetrieveCashDrawerShiftParams {
	o.SetShiftID(shiftID)
	return o
}

// SetShiftID adds the shiftId to the v1 retrieve cash drawer shift params
func (o *V1RetrieveCashDrawerShiftParams) SetShiftID(shiftID string) {
	o.ShiftID = shiftID
}

// WriteToRequest writes these params to a swagger request
func (o *V1RetrieveCashDrawerShiftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	// path param shift_id
	if err := r.SetPathParam("shift_id", o.ShiftID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}