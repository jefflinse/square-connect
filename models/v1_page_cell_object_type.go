// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1PageCellObjectType v1 page cell object type
//
// swagger:model V1PageCellObjectType
type V1PageCellObjectType string

const (

	// V1PageCellObjectTypeITEM captures enum value "ITEM"
	V1PageCellObjectTypeITEM V1PageCellObjectType = "ITEM"

	// V1PageCellObjectTypeDISCOUNT captures enum value "DISCOUNT"
	V1PageCellObjectTypeDISCOUNT V1PageCellObjectType = "DISCOUNT"

	// V1PageCellObjectTypeCATEGORY captures enum value "CATEGORY"
	V1PageCellObjectTypeCATEGORY V1PageCellObjectType = "CATEGORY"

	// V1PageCellObjectTypePLACEHOLDER captures enum value "PLACEHOLDER"
	V1PageCellObjectTypePLACEHOLDER V1PageCellObjectType = "PLACEHOLDER"
)

// for schema
var v1PageCellObjectTypeEnum []interface{}

func init() {
	var res []V1PageCellObjectType
	if err := json.Unmarshal([]byte(`["ITEM","DISCOUNT","CATEGORY","PLACEHOLDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PageCellObjectTypeEnum = append(v1PageCellObjectTypeEnum, v)
	}
}

func (m V1PageCellObjectType) validateV1PageCellObjectTypeEnum(path, location string, value V1PageCellObjectType) error {
	if err := validate.Enum(path, location, value, v1PageCellObjectTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 page cell object type
func (m V1PageCellObjectType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1PageCellObjectTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}