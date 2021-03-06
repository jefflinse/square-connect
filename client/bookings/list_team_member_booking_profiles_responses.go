// Code generated by go-swagger; DO NOT EDIT.

package bookings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// ListTeamMemberBookingProfilesReader is a Reader for the ListTeamMemberBookingProfiles structure.
type ListTeamMemberBookingProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamMemberBookingProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTeamMemberBookingProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTeamMemberBookingProfilesOK creates a ListTeamMemberBookingProfilesOK with default headers values
func NewListTeamMemberBookingProfilesOK() *ListTeamMemberBookingProfilesOK {
	return &ListTeamMemberBookingProfilesOK{}
}

/* ListTeamMemberBookingProfilesOK describes a response with status code 200, with default header values.

Success
*/
type ListTeamMemberBookingProfilesOK struct {
	Payload *models.ListTeamMemberBookingProfilesResponse
}

func (o *ListTeamMemberBookingProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v2/bookings/team-member-booking-profiles][%d] listTeamMemberBookingProfilesOK  %+v", 200, o.Payload)
}
func (o *ListTeamMemberBookingProfilesOK) GetPayload() *models.ListTeamMemberBookingProfilesResponse {
	return o.Payload
}

func (o *ListTeamMemberBookingProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListTeamMemberBookingProfilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
