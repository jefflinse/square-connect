// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// SearchOrdersReader is a Reader for the SearchOrders structure.
type SearchOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchOrdersOK creates a SearchOrdersOK with default headers values
func NewSearchOrdersOK() *SearchOrdersOK {
	return &SearchOrdersOK{}
}

/*SearchOrdersOK handles this case with default header values.

Success
*/
type SearchOrdersOK struct {
	Payload *models.SearchOrdersResponse
}

func (o *SearchOrdersOK) Error() string {
	return fmt.Sprintf("[POST /v2/orders/search][%d] searchOrdersOK  %+v", 200, o.Payload)
}

func (o *SearchOrdersOK) GetPayload() *models.SearchOrdersResponse {
	return o.Payload
}

func (o *SearchOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}