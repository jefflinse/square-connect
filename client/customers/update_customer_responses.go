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

// UpdateCustomerReader is a Reader for the UpdateCustomer structure.
type UpdateCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCustomerOK creates a UpdateCustomerOK with default headers values
func NewUpdateCustomerOK() *UpdateCustomerOK {
	return &UpdateCustomerOK{}
}

/*UpdateCustomerOK handles this case with default header values.

Success
*/
type UpdateCustomerOK struct {
	Payload *models.UpdateCustomerResponse
}

func (o *UpdateCustomerOK) Error() string {
	return fmt.Sprintf("[PUT /v2/customers/{customer_id}][%d] updateCustomerOK  %+v", 200, o.Payload)
}

func (o *UpdateCustomerOK) GetPayload() *models.UpdateCustomerResponse {
	return o.Payload
}

func (o *UpdateCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
