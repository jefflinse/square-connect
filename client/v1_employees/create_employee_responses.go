// Code generated by go-swagger; DO NOT EDIT.

package v1_employees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jefflinse/square-connect/models"
)

// CreateEmployeeReader is a Reader for the CreateEmployee structure.
type CreateEmployeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEmployeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEmployeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateEmployeeOK creates a CreateEmployeeOK with default headers values
func NewCreateEmployeeOK() *CreateEmployeeOK {
	return &CreateEmployeeOK{}
}

/*CreateEmployeeOK handles this case with default header values.

Success
*/
type CreateEmployeeOK struct {
	Payload *models.V1Employee
}

func (o *CreateEmployeeOK) Error() string {
	return fmt.Sprintf("[POST /v1/me/employees][%d] createEmployeeOK  %+v", 200, o.Payload)
}

func (o *CreateEmployeeOK) GetPayload() *models.V1Employee {
	return o.Payload
}

func (o *CreateEmployeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Employee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
