// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// GetDeviceCodeReader is a Reader for the GetDeviceCode structure.
type GetDeviceCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceCodeOK creates a GetDeviceCodeOK with default headers values
func NewGetDeviceCodeOK() *GetDeviceCodeOK {
	return &GetDeviceCodeOK{}
}

/* GetDeviceCodeOK describes a response with status code 200, with default header values.

Success
*/
type GetDeviceCodeOK struct {
	Payload *models.GetDeviceCodeResponse
}

func (o *GetDeviceCodeOK) Error() string {
	return fmt.Sprintf("[GET /v2/devices/codes/{id}][%d] getDeviceCodeOK  %+v", 200, o.Payload)
}
func (o *GetDeviceCodeOK) GetPayload() *models.GetDeviceCodeResponse {
	return o.Payload
}

func (o *GetDeviceCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDeviceCodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
