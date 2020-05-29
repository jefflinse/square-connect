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

// ListDiscountsReader is a Reader for the ListDiscounts structure.
type ListDiscountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDiscountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDiscountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDiscountsOK creates a ListDiscountsOK with default headers values
func NewListDiscountsOK() *ListDiscountsOK {
	return &ListDiscountsOK{}
}

/*ListDiscountsOK handles this case with default header values.

Success
*/
type ListDiscountsOK struct {
	Payload []*models.V1Discount
}

func (o *ListDiscountsOK) Error() string {
	return fmt.Sprintf("[GET /v1/{location_id}/discounts][%d] listDiscountsOK  %+v", 200, o.Payload)
}

func (o *ListDiscountsOK) GetPayload() []*models.V1Discount {
	return o.Payload
}

func (o *ListDiscountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
