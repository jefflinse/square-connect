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

// CreateTerminalCheckoutReader is a Reader for the CreateTerminalCheckout structure.
type CreateTerminalCheckoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTerminalCheckoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTerminalCheckoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTerminalCheckoutOK creates a CreateTerminalCheckoutOK with default headers values
func NewCreateTerminalCheckoutOK() *CreateTerminalCheckoutOK {
	return &CreateTerminalCheckoutOK{}
}

/*CreateTerminalCheckoutOK handles this case with default header values.

Success
*/
type CreateTerminalCheckoutOK struct {
	Payload *models.CreateTerminalCheckoutResponse
}

func (o *CreateTerminalCheckoutOK) Error() string {
	return fmt.Sprintf("[POST /v2/terminals/checkouts][%d] createTerminalCheckoutOK  %+v", 200, o.Payload)
}

func (o *CreateTerminalCheckoutOK) GetPayload() *models.CreateTerminalCheckoutResponse {
	return o.Payload
}

func (o *CreateTerminalCheckoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateTerminalCheckoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
