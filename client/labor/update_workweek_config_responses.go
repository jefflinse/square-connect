// Code generated by go-swagger; DO NOT EDIT.

package labor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// UpdateWorkweekConfigReader is a Reader for the UpdateWorkweekConfig structure.
type UpdateWorkweekConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWorkweekConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWorkweekConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateWorkweekConfigOK creates a UpdateWorkweekConfigOK with default headers values
func NewUpdateWorkweekConfigOK() *UpdateWorkweekConfigOK {
	return &UpdateWorkweekConfigOK{}
}

/* UpdateWorkweekConfigOK describes a response with status code 200, with default header values.

Success
*/
type UpdateWorkweekConfigOK struct {
	Payload *models.UpdateWorkweekConfigResponse
}

func (o *UpdateWorkweekConfigOK) Error() string {
	return fmt.Sprintf("[PUT /v2/labor/workweek-configs/{id}][%d] updateWorkweekConfigOK  %+v", 200, o.Payload)
}
func (o *UpdateWorkweekConfigOK) GetPayload() *models.UpdateWorkweekConfigResponse {
	return o.Payload
}

func (o *UpdateWorkweekConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateWorkweekConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
