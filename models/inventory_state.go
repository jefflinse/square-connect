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

// InventoryState Indicates the state of a tracked item quantity in the lifecycle of goods.
//
// swagger:model InventoryState
type InventoryState string

const (

	// InventoryStateCUSTOM captures enum value "CUSTOM"
	InventoryStateCUSTOM InventoryState = "CUSTOM"

	// InventoryStateINSTOCK captures enum value "IN_STOCK"
	InventoryStateINSTOCK InventoryState = "IN_STOCK"

	// InventoryStateSOLD captures enum value "SOLD"
	InventoryStateSOLD InventoryState = "SOLD"

	// InventoryStateRETURNEDBYCUSTOMER captures enum value "RETURNED_BY_CUSTOMER"
	InventoryStateRETURNEDBYCUSTOMER InventoryState = "RETURNED_BY_CUSTOMER"

	// InventoryStateRESERVEDFORSALE captures enum value "RESERVED_FOR_SALE"
	InventoryStateRESERVEDFORSALE InventoryState = "RESERVED_FOR_SALE"

	// InventoryStateSOLDONLINE captures enum value "SOLD_ONLINE"
	InventoryStateSOLDONLINE InventoryState = "SOLD_ONLINE"

	// InventoryStateORDEREDFROMVENDOR captures enum value "ORDERED_FROM_VENDOR"
	InventoryStateORDEREDFROMVENDOR InventoryState = "ORDERED_FROM_VENDOR"

	// InventoryStateRECEIVEDFROMVENDOR captures enum value "RECEIVED_FROM_VENDOR"
	InventoryStateRECEIVEDFROMVENDOR InventoryState = "RECEIVED_FROM_VENDOR"

	// InventoryStateINTRANSITTO captures enum value "IN_TRANSIT_TO"
	InventoryStateINTRANSITTO InventoryState = "IN_TRANSIT_TO"

	// InventoryStateNONE captures enum value "NONE"
	InventoryStateNONE InventoryState = "NONE"

	// InventoryStateWASTE captures enum value "WASTE"
	InventoryStateWASTE InventoryState = "WASTE"

	// InventoryStateUNLINKEDRETURN captures enum value "UNLINKED_RETURN"
	InventoryStateUNLINKEDRETURN InventoryState = "UNLINKED_RETURN"
)

// for schema
var inventoryStateEnum []interface{}

func init() {
	var res []InventoryState
	if err := json.Unmarshal([]byte(`["CUSTOM","IN_STOCK","SOLD","RETURNED_BY_CUSTOMER","RESERVED_FOR_SALE","SOLD_ONLINE","ORDERED_FROM_VENDOR","RECEIVED_FROM_VENDOR","IN_TRANSIT_TO","NONE","WASTE","UNLINKED_RETURN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryStateEnum = append(inventoryStateEnum, v)
	}
}

func (m InventoryState) validateInventoryStateEnum(path, location string, value InventoryState) error {
	if err := validate.Enum(path, location, value, inventoryStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this inventory state
func (m InventoryState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInventoryStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}