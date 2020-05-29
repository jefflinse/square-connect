// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DeletePageCellRequest v1 delete page cell request
//
// swagger:model V1DeletePageCellRequest
type V1DeletePageCellRequest struct {

	// The column of the cell to clear. Always an integer between 0 and 4, inclusive. Column 0 is the leftmost column.
	Column string `json:"column,omitempty"`

	// The row of the cell to clear. Always an integer between 0 and 4, inclusive. Row 0 is the top row.
	Row string `json:"row,omitempty"`
}

// Validate validates this v1 delete page cell request
func (m *V1DeletePageCellRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DeletePageCellRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeletePageCellRequest) UnmarshalBinary(b []byte) error {
	var res V1DeletePageCellRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
