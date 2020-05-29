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

// CatalogInfoReader is a Reader for the CatalogInfo structure.
type CatalogInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCatalogInfoOK creates a CatalogInfoOK with default headers values
func NewCatalogInfoOK() *CatalogInfoOK {
	return &CatalogInfoOK{}
}

/*CatalogInfoOK handles this case with default header values.

Success
*/
type CatalogInfoOK struct {
	Payload *models.CatalogInfoResponse
}

func (o *CatalogInfoOK) Error() string {
	return fmt.Sprintf("[GET /v2/catalog/info][%d] catalogInfoOK  %+v", 200, o.Payload)
}

func (o *CatalogInfoOK) GetPayload() *models.CatalogInfoResponse {
	return o.Payload
}

func (o *CatalogInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
