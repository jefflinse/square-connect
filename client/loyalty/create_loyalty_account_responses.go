// Code generated by go-swagger; DO NOT EDIT.

package loyalty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CreateLoyaltyAccountReader is a Reader for the CreateLoyaltyAccount structure.
type CreateLoyaltyAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLoyaltyAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLoyaltyAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLoyaltyAccountOK creates a CreateLoyaltyAccountOK with default headers values
func NewCreateLoyaltyAccountOK() *CreateLoyaltyAccountOK {
	return &CreateLoyaltyAccountOK{}
}

/* CreateLoyaltyAccountOK describes a response with status code 200, with default header values.

Success
*/
type CreateLoyaltyAccountOK struct {
	Payload *models.CreateLoyaltyAccountResponse
}

func (o *CreateLoyaltyAccountOK) Error() string {
	return fmt.Sprintf("[POST /v2/loyalty/accounts][%d] createLoyaltyAccountOK  %+v", 200, o.Payload)
}
func (o *CreateLoyaltyAccountOK) GetPayload() *models.CreateLoyaltyAccountResponse {
	return o.Payload
}

func (o *CreateLoyaltyAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateLoyaltyAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
