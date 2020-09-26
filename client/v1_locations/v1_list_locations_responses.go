// Code generated by go-swagger; DO NOT EDIT.

package v1_locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// V1ListLocationsReader is a Reader for the V1ListLocations structure.
type V1ListLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ListLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ListLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewV1ListLocationsOK creates a V1ListLocationsOK with default headers values
func NewV1ListLocationsOK() *V1ListLocationsOK {
	return &V1ListLocationsOK{}
}

/*V1ListLocationsOK handles this case with default header values.

Success
*/
type V1ListLocationsOK struct {
	Payload []*models.V1Merchant
}

func (o *V1ListLocationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/me/locations][%d] v1ListLocationsOK  %+v", 200, o.Payload)
}

func (o *V1ListLocationsOK) GetPayload() []*models.V1Merchant {
	return o.Payload
}

func (o *V1ListLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}