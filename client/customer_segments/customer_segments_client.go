// Code generated by go-swagger; DO NOT EDIT.

package customer_segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer segments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer segments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListCustomerSegments(params *ListCustomerSegmentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListCustomerSegmentsOK, error)

	RetrieveCustomerSegment(params *RetrieveCustomerSegmentParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveCustomerSegmentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListCustomerSegments lists customer segments

  Retrieves the list of customer segments of a business.
*/
func (a *Client) ListCustomerSegments(params *ListCustomerSegmentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListCustomerSegmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCustomerSegmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCustomerSegments",
		Method:             "GET",
		PathPattern:        "/v2/customers/segments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCustomerSegmentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCustomerSegmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCustomerSegments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveCustomerSegment retrieves customer segment

  Retrieves a specific customer segment as identified by the `segment_id` value.
*/
func (a *Client) RetrieveCustomerSegment(params *RetrieveCustomerSegmentParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveCustomerSegmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveCustomerSegmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveCustomerSegment",
		Method:             "GET",
		PathPattern:        "/v2/customers/segments/{segment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveCustomerSegmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveCustomerSegmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveCustomerSegment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
