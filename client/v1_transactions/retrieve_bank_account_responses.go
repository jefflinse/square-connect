// Code generated by go-swagger; DO NOT EDIT.

package v1_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// RetrieveBankAccountReader is a Reader for the RetrieveBankAccount structure.
type RetrieveBankAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveBankAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveBankAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveBankAccountOK creates a RetrieveBankAccountOK with default headers values
func NewRetrieveBankAccountOK() *RetrieveBankAccountOK {
	return &RetrieveBankAccountOK{}
}

/*RetrieveBankAccountOK handles this case with default header values.

Success
*/
type RetrieveBankAccountOK struct {
	Payload *models.V1BankAccount
}

func (o *RetrieveBankAccountOK) Error() string {
	return fmt.Sprintf("[GET /v1/{location_id}/bank-accounts/{bank_account_id}][%d] retrieveBankAccountOK  %+v", 200, o.Payload)
}

func (o *RetrieveBankAccountOK) GetPayload() *models.V1BankAccount {
	return o.Payload
}

func (o *RetrieveBankAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1BankAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}