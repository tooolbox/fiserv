// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CollectionsPromiseToPayAddResponse collections promise to pay add response
//
// swagger:model CollectionsPromiseToPayAddResponse
type CollectionsPromiseToPayAddResponse struct {

	//  Max length = 19, Account number populated from the input data.
	AcctNbr string `json:"acctNbr,omitempty"`

	//  Max length = 2, Source application code.
	ApplNbr string `json:"applNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 50, Free-form notation added to this account.
	MemoMsg string `json:"memoMsg,omitempty"`

	//  Max length = 4, PTP ACTION CD
	PtpActionCd string `json:"ptpActionCd,omitempty"`

	//  Max length = 17, Promise-To-Pay Amount: Next payment amount for a promise-to-pay.
	PtpAmt string `json:"ptpAmt,omitempty"`

	// Format: YYYYMMDD. PTP DT
	PtpDt string `json:"ptpDt,omitempty"`

	// Format: YYYYMMDD. REVIEW DT
	ReviewDt string `json:"reviewDt,omitempty"`

	//  Max length = 4, Review Time: Scheduled time of day when the collector will next review the account.
	ReviewTime string `json:"reviewTime,omitempty"`
}

// Validate validates this collections promise to pay add response
func (m *CollectionsPromiseToPayAddResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectionsPromiseToPayAddResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *CollectionsPromiseToPayAddResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectionsPromiseToPayAddResponse) UnmarshalBinary(b []byte) error {
	var res CollectionsPromiseToPayAddResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
