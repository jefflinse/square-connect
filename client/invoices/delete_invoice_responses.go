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

// DeleteInvoiceReader is a Reader for the DeleteInvoice structure.
type DeleteInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInvoiceOK creates a DeleteInvoiceOK with default headers values
func NewDeleteInvoiceOK() *DeleteInvoiceOK {
	return &DeleteInvoiceOK{}
}

/* DeleteInvoiceOK describes a response with status code 200, with default header values.

Success
*/
type DeleteInvoiceOK struct {
	Payload *models.DeleteInvoiceResponse
}

func (o *DeleteInvoiceOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/invoices/{invoice_id}][%d] deleteInvoiceOK  %+v", 200, o.Payload)
}
func (o *DeleteInvoiceOK) GetPayload() *models.DeleteInvoiceResponse {
	return o.Payload
}

func (o *DeleteInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
