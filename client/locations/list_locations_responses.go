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

// ListLocationsReader is a Reader for the ListLocations structure.
type ListLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListLocationsOK creates a ListLocationsOK with default headers values
func NewListLocationsOK() *ListLocationsOK {
	return &ListLocationsOK{}
}

/* ListLocationsOK describes a response with status code 200, with default header values.

Success
*/
type ListLocationsOK struct {
	Payload *models.ListLocationsResponse
}

func (o *ListLocationsOK) Error() string {
	return fmt.Sprintf("[GET /v2/locations][%d] listLocationsOK  %+v", 200, o.Payload)
}
func (o *ListLocationsOK) GetPayload() *models.ListLocationsResponse {
	return o.Payload
}

func (o *ListLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListLocationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
