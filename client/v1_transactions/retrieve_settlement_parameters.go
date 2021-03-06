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
)

// NewRetrieveSettlementParams creates a new RetrieveSettlementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveSettlementParams() *RetrieveSettlementParams {
	return &RetrieveSettlementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveSettlementParamsWithTimeout creates a new RetrieveSettlementParams object
// with the ability to set a timeout on a request.
func NewRetrieveSettlementParamsWithTimeout(timeout time.Duration) *RetrieveSettlementParams {
	return &RetrieveSettlementParams{
		timeout: timeout,
	}
}

// NewRetrieveSettlementParamsWithContext creates a new RetrieveSettlementParams object
// with the ability to set a context for a request.
func NewRetrieveSettlementParamsWithContext(ctx context.Context) *RetrieveSettlementParams {
	return &RetrieveSettlementParams{
		Context: ctx,
	}
}

// NewRetrieveSettlementParamsWithHTTPClient creates a new RetrieveSettlementParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveSettlementParamsWithHTTPClient(client *http.Client) *RetrieveSettlementParams {
	return &RetrieveSettlementParams{
		HTTPClient: client,
	}
}

/* RetrieveSettlementParams contains all the parameters to send to the API endpoint
   for the retrieve settlement operation.

   Typically these are written to a http.Request.
*/
type RetrieveSettlementParams struct {

	/* LocationID.

	   The ID of the settlements's associated location.
	*/
	LocationID string

	/* SettlementID.

	   The settlement's Square-issued ID. You obtain this value from Settlement objects returned by the List Settlements endpoint.
	*/
	SettlementID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve settlement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveSettlementParams) WithDefaults() *RetrieveSettlementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve settlement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveSettlementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve settlement params
func (o *RetrieveSettlementParams) WithTimeout(timeout time.Duration) *RetrieveSettlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve settlement params
func (o *RetrieveSettlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve settlement params
func (o *RetrieveSettlementParams) WithContext(ctx context.Context) *RetrieveSettlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve settlement params
func (o *RetrieveSettlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve settlement params
func (o *RetrieveSettlementParams) WithHTTPClient(client *http.Client) *RetrieveSettlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve settlement params
func (o *RetrieveSettlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationID adds the locationID to the retrieve settlement params
func (o *RetrieveSettlementParams) WithLocationID(locationID string) *RetrieveSettlementParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the retrieve settlement params
func (o *RetrieveSettlementParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithSettlementID adds the settlementID to the retrieve settlement params
func (o *RetrieveSettlementParams) WithSettlementID(settlementID string) *RetrieveSettlementParams {
	o.SetSettlementID(settlementID)
	return o
}

// SetSettlementID adds the settlementId to the retrieve settlement params
func (o *RetrieveSettlementParams) SetSettlementID(settlementID string) {
	o.SettlementID = settlementID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveSettlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	// path param settlement_id
	if err := r.SetPathParam("settlement_id", o.SettlementID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
