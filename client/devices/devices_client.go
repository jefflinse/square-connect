// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeviceCode(params *CreateDeviceCodeParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeviceCodeOK, error)

	GetDeviceCode(params *GetDeviceCodeParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCodeOK, error)

	ListDeviceCodes(params *ListDeviceCodesParams, authInfo runtime.ClientAuthInfoWriter) (*ListDeviceCodesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDeviceCode creates device code

  Creates a DeviceCode that can be used to login to a Square Terminal device to enter the connected
terminal mode.
*/
func (a *Client) CreateDeviceCode(params *CreateDeviceCodeParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeviceCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceCodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDeviceCode",
		Method:             "POST",
		PathPattern:        "/v2/devices/codes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceCodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeviceCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDeviceCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeviceCode gets device code

  Retrieves DeviceCode with the associated ID.
*/
func (a *Client) GetDeviceCode(params *GetDeviceCodeParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDeviceCode",
		Method:             "GET",
		PathPattern:        "/v2/devices/codes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeviceCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDeviceCodes lists device codes

  Lists all DeviceCodes associated with the merchant.
*/
func (a *Client) ListDeviceCodes(params *ListDeviceCodesParams, authInfo runtime.ClientAuthInfoWriter) (*ListDeviceCodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDeviceCodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDeviceCodes",
		Method:             "GET",
		PathPattern:        "/v2/devices/codes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDeviceCodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDeviceCodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListDeviceCodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
