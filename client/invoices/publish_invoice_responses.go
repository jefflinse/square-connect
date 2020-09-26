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

// PublishInvoiceReader is a Reader for the PublishInvoice structure.
type PublishInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublishInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublishInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPublishInvoiceOK creates a PublishInvoiceOK with default headers values
func NewPublishInvoiceOK() *PublishInvoiceOK {
	return &PublishInvoiceOK{}
}

/*PublishInvoiceOK handles this case with default header values.

Success
*/
type PublishInvoiceOK struct {
	Payload *models.PublishInvoiceResponse
}

func (o *PublishInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /v2/invoices/{invoice_id}/publish][%d] publishInvoiceOK  %+v", 200, o.Payload)
}

func (o *PublishInvoiceOK) GetPayload() *models.PublishInvoiceResponse {
	return o.Payload
}

func (o *PublishInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublishInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}