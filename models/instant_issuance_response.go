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

// InstantIssuanceResponse instant issuance response
//
// swagger:model InstantIssuanceResponse
type InstantIssuanceResponse struct {

	//  Max length = 19, Card number as received in the input message
	CardNbr string `json:"cardNbr,omitempty"`

	//  Max length = 4, Card sequence as received in the input message
	CardSeq string `json:"cardSeq,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	// REASON CODE DATA
	ReasonCodeData []*ReasonCodeDataForInstantIssuance1 `json:"reasonCodeData"`
}

// Validate validates this instant issuance response
func (m *InstantIssuanceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasonCodeData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstantIssuanceResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *InstantIssuanceResponse) validateReasonCodeData(formats strfmt.Registry) error {

	if swag.IsZero(m.ReasonCodeData) { // not required
		return nil
	}

	for i := 0; i < len(m.ReasonCodeData); i++ {
		if swag.IsZero(m.ReasonCodeData[i]) { // not required
			continue
		}

		if m.ReasonCodeData[i] != nil {
			if err := m.ReasonCodeData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasonCodeData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstantIssuanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstantIssuanceResponse) UnmarshalBinary(b []byte) error {
	var res InstantIssuanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
