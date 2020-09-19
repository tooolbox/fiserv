// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpndLmtsStatsForCardInquiry1 spnd lmts stats for card inquiry1
//
// swagger:model SpndLmtsStatsForCardInquiry1
type SpndLmtsStatsForCardInquiry1 struct {

	//  Max length = 17, Total monetary amount of the transactions for the spending category for this card cycle-to date
	SpndAmtCtd string `json:"spndAmtCtd,omitempty"`

	//  Max length = 17, Total monetary amount of the transactions for the spending category for this card life-to-date
	SpndAmtLtd string `json:"spndAmtLtd,omitempty"`

	//  Max length = 17, Total monetary amount of the transactions for the spending category for this card today
	SpndAmtToday string `json:"spndAmtToday,omitempty"`

	//  Max length = 17, Total monetary amount of the transactions for the spending category for this card week-to-date
	SpndAmtWtd string `json:"spndAmtWtd,omitempty"`

	//  Max length = 17, Total monetary amount of the transactions for the spending category for this card year-to-date
	SpndAmtYtd string `json:"spndAmtYtd,omitempty"`

	//  Max length = 9, Total number of transactions for the spending category for this Card cycle-to date
	SpndNbrCtd string `json:"spndNbrCtd,omitempty"`

	//  Max length = 9, Total number of transactions for the spending category for this Card life-to-date
	SpndNbrLtd string `json:"spndNbrLtd,omitempty"`

	//  Max length = 9, Total number of transactions for the spending category for this Card today
	SpndNbrToday string `json:"spndNbrToday,omitempty"`

	//  Max length = 9, Total number of transactions for the spending category for this Card week-to-date
	SpndNbrWtd string `json:"spndNbrWtd,omitempty"`

	//  Max length = 9, Total number of transactions for the spending category for this Card year-to-date
	SpndNbrYtd string `json:"spndNbrYtd,omitempty"`
}

// Validate validates this spnd lmts stats for card inquiry1
func (m *SpndLmtsStatsForCardInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpndLmtsStatsForCardInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpndLmtsStatsForCardInquiry1) UnmarshalBinary(b []byte) error {
	var res SpndLmtsStatsForCardInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
