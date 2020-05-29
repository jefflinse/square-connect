// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CreateRefundReader is a Reader for the CreateRefund structure.
type CreateRefundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRefundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRefundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRefundOK creates a CreateRefundOK with default headers values
func NewCreateRefundOK() *CreateRefundOK {
	return &CreateRefundOK{}
}

/*CreateRefundOK handles this case with default header values.

Success
*/
type CreateRefundOK struct {
	Payload *models.CreateRefundResponse
}

func (o *CreateRefundOK) Error() string {
	return fmt.Sprintf("[POST /v2/locations/{location_id}/transactions/{transaction_id}/refund][%d] createRefundOK  %+v", 200, o.Payload)
}

func (o *CreateRefundOK) GetPayload() *models.CreateRefundResponse {
	return o.Payload
}

func (o *CreateRefundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateRefundResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
