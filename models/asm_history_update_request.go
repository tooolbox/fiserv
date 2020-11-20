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

// AsmHistoryUpdateRequest asm history update request
//
// swagger:model AsmHistoryUpdateRequest
type AsmHistoryUpdateRequest struct {

	//  Max length = 4, A Four character code that identifies the action to be taken.
	// Max Length: 4
	// Min Length: 0
	ActionCode *string `json:"actionCode,omitempty"`

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr,omitempty"`

	//  Max length = 4, Card Sequence Number: Identification number assigned to Embosser record to distinguish between multiple cards issued with the same card. 1. Must be numeric and greater than 0 if provided 2. Must be between 0 and 99 if smart card 3. If not provided, and not a smart card, default to value of 0001
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSeq *string `json:"cardSeq,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Service Function Code. The values are: I = Inquire U = Update.
	// Max Length: 1
	// Min Length: 0
	FuncCode *string `json:"funcCode,omitempty"`

	//  Max length = 60, User-defined information.
	// Max Length: 60
	// Min Length: 0
	Memo1 *string `json:"memo1,omitempty"`

	//  Max length = 60, User-defined information.
	// Max Length: 60
	// Min Length: 0
	Memo2 *string `json:"memo2,omitempty"`

	//  Max length = 60, User-defined information.
	// Max Length: 60
	// Min Length: 0
	Memo3 *string `json:"memo3,omitempty"`

	//  Max length = 60, User-defined information.
	// Max Length: 60
	// Min Length: 0
	Memo4 *string `json:"memo4,omitempty"`

	//  Max length = 60, User-defined information.
	// Max Length: 60
	// Min Length: 0
	Memo5 *string `json:"memo5,omitempty"`
}

// Validate validates this asm history update request
func (m *AsmHistoryUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFuncCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo5(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AsmHistoryUpdateRequest) validateActionCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionCode) { // not required
		return nil
	}

	if err := validate.MinLength("actionCode", "body", string(*m.ActionCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("actionCode", "body", string(*m.ActionCode), 4); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateCardNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CardNbr) { // not required
		return nil
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateCardSeq(formats strfmt.Registry) error {

	if swag.IsZero(m.CardSeq) { // not required
		return nil
	}

	if err := validate.MinLength("cardSeq", "body", string(*m.CardSeq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardSeq", "body", string(*m.CardSeq), 4); err != nil {
		return err
	}

	if err := validate.Pattern("cardSeq", "body", string(*m.CardSeq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *AsmHistoryUpdateRequest) validateFuncCode(formats strfmt.Registry) error {

	if swag.IsZero(m.FuncCode) { // not required
		return nil
	}

	if err := validate.MinLength("funcCode", "body", string(*m.FuncCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("funcCode", "body", string(*m.FuncCode), 1); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateMemo1(formats strfmt.Registry) error {

	if swag.IsZero(m.Memo1) { // not required
		return nil
	}

	if err := validate.MinLength("memo1", "body", string(*m.Memo1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("memo1", "body", string(*m.Memo1), 60); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateMemo2(formats strfmt.Registry) error {

	if swag.IsZero(m.Memo2) { // not required
		return nil
	}

	if err := validate.MinLength("memo2", "body", string(*m.Memo2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("memo2", "body", string(*m.Memo2), 60); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateMemo3(formats strfmt.Registry) error {

	if swag.IsZero(m.Memo3) { // not required
		return nil
	}

	if err := validate.MinLength("memo3", "body", string(*m.Memo3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("memo3", "body", string(*m.Memo3), 60); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateMemo4(formats strfmt.Registry) error {

	if swag.IsZero(m.Memo4) { // not required
		return nil
	}

	if err := validate.MinLength("memo4", "body", string(*m.Memo4), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("memo4", "body", string(*m.Memo4), 60); err != nil {
		return err
	}

	return nil
}

func (m *AsmHistoryUpdateRequest) validateMemo5(formats strfmt.Registry) error {

	if swag.IsZero(m.Memo5) { // not required
		return nil
	}

	if err := validate.MinLength("memo5", "body", string(*m.Memo5), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("memo5", "body", string(*m.Memo5), 60); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AsmHistoryUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AsmHistoryUpdateRequest) UnmarshalBinary(b []byte) error {
	var res AsmHistoryUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}