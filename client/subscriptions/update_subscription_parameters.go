// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewUpdateSubscriptionParams creates a new UpdateSubscriptionParams object
// with the default values initialized.
func NewUpdateSubscriptionParams() *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSubscriptionParamsWithTimeout creates a new UpdateSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSubscriptionParamsWithTimeout(timeout time.Duration) *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{

		timeout: timeout,
	}
}

// NewUpdateSubscriptionParamsWithContext creates a new UpdateSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSubscriptionParamsWithContext(ctx context.Context) *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{

		Context: ctx,
	}
}

// NewUpdateSubscriptionParamsWithHTTPClient creates a new UpdateSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSubscriptionParamsWithHTTPClient(client *http.Client) *UpdateSubscriptionParams {
	var ()
	return &UpdateSubscriptionParams{
		HTTPClient: client,
	}
}

/*UpdateSubscriptionParams contains all the parameters to send to the API endpoint
for the update subscription operation typically these are written to a http.Request
*/
type UpdateSubscriptionParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.UpdateSubscriptionRequest
	/*SubscriptionID
	  The ID for the subscription to update.

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update subscription params
func (o *UpdateSubscriptionParams) WithTimeout(timeout time.Duration) *UpdateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update subscription params
func (o *UpdateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update subscription params
func (o *UpdateSubscriptionParams) WithContext(ctx context.Context) *UpdateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update subscription params
func (o *UpdateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update subscription params
func (o *UpdateSubscriptionParams) WithHTTPClient(client *http.Client) *UpdateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update subscription params
func (o *UpdateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update subscription params
func (o *UpdateSubscriptionParams) WithBody(body *models.UpdateSubscriptionRequest) *UpdateSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update subscription params
func (o *UpdateSubscriptionParams) SetBody(body *models.UpdateSubscriptionRequest) {
	o.Body = body
}

// WithSubscriptionID adds the subscriptionID to the update subscription params
func (o *UpdateSubscriptionParams) WithSubscriptionID(subscriptionID string) *UpdateSubscriptionParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the update subscription params
func (o *UpdateSubscriptionParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param subscription_id
	if err := r.SetPathParam("subscription_id", o.SubscriptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
