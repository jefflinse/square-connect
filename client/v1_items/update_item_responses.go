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

// UpdateItemReader is a Reader for the UpdateItem structure.
type UpdateItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateItemOK creates a UpdateItemOK with default headers values
func NewUpdateItemOK() *UpdateItemOK {
	return &UpdateItemOK{}
}

/*UpdateItemOK handles this case with default header values.

Success
*/
type UpdateItemOK struct {
	Payload *models.V1Item
}

func (o *UpdateItemOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{location_id}/items/{item_id}][%d] updateItemOK  %+v", 200, o.Payload)
}

func (o *UpdateItemOK) GetPayload() *models.V1Item {
	return o.Payload
}

func (o *UpdateItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
