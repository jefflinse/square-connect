// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CancelPaymentReader is a Reader for the CancelPayment structure.
type CancelPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelPaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelPaymentOK creates a CancelPaymentOK with default headers values
func NewCancelPaymentOK() *CancelPaymentOK {
	return &CancelPaymentOK{}
}

/* CancelPaymentOK describes a response with status code 200, with default header values.

Success
*/
type CancelPaymentOK struct {
	Payload *models.CancelPaymentResponse
}

func (o *CancelPaymentOK) Error() string {
	return fmt.Sprintf("[POST /v2/payments/{payment_id}/cancel][%d] cancelPaymentOK  %+v", 200, o.Payload)
}
func (o *CancelPaymentOK) GetPayload() *models.CancelPaymentResponse {
	return o.Payload
}

func (o *CancelPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelPaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
