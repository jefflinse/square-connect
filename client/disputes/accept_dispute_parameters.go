// Code generated by go-swagger; DO NOT EDIT.

package disputes

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

// NewAcceptDisputeParams creates a new AcceptDisputeParams object
// with the default values initialized.
func NewAcceptDisputeParams() *AcceptDisputeParams {
	var ()
	return &AcceptDisputeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptDisputeParamsWithTimeout creates a new AcceptDisputeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAcceptDisputeParamsWithTimeout(timeout time.Duration) *AcceptDisputeParams {
	var ()
	return &AcceptDisputeParams{

		timeout: timeout,
	}
}

// NewAcceptDisputeParamsWithContext creates a new AcceptDisputeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAcceptDisputeParamsWithContext(ctx context.Context) *AcceptDisputeParams {
	var ()
	return &AcceptDisputeParams{

		Context: ctx,
	}
}

// NewAcceptDisputeParamsWithHTTPClient creates a new AcceptDisputeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAcceptDisputeParamsWithHTTPClient(client *http.Client) *AcceptDisputeParams {
	var ()
	return &AcceptDisputeParams{
		HTTPClient: client,
	}
}

/*AcceptDisputeParams contains all the parameters to send to the API endpoint
for the accept dispute operation typically these are written to a http.Request
*/
type AcceptDisputeParams struct {

	/*DisputeID
	  ID of the dispute you want to accept.

	*/
	DisputeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the accept dispute params
func (o *AcceptDisputeParams) WithTimeout(timeout time.Duration) *AcceptDisputeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept dispute params
func (o *AcceptDisputeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept dispute params
func (o *AcceptDisputeParams) WithContext(ctx context.Context) *AcceptDisputeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept dispute params
func (o *AcceptDisputeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept dispute params
func (o *AcceptDisputeParams) WithHTTPClient(client *http.Client) *AcceptDisputeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept dispute params
func (o *AcceptDisputeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisputeID adds the disputeID to the accept dispute params
func (o *AcceptDisputeParams) WithDisputeID(disputeID string) *AcceptDisputeParams {
	o.SetDisputeID(disputeID)
	return o
}

// SetDisputeID adds the disputeId to the accept dispute params
func (o *AcceptDisputeParams) SetDisputeID(disputeID string) {
	o.DisputeID = disputeID
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptDisputeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dispute_id
	if err := r.SetPathParam("dispute_id", o.DisputeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
