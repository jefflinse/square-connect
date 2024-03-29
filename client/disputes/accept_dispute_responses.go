// Code generated by go-swagger; DO NOT EDIT.

package disputes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// AcceptDisputeReader is a Reader for the AcceptDispute structure.
type AcceptDisputeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AcceptDisputeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAcceptDisputeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAcceptDisputeOK creates a AcceptDisputeOK with default headers values
func NewAcceptDisputeOK() *AcceptDisputeOK {
	return &AcceptDisputeOK{}
}

/* AcceptDisputeOK describes a response with status code 200, with default header values.

Success
*/
type AcceptDisputeOK struct {
	Payload *models.AcceptDisputeResponse
}

func (o *AcceptDisputeOK) Error() string {
	return fmt.Sprintf("[POST /v2/disputes/{dispute_id}/accept][%d] acceptDisputeOK  %+v", 200, o.Payload)
}
func (o *AcceptDisputeOK) GetPayload() *models.AcceptDisputeResponse {
	return o.Payload
}

func (o *AcceptDisputeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AcceptDisputeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
