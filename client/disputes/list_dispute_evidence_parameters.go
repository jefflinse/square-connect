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

// NewListDisputeEvidenceParams creates a new ListDisputeEvidenceParams object
// with the default values initialized.
func NewListDisputeEvidenceParams() *ListDisputeEvidenceParams {
	var ()
	return &ListDisputeEvidenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDisputeEvidenceParamsWithTimeout creates a new ListDisputeEvidenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDisputeEvidenceParamsWithTimeout(timeout time.Duration) *ListDisputeEvidenceParams {
	var ()
	return &ListDisputeEvidenceParams{

		timeout: timeout,
	}
}

// NewListDisputeEvidenceParamsWithContext creates a new ListDisputeEvidenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDisputeEvidenceParamsWithContext(ctx context.Context) *ListDisputeEvidenceParams {
	var ()
	return &ListDisputeEvidenceParams{

		Context: ctx,
	}
}

// NewListDisputeEvidenceParamsWithHTTPClient creates a new ListDisputeEvidenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDisputeEvidenceParamsWithHTTPClient(client *http.Client) *ListDisputeEvidenceParams {
	var ()
	return &ListDisputeEvidenceParams{
		HTTPClient: client,
	}
}

/*ListDisputeEvidenceParams contains all the parameters to send to the API endpoint
for the list dispute evidence operation typically these are written to a http.Request
*/
type ListDisputeEvidenceParams struct {

	/*DisputeID
	  The ID of the dispute.

	*/
	DisputeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list dispute evidence params
func (o *ListDisputeEvidenceParams) WithTimeout(timeout time.Duration) *ListDisputeEvidenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dispute evidence params
func (o *ListDisputeEvidenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dispute evidence params
func (o *ListDisputeEvidenceParams) WithContext(ctx context.Context) *ListDisputeEvidenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dispute evidence params
func (o *ListDisputeEvidenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dispute evidence params
func (o *ListDisputeEvidenceParams) WithHTTPClient(client *http.Client) *ListDisputeEvidenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dispute evidence params
func (o *ListDisputeEvidenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisputeID adds the disputeID to the list dispute evidence params
func (o *ListDisputeEvidenceParams) WithDisputeID(disputeID string) *ListDisputeEvidenceParams {
	o.SetDisputeID(disputeID)
	return o
}

// SetDisputeID adds the disputeId to the list dispute evidence params
func (o *ListDisputeEvidenceParams) SetDisputeID(disputeID string) {
	o.DisputeID = disputeID
}

// WriteToRequest writes these params to a swagger request
func (o *ListDisputeEvidenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
