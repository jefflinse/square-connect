// Code generated by go-swagger; DO NOT EDIT.

package v1_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v1 transactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1 transactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListOrders(params *ListOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrdersOK, error)

	ListSettlements(params *ListSettlementsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSettlementsOK, error)

	RetrieveBankAccount(params *RetrieveBankAccountParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveBankAccountOK, error)

	RetrieveOrder(params *RetrieveOrderParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveOrderOK, error)

	RetrievePayment(params *RetrievePaymentParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePaymentOK, error)

	RetrieveSettlement(params *RetrieveSettlementParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSettlementOK, error)

	V1CreateRefund(params *V1CreateRefundParams, authInfo runtime.ClientAuthInfoWriter) (*V1CreateRefundOK, error)

	V1ListBankAccounts(params *V1ListBankAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListBankAccountsOK, error)

	V1ListPayments(params *V1ListPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListPaymentsOK, error)

	V1ListRefunds(params *V1ListRefundsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListRefundsOK, error)

	V1UpdateOrder(params *V1UpdateOrderParams, authInfo runtime.ClientAuthInfoWriter) (*V1UpdateOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListOrders lists orders

  Provides summary information for a merchant's online store orders.
*/
func (a *Client) ListOrders(params *ListOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListOrders",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSettlements lists settlements

  Provides summary information for all deposits and withdrawals
initiated by Square to a linked bank account during a date range. Date
ranges cannot exceed one year in length.

*Note**: the ListSettlements endpoint does not provide entry
information.
*/
func (a *Client) ListSettlements(params *ListSettlementsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSettlementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSettlementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListSettlements",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/settlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSettlementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSettlementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListSettlements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveBankAccount retrieves bank account

  Provides non-confidential details for a merchant's associated bank account. This endpoint does not provide full bank account numbers, and there is no way to obtain a full bank account number with the Connect API.
*/
func (a *Client) RetrieveBankAccount(params *RetrieveBankAccountParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveBankAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveBankAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveBankAccount",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/bank-accounts/{bank_account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveBankAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveBankAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveBankAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveOrder retrieves order

  Provides comprehensive information for a single online store order, including the order's history.
*/
func (a *Client) RetrieveOrder(params *RetrieveOrderParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveOrder",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/orders/{order_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrievePayment retrieves payment

  Provides comprehensive information for a single payment.
*/
func (a *Client) RetrievePayment(params *RetrievePaymentParams, authInfo runtime.ClientAuthInfoWriter) (*RetrievePaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrievePaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrievePayment",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/payments/{payment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrievePaymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrievePaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrievePayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveSettlement retrieves settlement

  Provides comprehensive information for a single settlement.

The returned `Settlement` objects include an `entries` field that lists
the transactions that contribute to the settlement total. Most
settlement entries correspond to a payment payout, but settlement
entries are also generated for less common events, like refunds, manual
adjustments, or chargeback holds.

Square initiates its regular deposits as indicated in the
[Deposit Options with Square](https://squareup.com/help/us/en/article/3807)
help article. Details for a regular deposit are usually not available
from Connect API endpoints before 10 p.m. PST the same day.

Square does not know when an initiated settlement **completes**, only
whether it has failed. A completed settlement is typically reflected in
a bank account within 3 business days, but in exceptional cases it may
take longer.
*/
func (a *Client) RetrieveSettlement(params *RetrieveSettlementParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSettlementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSettlementParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveSettlement",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/settlements/{settlement_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveSettlementReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveSettlementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveSettlement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  V1CreateRefund creates refund

  Issues a refund for a previously processed payment. You must issue
a refund within 60 days of the associated payment.

You cannot issue a partial refund for a split tender payment. You must
instead issue a full or partial refund for a particular tender, by
providing the applicable tender id to the V1CreateRefund endpoint.
Issuing a full refund for a split tender payment refunds all tenders
associated with the payment.

Issuing a refund for a card payment is not reversible. For development
purposes, you can create fake cash payments in Square Point of Sale and
refund them.
*/
func (a *Client) V1CreateRefund(params *V1CreateRefundParams, authInfo runtime.ClientAuthInfoWriter) (*V1CreateRefundOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1CreateRefundParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1CreateRefund",
		Method:             "POST",
		PathPattern:        "/v1/{location_id}/refunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V1CreateRefundReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1CreateRefundOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for V1CreateRefund: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  V1ListBankAccounts lists bank accounts

  Provides non-confidential details for all of a location's associated bank accounts. This endpoint does not provide full bank account numbers, and there is no way to obtain a full bank account number with the Connect API.
*/
func (a *Client) V1ListBankAccounts(params *V1ListBankAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListBankAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1ListBankAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1ListBankAccounts",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/bank-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V1ListBankAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1ListBankAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for V1ListBankAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  V1ListPayments lists payments

  Provides summary information for all payments taken for a given
Square account during a date range. Date ranges cannot exceed 1 year in
length. See Date ranges for details of inclusive and exclusive dates.

*Note**: Details for payments processed with Square Point of Sale while
in offline mode may not be transmitted to Square for up to 72 hours.
Offline payments have a `created_at` value that reflects the time the
payment was originally processed, not the time it was subsequently
transmitted to Square. Consequently, the ListPayments endpoint might
list an offline payment chronologically between online payments that
were seen in a previous request.
*/
func (a *Client) V1ListPayments(params *V1ListPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListPaymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1ListPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1ListPayments",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V1ListPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1ListPaymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for V1ListPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  V1ListRefunds lists refunds

  Provides the details for all refunds initiated by a merchant or any of the merchant's mobile staff during a date range. Date ranges cannot exceed one year in length.
*/
func (a *Client) V1ListRefunds(params *V1ListRefundsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListRefundsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1ListRefundsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1ListRefunds",
		Method:             "GET",
		PathPattern:        "/v1/{location_id}/refunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V1ListRefundsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1ListRefundsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for V1ListRefunds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  V1UpdateOrder updates order

  Updates the details of an online store order. Every update you perform on an order corresponds to one of three actions:
*/
func (a *Client) V1UpdateOrder(params *V1UpdateOrderParams, authInfo runtime.ClientAuthInfoWriter) (*V1UpdateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1UpdateOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1UpdateOrder",
		Method:             "PUT",
		PathPattern:        "/v1/{location_id}/orders/{order_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V1UpdateOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1UpdateOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for V1UpdateOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
