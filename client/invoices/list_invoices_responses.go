// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// ListInvoicesReader is a Reader for the ListInvoices structure.
type ListInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListInvoicesOK creates a ListInvoicesOK with default headers values
func NewListInvoicesOK() *ListInvoicesOK {
	return &ListInvoicesOK{}
}

/* ListInvoicesOK describes a response with status code 200, with default header values.

Success
*/
type ListInvoicesOK struct {
	Payload *models.ListInvoicesResponse
}

func (o *ListInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /v2/invoices][%d] listInvoicesOK  %+v", 200, o.Payload)
}
func (o *ListInvoicesOK) GetPayload() *models.ListInvoicesResponse {
	return o.Payload
}

func (o *ListInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListInvoicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
