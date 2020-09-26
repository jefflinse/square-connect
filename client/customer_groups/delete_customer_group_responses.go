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

// DeleteCustomerGroupReader is a Reader for the DeleteCustomerGroup structure.
type DeleteCustomerGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomerGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomerGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomerGroupOK creates a DeleteCustomerGroupOK with default headers values
func NewDeleteCustomerGroupOK() *DeleteCustomerGroupOK {
	return &DeleteCustomerGroupOK{}
}

/*DeleteCustomerGroupOK handles this case with default header values.

Success
*/
type DeleteCustomerGroupOK struct {
	Payload *models.DeleteCustomerGroupResponse
}

func (o *DeleteCustomerGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/customers/groups/{group_id}][%d] deleteCustomerGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteCustomerGroupOK) GetPayload() *models.DeleteCustomerGroupResponse {
	return o.Payload
}

func (o *DeleteCustomerGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCustomerGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}