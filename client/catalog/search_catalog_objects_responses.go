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

// SearchCatalogObjectsReader is a Reader for the SearchCatalogObjects structure.
type SearchCatalogObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCatalogObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchCatalogObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchCatalogObjectsOK creates a SearchCatalogObjectsOK with default headers values
func NewSearchCatalogObjectsOK() *SearchCatalogObjectsOK {
	return &SearchCatalogObjectsOK{}
}

/* SearchCatalogObjectsOK describes a response with status code 200, with default header values.

Success
*/
type SearchCatalogObjectsOK struct {
	Payload *models.SearchCatalogObjectsResponse
}

func (o *SearchCatalogObjectsOK) Error() string {
	return fmt.Sprintf("[POST /v2/catalog/search][%d] searchCatalogObjectsOK  %+v", 200, o.Payload)
}
func (o *SearchCatalogObjectsOK) GetPayload() *models.SearchCatalogObjectsResponse {
	return o.Payload
}

func (o *SearchCatalogObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchCatalogObjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
