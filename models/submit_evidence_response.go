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
)

// SubmitEvidenceResponse Defines the fields in a `SubmitEvidence` response.
// Example: {"dispute":{"amount_money":{"amount":2000,"currency":"USD"},"brand_dispute_id":"100000399240","card_brand":"visa","created_at":"2018-10-18T16:02:15.313Z","dispute_id":"EAZoK70gX3fyvibecLwIGB","disputed_payments":[{"payment_id":"2yeBUWJzllJTpmnSqtMRAL19taB"}],"due_at":"2018-11-01T00:00:00.000Z","evidence_ids":["CKWRhnZN0eMSUbh38BKmD"],"reason":"NO_KNOWLEDGE","state":"PROCESSING","updated_at":"2018-10-18T16:02:15.313Z"}}
//
// swagger:model SubmitEvidenceResponse
type SubmitEvidenceResponse struct {

	// The `Dispute` for which evidence was submitted.
	Dispute *Dispute `json:"dispute,omitempty"`

	// Information about errors encountered during the request.
	Errors []*Error `json:"errors"`
}

// Validate validates this submit evidence response
func (m *SubmitEvidenceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDispute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitEvidenceResponse) validateDispute(formats strfmt.Registry) error {
	if swag.IsZero(m.Dispute) { // not required
		return nil
	}

	if m.Dispute != nil {
		if err := m.Dispute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dispute")
			}
			return err
		}
	}

	return nil
}

func (m *SubmitEvidenceResponse) validateErrors(formats strfmt.Registry) error {
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

// ContextValidate validate this submit evidence response based on the context it is used
func (m *SubmitEvidenceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDispute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitEvidenceResponse) contextValidateDispute(ctx context.Context, formats strfmt.Registry) error {

	if m.Dispute != nil {
		if err := m.Dispute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dispute")
			}
			return err
		}
	}

	return nil
}

func (m *SubmitEvidenceResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubmitEvidenceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitEvidenceResponse) UnmarshalBinary(b []byte) error {
	var res SubmitEvidenceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
