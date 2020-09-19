// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstantIssuanceRestartResponse instant issuance restart response
//
// swagger:model InstantIssuanceRestartResponse
type InstantIssuanceRestartResponse struct {

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 19, Card number as received in the input message.
	CrdNbr string `json:"crdNbr,omitempty"`

	//  Max length = 4, Card sequence as received in the input message.
	CrdSeq string `json:"crdSeq,omitempty"`
}

// Validate validates this instant issuance restart response
func (m *InstantIssuanceRestartResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstantIssuanceRestartResponse) validateCommon(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *InstantIssuanceRestartResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstantIssuanceRestartResponse) UnmarshalBinary(b []byte) error {
	var res InstantIssuanceRestartResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
