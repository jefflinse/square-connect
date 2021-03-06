// Code generated by go-swagger; DO NOT EDIT.

package terminal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CancelTerminalCheckoutReader is a Reader for the CancelTerminalCheckout structure.
type CancelTerminalCheckoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelTerminalCheckoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelTerminalCheckoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelTerminalCheckoutOK creates a CancelTerminalCheckoutOK with default headers values
func NewCancelTerminalCheckoutOK() *CancelTerminalCheckoutOK {
	return &CancelTerminalCheckoutOK{}
}

/* CancelTerminalCheckoutOK describes a response with status code 200, with default header values.

Success
*/
type CancelTerminalCheckoutOK struct {
	Payload *models.CancelTerminalCheckoutResponse
}

func (o *CancelTerminalCheckoutOK) Error() string {
	return fmt.Sprintf("[POST /v2/terminals/checkouts/{checkout_id}/cancel][%d] cancelTerminalCheckoutOK  %+v", 200, o.Payload)
}
func (o *CancelTerminalCheckoutOK) GetPayload() *models.CancelTerminalCheckoutResponse {
	return o.Payload
}

func (o *CancelTerminalCheckoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelTerminalCheckoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
