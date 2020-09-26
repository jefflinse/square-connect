// Code generated by go-swagger; DO NOT EDIT.

package v1_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// RetrieveOrderReader is a Reader for the RetrieveOrder structure.
type RetrieveOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveOrderOK creates a RetrieveOrderOK with default headers values
func NewRetrieveOrderOK() *RetrieveOrderOK {
	return &RetrieveOrderOK{}
}

/*RetrieveOrderOK handles this case with default header values.

Success
*/
type RetrieveOrderOK struct {
	Payload *models.V1Order
}

func (o *RetrieveOrderOK) Error() string {
	return fmt.Sprintf("[GET /v1/{location_id}/orders/{order_id}][%d] retrieveOrderOK  %+v", 200, o.Payload)
}

func (o *RetrieveOrderOK) GetPayload() *models.V1Order {
	return o.Payload
}

func (o *RetrieveOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}