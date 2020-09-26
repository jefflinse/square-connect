// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// RetrieveCustomerGroupReader is a Reader for the RetrieveCustomerGroup structure.
type RetrieveCustomerGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveCustomerGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveCustomerGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveCustomerGroupOK creates a RetrieveCustomerGroupOK with default headers values
func NewRetrieveCustomerGroupOK() *RetrieveCustomerGroupOK {
	return &RetrieveCustomerGroupOK{}
}

/*RetrieveCustomerGroupOK handles this case with default header values.

Success
*/
type RetrieveCustomerGroupOK struct {
	Payload *models.RetrieveCustomerGroupResponse
}

func (o *RetrieveCustomerGroupOK) Error() string {
	return fmt.Sprintf("[GET /v2/customers/groups/{group_id}][%d] retrieveCustomerGroupOK  %+v", 200, o.Payload)
}

func (o *RetrieveCustomerGroupOK) GetPayload() *models.RetrieveCustomerGroupResponse {
	return o.Payload
}

func (o *RetrieveCustomerGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetrieveCustomerGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}