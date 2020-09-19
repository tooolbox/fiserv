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

// PaymentHistoryInquiryRequest payment history inquiry request
//
// swagger:model PaymentHistoryInquiryRequest
type PaymentHistoryInquiryRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero. This can be either Account number or Card number.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Foreign use indicator: Value indicates whether the incoming action is applied to the local or foreign account. Values are: 0 = Local account (default) 1 = Foreign account SPACE = defaults to 0
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	ForeignUse *string `json:"foreignUse,omitempty"`
}

// Validate validates this payment history inquiry request
func (m *PaymentHistoryInquiryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForeignUse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentHistoryInquiryRequest) validateAcctNbr(formats strfmt.Registry) error {

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

func (m *PaymentHistoryInquiryRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *PaymentHistoryInquiryRequest) validateForeignUse(formats strfmt.Registry) error {

	if swag.IsZero(m.ForeignUse) { // not required
		return nil
	}

	if err := validate.MinLength("foreignUse", "body", string(*m.ForeignUse), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("foreignUse", "body", string(*m.ForeignUse), 1); err != nil {
		return err
	}

	if err := validate.Pattern("foreignUse", "body", string(*m.ForeignUse), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentHistoryInquiryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentHistoryInquiryRequest) UnmarshalBinary(b []byte) error {
	var res PaymentHistoryInquiryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
