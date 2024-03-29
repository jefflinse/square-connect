// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CreateLocationReader is a Reader for the CreateLocation structure.
type CreateLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLocationOK creates a CreateLocationOK with default headers values
func NewCreateLocationOK() *CreateLocationOK {
	return &CreateLocationOK{}
}

/* CreateLocationOK describes a response with status code 200, with default header values.

Success
*/
type CreateLocationOK struct {
	Payload *models.CreateLocationResponse
}

func (o *CreateLocationOK) Error() string {
	return fmt.Sprintf("[POST /v2/locations][%d] createLocationOK  %+v", 200, o.Payload)
}
func (o *CreateLocationOK) GetPayload() *models.CreateLocationResponse {
	return o.Payload
}

func (o *CreateLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateLocationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
