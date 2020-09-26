// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Customer Represents a Square customer profile, which can have one or more
// cards on file associated with it.
//
// swagger:model Customer
type Customer struct {

	// The physical address associated with the customer profile.
	Address *Address `json:"address,omitempty"`

	// The birthday associated with the customer profile, in RFC-3339 format.
	// Year is optional, timezone and times are not allowed.
	// For example: `0000-09-01T00:00:00-00:00` indicates a birthday on September 1st.
	// `1998-09-01T00:00:00-00:00` indications a birthday on September 1st __1998__.
	Birthday string `json:"birthday,omitempty"`

	// Payment details of cards stored on file for the customer profile.
	Cards []*Card `json:"cards"`

	// A business name associated with the customer profile.
	CompanyName string `json:"company_name,omitempty"`

	// The timestamp when the customer profile was created, in RFC 3339 format.
	// Required: true
	CreatedAt *string `json:"created_at"`

	// A creation source represents the method used to create the
	// customer profile.
	// See [CustomerCreationSource](#type-customercreationsource) for possible values
	CreationSource string `json:"creation_source,omitempty"`

	// The email address associated with the customer profile.
	EmailAddress string `json:"email_address,omitempty"`

	// The family (i.e., last) name associated with the customer profile.
	FamilyName string `json:"family_name,omitempty"`

	// The given (i.e., first) name associated with the customer profile.
	GivenName string `json:"given_name,omitempty"`

	// The IDs of customer groups the customer belongs to.
	GroupIds []string `json:"group_ids"`

	// The customer groups and segments the customer belongs to. This deprecated field is replaced with dedicated `group_ids` for customer groups and `segment_ids` for customer segments.
	Groups []*CustomerGroupInfo `json:"groups"`

	// A unique Square-assigned ID for the customer profile.
	// Required: true
	ID *string `json:"id"`

	// A nickname for the customer profile.
	Nickname string `json:"nickname,omitempty"`

	// A custom note associated with the customer profile.
	Note string `json:"note,omitempty"`

	// The 11-digit phone number associated with the customer profile.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Represents general customer preferences.
	Preferences *CustomerPreferences `json:"preferences,omitempty"`

	// An optional, second ID used to associate the customer profile with an
	// entity in another system.
	ReferenceID string `json:"reference_id,omitempty"`

	// The IDs of segments the customer belongs to.
	SegmentIds []string `json:"segment_ids"`

	// The timestamp when the customer profile was last updated, in RFC 3339 format.
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this customer
func (m *Customer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Customer) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *Customer) validateCards(formats strfmt.Registry) error {

	if swag.IsZero(m.Cards) { // not required
		return nil
	}

	for i := 0; i < len(m.Cards); i++ {
		if swag.IsZero(m.Cards[i]) { // not required
			continue
		}

		if m.Cards[i] != nil {
			if err := m.Cards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Customer) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Customer) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validatePreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	if m.Preferences != nil {
		if err := m.Preferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

func (m *Customer) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Customer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customer) UnmarshalBinary(b []byte) error {
	var res Customer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}