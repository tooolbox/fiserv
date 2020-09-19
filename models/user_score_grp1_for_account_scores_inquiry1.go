// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserScoreGrp1ForAccountScoresInquiry1 user score grp1 for account scores inquiry1
//
// swagger:model UserScoreGrp1ForAccountScoresInquiry1
type UserScoreGrp1ForAccountScoresInquiry1 struct {

	//  Max length = 2, User Score 2 to 10. User score to be used as a decision key in TRIAD batch and online
	ScrGrp1Fld string `json:"scrGrp1Fld,omitempty"`
}

// Validate validates this user score grp1 for account scores inquiry1
func (m *UserScoreGrp1ForAccountScoresInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserScoreGrp1ForAccountScoresInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserScoreGrp1ForAccountScoresInquiry1) UnmarshalBinary(b []byte) error {
	var res UserScoreGrp1ForAccountScoresInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
