// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SearchCatalogItemsRequestStockLevel Defines supported stock levels of the item inventory.
//
// swagger:model SearchCatalogItemsRequestStockLevel
type SearchCatalogItemsRequestStockLevel string

const (

	// SearchCatalogItemsRequestStockLevelOUT captures enum value "OUT"
	SearchCatalogItemsRequestStockLevelOUT SearchCatalogItemsRequestStockLevel = "OUT"

	// SearchCatalogItemsRequestStockLevelLOW captures enum value "LOW"
	SearchCatalogItemsRequestStockLevelLOW SearchCatalogItemsRequestStockLevel = "LOW"
)

// for schema
var searchCatalogItemsRequestStockLevelEnum []interface{}

func init() {
	var res []SearchCatalogItemsRequestStockLevel
	if err := json.Unmarshal([]byte(`["OUT","LOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchCatalogItemsRequestStockLevelEnum = append(searchCatalogItemsRequestStockLevelEnum, v)
	}
}

func (m SearchCatalogItemsRequestStockLevel) validateSearchCatalogItemsRequestStockLevelEnum(path, location string, value SearchCatalogItemsRequestStockLevel) error {
	if err := validate.EnumCase(path, location, value, searchCatalogItemsRequestStockLevelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this search catalog items request stock level
func (m SearchCatalogItemsRequestStockLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSearchCatalogItemsRequestStockLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this search catalog items request stock level based on context it is used
func (m SearchCatalogItemsRequestStockLevel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
