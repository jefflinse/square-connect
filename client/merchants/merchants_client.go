// Code generated by go-swagger; DO NOT EDIT.

package merchants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new merchants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for merchants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListMerchants(params *ListMerchantsParams, authInfo runtime.ClientAuthInfoWriter) (*ListMerchantsOK, error)

	RetrieveMerchant(params *RetrieveMerchantParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveMerchantOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListMerchants lists merchants

  Returns `Merchant` information for a given access token.

If you don't know a `Merchant` ID, you can use this endpoint to retrieve the merchant ID for an access token.
You can specify your personal access token to get your own merchant information or specify an OAuth token
to get the information for the  merchant that granted you access.

If you know the merchant ID, you can also use the [RetrieveMerchant](#endpoint-merchants-retrievemerchant)
endpoint to get the merchant information.
*/
func (a *Client) ListMerchants(params *ListMerchantsParams, authInfo runtime.ClientAuthInfoWriter) (*ListMerchantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMerchantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListMerchants",
		Method:             "GET",
		PathPattern:        "/v2/merchants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListMerchantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMerchantsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListMerchants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveMerchant retrieves merchant

  Retrieve a `Merchant` object for the given `merchant_id`.
*/
func (a *Client) RetrieveMerchant(params *RetrieveMerchantParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveMerchantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveMerchantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveMerchant",
		Method:             "GET",
		PathPattern:        "/v2/merchants/{merchant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveMerchantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveMerchantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveMerchant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
