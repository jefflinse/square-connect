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

// NewDeleteModifierListParams creates a new DeleteModifierListParams object
// with the default values initialized.
func NewDeleteModifierListParams() *DeleteModifierListParams {
	var ()
	return &DeleteModifierListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModifierListParamsWithTimeout creates a new DeleteModifierListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteModifierListParamsWithTimeout(timeout time.Duration) *DeleteModifierListParams {
	var ()
	return &DeleteModifierListParams{

		timeout: timeout,
	}
}

// NewDeleteModifierListParamsWithContext creates a new DeleteModifierListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteModifierListParamsWithContext(ctx context.Context) *DeleteModifierListParams {
	var ()
	return &DeleteModifierListParams{

		Context: ctx,
	}
}

// NewDeleteModifierListParamsWithHTTPClient creates a new DeleteModifierListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteModifierListParamsWithHTTPClient(client *http.Client) *DeleteModifierListParams {
	var ()
	return &DeleteModifierListParams{
		HTTPClient: client,
	}
}

/*DeleteModifierListParams contains all the parameters to send to the API endpoint
for the delete modifier list operation typically these are written to a http.Request
*/
type DeleteModifierListParams struct {

	/*LocationID
	  The ID of the item's associated location.

	*/
	LocationID string
	/*ModifierListID
	  The ID of the modifier list to delete.

	*/
	ModifierListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete modifier list params
func (o *DeleteModifierListParams) WithTimeout(timeout time.Duration) *DeleteModifierListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete modifier list params
func (o *DeleteModifierListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete modifier list params
func (o *DeleteModifierListParams) WithContext(ctx context.Context) *DeleteModifierListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete modifier list params
func (o *DeleteModifierListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete modifier list params
func (o *DeleteModifierListParams) WithHTTPClient(client *http.Client) *DeleteModifierListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete modifier list params
func (o *DeleteModifierListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationID adds the locationID to the delete modifier list params
func (o *DeleteModifierListParams) WithLocationID(locationID string) *DeleteModifierListParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the delete modifier list params
func (o *DeleteModifierListParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithModifierListID adds the modifierListID to the delete modifier list params
func (o *DeleteModifierListParams) WithModifierListID(modifierListID string) *DeleteModifierListParams {
	o.SetModifierListID(modifierListID)
	return o
}

// SetModifierListID adds the modifierListId to the delete modifier list params
func (o *DeleteModifierListParams) SetModifierListID(modifierListID string) {
	o.ModifierListID = modifierListID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModifierListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param location_id
	if err := r.SetPathParam("location_id", o.LocationID); err != nil {
		return err
	}

	// path param modifier_list_id
	if err := r.SetPathParam("modifier_list_id", o.ModifierListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
