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

// PinBlockInquiryRequest pin block inquiry request
//
// swagger:model PinBlockInquiryRequest
type PinBlockInquiryRequest struct {

	//  Max length = 3, Client association id. This particular process of PIN by XML uses the Association 'XMl'.
	// Required: true
	// Max Length: 3
	// Min Length: 0
	Association *string `json:"association"`

	//  Max length = 9, Bank sort code (ARMB05) of the cardholder:  Standard convenience chequebook Bank sort code assigned to the customer   (for additional validation only).
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	BankSortCd *string `json:"bankSortCd,omitempty"`

	//  Max length = 19, Card number for which the PIN is required to be retrieved
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr"`

	//  Max length = 4, Card verification value like CVV2/CVC2/CSC etc. depending on the scheme.
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSecCd *string `json:"cardSecCd,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 19, Current (DD) bank account number  of the cardholder(for additional validation only).
	// Max Length: 19
	// Min Length: 0
	CurrAcctNbr *string `json:"currAcctNbr,omitempty"`

	// Format: YYYYMMDD. Customer's date of birth.
	Dob string `json:"dob,omitempty"`

	//  Max length = 26, Embossed Name: Cardholder name as it is embossed on the card.
	// Max Length: 26
	// Min Length: 0
	EmbName *string `json:"embName,omitempty"`

	// Format: YYMM. Card Expiry date.
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ExpDate *string `json:"expDate,omitempty"`

	//  Max length = 20, Cardholder home phone number with or without country code depending on the issuing platform.
	// Max Length: 20
	// Min Length: 0
	HomePhoneNbr *string `json:"homePhoneNbr,omitempty"`

	//  Max length = 20, Memorable word: Memorable word used for SMS PIN Data.   Values are: - A to Z - 0-9 - Space - Underscore - Full stop Mandatory if customer chooses PIN by SMS mode or logo level preference is SMS.
	// Max Length: 20
	// Min Length: 0
	MemPinWord *string `json:"memPinWord,omitempty"`

	//  Max length = 20, Cardholder mobile phone number with or without country code depending on the issuing platform.
	// Max Length: 20
	// Min Length: 0
	MobileNbr *string `json:"mobileNbr,omitempty"`

	//  Max length = 10, Postal Code: Postal code portion of the mailing address.
	// Max Length: 10
	// Min Length: 0
	PostCode *string `json:"postCode,omitempty"`

	//  Max length = 25, Social Security number (national Insurance number): field that identifies the SSN ,tax identification number or another user-defined identification number of the cardholder.
	// Max Length: 25
	// Min Length: 0
	Ssn *string `json:"ssn,omitempty"`

	//  Max length = 19, Unique ID: It identifies a customer uniquely.
	// Max Length: 19
	// Min Length: 0
	UniqueID *string `json:"uniqueId,omitempty"`
}

// Validate validates this pin block inquiry request
func (m *PinBlockInquiryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankSortCd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSecCd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmbName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomePhoneNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemPinWord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobileNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniqueID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PinBlockInquiryRequest) validateAssociation(formats strfmt.Registry) error {

	if err := validate.Required("association", "body", m.Association); err != nil {
		return err
	}

	if err := validate.MinLength("association", "body", string(*m.Association), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("association", "body", string(*m.Association), 3); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateBankSortCd(formats strfmt.Registry) error {

	if swag.IsZero(m.BankSortCd) { // not required
		return nil
	}

	if err := validate.MinLength("bankSortCd", "body", string(*m.BankSortCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("bankSortCd", "body", string(*m.BankSortCd), 9); err != nil {
		return err
	}

	if err := validate.Pattern("bankSortCd", "body", string(*m.BankSortCd), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateCardNbr(formats strfmt.Registry) error {

	if err := validate.Required("cardNbr", "body", m.CardNbr); err != nil {
		return err
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateCardSecCd(formats strfmt.Registry) error {

	if swag.IsZero(m.CardSecCd) { // not required
		return nil
	}

	if err := validate.MinLength("cardSecCd", "body", string(*m.CardSecCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardSecCd", "body", string(*m.CardSecCd), 4); err != nil {
		return err
	}

	if err := validate.Pattern("cardSecCd", "body", string(*m.CardSecCd), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *PinBlockInquiryRequest) validateCurrAcctNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrAcctNbr) { // not required
		return nil
	}

	if err := validate.MinLength("currAcctNbr", "body", string(*m.CurrAcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("currAcctNbr", "body", string(*m.CurrAcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateEmbName(formats strfmt.Registry) error {

	if swag.IsZero(m.EmbName) { // not required
		return nil
	}

	if err := validate.MinLength("embName", "body", string(*m.EmbName), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("embName", "body", string(*m.EmbName), 26); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateExpDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpDate) { // not required
		return nil
	}

	if err := validate.MinLength("expDate", "body", string(*m.ExpDate), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("expDate", "body", string(*m.ExpDate), 4); err != nil {
		return err
	}

	if err := validate.Pattern("expDate", "body", string(*m.ExpDate), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateHomePhoneNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.HomePhoneNbr) { // not required
		return nil
	}

	if err := validate.MinLength("homePhoneNbr", "body", string(*m.HomePhoneNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("homePhoneNbr", "body", string(*m.HomePhoneNbr), 20); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateMemPinWord(formats strfmt.Registry) error {

	if swag.IsZero(m.MemPinWord) { // not required
		return nil
	}

	if err := validate.MinLength("memPinWord", "body", string(*m.MemPinWord), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("memPinWord", "body", string(*m.MemPinWord), 20); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateMobileNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.MobileNbr) { // not required
		return nil
	}

	if err := validate.MinLength("mobileNbr", "body", string(*m.MobileNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("mobileNbr", "body", string(*m.MobileNbr), 20); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validatePostCode(formats strfmt.Registry) error {

	if swag.IsZero(m.PostCode) { // not required
		return nil
	}

	if err := validate.MinLength("postCode", "body", string(*m.PostCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("postCode", "body", string(*m.PostCode), 10); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateSsn(formats strfmt.Registry) error {

	if swag.IsZero(m.Ssn) { // not required
		return nil
	}

	if err := validate.MinLength("ssn", "body", string(*m.Ssn), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("ssn", "body", string(*m.Ssn), 25); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockInquiryRequest) validateUniqueID(formats strfmt.Registry) error {

	if swag.IsZero(m.UniqueID) { // not required
		return nil
	}

	if err := validate.MinLength("uniqueId", "body", string(*m.UniqueID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("uniqueId", "body", string(*m.UniqueID), 19); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PinBlockInquiryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PinBlockInquiryRequest) UnmarshalBinary(b []byte) error {
	var res PinBlockInquiryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
