// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new locations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for locations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLocation(params *CreateLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateLocationOK, error)

	ListLocations(params *ListLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListLocationsOK, error)

	RetrieveLocation(params *RetrieveLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveLocationOK, error)

	UpdateLocation(params *UpdateLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLocationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateLocation creates location

  Creates a location.
*/
func (a *Client) CreateLocation(params *CreateLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateLocation",
		Method:             "POST",
		PathPattern:        "/v2/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLocationReader{formats: a.formats},
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
	success, ok := result.(*CreateLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListLocations lists locations

  Provides information of all locations of a business.

Many Square API endpoints require a `location_id` parameter.
The `id` field of the [`Location`](#type-location) objects returned by this
endpoint correspond to that `location_id` parameter.
*/
func (a *Client) ListLocations(params *ListLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListLocations",
		Method:             "GET",
		PathPattern:        "/v2/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListLocationsReader{formats: a.formats},
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
	success, ok := result.(*ListLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListLocations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveLocation retrieves location

  Retrieves details of a location. You can specify "main"
as the location ID to retrieve details of the
main location.
*/
func (a *Client) RetrieveLocation(params *RetrieveLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveLocation",
		Method:             "GET",
		PathPattern:        "/v2/locations/{location_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveLocationReader{formats: a.formats},
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
	success, ok := result.(*RetrieveLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLocation updates location

  Updates a location.
*/
func (a *Client) UpdateLocation(params *UpdateLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateLocation",
		Method:             "PUT",
		PathPattern:        "/v2/locations/{location_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLocationReader{formats: a.formats},
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
	success, ok := result.(*UpdateLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
