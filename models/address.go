// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Address Represents a physical address.
//
// swagger:model Address
type Address struct {

	// The first line of the address.
	//
	// Fields that start with `address_line` provide the address's most specific
	// details, like street number, street name, and building name. They do *not*
	// provide less specific details like city, state/province, or country (these
	// details are provided in other fields).
	AddressLine1 string `json:"address_line_1,omitempty"`

	// The second line of the address, if any.
	AddressLine2 string `json:"address_line_2,omitempty"`

	// The third line of the address, if any.
	AddressLine3 string `json:"address_line_3,omitempty"`

	// A civil entity within the address's country. In the US, this
	// is the state.
	AdministrativeDistrictLevel1 string `json:"administrative_district_level_1,omitempty"`

	// A civil entity within the address's `administrative_district_level_1`.
	// In the US, this is the county.
	AdministrativeDistrictLevel2 string `json:"administrative_district_level_2,omitempty"`

	// A civil entity within the address's `administrative_district_level_2`,
	// if any.
	AdministrativeDistrictLevel3 string `json:"administrative_district_level_3,omitempty"`

	// The address's country, in ISO 3166-1-alpha-2 format.
	// See [Country](#type-country) for possible values
	Country string `json:"country,omitempty"`

	// Optional first name when it's representing recipient.
	FirstName string `json:"first_name,omitempty"`

	// Optional last name when it's representing recipient.
	LastName string `json:"last_name,omitempty"`

	// The city or town of the address.
	Locality string `json:"locality,omitempty"`

	// Optional organization name when it's representing recipient.
	Organization string `json:"organization,omitempty"`

	// The address's postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// A civil region within the address's `locality`, if any.
	Sublocality string `json:"sublocality,omitempty"`

	// A civil region within the address's `sublocality`, if any.
	Sublocality2 string `json:"sublocality_2,omitempty"`

	// A civil region within the address's `sublocality_2`, if any.
	Sublocality3 string `json:"sublocality_3,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}