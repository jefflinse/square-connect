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

// UpdateVariationReader is a Reader for the UpdateVariation structure.
type UpdateVariationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVariationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVariationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateVariationOK creates a UpdateVariationOK with default headers values
func NewUpdateVariationOK() *UpdateVariationOK {
	return &UpdateVariationOK{}
}

/*UpdateVariationOK handles this case with default header values.

Success
*/
type UpdateVariationOK struct {
	Payload *models.V1Variation
}

func (o *UpdateVariationOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{location_id}/items/{item_id}/variations/{variation_id}][%d] updateVariationOK  %+v", 200, o.Payload)
}

func (o *UpdateVariationOK) GetPayload() *models.V1Variation {
	return o.Payload
}

func (o *UpdateVariationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Variation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
