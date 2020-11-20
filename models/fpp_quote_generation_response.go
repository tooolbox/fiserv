// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FppQuoteGenerationResponse fpp quote generation response
//
// swagger:model FppQuoteGenerationResponse
type FppQuoteGenerationResponse struct {

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 7, Loan AER. AER rate on flexible payment plan
	FppCurrAer string `json:"fppCurrAer,omitempty"`

	//  Max length = 17, Loan total interests. Total estimated Interest amount that customer would have paid by the end of the  TERM
	FppEstInt string `json:"fppEstInt,omitempty"`

	//  Max length = 17, Installment Amount. This is the fixed instalment amount that customer will have to pay.
	FppFixedPmtAmt string `json:"fppFixedPmtAmt,omitempty"`

	//  Max length = 7, Monthly Rate. Simple monthly rate of flexible payment plan
	FppMthlyRate string `json:"fppMthlyRate,omitempty"`

	//  Max length = 3, Number of Payments. Total number of payment that customer will have to make to payout the FPP balance
	FppNbrOfPmts string `json:"fppNbrOfPmts,omitempty"`

	// Format: YYYYMMDD. Start Date. Date on which the Credit Plan Segment record was opened, or generated for the account This is calculated as account accrue through date + 1  CCYYMMDD format
	FppPlanStartDate string `json:"fppPlanStartDate,omitempty"`

	//  Max length = 17, Loan total Cost. Total amount that customer would have paid by the end of the TERM
	FppTotalCost string `json:"fppTotalCost,omitempty"`
}

// Validate validates this fpp quote generation response
func (m *FppQuoteGenerationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FppQuoteGenerationResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *FppQuoteGenerationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FppQuoteGenerationResponse) UnmarshalBinary(b []byte) error {
	var res FppQuoteGenerationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}