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

// UpsertCatalogObjectReader is a Reader for the UpsertCatalogObject structure.
type UpsertCatalogObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertCatalogObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpsertCatalogObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpsertCatalogObjectOK creates a UpsertCatalogObjectOK with default headers values
func NewUpsertCatalogObjectOK() *UpsertCatalogObjectOK {
	return &UpsertCatalogObjectOK{}
}

/*UpsertCatalogObjectOK handles this case with default header values.

Success
*/
type UpsertCatalogObjectOK struct {
	Payload *models.UpsertCatalogObjectResponse
}

func (o *UpsertCatalogObjectOK) Error() string {
	return fmt.Sprintf("[POST /v2/catalog/object][%d] upsertCatalogObjectOK  %+v", 200, o.Payload)
}

func (o *UpsertCatalogObjectOK) GetPayload() *models.UpsertCatalogObjectResponse {
	return o.Payload
}

func (o *UpsertCatalogObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpsertCatalogObjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
