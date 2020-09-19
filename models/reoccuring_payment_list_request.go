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

// ReoccuringPaymentListRequest reoccuring payment list request
//
// swagger:model reoccuringPaymentListRequest
type ReoccuringPaymentListRequest struct {

	// 8 digit Bank Account Number. Mandatory if CardNumberPaid is not used.
	// Max Length: 8
	// Min Length: 8
	// Pattern: ^[0-9]+$
	BankAccountNumber string `json:"bankAccountNumber,omitempty"`

	// Populate with the bank reference number if appropriate. Optional field if BankAccountNumber and BankSortCode are used.
	// Max Length: 16
	// Min Length: 16
	// Pattern: ^[0-9]+$
	BankReferenceNumber string `json:"bankReferenceNumber,omitempty"`

	// 6 digit Bank sort code. Mandatory if CardNumberPaid is not used.
	// Max Length: 6
	// Min Length: 6
	// Pattern: ^[0-9]+$
	BankSortCode string `json:"bankSortCode,omitempty"`

	// 16 digit Card Number. Mandatory if BankAccountNumber, BankSortCode, and optional BankReferenceNumber is not used.
	// Max Length: 16
	// Min Length: 16
	// Pattern: ^[0-9]+$
	CardNumberPaid string `json:"cardNumberPaid,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	// Populate with the FirstPay userID to target a specific region. This is a mandatory field.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this reoccuring payment list request
func (m *ReoccuringPaymentListRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankReferenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankSortCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNumberPaid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReoccuringPaymentListRequest) validateBankAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccountNumber) { // not required
		return nil
	}

	if err := validate.MinLength("bankAccountNumber", "body", string(m.BankAccountNumber), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("bankAccountNumber", "body", string(m.BankAccountNumber), 8); err != nil {
		return err
	}

	if err := validate.Pattern("bankAccountNumber", "body", string(m.BankAccountNumber), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ReoccuringPaymentListRequest) validateBankReferenceNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.BankReferenceNumber) { // not required
		return nil
	}

	if err := validate.MinLength("bankReferenceNumber", "body", string(m.BankReferenceNumber), 16); err != nil {
		return err
	}

	if err := validate.MaxLength("bankReferenceNumber", "body", string(m.BankReferenceNumber), 16); err != nil {
		return err
	}

	if err := validate.Pattern("bankReferenceNumber", "body", string(m.BankReferenceNumber), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ReoccuringPaymentListRequest) validateBankSortCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankSortCode) { // not required
		return nil
	}

	if err := validate.MinLength("bankSortCode", "body", string(m.BankSortCode), 6); err != nil {
		return err
	}

	if err := validate.MaxLength("bankSortCode", "body", string(m.BankSortCode), 6); err != nil {
		return err
	}

	if err := validate.Pattern("bankSortCode", "body", string(m.BankSortCode), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ReoccuringPaymentListRequest) validateCardNumberPaid(formats strfmt.Registry) error {

	if swag.IsZero(m.CardNumberPaid) { // not required
		return nil
	}

	if err := validate.MinLength("cardNumberPaid", "body", string(m.CardNumberPaid), 16); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNumberPaid", "body", string(m.CardNumberPaid), 16); err != nil {
		return err
	}

	if err := validate.Pattern("cardNumberPaid", "body", string(m.CardNumberPaid), `^[0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ReoccuringPaymentListRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *ReoccuringPaymentListRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReoccuringPaymentListRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReoccuringPaymentListRequest) UnmarshalBinary(b []byte) error {
	var res ReoccuringPaymentListRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
