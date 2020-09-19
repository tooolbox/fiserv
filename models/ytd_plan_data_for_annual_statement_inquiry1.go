// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// YtdPlanDataForAnnualStatementInquiry1 ytd plan data for annual statement inquiry1
//
// swagger:model YtdPlanDataForAnnualStatementInquiry1
type YtdPlanDataForAnnualStatementInquiry1 struct {

	//  Max length = 7, Annual year-to-date total Retail Interest amount and reversals for all six plans from the ACCS annual summary.
	YtdIntRtl string `json:"ytdIntRtl,omitempty"`

	//  Max length = 7, Annual year-to-date total overseas fee amount and reversals for all six plans from the ACCS annual summary.
	YtdOseasFee string `json:"ytdOseasFee,omitempty"`

	//  Max length = 7, Annual year-to-date total Principal amount and reversals for all six plans from the ACCS annual summary.
	YtdPrin string `json:"ytdPrin,omitempty"`

	//  Max length = 7, Annual year-to-date total other Principal amount and reversals for all six plans from the ACCS annual summary.
	YtdPrinOth string `json:"ytdPrinOth,omitempty"`

	//  Max length = 7, Annual year-to-date total SVC amount and reversals for all six plans from the ACCS annual summary.
	YtdSvc string `json:"ytdSvc,omitempty"`
}

// Validate validates this ytd plan data for annual statement inquiry1
func (m *YtdPlanDataForAnnualStatementInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *YtdPlanDataForAnnualStatementInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *YtdPlanDataForAnnualStatementInquiry1) UnmarshalBinary(b []byte) error {
	var res YtdPlanDataForAnnualStatementInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
