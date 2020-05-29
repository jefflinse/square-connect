// Code generated by go-swagger; DO NOT EDIT.

package bank_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// GetBankAccountByV1IDReader is a Reader for the GetBankAccountByV1ID structure.
type GetBankAccountByV1IDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBankAccountByV1IDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBankAccountByV1IDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBankAccountByV1IDOK creates a GetBankAccountByV1IDOK with default headers values
func NewGetBankAccountByV1IDOK() *GetBankAccountByV1IDOK {
	return &GetBankAccountByV1IDOK{}
}

/*GetBankAccountByV1IDOK handles this case with default header values.

Success
*/
type GetBankAccountByV1IDOK struct {
	Payload *models.GetBankAccountByV1IDResponse
}

func (o *GetBankAccountByV1IDOK) Error() string {
	return fmt.Sprintf("[GET /v2/bank-accounts/by-v1-id/{v1_bank_account_id}][%d] getBankAccountByV1IdOK  %+v", 200, o.Payload)
}

func (o *GetBankAccountByV1IDOK) GetPayload() *models.GetBankAccountByV1IDResponse {
	return o.Payload
}

func (o *GetBankAccountByV1IDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetBankAccountByV1IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
