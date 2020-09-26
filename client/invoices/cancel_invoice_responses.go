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

// CancelInvoiceReader is a Reader for the CancelInvoice structure.
type CancelInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelInvoiceOK creates a CancelInvoiceOK with default headers values
func NewCancelInvoiceOK() *CancelInvoiceOK {
	return &CancelInvoiceOK{}
}

/*CancelInvoiceOK handles this case with default header values.

Success
*/
type CancelInvoiceOK struct {
	Payload *models.CancelInvoiceResponse
}

func (o *CancelInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /v2/invoices/{invoice_id}/cancel][%d] cancelInvoiceOK  %+v", 200, o.Payload)
}

func (o *CancelInvoiceOK) GetPayload() *models.CancelInvoiceResponse {
	return o.Payload
}

func (o *CancelInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelInvoiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
