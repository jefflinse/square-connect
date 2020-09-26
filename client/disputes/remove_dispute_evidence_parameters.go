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

// NewRemoveDisputeEvidenceParams creates a new RemoveDisputeEvidenceParams object
// with the default values initialized.
func NewRemoveDisputeEvidenceParams() *RemoveDisputeEvidenceParams {
	var ()
	return &RemoveDisputeEvidenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveDisputeEvidenceParamsWithTimeout creates a new RemoveDisputeEvidenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveDisputeEvidenceParamsWithTimeout(timeout time.Duration) *RemoveDisputeEvidenceParams {
	var ()
	return &RemoveDisputeEvidenceParams{

		timeout: timeout,
	}
}

// NewRemoveDisputeEvidenceParamsWithContext creates a new RemoveDisputeEvidenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveDisputeEvidenceParamsWithContext(ctx context.Context) *RemoveDisputeEvidenceParams {
	var ()
	return &RemoveDisputeEvidenceParams{

		Context: ctx,
	}
}

// NewRemoveDisputeEvidenceParamsWithHTTPClient creates a new RemoveDisputeEvidenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveDisputeEvidenceParamsWithHTTPClient(client *http.Client) *RemoveDisputeEvidenceParams {
	var ()
	return &RemoveDisputeEvidenceParams{
		HTTPClient: client,
	}
}

/*RemoveDisputeEvidenceParams contains all the parameters to send to the API endpoint
for the remove dispute evidence operation typically these are written to a http.Request
*/
type RemoveDisputeEvidenceParams struct {

	/*DisputeID
	  The ID of the dispute you want to remove evidence from.

	*/
	DisputeID string
	/*EvidenceID
	  The ID of the evidence you want to remove.

	*/
	EvidenceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) WithTimeout(timeout time.Duration) *RemoveDisputeEvidenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) WithContext(ctx context.Context) *RemoveDisputeEvidenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) WithHTTPClient(client *http.Client) *RemoveDisputeEvidenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisputeID adds the disputeID to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) WithDisputeID(disputeID string) *RemoveDisputeEvidenceParams {
	o.SetDisputeID(disputeID)
	return o
}

// SetDisputeID adds the disputeId to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) SetDisputeID(disputeID string) {
	o.DisputeID = disputeID
}

// WithEvidenceID adds the evidenceID to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) WithEvidenceID(evidenceID string) *RemoveDisputeEvidenceParams {
	o.SetEvidenceID(evidenceID)
	return o
}

// SetEvidenceID adds the evidenceId to the remove dispute evidence params
func (o *RemoveDisputeEvidenceParams) SetEvidenceID(evidenceID string) {
	o.EvidenceID = evidenceID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveDisputeEvidenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dispute_id
	if err := r.SetPathParam("dispute_id", o.DisputeID); err != nil {
		return err
	}

	// path param evidence_id
	if err := r.SetPathParam("evidence_id", o.EvidenceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}