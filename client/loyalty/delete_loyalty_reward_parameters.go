// Code generated by go-swagger; DO NOT EDIT.

package loyalty

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

// NewDeleteLoyaltyRewardParams creates a new DeleteLoyaltyRewardParams object
// with the default values initialized.
func NewDeleteLoyaltyRewardParams() *DeleteLoyaltyRewardParams {
	var ()
	return &DeleteLoyaltyRewardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLoyaltyRewardParamsWithTimeout creates a new DeleteLoyaltyRewardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLoyaltyRewardParamsWithTimeout(timeout time.Duration) *DeleteLoyaltyRewardParams {
	var ()
	return &DeleteLoyaltyRewardParams{

		timeout: timeout,
	}
}

// NewDeleteLoyaltyRewardParamsWithContext creates a new DeleteLoyaltyRewardParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLoyaltyRewardParamsWithContext(ctx context.Context) *DeleteLoyaltyRewardParams {
	var ()
	return &DeleteLoyaltyRewardParams{

		Context: ctx,
	}
}

// NewDeleteLoyaltyRewardParamsWithHTTPClient creates a new DeleteLoyaltyRewardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLoyaltyRewardParamsWithHTTPClient(client *http.Client) *DeleteLoyaltyRewardParams {
	var ()
	return &DeleteLoyaltyRewardParams{
		HTTPClient: client,
	}
}

/*DeleteLoyaltyRewardParams contains all the parameters to send to the API endpoint
for the delete loyalty reward operation typically these are written to a http.Request
*/
type DeleteLoyaltyRewardParams struct {

	/*RewardID
	  The ID of the `loyalty reward` to delete.

	*/
	RewardID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) WithTimeout(timeout time.Duration) *DeleteLoyaltyRewardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) WithContext(ctx context.Context) *DeleteLoyaltyRewardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) WithHTTPClient(client *http.Client) *DeleteLoyaltyRewardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRewardID adds the rewardID to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) WithRewardID(rewardID string) *DeleteLoyaltyRewardParams {
	o.SetRewardID(rewardID)
	return o
}

// SetRewardID adds the rewardId to the delete loyalty reward params
func (o *DeleteLoyaltyRewardParams) SetRewardID(rewardID string) {
	o.RewardID = rewardID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLoyaltyRewardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param reward_id
	if err := r.SetPathParam("reward_id", o.RewardID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}