// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscriptions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscriptions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelSubscription(params *CancelSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelSubscriptionOK, error)

	CreateSubscription(params *CreateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSubscriptionOK, error)

	ListSubscriptionEvents(params *ListSubscriptionEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSubscriptionEventsOK, error)

	RetrieveSubscription(params *RetrieveSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveSubscriptionOK, error)

	SearchSubscriptions(params *SearchSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchSubscriptionsOK, error)

	UpdateSubscription(params *UpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSubscriptionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelSubscription cancels subscription

  Sets the `canceled_date` field to the end of the active billing period.
After this date, the status changes from ACTIVE to CANCELED.
*/
func (a *Client) CancelSubscription(params *CancelSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CancelSubscription",
		Method:             "POST",
		PathPattern:        "/v2/subscriptions/{subscription_id}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*CancelSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CancelSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSubscription creates subscription

  Creates a subscription for a customer to a subscription plan.

If you provide a card on file in the request, Square charges the card for
the subscription. Otherwise, Square bills an invoice to the customer's email
address. The subscription starts immediately, unless the request includes
the optional `start_date`. Each individual subscription is associated with a particular location.
*/
func (a *Client) CreateSubscription(params *CreateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSubscription",
		Method:             "POST",
		PathPattern:        "/v2/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*CreateSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListSubscriptionEvents lists subscription events

  Lists all events for a specific subscription.
In the current implementation, only `START_SUBSCRIPTION` and `STOP_SUBSCRIPTION` (when the subscription was canceled) events are returned.
*/
func (a *Client) ListSubscriptionEvents(params *ListSubscriptionEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSubscriptionEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSubscriptionEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSubscriptionEvents",
		Method:             "GET",
		PathPattern:        "/v2/subscriptions/{subscription_id}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSubscriptionEventsReader{formats: a.formats},
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
	success, ok := result.(*ListSubscriptionEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListSubscriptionEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveSubscription retrieves subscription

  Retrieves a subscription.
*/
func (a *Client) RetrieveSubscription(params *RetrieveSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveSubscription",
		Method:             "GET",
		PathPattern:        "/v2/subscriptions/{subscription_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*RetrieveSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchSubscriptions searches subscriptions

  Searches for subscriptions.
Results are ordered chronologically by subscription creation date. If
the request specifies more than one location ID,
the endpoint orders the result
by location ID, and then by creation date within each location. If no locations are given
in the query, all locations are searched.

You can also optionally specify `customer_ids` to search by customer.
If left unset, all customers
associated with the specified locations are returned.
If the request specifies customer IDs, the endpoint orders results
first by location, within location by customer ID, and within
customer by subscription creation date.

For more information, see
[Retrieve subscriptions](/docs/subscriptions-api/overview#retrieve-subscriptions).
*/
func (a *Client) SearchSubscriptions(params *SearchSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchSubscriptionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchSubscriptions",
		Method:             "POST",
		PathPattern:        "/v2/subscriptions/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchSubscriptionsReader{formats: a.formats},
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
	success, ok := result.(*SearchSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchSubscriptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSubscription updates subscription

  Updates a subscription. You can set, modify, and clear the
`subscription` field values.
*/
func (a *Client) UpdateSubscription(params *UpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSubscription",
		Method:             "PUT",
		PathPattern:        "/v2/subscriptions/{subscription_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*UpdateSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
