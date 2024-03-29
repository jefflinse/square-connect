// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// PayOrderReader is a Reader for the PayOrder structure.
type PayOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PayOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPayOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPayOrderOK creates a PayOrderOK with default headers values
func NewPayOrderOK() *PayOrderOK {
	return &PayOrderOK{}
}

/* PayOrderOK describes a response with status code 200, with default header values.

Success
*/
type PayOrderOK struct {
	Payload *models.PayOrderResponse
}

func (o *PayOrderOK) Error() string {
	return fmt.Sprintf("[POST /v2/orders/{order_id}/pay][%d] payOrderOK  %+v", 200, o.Payload)
}
func (o *PayOrderOK) GetPayload() *models.PayOrderResponse {
	return o.Payload
}

func (o *PayOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
