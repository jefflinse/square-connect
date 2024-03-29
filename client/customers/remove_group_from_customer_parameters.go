// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewRemoveGroupFromCustomerParams creates a new RemoveGroupFromCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveGroupFromCustomerParams() *RemoveGroupFromCustomerParams {
	return &RemoveGroupFromCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveGroupFromCustomerParamsWithTimeout creates a new RemoveGroupFromCustomerParams object
// with the ability to set a timeout on a request.
func NewRemoveGroupFromCustomerParamsWithTimeout(timeout time.Duration) *RemoveGroupFromCustomerParams {
	return &RemoveGroupFromCustomerParams{
		timeout: timeout,
	}
}

// NewRemoveGroupFromCustomerParamsWithContext creates a new RemoveGroupFromCustomerParams object
// with the ability to set a context for a request.
func NewRemoveGroupFromCustomerParamsWithContext(ctx context.Context) *RemoveGroupFromCustomerParams {
	return &RemoveGroupFromCustomerParams{
		Context: ctx,
	}
}

// NewRemoveGroupFromCustomerParamsWithHTTPClient creates a new RemoveGroupFromCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveGroupFromCustomerParamsWithHTTPClient(client *http.Client) *RemoveGroupFromCustomerParams {
	return &RemoveGroupFromCustomerParams{
		HTTPClient: client,
	}
}

/* RemoveGroupFromCustomerParams contains all the parameters to send to the API endpoint
   for the remove group from customer operation.

   Typically these are written to a http.Request.
*/
type RemoveGroupFromCustomerParams struct {

	/* CustomerID.

	   The ID of the customer to remove from the group.
	*/
	CustomerID string

	/* GroupID.

	   The ID of the customer group to remove the customer from.
	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove group from customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveGroupFromCustomerParams) WithDefaults() *RemoveGroupFromCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove group from customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveGroupFromCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) WithTimeout(timeout time.Duration) *RemoveGroupFromCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) WithContext(ctx context.Context) *RemoveGroupFromCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) WithHTTPClient(client *http.Client) *RemoveGroupFromCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) WithCustomerID(customerID string) *RemoveGroupFromCustomerParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithGroupID adds the groupID to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) WithGroupID(groupID string) *RemoveGroupFromCustomerParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the remove group from customer params
func (o *RemoveGroupFromCustomerParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveGroupFromCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
		return err
	}

	// path param group_id
	if err := r.SetPathParam("group_id", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
