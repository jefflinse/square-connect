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

// RetrieveCatalogObjectReader is a Reader for the RetrieveCatalogObject structure.
type RetrieveCatalogObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCatalogObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveCatalogObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrieveCatalogObjectOK creates a RetrieveCatalogObjectOK with default headers values
func NewRetrieveCatalogObjectOK() *RetrieveCatalogObjectOK {
	return &RetrieveCatalogObjectOK{}
}

/* RetrieveCatalogObjectOK describes a response with status code 200, with default header values.

Success
*/
type RetrieveCatalogObjectOK struct {
	Payload *models.RetrieveCatalogObjectResponse
}

func (o *RetrieveCatalogObjectOK) Error() string {
	return fmt.Sprintf("[GET /v2/catalog/object/{object_id}][%d] retrieveCatalogObjectOK  %+v", 200, o.Payload)
}
func (o *RetrieveCatalogObjectOK) GetPayload() *models.RetrieveCatalogObjectResponse {
	return o.Payload
}

func (o *RetrieveCatalogObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetrieveCatalogObjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
