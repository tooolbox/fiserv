// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentHistoryInquiryResponse payment history inquiry response
//
// swagger:model PaymentHistoryInquiryResponse
type PaymentHistoryInquiryResponse struct {

	//  Max length = 19, Account Number of the Customers account or the Customers Card Number. Must be numeric and greater than zero.
	AcctNbr string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Current balance: Total current balance of all credit plan segments associated with the account.
	CurrBal string `json:"currBal,omitempty"`

	//  Max length = 3, CURR CODE
	CurrCode string `json:"currCode,omitempty"`

	//  Max length = 1, CURR NOD
	CurrNod string `json:"currNod,omitempty"`

	//  Max length = 1, FOREIGN USE
	ForeignUse string `json:"foreignUse,omitempty"`

	//  Max length = 3, LOGO
	Logo string `json:"logo,omitempty"`

	//  Max length = 2, Valid Values:  0-6
	NbrOfPayments string `json:"nbrOfPayments,omitempty"`

	// OCCURS SIX (6) TIMES
	PmtHistoryTableData []*PmtHistoryTableDataForPaymentHistoryInquiry1 `json:"pmtHistoryTableData"`
}

// Validate validates this payment history inquiry response
func (m *PaymentHistoryInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmtHistoryTableData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentHistoryInquiryResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *PaymentHistoryInquiryResponse) validatePmtHistoryTableData(formats strfmt.Registry) error {

	if swag.IsZero(m.PmtHistoryTableData) { // not required
		return nil
	}

	for i := 0; i < len(m.PmtHistoryTableData); i++ {
		if swag.IsZero(m.PmtHistoryTableData[i]) { // not required
			continue
		}

		if m.PmtHistoryTableData[i] != nil {
			if err := m.PmtHistoryTableData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pmtHistoryTableData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentHistoryInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentHistoryInquiryResponse) UnmarshalBinary(b []byte) error {
	var res PaymentHistoryInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
