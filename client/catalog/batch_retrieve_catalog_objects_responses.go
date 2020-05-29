// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// BatchRetrieveCatalogObjectsReader is a Reader for the BatchRetrieveCatalogObjects structure.
type BatchRetrieveCatalogObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchRetrieveCatalogObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchRetrieveCatalogObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBatchRetrieveCatalogObjectsOK creates a BatchRetrieveCatalogObjectsOK with default headers values
func NewBatchRetrieveCatalogObjectsOK() *BatchRetrieveCatalogObjectsOK {
	return &BatchRetrieveCatalogObjectsOK{}
}

/*BatchRetrieveCatalogObjectsOK handles this case with default header values.

Success
*/
type BatchRetrieveCatalogObjectsOK struct {
	Payload *models.BatchRetrieveCatalogObjectsResponse
}

func (o *BatchRetrieveCatalogObjectsOK) Error() string {
	return fmt.Sprintf("[POST /v2/catalog/batch-retrieve][%d] batchRetrieveCatalogObjectsOK  %+v", 200, o.Payload)
}

func (o *BatchRetrieveCatalogObjectsOK) GetPayload() *models.BatchRetrieveCatalogObjectsResponse {
	return o.Payload
}

func (o *BatchRetrieveCatalogObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchRetrieveCatalogObjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
