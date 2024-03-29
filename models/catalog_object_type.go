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

// CatalogObjectType Possible types of CatalogObjects returned from the Catalog, each
// containing type-specific properties in the `*_data` field corresponding to the object type.
//
// swagger:model CatalogObjectType
type CatalogObjectType string

const (

	// CatalogObjectTypeITEM captures enum value "ITEM"
	CatalogObjectTypeITEM CatalogObjectType = "ITEM"

	// CatalogObjectTypeIMAGE captures enum value "IMAGE"
	CatalogObjectTypeIMAGE CatalogObjectType = "IMAGE"

	// CatalogObjectTypeCATEGORY captures enum value "CATEGORY"
	CatalogObjectTypeCATEGORY CatalogObjectType = "CATEGORY"

	// CatalogObjectTypeITEMVARIATION captures enum value "ITEM_VARIATION"
	CatalogObjectTypeITEMVARIATION CatalogObjectType = "ITEM_VARIATION"

	// CatalogObjectTypeTAX captures enum value "TAX"
	CatalogObjectTypeTAX CatalogObjectType = "TAX"

	// CatalogObjectTypeDISCOUNT captures enum value "DISCOUNT"
	CatalogObjectTypeDISCOUNT CatalogObjectType = "DISCOUNT"

	// CatalogObjectTypeMODIFIERLIST captures enum value "MODIFIER_LIST"
	CatalogObjectTypeMODIFIERLIST CatalogObjectType = "MODIFIER_LIST"

	// CatalogObjectTypeMODIFIER captures enum value "MODIFIER"
	CatalogObjectTypeMODIFIER CatalogObjectType = "MODIFIER"

	// CatalogObjectTypePRICINGRULE captures enum value "PRICING_RULE"
	CatalogObjectTypePRICINGRULE CatalogObjectType = "PRICING_RULE"

	// CatalogObjectTypePRODUCTSET captures enum value "PRODUCT_SET"
	CatalogObjectTypePRODUCTSET CatalogObjectType = "PRODUCT_SET"

	// CatalogObjectTypeTIMEPERIOD captures enum value "TIME_PERIOD"
	CatalogObjectTypeTIMEPERIOD CatalogObjectType = "TIME_PERIOD"

	// CatalogObjectTypeMEASUREMENTUNIT captures enum value "MEASUREMENT_UNIT"
	CatalogObjectTypeMEASUREMENTUNIT CatalogObjectType = "MEASUREMENT_UNIT"

	// CatalogObjectTypeSUBSCRIPTIONPLAN captures enum value "SUBSCRIPTION_PLAN"
	CatalogObjectTypeSUBSCRIPTIONPLAN CatalogObjectType = "SUBSCRIPTION_PLAN"

	// CatalogObjectTypeITEMOPTION captures enum value "ITEM_OPTION"
	CatalogObjectTypeITEMOPTION CatalogObjectType = "ITEM_OPTION"

	// CatalogObjectTypeITEMOPTIONVAL captures enum value "ITEM_OPTION_VAL"
	CatalogObjectTypeITEMOPTIONVAL CatalogObjectType = "ITEM_OPTION_VAL"

	// CatalogObjectTypeCUSTOMATTRIBUTEDEFINITION captures enum value "CUSTOM_ATTRIBUTE_DEFINITION"
	CatalogObjectTypeCUSTOMATTRIBUTEDEFINITION CatalogObjectType = "CUSTOM_ATTRIBUTE_DEFINITION"

	// CatalogObjectTypeQUICKAMOUNTSSETTINGS captures enum value "QUICK_AMOUNTS_SETTINGS"
	CatalogObjectTypeQUICKAMOUNTSSETTINGS CatalogObjectType = "QUICK_AMOUNTS_SETTINGS"
)

// for schema
var catalogObjectTypeEnum []interface{}

func init() {
	var res []CatalogObjectType
	if err := json.Unmarshal([]byte(`["ITEM","IMAGE","CATEGORY","ITEM_VARIATION","TAX","DISCOUNT","MODIFIER_LIST","MODIFIER","PRICING_RULE","PRODUCT_SET","TIME_PERIOD","MEASUREMENT_UNIT","SUBSCRIPTION_PLAN","ITEM_OPTION","ITEM_OPTION_VAL","CUSTOM_ATTRIBUTE_DEFINITION","QUICK_AMOUNTS_SETTINGS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogObjectTypeEnum = append(catalogObjectTypeEnum, v)
	}
}

func (m CatalogObjectType) validateCatalogObjectTypeEnum(path, location string, value CatalogObjectType) error {
	if err := validate.EnumCase(path, location, value, catalogObjectTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this catalog object type
func (m CatalogObjectType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCatalogObjectTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this catalog object type based on context it is used
func (m CatalogObjectType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
