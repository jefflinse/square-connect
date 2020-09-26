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

// UpdatePageCellReader is a Reader for the UpdatePageCell structure.
type UpdatePageCellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePageCellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePageCellOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePageCellOK creates a UpdatePageCellOK with default headers values
func NewUpdatePageCellOK() *UpdatePageCellOK {
	return &UpdatePageCellOK{}
}

/*UpdatePageCellOK handles this case with default header values.

Success
*/
type UpdatePageCellOK struct {
	Payload *models.V1Page
}

func (o *UpdatePageCellOK) Error() string {
	return fmt.Sprintf("[PUT /v1/{location_id}/pages/{page_id}/cells][%d] updatePageCellOK  %+v", 200, o.Payload)
}

func (o *UpdatePageCellOK) GetPayload() *models.V1Page {
	return o.Payload
}

func (o *UpdatePageCellOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Page)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}