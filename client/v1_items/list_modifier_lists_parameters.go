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

// NewListModifierListsParams creates a new ListModifierListsParams object
// with the default values initialized.
func NewListModifierListsParams() *ListModifierListsParams {
	var ()
	return &ListModifierListsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListModifierListsParamsWithTimeout creates a new ListModifierListsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListModifierListsParamsWithTimeout(timeout time.Duration) *ListModifierListsParams {
	var ()
	return &ListModifierListsParams{

		timeout: timeout,
	}
}

// NewListModifierListsParamsWithContext creates a new ListModifierListsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListModifierListsParamsWithContext(ctx context.Context) *ListModifierListsParams {
	var ()
	return &ListModifierListsParams{

		Context: ctx,
	}
}

// NewListModifierListsParamsWithHTTPClient creates a new ListModifierListsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListModifierListsParamsWithHTTPClient(client *http.Client) *ListModifierListsParams {
	var ()
	return &ListModifierListsParams{
		HTTPClient: client,
	}
}

/*ListModifierListsParams contains all the parameters to send to the API endpoint
for the list modifier lists operation typically these are written to a http.Request
*/
type ListModifierListsParams struct {

	/*LocationID
	  The ID of the location to list modifier lists for.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list modifier lists params
func (o *ListModifierListsParams) WithTimeout(timeout time.Duration) *ListModifierListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list modifier lists params
func (o *ListModifierListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list modifier lists params
func (o *ListModifierListsParams) WithContext(ctx context.Context) *ListModifierListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list modifier lists params
func (o *ListModifierListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list modifier lists params
func (o *ListModifierListsParams) WithHTTPClient(client *http.Client) *ListModifierListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list modifier lists params
func (o *ListModifierListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationID adds the locationID to the list modifier lists params
func (o *ListModifierListsParams) WithLocationID(locationID string) *ListModifierListsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list modifier lists params
func (o *ListModifierListsParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ListModifierListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}