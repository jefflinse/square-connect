// Code generated by go-swagger; DO NOT EDIT.

package disputes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new disputes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for disputes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptDispute(params *AcceptDisputeParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptDisputeOK, error)

	CreateDisputeEvidenceText(params *CreateDisputeEvidenceTextParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDisputeEvidenceTextOK, error)

	ListDisputeEvidence(params *ListDisputeEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*ListDisputeEvidenceOK, error)

	ListDisputes(params *ListDisputesParams, authInfo runtime.ClientAuthInfoWriter) (*ListDisputesOK, error)

	RemoveDisputeEvidence(params *RemoveDisputeEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveDisputeEvidenceOK, error)

	RetrieveDispute(params *RetrieveDisputeParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveDisputeOK, error)

	RetrieveDisputeEvidence(params *RetrieveDisputeEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveDisputeEvidenceOK, error)

	SubmitEvidence(params *SubmitEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*SubmitEvidenceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AcceptDispute accepts dispute

  Accepts loss on a dispute. Square returns
the disputed amount to the cardholder and updates the
dispute state to ACCEPTED.

Square debits the disputed amount from the seller’s Square
account. If the Square account balance does not have
sufficient funds, Square debits the associated bank account.
*/
func (a *Client) AcceptDispute(params *AcceptDisputeParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptDisputeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptDisputeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AcceptDispute",
		Method:             "POST",
		PathPattern:        "/v2/disputes/{dispute_id}/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptDisputeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptDisputeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AcceptDispute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDisputeEvidenceText creates dispute evidence text

  Uploads text to use as evidence for a dispute challenge.
*/
func (a *Client) CreateDisputeEvidenceText(params *CreateDisputeEvidenceTextParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDisputeEvidenceTextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDisputeEvidenceTextParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDisputeEvidenceText",
		Method:             "POST",
		PathPattern:        "/v2/disputes/{dispute_id}/evidence_text",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDisputeEvidenceTextReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDisputeEvidenceTextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDisputeEvidenceText: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDisputeEvidence lists dispute evidence

  Returns a list of evidence associated with a dispute.
*/
func (a *Client) ListDisputeEvidence(params *ListDisputeEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*ListDisputeEvidenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDisputeEvidenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDisputeEvidence",
		Method:             "GET",
		PathPattern:        "/v2/disputes/{dispute_id}/evidence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDisputeEvidenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDisputeEvidenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListDisputeEvidence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDisputes lists disputes

  Returns a list of disputes associated
with a particular account.
*/
func (a *Client) ListDisputes(params *ListDisputesParams, authInfo runtime.ClientAuthInfoWriter) (*ListDisputesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDisputesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDisputes",
		Method:             "GET",
		PathPattern:        "/v2/disputes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDisputesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDisputesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListDisputes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveDisputeEvidence removes dispute evidence

  Removes specified evidence from a dispute.

Square does not send the bank any evidence that
is removed. Also, you cannot remove evidence after
submitting it to the bank using [SubmitEvidence](/reference/square/disputes-api/submit-evidence).
*/
func (a *Client) RemoveDisputeEvidence(params *RemoveDisputeEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveDisputeEvidenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveDisputeEvidenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveDisputeEvidence",
		Method:             "DELETE",
		PathPattern:        "/v2/disputes/{dispute_id}/evidence/{evidence_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveDisputeEvidenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveDisputeEvidenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RemoveDisputeEvidence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveDispute retrieves dispute

  Returns details of a specific dispute.
*/
func (a *Client) RetrieveDispute(params *RetrieveDisputeParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveDisputeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveDisputeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveDispute",
		Method:             "GET",
		PathPattern:        "/v2/disputes/{dispute_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveDisputeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveDisputeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveDispute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveDisputeEvidence retrieves dispute evidence

  Returns the specific evidence metadata associated with a specific dispute.

You must maintain a copy of the evidence you upload if you want to
reference it later. You cannot download the evidence
after you upload it.
*/
func (a *Client) RetrieveDisputeEvidence(params *RetrieveDisputeEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveDisputeEvidenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveDisputeEvidenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RetrieveDisputeEvidence",
		Method:             "GET",
		PathPattern:        "/v2/disputes/{dispute_id}/evidence/{evidence_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveDisputeEvidenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveDisputeEvidenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveDisputeEvidence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubmitEvidence submits evidence

  Submits evidence to the cardholder's bank.

Before submitting evidence, Square compiles all available evidence. This includes
evidence uploaded using the
[CreateDisputeEvidenceFile](/reference/square/disputes-api/create-dispute-evidence-file) and
[CreateDisputeEvidenceText](/reference/square/disputes-api/create-dispute-evidence-text) endpoints,
and evidence automatically provided by Square, when
available.
*/
func (a *Client) SubmitEvidence(params *SubmitEvidenceParams, authInfo runtime.ClientAuthInfoWriter) (*SubmitEvidenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitEvidenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SubmitEvidence",
		Method:             "POST",
		PathPattern:        "/v2/disputes/{dispute_id}/submit-evidence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubmitEvidenceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubmitEvidenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SubmitEvidence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
