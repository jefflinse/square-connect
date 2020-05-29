// Code generated by go-swagger; DO NOT EDIT.

package v1_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// ListCategoriesReader is a Reader for the ListCategories structure.
type ListCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCategoriesOK creates a ListCategoriesOK with default headers values
func NewListCategoriesOK() *ListCategoriesOK {
	return &ListCategoriesOK{}
}

/*ListCategoriesOK handles this case with default header values.

Success
*/
type ListCategoriesOK struct {
	Payload []*models.V1Category
}

func (o *ListCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /v1/{location_id}/categories][%d] listCategoriesOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesOK) GetPayload() []*models.V1Category {
	return o.Payload
}

func (o *ListCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
