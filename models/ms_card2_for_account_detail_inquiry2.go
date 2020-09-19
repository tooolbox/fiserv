// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsCard2ForAccountDetailInquiry2 ms card2 for account detail inquiry2
//
// swagger:model MsCard2ForAccountDetailInquiry2
type MsCard2ForAccountDetailInquiry2 struct {

	//  Max length = 3, Multi Scheme Card Logo
	MsCardLogo string `json:"msCardLogo,omitempty"`

	//  Max length = 5, Multi Scheme Card Active
	MsCardNbrActive string `json:"msCardNbrActive,omitempty"`
}

// Validate validates this ms card2 for account detail inquiry2
func (m *MsCard2ForAccountDetailInquiry2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsCard2ForAccountDetailInquiry2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsCard2ForAccountDetailInquiry2) UnmarshalBinary(b []byte) error {
	var res MsCard2ForAccountDetailInquiry2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
