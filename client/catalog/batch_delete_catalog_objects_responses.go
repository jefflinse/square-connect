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

// BatchDeleteCatalogObjectsReader is a Reader for the BatchDeleteCatalogObjects structure.
type BatchDeleteCatalogObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchDeleteCatalogObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchDeleteCatalogObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBatchDeleteCatalogObjectsOK creates a BatchDeleteCatalogObjectsOK with default headers values
func NewBatchDeleteCatalogObjectsOK() *BatchDeleteCatalogObjectsOK {
	return &BatchDeleteCatalogObjectsOK{}
}

/* BatchDeleteCatalogObjectsOK describes a response with status code 200, with default header values.

Success
*/
type BatchDeleteCatalogObjectsOK struct {
	Payload *models.BatchDeleteCatalogObjectsResponse
}

func (o *BatchDeleteCatalogObjectsOK) Error() string {
	return fmt.Sprintf("[POST /v2/catalog/batch-delete][%d] batchDeleteCatalogObjectsOK  %+v", 200, o.Payload)
}
func (o *BatchDeleteCatalogObjectsOK) GetPayload() *models.BatchDeleteCatalogObjectsResponse {
	return o.Payload
}

func (o *BatchDeleteCatalogObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchDeleteCatalogObjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
