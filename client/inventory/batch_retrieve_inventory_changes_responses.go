// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// BatchRetrieveInventoryChangesReader is a Reader for the BatchRetrieveInventoryChanges structure.
type BatchRetrieveInventoryChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchRetrieveInventoryChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchRetrieveInventoryChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBatchRetrieveInventoryChangesOK creates a BatchRetrieveInventoryChangesOK with default headers values
func NewBatchRetrieveInventoryChangesOK() *BatchRetrieveInventoryChangesOK {
	return &BatchRetrieveInventoryChangesOK{}
}

/*BatchRetrieveInventoryChangesOK handles this case with default header values.

Success
*/
type BatchRetrieveInventoryChangesOK struct {
	Payload *models.BatchRetrieveInventoryChangesResponse
}

func (o *BatchRetrieveInventoryChangesOK) Error() string {
	return fmt.Sprintf("[POST /v2/inventory/batch-retrieve-changes][%d] batchRetrieveInventoryChangesOK  %+v", 200, o.Payload)
}

func (o *BatchRetrieveInventoryChangesOK) GetPayload() *models.BatchRetrieveInventoryChangesResponse {
	return o.Payload
}

func (o *BatchRetrieveInventoryChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchRetrieveInventoryChangesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
