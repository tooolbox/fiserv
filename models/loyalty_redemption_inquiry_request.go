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

// LoyaltyRedemptionInquiryRequest loyalty redemption inquiry request
//
// swagger:model LoyaltyRedemptionInquiryRequest
type LoyaltyRedemptionInquiryRequest struct {

	//  Max length = 19, LMS Account number involved in Points Scheme
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr,omitempty"`

	// Format: YYYYMMDD. Date of Batch Number 99996 on which the inquiry is made
	BtchDte string `json:"btchDte,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 19, Card number used to obtain associated Account Nbr and Scheme Id from GMXC Card Strip file.
	// Max Length: 19
	// Min Length: 0
	CrdNbr *string `json:"crdNbr,omitempty"`

	//  Max length = 1, Service Function Code used in conjunction with Sequence Number to influence output transaction list. Values are :  Spaces=first 10 records, 'N'= Next 10 records, 'P'=Previous 10 records
	// Max Length: 1
	// Min Length: 0
	FuncCd *string `json:"funcCd,omitempty"`

	//  Max length = 5, Scheme ID that identifies the points scheme record assigned to the points account.
	// Max Length: 5
	// Min Length: 0
	SchmID *string `json:"schmId,omitempty"`

	//  Max length = 3, Sequence number of transaction within a Batch 99996, start point for output data collation.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	SeqNbr *string `json:"seqNbr,omitempty"`
}

// Validate validates this loyalty redemption inquiry request
func (m *LoyaltyRedemptionInquiryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrdNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFuncCd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchmID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeqNbr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyRedemptionInquiryRequest) validateAcctNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctNbr) { // not required
		return nil
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyRedemptionInquiryRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *LoyaltyRedemptionInquiryRequest) validateCrdNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CrdNbr) { // not required
		return nil
	}

	if err := validate.MinLength("crdNbr", "body", string(*m.CrdNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("crdNbr", "body", string(*m.CrdNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyRedemptionInquiryRequest) validateFuncCd(formats strfmt.Registry) error {

	if swag.IsZero(m.FuncCd) { // not required
		return nil
	}

	if err := validate.MinLength("funcCd", "body", string(*m.FuncCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("funcCd", "body", string(*m.FuncCd), 1); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyRedemptionInquiryRequest) validateSchmID(formats strfmt.Registry) error {

	if swag.IsZero(m.SchmID) { // not required
		return nil
	}

	if err := validate.MinLength("schmId", "body", string(*m.SchmID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("schmId", "body", string(*m.SchmID), 5); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyRedemptionInquiryRequest) validateSeqNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.SeqNbr) { // not required
		return nil
	}

	if err := validate.MinLength("seqNbr", "body", string(*m.SeqNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("seqNbr", "body", string(*m.SeqNbr), 3); err != nil {
		return err
	}

	if err := validate.Pattern("seqNbr", "body", string(*m.SeqNbr), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyRedemptionInquiryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyRedemptionInquiryRequest) UnmarshalBinary(b []byte) error {
	var res LoyaltyRedemptionInquiryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
