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

// NewListTimecardEventsParams creates a new ListTimecardEventsParams object
// with the default values initialized.
func NewListTimecardEventsParams() *ListTimecardEventsParams {
	var ()
	return &ListTimecardEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTimecardEventsParamsWithTimeout creates a new ListTimecardEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTimecardEventsParamsWithTimeout(timeout time.Duration) *ListTimecardEventsParams {
	var ()
	return &ListTimecardEventsParams{

		timeout: timeout,
	}
}

// NewListTimecardEventsParamsWithContext creates a new ListTimecardEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTimecardEventsParamsWithContext(ctx context.Context) *ListTimecardEventsParams {
	var ()
	return &ListTimecardEventsParams{

		Context: ctx,
	}
}

// NewListTimecardEventsParamsWithHTTPClient creates a new ListTimecardEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTimecardEventsParamsWithHTTPClient(client *http.Client) *ListTimecardEventsParams {
	var ()
	return &ListTimecardEventsParams{
		HTTPClient: client,
	}
}

/*ListTimecardEventsParams contains all the parameters to send to the API endpoint
for the list timecard events operation typically these are written to a http.Request
*/
type ListTimecardEventsParams struct {

	/*TimecardID
	  The ID of the timecard to list events for.

	*/
	TimecardID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list timecard events params
func (o *ListTimecardEventsParams) WithTimeout(timeout time.Duration) *ListTimecardEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list timecard events params
func (o *ListTimecardEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list timecard events params
func (o *ListTimecardEventsParams) WithContext(ctx context.Context) *ListTimecardEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list timecard events params
func (o *ListTimecardEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list timecard events params
func (o *ListTimecardEventsParams) WithHTTPClient(client *http.Client) *ListTimecardEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list timecard events params
func (o *ListTimecardEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimecardID adds the timecardID to the list timecard events params
func (o *ListTimecardEventsParams) WithTimecardID(timecardID string) *ListTimecardEventsParams {
	o.SetTimecardID(timecardID)
	return o
}

// SetTimecardID adds the timecardId to the list timecard events params
func (o *ListTimecardEventsParams) SetTimecardID(timecardID string) {
	o.TimecardID = timecardID
}

// WriteToRequest writes these params to a swagger request
func (o *ListTimecardEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param timecard_id
	if err := r.SetPathParam("timecard_id", o.TimecardID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
