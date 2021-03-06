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

// AddGroupToCustomerReader is a Reader for the AddGroupToCustomer structure.
type AddGroupToCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGroupToCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddGroupToCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddGroupToCustomerOK creates a AddGroupToCustomerOK with default headers values
func NewAddGroupToCustomerOK() *AddGroupToCustomerOK {
	return &AddGroupToCustomerOK{}
}

/* AddGroupToCustomerOK describes a response with status code 200, with default header values.

Success
*/
type AddGroupToCustomerOK struct {
	Payload *models.AddGroupToCustomerResponse
}

func (o *AddGroupToCustomerOK) Error() string {
	return fmt.Sprintf("[PUT /v2/customers/{customer_id}/groups/{group_id}][%d] addGroupToCustomerOK  %+v", 200, o.Payload)
}
func (o *AddGroupToCustomerOK) GetPayload() *models.AddGroupToCustomerResponse {
	return o.Payload
}

func (o *AddGroupToCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddGroupToCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
