// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchOrdersFilter Filtering criteria to use for a SearchOrders request. Multiple filters
// will be ANDed together.
//
// swagger:model SearchOrdersFilter
type SearchOrdersFilter struct {

	// Filter by customers associated with the order.
	CustomerFilter *SearchOrdersCustomerFilter `json:"customer_filter,omitempty"`

	// Filter for results within a time range.
	//
	// __Important:__ If you filter for orders by time range, you must set SearchOrdersSort
	// to sort by the same field.
	// [Learn more about filtering orders by time range](https://developer.squareup.com/docs/orders-api/manage-orders#important-note-on-filtering-orders-by-time-range)
	DateTimeFilter *SearchOrdersDateTimeFilter `json:"date_time_filter,omitempty"`

	// Filter by fulfillment type or state.
	FulfillmentFilter *SearchOrdersFulfillmentFilter `json:"fulfillment_filter,omitempty"`

	// Filter by source of order.
	SourceFilter *SearchOrdersSourceFilter `json:"source_filter,omitempty"`

	// Filter by ``OrderState``.
	StateFilter *SearchOrdersStateFilter `json:"state_filter,omitempty"`
}

// Validate validates this search orders filter
func (m *SearchOrdersFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTimeFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchOrdersFilter) validateCustomerFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerFilter) { // not required
		return nil
	}

	if m.CustomerFilter != nil {
		if err := m.CustomerFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) validateDateTimeFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.DateTimeFilter) { // not required
		return nil
	}

	if m.DateTimeFilter != nil {
		if err := m.DateTimeFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date_time_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) validateFulfillmentFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentFilter) { // not required
		return nil
	}

	if m.FulfillmentFilter != nil {
		if err := m.FulfillmentFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fulfillment_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) validateSourceFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceFilter) { // not required
		return nil
	}

	if m.SourceFilter != nil {
		if err := m.SourceFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) validateStateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.StateFilter) { // not required
		return nil
	}

	if m.StateFilter != nil {
		if err := m.StateFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_filter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search orders filter based on the context it is used
func (m *SearchOrdersFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomerFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateTimeFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchOrdersFilter) contextValidateCustomerFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomerFilter != nil {
		if err := m.CustomerFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) contextValidateDateTimeFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.DateTimeFilter != nil {
		if err := m.DateTimeFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date_time_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) contextValidateFulfillmentFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.FulfillmentFilter != nil {
		if err := m.FulfillmentFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fulfillment_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) contextValidateSourceFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceFilter != nil {
		if err := m.SourceFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_filter")
			}
			return err
		}
	}

	return nil
}

func (m *SearchOrdersFilter) contextValidateStateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.StateFilter != nil {
		if err := m.StateFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_filter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchOrdersFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchOrdersFilter) UnmarshalBinary(b []byte) error {
	var res SearchOrdersFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
