// Code generated by go-swagger; DO NOT EDIT.

package bank_accounts

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

// NewGetBankAccountParams creates a new GetBankAccountParams object
// with the default values initialized.
func NewGetBankAccountParams() *GetBankAccountParams {
	var ()
	return &GetBankAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBankAccountParamsWithTimeout creates a new GetBankAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBankAccountParamsWithTimeout(timeout time.Duration) *GetBankAccountParams {
	var ()
	return &GetBankAccountParams{

		timeout: timeout,
	}
}

// NewGetBankAccountParamsWithContext creates a new GetBankAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBankAccountParamsWithContext(ctx context.Context) *GetBankAccountParams {
	var ()
	return &GetBankAccountParams{

		Context: ctx,
	}
}

// NewGetBankAccountParamsWithHTTPClient creates a new GetBankAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBankAccountParamsWithHTTPClient(client *http.Client) *GetBankAccountParams {
	var ()
	return &GetBankAccountParams{
		HTTPClient: client,
	}
}

/*GetBankAccountParams contains all the parameters to send to the API endpoint
for the get bank account operation typically these are written to a http.Request
*/
type GetBankAccountParams struct {

	/*BankAccountID
	  Square-issued ID of the desired `BankAccount`.

	*/
	BankAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bank account params
func (o *GetBankAccountParams) WithTimeout(timeout time.Duration) *GetBankAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bank account params
func (o *GetBankAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bank account params
func (o *GetBankAccountParams) WithContext(ctx context.Context) *GetBankAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bank account params
func (o *GetBankAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bank account params
func (o *GetBankAccountParams) WithHTTPClient(client *http.Client) *GetBankAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bank account params
func (o *GetBankAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBankAccountID adds the bankAccountID to the get bank account params
func (o *GetBankAccountParams) WithBankAccountID(bankAccountID string) *GetBankAccountParams {
	o.SetBankAccountID(bankAccountID)
	return o
}

// SetBankAccountID adds the bankAccountId to the get bank account params
func (o *GetBankAccountParams) SetBankAccountID(bankAccountID string) {
	o.BankAccountID = bankAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBankAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bank_account_id
	if err := r.SetPathParam("bank_account_id", o.BankAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}