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

// SearchCatalogItemsReader is a Reader for the SearchCatalogItems structure.
type SearchCatalogItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCatalogItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchCatalogItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchCatalogItemsOK creates a SearchCatalogItemsOK with default headers values
func NewSearchCatalogItemsOK() *SearchCatalogItemsOK {
	return &SearchCatalogItemsOK{}
}

/*SearchCatalogItemsOK handles this case with default header values.

Success
*/
type SearchCatalogItemsOK struct {
	Payload *models.SearchCatalogItemsResponse
}

func (o *SearchCatalogItemsOK) Error() string {
	return fmt.Sprintf("[POST /v2/catalog/search-catalog-items][%d] searchCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *SearchCatalogItemsOK) GetPayload() *models.SearchCatalogItemsResponse {
	return o.Payload
}

func (o *SearchCatalogItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}