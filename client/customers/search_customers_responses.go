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

// SearchCustomersReader is a Reader for the SearchCustomers structure.
type SearchCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchCustomersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchCustomersOK creates a SearchCustomersOK with default headers values
func NewSearchCustomersOK() *SearchCustomersOK {
	return &SearchCustomersOK{}
}

/* SearchCustomersOK describes a response with status code 200, with default header values.

Success
*/
type SearchCustomersOK struct {
	Payload *models.SearchCustomersResponse
}

func (o *SearchCustomersOK) Error() string {
	return fmt.Sprintf("[POST /v2/customers/search][%d] searchCustomersOK  %+v", 200, o.Payload)
}
func (o *SearchCustomersOK) GetPayload() *models.SearchCustomersResponse {
	return o.Payload
}

func (o *SearchCustomersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchCustomersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
