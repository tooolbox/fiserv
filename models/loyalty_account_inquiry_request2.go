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

// LoyaltyAccountInquiryRequest2 loyalty account inquiry request2
//
// swagger:model LoyaltyAccountInquiryRequest2
type LoyaltyAccountInquiryRequest2 struct {

	//  Max length = 19, LMS Account Number. If account number is not entered  LMS will auto generate the account number.  The account number should be present in account demographic records if trying to update add pending account.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	Account *string `json:"account"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 5, Scheme ID that identifies the points scheme record assigned to the points account.
	// Required: true
	// Max Length: 5
	// Min Length: 0
	Schm *string `json:"schm"`
}

// Validate validates this loyalty account inquiry request2
func (m *LoyaltyAccountInquiryRequest2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoyaltyAccountInquiryRequest2) validateAccount(formats strfmt.Registry) error {

	if err := validate.Required("account", "body", m.Account); err != nil {
		return err
	}

	if err := validate.MinLength("account", "body", string(*m.Account), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("account", "body", string(*m.Account), 19); err != nil {
		return err
	}

	return nil
}

func (m *LoyaltyAccountInquiryRequest2) validateCommon(formats strfmt.Registry) error {

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

func (m *LoyaltyAccountInquiryRequest2) validateSchm(formats strfmt.Registry) error {

	if err := validate.Required("schm", "body", m.Schm); err != nil {
		return err
	}

	if err := validate.MinLength("schm", "body", string(*m.Schm), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("schm", "body", string(*m.Schm), 5); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltyAccountInquiryRequest2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltyAccountInquiryRequest2) UnmarshalBinary(b []byte) error {
	var res LoyaltyAccountInquiryRequest2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
