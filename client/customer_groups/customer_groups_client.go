// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCustomerGroup(params *CreateCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCustomerGroupOK, error)

	DeleteCustomerGroup(params *DeleteCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCustomerGroupOK, error)

	ListCustomerGroups(params *ListCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListCustomerGroupsOK, error)

	RetrieveCustomerGroup(params *RetrieveCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveCustomerGroupOK, error)

	UpdateCustomerGroup(params *UpdateCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCustomerGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCustomerGroup creates customer group

  Creates a new customer group for a business.

The request must include the `name` value of the group.
*/
func (a *Client) CreateCustomerGroup(params *CreateCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCustomerGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCustomerGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCustomerGroup",
		Method:             "POST",
		PathPattern:        "/v2/customers/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCustomerGroupReader{formats: a.formats},
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
	success, ok := result.(*CreateCustomerGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateCustomerGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCustomerGroup deletes customer group

  Deletes a customer group as identified by the `group_id` value.
*/
func (a *Client) DeleteCustomerGroup(params *DeleteCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCustomerGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomerGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCustomerGroup",
		Method:             "DELETE",
		PathPattern:        "/v2/customers/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCustomerGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteCustomerGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCustomerGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCustomerGroups lists customer groups

  Retrieves the list of customer groups of a business.
*/
func (a *Client) ListCustomerGroups(params *ListCustomerGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListCustomerGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCustomerGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListCustomerGroups",
		Method:             "GET",
		PathPattern:        "/v2/customers/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCustomerGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListCustomerGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCustomerGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveCustomerGroup retrieves customer group

  Retrieves a specific customer group as identified by the `group_id` value.
*/
func (a *Client) RetrieveCustomerGroup(params *RetrieveCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveCustomerGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveCustomerGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveCustomerGroup",
		Method:             "GET",
		PathPattern:        "/v2/customers/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveCustomerGroupReader{formats: a.formats},
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
	success, ok := result.(*RetrieveCustomerGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveCustomerGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCustomerGroup updates customer group

  Updates a customer group as identified by the `group_id` value.
*/
func (a *Client) UpdateCustomerGroup(params *UpdateCustomerGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCustomerGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCustomerGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCustomerGroup",
		Method:             "PUT",
		PathPattern:        "/v2/customers/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCustomerGroupReader{formats: a.formats},
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
	success, ok := result.(*UpdateCustomerGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateCustomerGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
