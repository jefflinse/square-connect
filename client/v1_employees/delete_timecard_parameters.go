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

// NewDeleteTimecardParams creates a new DeleteTimecardParams object
// with the default values initialized.
func NewDeleteTimecardParams() *DeleteTimecardParams {
	var ()
	return &DeleteTimecardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTimecardParamsWithTimeout creates a new DeleteTimecardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTimecardParamsWithTimeout(timeout time.Duration) *DeleteTimecardParams {
	var ()
	return &DeleteTimecardParams{

		timeout: timeout,
	}
}

// NewDeleteTimecardParamsWithContext creates a new DeleteTimecardParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTimecardParamsWithContext(ctx context.Context) *DeleteTimecardParams {
	var ()
	return &DeleteTimecardParams{

		Context: ctx,
	}
}

// NewDeleteTimecardParamsWithHTTPClient creates a new DeleteTimecardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTimecardParamsWithHTTPClient(client *http.Client) *DeleteTimecardParams {
	var ()
	return &DeleteTimecardParams{
		HTTPClient: client,
	}
}

/*DeleteTimecardParams contains all the parameters to send to the API endpoint
for the delete timecard operation typically these are written to a http.Request
*/
type DeleteTimecardParams struct {

	/*TimecardID
	  The ID of the timecard to delete.

	*/
	TimecardID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete timecard params
func (o *DeleteTimecardParams) WithTimeout(timeout time.Duration) *DeleteTimecardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete timecard params
func (o *DeleteTimecardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete timecard params
func (o *DeleteTimecardParams) WithContext(ctx context.Context) *DeleteTimecardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete timecard params
func (o *DeleteTimecardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete timecard params
func (o *DeleteTimecardParams) WithHTTPClient(client *http.Client) *DeleteTimecardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete timecard params
func (o *DeleteTimecardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimecardID adds the timecardID to the delete timecard params
func (o *DeleteTimecardParams) WithTimecardID(timecardID string) *DeleteTimecardParams {
	o.SetTimecardID(timecardID)
	return o
}

// SetTimecardID adds the timecardId to the delete timecard params
func (o *DeleteTimecardParams) SetTimecardID(timecardID string) {
	o.TimecardID = timecardID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTimecardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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