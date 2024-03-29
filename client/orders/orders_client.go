// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BatchRetrieveOrders(params *BatchRetrieveOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchRetrieveOrdersOK, error)

	CalculateOrder(params *CalculateOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CalculateOrderOK, error)

	CreateOrder(params *CreateOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrderOK, error)

	PayOrder(params *PayOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PayOrderOK, error)

	RetrieveOrder(params *RetrieveOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveOrderOK, error)

	SearchOrders(params *SearchOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOrdersOK, error)

	UpdateOrder(params *UpdateOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BatchRetrieveOrders batches retrieve orders

  Retrieves a set of [Order](#type-order)s by their IDs.

If a given Order ID does not exist, the ID is ignored instead of generating an error.
*/
func (a *Client) BatchRetrieveOrders(params *BatchRetrieveOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BatchRetrieveOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchRetrieveOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BatchRetrieveOrders",
		Method:             "POST",
		PathPattern:        "/v2/orders/batch-retrieve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchRetrieveOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchRetrieveOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchRetrieveOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CalculateOrder calculates order

  Calculates an [Order](#type-order).
*/
func (a *Client) CalculateOrder(params *CalculateOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CalculateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCalculateOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CalculateOrder",
		Method:             "POST",
		PathPattern:        "/v2/orders/calculate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CalculateOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CalculateOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CalculateOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateOrder creates order

  Creates a new [Order](#type-order) which can include information on products for
purchase and settings to apply to the purchase.

To pay for a created order, please refer to the [Pay for Orders](/orders-api/pay-for-orders)
guide.

You can modify open orders using the [UpdateOrder](#endpoint-orders-updateorder) endpoint.
*/
func (a *Client) CreateOrder(params *CreateOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateOrder",
		Method:             "POST",
		PathPattern:        "/v2/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PayOrder pays order

  Pay for an [order](#type-order) using one or more approved [payments](#type-payment),
or settle an order with a total of `0`.

The total of the `payment_ids` listed in the request must be equal to the order
total. Orders with a total amount of `0` can be marked as paid by specifying an empty
array of `payment_ids` in the request.

To be used with PayOrder, a payment must:

- Reference the order by specifying the `order_id` when [creating the payment](#endpoint-payments-createpayment).
Any approved payments that reference the same `order_id` not specified in the
`payment_ids` will be canceled.
- Be approved with [delayed capture](/payments-api/take-payments#delayed-capture).
Using a delayed capture payment with PayOrder will complete the approved payment.
*/
func (a *Client) PayOrder(params *PayOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PayOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPayOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PayOrder",
		Method:             "POST",
		PathPattern:        "/v2/orders/{order_id}/pay",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PayOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PayOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PayOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveOrder retrieves order

  Retrieves an [Order](#type-order) by ID.
*/
func (a *Client) RetrieveOrder(params *RetrieveOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveOrder",
		Method:             "GET",
		PathPattern:        "/v2/orders/{order_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
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
  SearchOrders searches orders

  Search all orders for one or more locations. Orders include all sales,
returns, and exchanges regardless of how or when they entered the Square
Ecosystem (e.g. Point of Sale, Invoices, Connect APIs, etc).

SearchOrders requests need to specify which locations to search and define a
[`SearchOrdersQuery`](#type-searchordersquery) object which controls
how to sort or filter the results. Your SearchOrdersQuery can:

  Set filter criteria.
  Set sort order.
  Determine whether to return results as complete Order objects, or as
[OrderEntry](#type-orderentry) objects.

Note that details for orders processed with Square Point of Sale while in
offline mode may not be transmitted to Square for up to 72 hours. Offline
orders have a `created_at` value that reflects the time the order was created,
not the time it was subsequently transmitted to Square.
*/
func (a *Client) SearchOrders(params *SearchOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchOrders",
		Method:             "POST",
		PathPattern:        "/v2/orders/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateOrder updates order

  Updates an open [Order](#type-order) by adding, replacing, or deleting
fields. Orders with a `COMPLETED` or `CANCELED` state cannot be updated.

An UpdateOrder request requires the following:

- The `order_id` in the endpoint path, identifying the order to update.
- The latest `version` of the order to update.
- The [sparse order](/orders-api/manage-orders#sparse-order-objects)
containing only the fields to update and the version the update is
being applied to.
- If deleting fields, the [dot notation paths](/orders-api/manage-orders#on-dot-notation)
identifying fields to clear.

To pay for an order, please refer to the [Pay for Orders](/orders-api/pay-for-orders) guide.
*/
func (a *Client) UpdateOrder(params *UpdateOrderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateOrder",
		Method:             "PUT",
		PathPattern:        "/v2/orders/{order_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
