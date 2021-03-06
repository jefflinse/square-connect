// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CancelPayment(params *CancelPaymentParams, authInfo runtime.ClientAuthInfoWriter) (*CancelPaymentOK, error)

	CancelPaymentByIdempotencyKey(params *CancelPaymentByIdempotencyKeyParams, authInfo runtime.ClientAuthInfoWriter) (*CancelPaymentByIdempotencyKeyOK, error)

	CompletePayment(params *CompletePaymentParams, authInfo runtime.ClientAuthInfoWriter) (*CompletePaymentOK, error)

	CreatePayment(params *CreatePaymentParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePaymentOK, error)

	GetPayment(params *GetPaymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentOK, error)

	ListPayments(params *ListPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPaymentsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelPayment cancels payment

  Cancels (voids) a payment. If you set `autocomplete` to false when creating a payment,
you can cancel the payment using this endpoint.
*/
func (a *Client) CancelPayment(params *CancelPaymentParams, authInfo runtime.ClientAuthInfoWriter) (*CancelPaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelPaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CancelPayment",
		Method:             "POST",
		PathPattern:        "/v2/payments/{payment_id}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelPaymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CancelPaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CancelPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CancelPaymentByIdempotencyKey cancels payment by idempotency key

  Cancels (voids) a payment identified by the idempotency key that is specified in the
request.

Use this method when status of a CreatePayment request is unknown. For example, after you send a
CreatePayment request a network error occurs and you don't get a response. In this case, you can
direct Square to cancel the payment using this endpoint. In the request, you provide the same
idempotency key that you provided in your CreatePayment request you want  to cancel. After
cancelling the payment, you can submit your CreatePayment request again.

Note that if no payment with the specified idempotency key is found, no action is taken, the end
point returns successfully.
*/
func (a *Client) CancelPaymentByIdempotencyKey(params *CancelPaymentByIdempotencyKeyParams, authInfo runtime.ClientAuthInfoWriter) (*CancelPaymentByIdempotencyKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelPaymentByIdempotencyKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CancelPaymentByIdempotencyKey",
		Method:             "POST",
		PathPattern:        "/v2/payments/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelPaymentByIdempotencyKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CancelPaymentByIdempotencyKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CancelPaymentByIdempotencyKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompletePayment completes payment

  Completes (captures) a payment.

By default, payments are set to complete immediately after they are created.
If you set autocomplete to false when creating a payment, you can complete (capture)
the payment using this endpoint.
*/
func (a *Client) CompletePayment(params *CompletePaymentParams, authInfo runtime.ClientAuthInfoWriter) (*CompletePaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompletePaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompletePayment",
		Method:             "POST",
		PathPattern:        "/v2/payments/{payment_id}/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompletePaymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompletePaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompletePayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreatePayment creates payment

  Charges a payment source, for example, a card
represented by customer's card on file or a card nonce. In addition
to the payment source, the request must also include the
amount to accept for the payment.

There are several optional parameters that you can include in the request.
For example, tip money, whether to autocomplete the payment, or a reference ID
to correlate this payment with another system.

The `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission is required
to enable application fees.
*/
func (a *Client) CreatePayment(params *CreatePaymentParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePayment",
		Method:             "POST",
		PathPattern:        "/v2/payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePaymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreatePayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPayment gets payment

  Retrieves details for a specific Payment.
*/
func (a *Client) GetPayment(params *GetPaymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayment",
		Method:             "GET",
		PathPattern:        "/v2/payments/{payment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPayments lists payments

  Retrieves a list of payments taken by the account making the request.

Max results per page: 100
*/
func (a *Client) ListPayments(params *ListPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPaymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPayments",
		Method:             "GET",
		PathPattern:        "/v2/payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPaymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
