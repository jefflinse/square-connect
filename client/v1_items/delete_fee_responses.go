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

// DeleteFeeReader is a Reader for the DeleteFee structure.
type DeleteFeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFeeOK creates a DeleteFeeOK with default headers values
func NewDeleteFeeOK() *DeleteFeeOK {
	return &DeleteFeeOK{}
}

/*DeleteFeeOK handles this case with default header values.

Success
*/
type DeleteFeeOK struct {
	Payload *models.V1Fee
}

func (o *DeleteFeeOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/{location_id}/fees/{fee_id}][%d] deleteFeeOK  %+v", 200, o.Payload)
}

func (o *DeleteFeeOK) GetPayload() *models.V1Fee {
	return o.Payload
}

func (o *DeleteFeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Fee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}