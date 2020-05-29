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

// CompletePaymentReader is a Reader for the CompletePayment structure.
type CompletePaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompletePaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompletePaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompletePaymentOK creates a CompletePaymentOK with default headers values
func NewCompletePaymentOK() *CompletePaymentOK {
	return &CompletePaymentOK{}
}

/*CompletePaymentOK handles this case with default header values.

Success
*/
type CompletePaymentOK struct {
	Payload *models.CompletePaymentResponse
}

func (o *CompletePaymentOK) Error() string {
	return fmt.Sprintf("[POST /v2/payments/{payment_id}/complete][%d] completePaymentOK  %+v", 200, o.Payload)
}

func (o *CompletePaymentOK) GetPayload() *models.CompletePaymentResponse {
	return o.Payload
}

func (o *CompletePaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompletePaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
