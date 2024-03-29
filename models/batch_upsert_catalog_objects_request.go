// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BatchUpsertCatalogObjectsRequest batch upsert catalog objects request
// Example: {"request_body":{"batches":[{"objects":[{"id":"#Tea","item_data":{"category_id":"#Beverages","description":"Hot Leaf Juice","name":"Tea","tax_ids":["#SalesTax"],"variations":[{"id":"#Tea_Mug","item_variation_data":{"item_id":"#Tea","name":"Mug","price_money":{"amount":150,"currency":"USD"},"pricing_type":"FIXED_PRICING"},"present_at_all_locations":true,"type":"ITEM_VARIATION"}]},"present_at_all_locations":true,"type":"ITEM"},{"id":"#Coffee","item_data":{"category_id":"#Beverages","description":"Hot Bean Juice","name":"Coffee","tax_ids":["#SalesTax"],"variations":[{"id":"#Coffee_Regular","item_variation_data":{"item_id":"#Coffee","name":"Regular","price_money":{"amount":250,"currency":"USD"},"pricing_type":"FIXED_PRICING"},"present_at_all_locations":true,"type":"ITEM_VARIATION"},{"id":"#Coffee_Large","item_variation_data":{"item_id":"#Coffee","name":"Large","price_money":{"amount":350,"currency":"USD"},"pricing_type":"FIXED_PRICING"},"present_at_all_locations":true,"type":"ITEM_VARIATION"}]},"present_at_all_locations":true,"type":"ITEM"},{"category_data":{"name":"Beverages"},"id":"#Beverages","present_at_all_locations":true,"type":"CATEGORY"},{"id":"#SalesTax","present_at_all_locations":true,"tax_data":{"applies_to_custom_amounts":true,"calculation_phase":"TAX_SUBTOTAL_PHASE","enabled":true,"inclusion_type":"ADDITIVE","name":"Sales Tax","percentage":"5.0"},"type":"TAX"}]}],"idempotency_key":"789ff020-f723-43a9-b4b5-43b5dc1fa3dc"}}
//
// swagger:model BatchUpsertCatalogObjectsRequest
type BatchUpsertCatalogObjectsRequest struct {

	// A batch of CatalogObjects to be inserted/updated atomically.
	// The objects within a batch will be inserted in an all-or-nothing fashion, i.e., if an error occurs
	// attempting to insert or update an object within a batch, the entire batch will be rejected. However, an error
	// in one batch will not affect other batches within the same request.
	//
	// For each object, its `updated_at` field is ignored and replaced with a current [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates), and its
	// `is_deleted` field must not be set to `true`.
	//
	// To modify an existing object, supply its ID. To create a new object, use an ID starting
	// with `#`. These IDs may be used to create relationships between an object and attributes of
	// other objects that reference it. For example, you can create a CatalogItem with
	// ID `#ABC` and a CatalogItemVariation with its `item_id` attribute set to
	// `#ABC` in order to associate the CatalogItemVariation with its parent
	// CatalogItem.
	//
	// Any `#`-prefixed IDs are valid only within a single atomic batch, and will be replaced by server-generated IDs.
	//
	// Each batch may contain up to 1,000 objects. The total number of objects across all batches for a single request
	// may not exceed 10,000. If either of these limits is violated, an error will be returned and no objects will
	// be inserted or updated.
	// Required: true
	Batches []*CatalogObjectBatch `json:"batches"`

	// A value you specify that uniquely identifies this
	// request among all your requests. A common way to create
	// a valid idempotency key is to use a Universally unique
	// identifier (UUID).
	//
	// If you're unsure whether a particular request was successful,
	// you can reattempt it with the same idempotency key without
	// worrying about creating duplicate objects.
	//
	// See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	// Required: true
	// Min Length: 1
	IdempotencyKey *string `json:"idempotency_key"`
}

// Validate validates this batch upsert catalog objects request
func (m *BatchUpsertCatalogObjectsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBatches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdempotencyKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchUpsertCatalogObjectsRequest) validateBatches(formats strfmt.Registry) error {

	if err := validate.Required("batches", "body", m.Batches); err != nil {
		return err
	}

	for i := 0; i < len(m.Batches); i++ {
		if swag.IsZero(m.Batches[i]) { // not required
			continue
		}

		if m.Batches[i] != nil {
			if err := m.Batches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("batches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BatchUpsertCatalogObjectsRequest) validateIdempotencyKey(formats strfmt.Registry) error {

	if err := validate.Required("idempotency_key", "body", m.IdempotencyKey); err != nil {
		return err
	}

	if err := validate.MinLength("idempotency_key", "body", *m.IdempotencyKey, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this batch upsert catalog objects request based on the context it is used
func (m *BatchUpsertCatalogObjectsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBatches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchUpsertCatalogObjectsRequest) contextValidateBatches(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Batches); i++ {

		if m.Batches[i] != nil {
			if err := m.Batches[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("batches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchUpsertCatalogObjectsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchUpsertCatalogObjectsRequest) UnmarshalBinary(b []byte) error {
	var res BatchUpsertCatalogObjectsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
