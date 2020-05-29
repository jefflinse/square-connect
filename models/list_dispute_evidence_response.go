// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListDisputeEvidenceResponse Defines fields in a ListDisputeEvidence response.
//
// swagger:model ListDisputeEvidenceResponse
type ListDisputeEvidenceResponse struct {

	// Information on errors encountered during the request.
	Errors []*Error `json:"errors"`

	// The list of evidence previously uploaded to the specified dispute.
	Evidence []*DisputeEvidence `json:"evidence"`
}

// Validate validates this list dispute evidence response
func (m *ListDisputeEvidenceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvidence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDisputeEvidenceResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListDisputeEvidenceResponse) validateEvidence(formats strfmt.Registry) error {

	if swag.IsZero(m.Evidence) { // not required
		return nil
	}

	for i := 0; i < len(m.Evidence); i++ {
		if swag.IsZero(m.Evidence[i]) { // not required
			continue
		}

		if m.Evidence[i] != nil {
			if err := m.Evidence[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evidence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDisputeEvidenceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDisputeEvidenceResponse) UnmarshalBinary(b []byte) error {
	var res ListDisputeEvidenceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
