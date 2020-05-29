// Code generated by go-swagger; DO NOT EDIT.

package disputes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// RetrieveDisputeEvidenceReader is a Reader for the RetrieveDisputeEvidence structure.
type RetrieveDisputeEvidenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveDisputeEvidenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveDisputeEvidenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveDisputeEvidenceOK creates a RetrieveDisputeEvidenceOK with default headers values
func NewRetrieveDisputeEvidenceOK() *RetrieveDisputeEvidenceOK {
	return &RetrieveDisputeEvidenceOK{}
}

/*RetrieveDisputeEvidenceOK handles this case with default header values.

Success
*/
type RetrieveDisputeEvidenceOK struct {
	Payload *models.RetrieveDisputeEvidenceResponse
}

func (o *RetrieveDisputeEvidenceOK) Error() string {
	return fmt.Sprintf("[GET /v2/disputes/{dispute_id}/evidence/{evidence_id}][%d] retrieveDisputeEvidenceOK  %+v", 200, o.Payload)
}

func (o *RetrieveDisputeEvidenceOK) GetPayload() *models.RetrieveDisputeEvidenceResponse {
	return o.Payload
}

func (o *RetrieveDisputeEvidenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetrieveDisputeEvidenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
