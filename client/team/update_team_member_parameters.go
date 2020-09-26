// Code generated by go-swagger; DO NOT EDIT.

package team

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

// NewUpdateTeamMemberParams creates a new UpdateTeamMemberParams object
// with the default values initialized.
func NewUpdateTeamMemberParams() *UpdateTeamMemberParams {
	var ()
	return &UpdateTeamMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTeamMemberParamsWithTimeout creates a new UpdateTeamMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTeamMemberParamsWithTimeout(timeout time.Duration) *UpdateTeamMemberParams {
	var ()
	return &UpdateTeamMemberParams{

		timeout: timeout,
	}
}

// NewUpdateTeamMemberParamsWithContext creates a new UpdateTeamMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTeamMemberParamsWithContext(ctx context.Context) *UpdateTeamMemberParams {
	var ()
	return &UpdateTeamMemberParams{

		Context: ctx,
	}
}

// NewUpdateTeamMemberParamsWithHTTPClient creates a new UpdateTeamMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTeamMemberParamsWithHTTPClient(client *http.Client) *UpdateTeamMemberParams {
	var ()
	return &UpdateTeamMemberParams{
		HTTPClient: client,
	}
}

/*UpdateTeamMemberParams contains all the parameters to send to the API endpoint
for the update team member operation typically these are written to a http.Request
*/
type UpdateTeamMemberParams struct {

	/*Body
	  An object containing the fields to POST for the request.

	See the corresponding object definition for field details.

	*/
	Body *models.UpdateTeamMemberRequest
	/*TeamMemberID
	  The ID of the team member to update.

	*/
	TeamMemberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update team member params
func (o *UpdateTeamMemberParams) WithTimeout(timeout time.Duration) *UpdateTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update team member params
func (o *UpdateTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update team member params
func (o *UpdateTeamMemberParams) WithContext(ctx context.Context) *UpdateTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update team member params
func (o *UpdateTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update team member params
func (o *UpdateTeamMemberParams) WithHTTPClient(client *http.Client) *UpdateTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update team member params
func (o *UpdateTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update team member params
func (o *UpdateTeamMemberParams) WithBody(body *models.UpdateTeamMemberRequest) *UpdateTeamMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update team member params
func (o *UpdateTeamMemberParams) SetBody(body *models.UpdateTeamMemberRequest) {
	o.Body = body
}

// WithTeamMemberID adds the teamMemberID to the update team member params
func (o *UpdateTeamMemberParams) WithTeamMemberID(teamMemberID string) *UpdateTeamMemberParams {
	o.SetTeamMemberID(teamMemberID)
	return o
}

// SetTeamMemberID adds the teamMemberId to the update team member params
func (o *UpdateTeamMemberParams) SetTeamMemberID(teamMemberID string) {
	o.TeamMemberID = teamMemberID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param team_member_id
	if err := r.SetPathParam("team_member_id", o.TeamMemberID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}