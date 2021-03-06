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

// RetrieveSettlementReader is a Reader for the RetrieveSettlement structure.
type RetrieveSettlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveSettlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveSettlementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveSettlementOK creates a RetrieveSettlementOK with default headers values
func NewRetrieveSettlementOK() *RetrieveSettlementOK {
	return &RetrieveSettlementOK{}
}

/*RetrieveSettlementOK handles this case with default header values.

Success
*/
type RetrieveSettlementOK struct {
	Payload *models.V1Settlement
}

func (o *RetrieveSettlementOK) Error() string {
	return fmt.Sprintf("[GET /v1/{location_id}/settlements/{settlement_id}][%d] retrieveSettlementOK  %+v", 200, o.Payload)
}

func (o *RetrieveSettlementOK) GetPayload() *models.V1Settlement {
	return o.Payload
}

func (o *RetrieveSettlementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Settlement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
