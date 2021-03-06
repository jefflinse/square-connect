// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// UpdateLocationReader is a Reader for the UpdateLocation structure.
type UpdateLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLocationOK creates a UpdateLocationOK with default headers values
func NewUpdateLocationOK() *UpdateLocationOK {
	return &UpdateLocationOK{}
}

/* UpdateLocationOK describes a response with status code 200, with default header values.

Success
*/
type UpdateLocationOK struct {
	Payload *models.UpdateLocationResponse
}

func (o *UpdateLocationOK) Error() string {
	return fmt.Sprintf("[PUT /v2/locations/{location_id}][%d] updateLocationOK  %+v", 200, o.Payload)
}
func (o *UpdateLocationOK) GetPayload() *models.UpdateLocationResponse {
	return o.Payload
}

func (o *UpdateLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateLocationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
