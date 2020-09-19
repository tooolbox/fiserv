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

// MemoInquiryRequest memo inquiry request
//
// swagger:model MemoInquiryRequest
type MemoInquiryRequest struct {

	//  Max length = 19, The Account or Card Number for this request.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr"`

	// common
	Common *Header `json:"common,omitempty"`

	// Format: YYYYMMDD. Populate date from MEMO_DATE from the previous call's response
	ContDate string `json:"contDate,omitempty"`

	//  Max length = 1, Populate form REC_TYPE from the previous call's response
	// Max Length: 1
	// Min Length: 0
	ContRecType *string `json:"contRecType,omitempty"`

	//  Max length = 6, Populate time from MEMO_TIME from the previous call's response
	// Max Length: 6
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ContTime *string `json:"contTime,omitempty"`

	//  Max length = 2, Number of memos to be returned with the output service.  Maximum number allowed is 50.
	// Required: true
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	NbrRec *string `json:"nbrRec"`

	// sequence
	// Max Length: 1
	// Min Length: 0
	Sequence *string `json:"sequence,omitempty"`

	//  Max length = 1, Indicates if this is an initial request, or a continuation of a previous request. Space = initial N = next Defaults to space
	// Max Length: 1
	// Min Length: 0
	SvcFuncCd *string `json:"svcFuncCd,omitempty"`
}

// Validate validates this memo inquiry request
func (m *MemoInquiryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContRecType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNbrRec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcFuncCd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemoInquiryRequest) validateAcctNbr(formats strfmt.Registry) error {

	if err := validate.Required("acctNbr", "body", m.AcctNbr); err != nil {
		return err
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *MemoInquiryRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *MemoInquiryRequest) validateContRecType(formats strfmt.Registry) error {

	if swag.IsZero(m.ContRecType) { // not required
		return nil
	}

	if err := validate.MinLength("contRecType", "body", string(*m.ContRecType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("contRecType", "body", string(*m.ContRecType), 1); err != nil {
		return err
	}

	return nil
}

func (m *MemoInquiryRequest) validateContTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ContTime) { // not required
		return nil
	}

	if err := validate.MinLength("contTime", "body", string(*m.ContTime), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("contTime", "body", string(*m.ContTime), 6); err != nil {
		return err
	}

	if err := validate.Pattern("contTime", "body", string(*m.ContTime), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MemoInquiryRequest) validateNbrRec(formats strfmt.Registry) error {

	if err := validate.Required("nbrRec", "body", m.NbrRec); err != nil {
		return err
	}

	if err := validate.MinLength("nbrRec", "body", string(*m.NbrRec), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("nbrRec", "body", string(*m.NbrRec), 2); err != nil {
		return err
	}

	if err := validate.Pattern("nbrRec", "body", string(*m.NbrRec), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MemoInquiryRequest) validateSequence(formats strfmt.Registry) error {

	if swag.IsZero(m.Sequence) { // not required
		return nil
	}

	if err := validate.MinLength("sequence", "body", string(*m.Sequence), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("sequence", "body", string(*m.Sequence), 1); err != nil {
		return err
	}

	return nil
}

func (m *MemoInquiryRequest) validateSvcFuncCd(formats strfmt.Registry) error {

	if swag.IsZero(m.SvcFuncCd) { // not required
		return nil
	}

	if err := validate.MinLength("svcFuncCd", "body", string(*m.SvcFuncCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("svcFuncCd", "body", string(*m.SvcFuncCd), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemoInquiryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemoInquiryRequest) UnmarshalBinary(b []byte) error {
	var res MemoInquiryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
