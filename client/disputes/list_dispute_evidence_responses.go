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

// ListDisputeEvidenceReader is a Reader for the ListDisputeEvidence structure.
type ListDisputeEvidenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDisputeEvidenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDisputeEvidenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListDisputeEvidenceOK creates a ListDisputeEvidenceOK with default headers values
func NewListDisputeEvidenceOK() *ListDisputeEvidenceOK {
	return &ListDisputeEvidenceOK{}
}

/* ListDisputeEvidenceOK describes a response with status code 200, with default header values.

Success
*/
type ListDisputeEvidenceOK struct {
	Payload *models.ListDisputeEvidenceResponse
}

func (o *ListDisputeEvidenceOK) Error() string {
	return fmt.Sprintf("[GET /v2/disputes/{dispute_id}/evidence][%d] listDisputeEvidenceOK  %+v", 200, o.Payload)
}
func (o *ListDisputeEvidenceOK) GetPayload() *models.ListDisputeEvidenceResponse {
	return o.Payload
}

func (o *ListDisputeEvidenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDisputeEvidenceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
