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

// ListDisputesReader is a Reader for the ListDisputes structure.
type ListDisputesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDisputesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDisputesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListDisputesOK creates a ListDisputesOK with default headers values
func NewListDisputesOK() *ListDisputesOK {
	return &ListDisputesOK{}
}

/* ListDisputesOK describes a response with status code 200, with default header values.

Success
*/
type ListDisputesOK struct {
	Payload *models.ListDisputesResponse
}

func (o *ListDisputesOK) Error() string {
	return fmt.Sprintf("[GET /v2/disputes][%d] listDisputesOK  %+v", 200, o.Payload)
}
func (o *ListDisputesOK) GetPayload() *models.ListDisputesResponse {
	return o.Payload
}

func (o *ListDisputesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDisputesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
