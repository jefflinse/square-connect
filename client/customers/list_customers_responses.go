// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// ListCustomersReader is a Reader for the ListCustomers structure.
type ListCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCustomersOK creates a ListCustomersOK with default headers values
func NewListCustomersOK() *ListCustomersOK {
	return &ListCustomersOK{}
}

/* ListCustomersOK describes a response with status code 200, with default header values.

Success
*/
type ListCustomersOK struct {
	Payload *models.ListCustomersResponse
}

func (o *ListCustomersOK) Error() string {
	return fmt.Sprintf("[GET /v2/customers][%d] listCustomersOK  %+v", 200, o.Payload)
}
func (o *ListCustomersOK) GetPayload() *models.ListCustomersResponse {
	return o.Payload
}

func (o *ListCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCustomersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
