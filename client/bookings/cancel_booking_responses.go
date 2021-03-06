// Code generated by go-swagger; DO NOT EDIT.

package bookings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CancelBookingReader is a Reader for the CancelBooking structure.
type CancelBookingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelBookingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelBookingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelBookingOK creates a CancelBookingOK with default headers values
func NewCancelBookingOK() *CancelBookingOK {
	return &CancelBookingOK{}
}

/* CancelBookingOK describes a response with status code 200, with default header values.

Success
*/
type CancelBookingOK struct {
	Payload *models.CancelBookingResponse
}

func (o *CancelBookingOK) Error() string {
	return fmt.Sprintf("[POST /v2/bookings/{booking_id}/cancel][%d] cancelBookingOK  %+v", 200, o.Payload)
}
func (o *CancelBookingOK) GetPayload() *models.CancelBookingResponse {
	return o.Payload
}

func (o *CancelBookingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelBookingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}