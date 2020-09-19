// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDisabilityIndicatorRequest update disability indicator request
//
// swagger:model UpdateDisabilityIndicatorRequest
type UpdateDisabilityIndicatorRequest struct {

	//  Max length = 19, CARD NBR
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr,omitempty"`

	//  Max length = 4, Identification number assigned to an Embosser record to distinguish between multiple cards issued with the same card
	// Required: true
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSeq *string `json:"cardSeq"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, The disability indicator is used to identify customers that have a requirement for correspondence to be sent in a particular format. The valid values are :                          AUDIO - 'A' BRAILLE - 'B'.  HARD-OF-HERING - 'H'.  LARGE-PRINT - 'L'.  CDROM - 'C'.  PLAIN-PRINT - 'P'.
	// Required: true
	// Max Length: 1
	// Min Length: 0
	DisabilityInd *string `json:"disabilityInd"`
}

// Validate validates this update disability indicator request
func (m *UpdateDisabilityIndicatorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabilityInd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDisabilityIndicatorRequest) validateCardNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CardNbr) { // not required
		return nil
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDisabilityIndicatorRequest) validateCardSeq(formats strfmt.Registry) error {

	if err := validate.Required("cardSeq", "body", m.CardSeq); err != nil {
		return err
	}

	if err := validate.MinLength("cardSeq", "body", string(*m.CardSeq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardSeq", "body", string(*m.CardSeq), 4); err != nil {
		return err
	}

	if err := validate.Pattern("cardSeq", "body", string(*m.CardSeq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDisabilityIndicatorRequest) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(m.Common) { // not required
		return nil
	}

	if m.Common != nil {
		if err := m.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("common")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateDisabilityIndicatorRequest) validateDisabilityInd(formats strfmt.Registry) error {

	if err := validate.Required("disabilityInd", "body", m.DisabilityInd); err != nil {
		return err
	}

	if err := validate.MinLength("disabilityInd", "body", string(*m.DisabilityInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("disabilityInd", "body", string(*m.DisabilityInd), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDisabilityIndicatorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDisabilityIndicatorRequest) UnmarshalBinary(b []byte) error {
	var res UpdateDisabilityIndicatorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
