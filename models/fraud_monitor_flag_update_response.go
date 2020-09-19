// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FraudMonitorFlagUpdateResponse fraud monitor flag update response
//
// swagger:model FraudMonitorFlagUpdateResponse
type FraudMonitorFlagUpdateResponse struct {

	//  Max length = 19, The card number as received in the input message
	CardNo string `json:"cardNo,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Monitoring Expiry Date. The value as received in the input message (CCYYMMDD)
	FrdMonExpDte string `json:"frdMonExpDte,omitempty"`

	//  Max length = 2, Monitoring Flag. The value as received in the input message
	FrdMonFlg string `json:"frdMonFlg,omitempty"`
}

// Validate validates this fraud monitor flag update response
func (m *FraudMonitorFlagUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FraudMonitorFlagUpdateResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *FraudMonitorFlagUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FraudMonitorFlagUpdateResponse) UnmarshalBinary(b []byte) error {
	var res FraudMonitorFlagUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
