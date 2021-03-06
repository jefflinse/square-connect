// Code generated by go-swagger; DO NOT EDIT.

package bookings

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

// NewCancelBookingParams creates a new CancelBookingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelBookingParams() *CancelBookingParams {
	return &CancelBookingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelBookingParamsWithTimeout creates a new CancelBookingParams object
// with the ability to set a timeout on a request.
func NewCancelBookingParamsWithTimeout(timeout time.Duration) *CancelBookingParams {
	return &CancelBookingParams{
		timeout: timeout,
	}
}

// NewCancelBookingParamsWithContext creates a new CancelBookingParams object
// with the ability to set a context for a request.
func NewCancelBookingParamsWithContext(ctx context.Context) *CancelBookingParams {
	return &CancelBookingParams{
		Context: ctx,
	}
}

// NewCancelBookingParamsWithHTTPClient creates a new CancelBookingParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelBookingParamsWithHTTPClient(client *http.Client) *CancelBookingParams {
	return &CancelBookingParams{
		HTTPClient: client,
	}
}

/* CancelBookingParams contains all the parameters to send to the API endpoint
   for the cancel booking operation.

   Typically these are written to a http.Request.
*/
type CancelBookingParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.CancelBookingRequest

	/* BookingID.

	   The ID of the `Booking` object representing the to-be-cancelled booking.
	*/
	BookingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel booking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelBookingParams) WithDefaults() *CancelBookingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel booking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelBookingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel booking params
func (o *CancelBookingParams) WithTimeout(timeout time.Duration) *CancelBookingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel booking params
func (o *CancelBookingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel booking params
func (o *CancelBookingParams) WithContext(ctx context.Context) *CancelBookingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel booking params
func (o *CancelBookingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel booking params
func (o *CancelBookingParams) WithHTTPClient(client *http.Client) *CancelBookingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel booking params
func (o *CancelBookingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cancel booking params
func (o *CancelBookingParams) WithBody(body *models.CancelBookingRequest) *CancelBookingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cancel booking params
func (o *CancelBookingParams) SetBody(body *models.CancelBookingRequest) {
	o.Body = body
}

// WithBookingID adds the bookingID to the cancel booking params
func (o *CancelBookingParams) WithBookingID(bookingID string) *CancelBookingParams {
	o.SetBookingID(bookingID)
	return o
}

// SetBookingID adds the bookingId to the cancel booking params
func (o *CancelBookingParams) SetBookingID(bookingID string) {
	o.BookingID = bookingID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelBookingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param booking_id
	if err := r.SetPathParam("booking_id", o.BookingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
