// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewUpdateInvoiceParams creates a new UpdateInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateInvoiceParams() *UpdateInvoiceParams {
	return &UpdateInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInvoiceParamsWithTimeout creates a new UpdateInvoiceParams object
// with the ability to set a timeout on a request.
func NewUpdateInvoiceParamsWithTimeout(timeout time.Duration) *UpdateInvoiceParams {
	return &UpdateInvoiceParams{
		timeout: timeout,
	}
}

// NewUpdateInvoiceParamsWithContext creates a new UpdateInvoiceParams object
// with the ability to set a context for a request.
func NewUpdateInvoiceParamsWithContext(ctx context.Context) *UpdateInvoiceParams {
	return &UpdateInvoiceParams{
		Context: ctx,
	}
}

// NewUpdateInvoiceParamsWithHTTPClient creates a new UpdateInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateInvoiceParamsWithHTTPClient(client *http.Client) *UpdateInvoiceParams {
	return &UpdateInvoiceParams{
		HTTPClient: client,
	}
}

/* UpdateInvoiceParams contains all the parameters to send to the API endpoint
   for the update invoice operation.

   Typically these are written to a http.Request.
*/
type UpdateInvoiceParams struct {

	/* Body.

	     An object containing the fields to POST for the request.

	See the corresponding object definition for field details.
	*/
	Body *models.UpdateInvoiceRequest

	/* InvoiceID.

	   The ID of the invoice to update.
	*/
	InvoiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInvoiceParams) WithDefaults() *UpdateInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update invoice params
func (o *UpdateInvoiceParams) WithTimeout(timeout time.Duration) *UpdateInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update invoice params
func (o *UpdateInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update invoice params
func (o *UpdateInvoiceParams) WithContext(ctx context.Context) *UpdateInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update invoice params
func (o *UpdateInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update invoice params
func (o *UpdateInvoiceParams) WithHTTPClient(client *http.Client) *UpdateInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update invoice params
func (o *UpdateInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update invoice params
func (o *UpdateInvoiceParams) WithBody(body *models.UpdateInvoiceRequest) *UpdateInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update invoice params
func (o *UpdateInvoiceParams) SetBody(body *models.UpdateInvoiceRequest) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the update invoice params
func (o *UpdateInvoiceParams) WithInvoiceID(invoiceID string) *UpdateInvoiceParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the update invoice params
func (o *UpdateInvoiceParams) SetInvoiceID(invoiceID string) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param invoice_id
	if err := r.SetPathParam("invoice_id", o.InvoiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
