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

// NewGetBankAccountByV1IDParams creates a new GetBankAccountByV1IDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBankAccountByV1IDParams() *GetBankAccountByV1IDParams {
	return &GetBankAccountByV1IDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBankAccountByV1IDParamsWithTimeout creates a new GetBankAccountByV1IDParams object
// with the ability to set a timeout on a request.
func NewGetBankAccountByV1IDParamsWithTimeout(timeout time.Duration) *GetBankAccountByV1IDParams {
	return &GetBankAccountByV1IDParams{
		timeout: timeout,
	}
}

// NewGetBankAccountByV1IDParamsWithContext creates a new GetBankAccountByV1IDParams object
// with the ability to set a context for a request.
func NewGetBankAccountByV1IDParamsWithContext(ctx context.Context) *GetBankAccountByV1IDParams {
	return &GetBankAccountByV1IDParams{
		Context: ctx,
	}
}

// NewGetBankAccountByV1IDParamsWithHTTPClient creates a new GetBankAccountByV1IDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBankAccountByV1IDParamsWithHTTPClient(client *http.Client) *GetBankAccountByV1IDParams {
	return &GetBankAccountByV1IDParams{
		HTTPClient: client,
	}
}

/* GetBankAccountByV1IDParams contains all the parameters to send to the API endpoint
   for the get bank account by v1 Id operation.

   Typically these are written to a http.Request.
*/
type GetBankAccountByV1IDParams struct {

	/* V1BankAccountID.

	     Connect V1 ID of the desired `BankAccount`. For more information, see
	[Retrieve a bank account by using an ID issued by V1 Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api#retrieve-a-bank-account-by-using-an-id-issued-by-v1-bank-accounts-api).
	*/
	V1BankAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bank account by v1 Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBankAccountByV1IDParams) WithDefaults() *GetBankAccountByV1IDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bank account by v1 Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBankAccountByV1IDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) WithTimeout(timeout time.Duration) *GetBankAccountByV1IDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) WithContext(ctx context.Context) *GetBankAccountByV1IDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) WithHTTPClient(client *http.Client) *GetBankAccountByV1IDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1BankAccountID adds the v1BankAccountID to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) WithV1BankAccountID(v1BankAccountID string) *GetBankAccountByV1IDParams {
	o.SetV1BankAccountID(v1BankAccountID)
	return o
}

// SetV1BankAccountID adds the v1BankAccountId to the get bank account by v1 Id params
func (o *GetBankAccountByV1IDParams) SetV1BankAccountID(v1BankAccountID string) {
	o.V1BankAccountID = v1BankAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBankAccountByV1IDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param v1_bank_account_id
	if err := r.SetPathParam("v1_bank_account_id", o.V1BankAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
