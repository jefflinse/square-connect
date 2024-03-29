// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// ListCustomerGroupsReader is a Reader for the ListCustomerGroups structure.
type ListCustomerGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCustomerGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCustomerGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCustomerGroupsOK creates a ListCustomerGroupsOK with default headers values
func NewListCustomerGroupsOK() *ListCustomerGroupsOK {
	return &ListCustomerGroupsOK{}
}

/* ListCustomerGroupsOK describes a response with status code 200, with default header values.

Success
*/
type ListCustomerGroupsOK struct {
	Payload *models.ListCustomerGroupsResponse
}

func (o *ListCustomerGroupsOK) Error() string {
	return fmt.Sprintf("[GET /v2/customers/groups][%d] listCustomerGroupsOK  %+v", 200, o.Payload)
}
func (o *ListCustomerGroupsOK) GetPayload() *models.ListCustomerGroupsResponse {
	return o.Payload
}

func (o *ListCustomerGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCustomerGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
