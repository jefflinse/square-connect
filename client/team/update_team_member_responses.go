// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// UpdateTeamMemberReader is a Reader for the UpdateTeamMember structure.
type UpdateTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTeamMemberOK creates a UpdateTeamMemberOK with default headers values
func NewUpdateTeamMemberOK() *UpdateTeamMemberOK {
	return &UpdateTeamMemberOK{}
}

/*UpdateTeamMemberOK handles this case with default header values.

Success
*/
type UpdateTeamMemberOK struct {
	Payload *models.UpdateTeamMemberResponse
}

func (o *UpdateTeamMemberOK) Error() string {
	return fmt.Sprintf("[PUT /v2/team-members/{team_member_id}][%d] updateTeamMemberOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamMemberOK) GetPayload() *models.UpdateTeamMemberResponse {
	return o.Payload
}

func (o *UpdateTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateTeamMemberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}