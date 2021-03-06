// Code generated by go-swagger; DO NOT EDIT.

package refunds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new refunds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for refunds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetPaymentRefund(params *GetPaymentRefundParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentRefundOK, error)

	ListPaymentRefunds(params *ListPaymentRefundsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPaymentRefundsOK, error)

	RefundPayment(params *RefundPaymentParams, authInfo runtime.ClientAuthInfoWriter) (*RefundPaymentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPaymentRefund gets payment refund

  Retrieves a specific `Refund` using the `refund_id`.
*/
func (a *Client) GetPaymentRefund(params *GetPaymentRefundParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentRefundOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentRefundParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRefund",
		Method:             "GET",
		PathPattern:        "/v2/refunds/{refund_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentRefundReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentRefundOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPaymentRefund: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPaymentRefunds lists payment refunds

  Retrieves a list of refunds for the account making the request.

Max results per page: 100
*/
func (a *Client) ListPaymentRefunds(params *ListPaymentRefundsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPaymentRefundsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPaymentRefundsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPaymentRefunds",
		Method:             "GET",
		PathPattern:        "/v2/refunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPaymentRefundsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPaymentRefundsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListPaymentRefunds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RefundPayment refunds payment

  Refunds a payment. You can refund the entire payment amount or a
portion of it.
*/
func (a *Client) RefundPayment(params *RefundPaymentParams, authInfo runtime.ClientAuthInfoWriter) (*RefundPaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefundPaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RefundPayment",
		Method:             "POST",
		PathPattern:        "/v2/refunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RefundPaymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefundPaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RefundPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
