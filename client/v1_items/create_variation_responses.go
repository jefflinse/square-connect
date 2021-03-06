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

// CreateVariationReader is a Reader for the CreateVariation structure.
type CreateVariationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVariationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVariationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVariationOK creates a CreateVariationOK with default headers values
func NewCreateVariationOK() *CreateVariationOK {
	return &CreateVariationOK{}
}

/*CreateVariationOK handles this case with default header values.

Success
*/
type CreateVariationOK struct {
	Payload *models.V1Variation
}

func (o *CreateVariationOK) Error() string {
	return fmt.Sprintf("[POST /v1/{location_id}/items/{item_id}/variations][%d] createVariationOK  %+v", 200, o.Payload)
}

func (o *CreateVariationOK) GetPayload() *models.V1Variation {
	return o.Payload
}

func (o *CreateVariationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Variation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
