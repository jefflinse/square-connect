// Code generated by go-swagger; DO NOT EDIT.

package bookings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bookings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bookings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelBooking(params *CancelBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelBookingOK, error)

	CreateBooking(params *CreateBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBookingOK, error)

	ListTeamMemberBookingProfiles(params *ListTeamMemberBookingProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTeamMemberBookingProfilesOK, error)

	RetrieveBooking(params *RetrieveBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveBookingOK, error)

	RetrieveBusinessBookingProfile(params *RetrieveBusinessBookingProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveBusinessBookingProfileOK, error)

	RetrieveTeamMemberBookingProfile(params *RetrieveTeamMemberBookingProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveTeamMemberBookingProfileOK, error)

	SearchAvailability(params *SearchAvailabilityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchAvailabilityOK, error)

	UpdateBooking(params *UpdateBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBookingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelBooking cancels booking

  Cancels an existing booking.
*/
func (a *Client) CancelBooking(params *CancelBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelBookingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelBookingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CancelBooking",
		Method:             "POST",
		PathPattern:        "/v2/bookings/{booking_id}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelBookingReader{formats: a.formats},
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
	success, ok := result.(*CancelBookingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CancelBooking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateBooking creates booking

  Creates a booking.
*/
func (a *Client) CreateBooking(params *CreateBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateBookingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBookingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBooking",
		Method:             "POST",
		PathPattern:        "/v2/bookings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBookingReader{formats: a.formats},
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
	success, ok := result.(*CreateBookingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateBooking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTeamMemberBookingProfiles lists team member booking profiles

  Lists booking profiles for team members.
*/
func (a *Client) ListTeamMemberBookingProfiles(params *ListTeamMemberBookingProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTeamMemberBookingProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTeamMemberBookingProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTeamMemberBookingProfiles",
		Method:             "GET",
		PathPattern:        "/v2/bookings/team-member-booking-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTeamMemberBookingProfilesReader{formats: a.formats},
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
	success, ok := result.(*ListTeamMemberBookingProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListTeamMemberBookingProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveBooking retrieves booking

  Retrieves a booking.
*/
func (a *Client) RetrieveBooking(params *RetrieveBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveBookingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveBookingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveBooking",
		Method:             "GET",
		PathPattern:        "/v2/bookings/{booking_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveBookingReader{formats: a.formats},
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
	success, ok := result.(*RetrieveBookingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveBooking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveBusinessBookingProfile retrieves business booking profile

  Retrieves a seller's booking profile.
*/
func (a *Client) RetrieveBusinessBookingProfile(params *RetrieveBusinessBookingProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveBusinessBookingProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveBusinessBookingProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveBusinessBookingProfile",
		Method:             "GET",
		PathPattern:        "/v2/bookings/business-booking-profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveBusinessBookingProfileReader{formats: a.formats},
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
	success, ok := result.(*RetrieveBusinessBookingProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveBusinessBookingProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveTeamMemberBookingProfile retrieves team member booking profile

  Retrieves a team member's booking profile.
*/
func (a *Client) RetrieveTeamMemberBookingProfile(params *RetrieveTeamMemberBookingProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RetrieveTeamMemberBookingProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveTeamMemberBookingProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RetrieveTeamMemberBookingProfile",
		Method:             "GET",
		PathPattern:        "/v2/bookings/team-member-booking-profiles/{team_member_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveTeamMemberBookingProfileReader{formats: a.formats},
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
	success, ok := result.(*RetrieveTeamMemberBookingProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RetrieveTeamMemberBookingProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchAvailability searches availability

  Searches for availabilities for booking.
*/
func (a *Client) SearchAvailability(params *SearchAvailabilityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SearchAvailabilityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchAvailabilityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchAvailability",
		Method:             "POST",
		PathPattern:        "/v2/bookings/availability/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchAvailabilityReader{formats: a.formats},
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
	success, ok := result.(*SearchAvailabilityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchAvailability: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBooking updates booking

  Updates a booking.
*/
func (a *Client) UpdateBooking(params *UpdateBookingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBookingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBookingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBooking",
		Method:             "PUT",
		PathPattern:        "/v2/bookings/{booking_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBookingReader{formats: a.formats},
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
	success, ok := result.(*UpdateBookingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBooking: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
