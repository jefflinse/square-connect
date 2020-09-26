// Code generated by go-swagger; DO NOT EDIT.

package v1_locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v1 locations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1 locations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	RetrieveBusiness(params *RetrieveBusinessParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveBusinessOK, error)

	V1ListLocations(params *V1ListLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListLocationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  RetrieveBusiness retrieves business

  Get the general information for a business.
*/
func (a *Client) RetrieveBusiness(params *RetrieveBusinessParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveBusinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveBusinessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveBusiness",
		Method:             "GET",
		PathPattern:        "/v1/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveBusinessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveBusinessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveBusiness: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  V1ListLocations lists locations

  Provides details for all business locations associated with a Square
account, including the Square-assigned object ID for the location.
*/
func (a *Client) V1ListLocations(params *V1ListLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*V1ListLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1ListLocationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1ListLocations",
		Method:             "GET",
		PathPattern:        "/v1/me/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &V1ListLocationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*V1ListLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for V1ListLocations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
