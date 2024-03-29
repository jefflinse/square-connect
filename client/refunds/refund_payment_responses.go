// Code generated by go-swagger; DO NOT EDIT.

package refunds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// RefundPaymentReader is a Reader for the RefundPayment structure.
type RefundPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefundPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefundPaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefundPaymentOK creates a RefundPaymentOK with default headers values
func NewRefundPaymentOK() *RefundPaymentOK {
	return &RefundPaymentOK{}
}

/* RefundPaymentOK describes a response with status code 200, with default header values.

Success
*/
type RefundPaymentOK struct {
	Payload *models.RefundPaymentResponse
}

func (o *RefundPaymentOK) Error() string {
	return fmt.Sprintf("[POST /v2/refunds][%d] refundPaymentOK  %+v", 200, o.Payload)
}
func (o *RefundPaymentOK) GetPayload() *models.RefundPaymentResponse {
	return o.Payload
}

func (o *RefundPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RefundPaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
