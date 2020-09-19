// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CardTransferUpdateResponse card transfer update response
//
// swagger:model CardTransferUpdateResponse
type CardTransferUpdateResponse struct {

	//  Max length = 1, BLOCK CODE
	BlockCode string `json:"blockCode,omitempty"`

	//  Max length = 19, Unique Card Number embossed on card plastic.
	CardNbr string `json:"cardNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 8, Effective Date of the card transfer.  A value is pre-populated in this field only when the transfer FUNCTION is R (reversal) and the card transfer was generated earlier in the same processing day.
	EffDate string `json:"effDate,omitempty"`

	//  Max length = 19, Transfer To Fraud Account: Fraud Account number to which the card was or is transferred today.  A value is pre-populated in this field only when the Transfer FUNCTION is R (reversal) and the card transfer was generated earlier in the same processing day.
	XfrToAcct string `json:"xfrToAcct,omitempty"`

	//  Max length = 19, Transfer-to Customer Number: Customer Name/Address Record associated with the Fraud Account number to which the card was or is transferred today.  A value is pre-populated in this field only when the transfer FUNCTION is R (reversal) and the card transfer was generated earlier in the same processing day.
	XfrToCust string `json:"xfrToCust,omitempty"`
}

// Validate validates this card transfer update response
func (m *CardTransferUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardTransferUpdateResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *CardTransferUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardTransferUpdateResponse) UnmarshalBinary(b []byte) error {
	var res CardTransferUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
