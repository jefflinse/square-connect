// Code generated by go-swagger; DO NOT EDIT.

package v1_transactions

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
	"github.com/go-openapi/swag"
)

// NewV1ListPaymentsParams creates a new V1ListPaymentsParams object
// with the default values initialized.
func NewV1ListPaymentsParams() *V1ListPaymentsParams {
	var ()
	return &V1ListPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ListPaymentsParamsWithTimeout creates a new V1ListPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ListPaymentsParamsWithTimeout(timeout time.Duration) *V1ListPaymentsParams {
	var ()
	return &V1ListPaymentsParams{

		timeout: timeout,
	}
}

// NewV1ListPaymentsParamsWithContext creates a new V1ListPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ListPaymentsParamsWithContext(ctx context.Context) *V1ListPaymentsParams {
	var ()
	return &V1ListPaymentsParams{

		Context: ctx,
	}
}

// NewV1ListPaymentsParamsWithHTTPClient creates a new V1ListPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ListPaymentsParamsWithHTTPClient(client *http.Client) *V1ListPaymentsParams {
	var ()
	return &V1ListPaymentsParams{
		HTTPClient: client,
	}
}

/*V1ListPaymentsParams contains all the parameters to send to the API endpoint
for the v1 list payments operation typically these are written to a http.Request
*/
type V1ListPaymentsParams struct {

	/*BatchToken
	  A pagination cursor to retrieve the next set of results for your
	original query to the endpoint.

	*/
	BatchToken *string
	/*BeginTime
	  The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year.

	*/
	BeginTime *string
	/*EndTime
	  The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time.

	*/
	EndTime *string
	/*IncludePartial
	  Indicates whether or not to include partial payments in the response. Partial payments will have the tenders collected so far, but the itemizations will be empty until the payment is completed.

	*/
	IncludePartial *bool
	/*Limit
	  The maximum number of payments to return in a single response. This value cannot exceed 200.

	*/
	Limit *int64
	/*LocationID
	  The ID of the location to list payments for. If you specify me, this endpoint returns payments aggregated from all of the business's locations.

	*/
	LocationID string
	/*Order
	  The order in which payments are listed in the response.

	*/
	Order *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 list payments params
func (o *V1ListPaymentsParams) WithTimeout(timeout time.Duration) *V1ListPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 list payments params
func (o *V1ListPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 list payments params
func (o *V1ListPaymentsParams) WithContext(ctx context.Context) *V1ListPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 list payments params
func (o *V1ListPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 list payments params
func (o *V1ListPaymentsParams) WithHTTPClient(client *http.Client) *V1ListPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 list payments params
func (o *V1ListPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatchToken adds the batchToken to the v1 list payments params
func (o *V1ListPaymentsParams) WithBatchToken(batchToken *string) *V1ListPaymentsParams {
	o.SetBatchToken(batchToken)
	return o
}

// SetBatchToken adds the batchToken to the v1 list payments params
func (o *V1ListPaymentsParams) SetBatchToken(batchToken *string) {
	o.BatchToken = batchToken
}

// WithBeginTime adds the beginTime to the v1 list payments params
func (o *V1ListPaymentsParams) WithBeginTime(beginTime *string) *V1ListPaymentsParams {
	o.SetBeginTime(beginTime)
	return o
}

// SetBeginTime adds the beginTime to the v1 list payments params
func (o *V1ListPaymentsParams) SetBeginTime(beginTime *string) {
	o.BeginTime = beginTime
}

// WithEndTime adds the endTime to the v1 list payments params
func (o *V1ListPaymentsParams) WithEndTime(endTime *string) *V1ListPaymentsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the v1 list payments params
func (o *V1ListPaymentsParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithIncludePartial adds the includePartial to the v1 list payments params
func (o *V1ListPaymentsParams) WithIncludePartial(includePartial *bool) *V1ListPaymentsParams {
	o.SetIncludePartial(includePartial)
	return o
}

// SetIncludePartial adds the includePartial to the v1 list payments params
func (o *V1ListPaymentsParams) SetIncludePartial(includePartial *bool) {
	o.IncludePartial = includePartial
}

// WithLimit adds the limit to the v1 list payments params
func (o *V1ListPaymentsParams) WithLimit(limit *int64) *V1ListPaymentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 list payments params
func (o *V1ListPaymentsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocationID adds the locationID to the v1 list payments params
func (o *V1ListPaymentsParams) WithLocationID(locationID string) *V1ListPaymentsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the v1 list payments params
func (o *V1ListPaymentsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithOrder adds the order to the v1 list payments params
func (o *V1ListPaymentsParams) WithOrder(order *string) *V1ListPaymentsParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the v1 list payments params
func (o *V1ListPaymentsParams) SetOrder(order *string) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *V1ListPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BatchToken != nil {

		// query param batch_token
		var qrBatchToken string
		if o.BatchToken != nil {
			qrBatchToken = *o.BatchToken
		}
		qBatchToken := qrBatchToken
		if qBatchToken != "" {
			if err := r.SetQueryParam("batch_token", qBatchToken); err != nil {
				return err
			}
		}

	}

	if o.BeginTime != nil {

		// query param begin_time
		var qrBeginTime string
		if o.BeginTime != nil {
			qrBeginTime = *o.BeginTime
		}
		qBeginTime := qrBeginTime
		if qBeginTime != "" {
			if err := r.SetQueryParam("begin_time", qBeginTime); err != nil {
				return err
			}
		}

	}

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime string
		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {
			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}

	}

	if o.IncludePartial != nil {

		// query param include_partial
		var qrIncludePartial bool
		if o.IncludePartial != nil {
			qrIncludePartial = *o.IncludePartial
		}
		qIncludePartial := swag.FormatBool(qrIncludePartial)
		if qIncludePartial != "" {
			if err := r.SetQueryParam("include_partial", qIncludePartial); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
