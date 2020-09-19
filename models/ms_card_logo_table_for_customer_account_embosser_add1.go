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

// MsCardLogoTableForCustomerAccountEmbosserAdd1 ms card logo table for customer account embosser add1
//
// swagger:model MsCardLogoTableForCustomerAccountEmbosserAdd1
type MsCardLogoTableForCustomerAccountEmbosserAdd1 struct {

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Max Length: 19
	// Min Length: 0
	CardNumber *string `json:"cardNumber,omitempty"`

	//  Max length = 5, Field that defines the cardholder's affiliation or affinity group. The values are user-defined. If EmblemID is not provided it will be defaulted from the Card Logo setting
	// Max Length: 5
	// Min Length: 0
	// Pattern: ^[0-9]*$
	EmblemID *string `json:"emblemId,omitempty"`

	//  Max length = 3, At least 1 CARD LOGO should be entered when EmbosserRecordAction = 'A'. When adding a card to an account, the Card Logo is automatically added to the Card Details table on the Account record
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	MultiSchemeCardLogo *string `json:"multiSchemeCardLogo,omitempty"`

	//  Max length = 4, Card Sequence Number: Identification number assigned to Embosser record to distinguish between multiple cards issued with the same card.1. Must be numeric and greater than 0 if provided2. Must be between 0 and 99 if smart card3. If not provided, and not a smart card, default to value of 0001
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	RecordNumber *string `json:"recordNumber,omitempty"`
}

// Validate validates this ms card logo table for customer account embosser add1
func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmblemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiSchemeCardLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) validateCardNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.CardNumber) { // not required
		return nil
	}

	if err := validate.MinLength("cardNumber", "body", string(*m.CardNumber), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNumber", "body", string(*m.CardNumber), 19); err != nil {
		return err
	}

	return nil
}

func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) validateEmblemID(formats strfmt.Registry) error {

	if swag.IsZero(m.EmblemID) { // not required
		return nil
	}

	if err := validate.MinLength("emblemId", "body", string(*m.EmblemID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("emblemId", "body", string(*m.EmblemID), 5); err != nil {
		return err
	}

	if err := validate.Pattern("emblemId", "body", string(*m.EmblemID), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) validateMultiSchemeCardLogo(formats strfmt.Registry) error {

	if swag.IsZero(m.MultiSchemeCardLogo) { // not required
		return nil
	}

	if err := validate.MinLength("multiSchemeCardLogo", "body", string(*m.MultiSchemeCardLogo), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("multiSchemeCardLogo", "body", string(*m.MultiSchemeCardLogo), 3); err != nil {
		return err
	}

	if err := validate.Pattern("multiSchemeCardLogo", "body", string(*m.MultiSchemeCardLogo), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) validateRecordNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordNumber) { // not required
		return nil
	}

	if err := validate.MinLength("recordNumber", "body", string(*m.RecordNumber), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("recordNumber", "body", string(*m.RecordNumber), 4); err != nil {
		return err
	}

	if err := validate.Pattern("recordNumber", "body", string(*m.RecordNumber), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsCardLogoTableForCustomerAccountEmbosserAdd1) UnmarshalBinary(b []byte) error {
	var res MsCardLogoTableForCustomerAccountEmbosserAdd1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
