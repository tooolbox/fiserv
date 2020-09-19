// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReasonCodeDataForInstantIssuance1 reason code data for instant issuance1
//
// swagger:model ReasonCodeDataForInstantIssuance1
type ReasonCodeDataForInstantIssuance1 struct {

	//  Max length = 10, REASON CODE
	ReasonCode string `json:"reasonCode,omitempty"`
}

// Validate validates this reason code data for instant issuance1
func (m *ReasonCodeDataForInstantIssuance1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReasonCodeDataForInstantIssuance1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReasonCodeDataForInstantIssuance1) UnmarshalBinary(b []byte) error {
	var res ReasonCodeDataForInstantIssuance1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
