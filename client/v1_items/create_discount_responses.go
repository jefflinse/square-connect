// Code generated by go-swagger; DO NOT EDIT.

package v1_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CreateDiscountReader is a Reader for the CreateDiscount structure.
type CreateDiscountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDiscountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDiscountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDiscountOK creates a CreateDiscountOK with default headers values
func NewCreateDiscountOK() *CreateDiscountOK {
	return &CreateDiscountOK{}
}

/*CreateDiscountOK handles this case with default header values.

Success
*/
type CreateDiscountOK struct {
	Payload *models.V1Discount
}

func (o *CreateDiscountOK) Error() string {
	return fmt.Sprintf("[POST /v1/{location_id}/discounts][%d] createDiscountOK  %+v", 200, o.Payload)
}

func (o *CreateDiscountOK) GetPayload() *models.V1Discount {
	return o.Payload
}

func (o *CreateDiscountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Discount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
