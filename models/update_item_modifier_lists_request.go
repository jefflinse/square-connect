// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateItemModifierListsRequest update item modifier lists request
// Example: {"request_body":{"item_ids":["H42BRLUJ5KTZTTMPVSLFAACQ","2JXOBJIHCWBQ4NZ3RIXQGJA6"],"modifier_lists_to_disable":["7WRC16CJZDVLSNDQ35PP6YAD"],"modifier_lists_to_enable":["H42BRLUJ5KTZTTMPVSLFAACQ","2JXOBJIHCWBQ4NZ3RIXQGJA6"]}}
//
// swagger:model UpdateItemModifierListsRequest
type UpdateItemModifierListsRequest struct {

	// The IDs of the catalog items associated with the CatalogModifierList objects being updated.
	// Required: true
	ItemIds []string `json:"item_ids"`

	// The IDs of the CatalogModifierList objects to disable for the CatalogItem.
	ModifierListsToDisable []string `json:"modifier_lists_to_disable"`

	// The IDs of the CatalogModifierList objects to enable for the CatalogItem.
	ModifierListsToEnable []string `json:"modifier_lists_to_enable"`
}

// Validate validates this update item modifier lists request
func (m *UpdateItemModifierListsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateItemModifierListsRequest) validateItemIds(formats strfmt.Registry) error {

	if err := validate.Required("item_ids", "body", m.ItemIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update item modifier lists request based on context it is used
func (m *UpdateItemModifierListsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateItemModifierListsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateItemModifierListsRequest) UnmarshalBinary(b []byte) error {
	var res UpdateItemModifierListsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
