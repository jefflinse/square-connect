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

// AdjustInventoryReader is a Reader for the AdjustInventory structure.
type AdjustInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdjustInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdjustInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdjustInventoryOK creates a AdjustInventoryOK with default headers values
func NewAdjustInventoryOK() *AdjustInventoryOK {
	return &AdjustInventoryOK{}
}

/*AdjustInventoryOK handles this case with default header values.

Success
*/
type AdjustInventoryOK struct {
	Payload *models.V1InventoryEntry
}

func (o *AdjustInventoryOK) Error() string {
	return fmt.Sprintf("[POST /v1/{location_id}/inventory/{variation_id}][%d] adjustInventoryOK  %+v", 200, o.Payload)
}

func (o *AdjustInventoryOK) GetPayload() *models.V1InventoryEntry {
	return o.Payload
}

func (o *AdjustInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1InventoryEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}