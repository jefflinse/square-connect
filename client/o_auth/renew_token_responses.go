// Code generated by go-swagger; DO NOT EDIT.

package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// RenewTokenReader is a Reader for the RenewToken structure.
type RenewTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenewTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenewTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenewTokenOK creates a RenewTokenOK with default headers values
func NewRenewTokenOK() *RenewTokenOK {
	return &RenewTokenOK{}
}

/*RenewTokenOK handles this case with default header values.

Success
*/
type RenewTokenOK struct {
	Payload *models.RenewTokenResponse
}

func (o *RenewTokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/clients/{client_id}/access-token/renew][%d] renewTokenOK  %+v", 200, o.Payload)
}

func (o *RenewTokenOK) GetPayload() *models.RenewTokenResponse {
	return o.Payload
}

func (o *RenewTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RenewTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}