// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListOrdersRequest v1 list orders request
//
// swagger:model V1ListOrdersRequest
type V1ListOrdersRequest struct {

	// A pagination cursor to retrieve the next set of results for your
	// original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`

	// The maximum number of payments to return in a single response. This value cannot exceed 200.
	Limit int64 `json:"limit,omitempty"`

	// The order in which payments are listed in the response.
	// See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`
}

// Validate validates this v1 list orders request
func (m *V1ListOrdersRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 list orders request based on context it is used
func (m *V1ListOrdersRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListOrdersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListOrdersRequest) UnmarshalBinary(b []byte) error {
	var res V1ListOrdersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
