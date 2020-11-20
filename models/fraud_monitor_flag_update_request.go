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

// FraudMonitorFlagUpdateRequest fraud monitor flag update request
//
// swagger:model FraudMonitorFlagUpdateRequest
type FraudMonitorFlagUpdateRequest struct {

	//  Max length = 19, The Card Number for this request.  Unique Card Number embossed on card plastic.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNum *string `json:"cardNum"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. The date the Flag is to be removed from the card (CCYYMMDD)
	FrdMonExpDte string `json:"frdMonExpDte,omitempty"`

	//  Max length = 2, The value to set for the Fraud Monitor Flag
	// Max Length: 2
	// Min Length: 0
	FrdMonFlg *string `json:"frdMonFlg,omitempty"`

	//  Max length = 4, Embosser Sequence number. Card sequence as received in the input message  If not provided, a value of '0001' is assumed.
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	SeqNo *string `json:"seqNo,omitempty"`
}

// Validate validates this fraud monitor flag update request
func (m *FraudMonitorFlagUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrdMonFlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeqNo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FraudMonitorFlagUpdateRequest) validateCardNum(formats strfmt.Registry) error {

	if err := validate.Required("cardNum", "body", m.CardNum); err != nil {
		return err
	}

	if err := validate.MinLength("cardNum", "body", string(*m.CardNum), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNum", "body", string(*m.CardNum), 19); err != nil {
		return err
	}

	return nil
}

func (m *FraudMonitorFlagUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *FraudMonitorFlagUpdateRequest) validateFrdMonFlg(formats strfmt.Registry) error {

	if swag.IsZero(m.FrdMonFlg) { // not required
		return nil
	}

	if err := validate.MinLength("frdMonFlg", "body", string(*m.FrdMonFlg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("frdMonFlg", "body", string(*m.FrdMonFlg), 2); err != nil {
		return err
	}

	return nil
}

func (m *FraudMonitorFlagUpdateRequest) validateSeqNo(formats strfmt.Registry) error {

	if swag.IsZero(m.SeqNo) { // not required
		return nil
	}

	if err := validate.MinLength("seqNo", "body", string(*m.SeqNo), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("seqNo", "body", string(*m.SeqNo), 4); err != nil {
		return err
	}

	if err := validate.Pattern("seqNo", "body", string(*m.SeqNo), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FraudMonitorFlagUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FraudMonitorFlagUpdateRequest) UnmarshalBinary(b []byte) error {
	var res FraudMonitorFlagUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}