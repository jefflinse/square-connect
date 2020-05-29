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

// DeleteModifierListReader is a Reader for the DeleteModifierList structure.
type DeleteModifierListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModifierListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteModifierListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteModifierListOK creates a DeleteModifierListOK with default headers values
func NewDeleteModifierListOK() *DeleteModifierListOK {
	return &DeleteModifierListOK{}
}

/*DeleteModifierListOK handles this case with default header values.

Success
*/
type DeleteModifierListOK struct {
	Payload *models.V1ModifierList
}

func (o *DeleteModifierListOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/{location_id}/modifier-lists/{modifier_list_id}][%d] deleteModifierListOK  %+v", 200, o.Payload)
}

func (o *DeleteModifierListOK) GetPayload() *models.V1ModifierList {
	return o.Payload
}

func (o *DeleteModifierListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ModifierList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
