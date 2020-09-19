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

// BalanceUpdateRequest balance update request
//
// swagger:model BalanceUpdateRequest
type BalanceUpdateRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero. This can be either Account number or Card number.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	Account *string `json:"account"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Account spending/credit limit: current credit limit assigned to the account.
	// Required: true
	// Pattern: ^(-)?[0-9]{1,17}$
	Crlim *string `json:"crlim"`

	//  Max length = 1, Dual Indicator: Dual Account Flag of account.Valid values are: L = Local account (default) F = Foreign account If dual currency is not being used or if the field is left blank, the value is L.
	// Max Length: 1
	// Min Length: 0
	DualInd *string `json:"dualInd,omitempty"`

	//  Max length = 1, RESET HOLD AMT
	// Max Length: 1
	// Min Length: 0
	ResetHoldAmt *string `json:"resetHoldAmt,omitempty"`
}

// Validate validates this balance update request
func (m *BalanceUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrlim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResetHoldAmt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BalanceUpdateRequest) validateAccount(formats strfmt.Registry) error {

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

func (m *BalanceUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *BalanceUpdateRequest) validateCrlim(formats strfmt.Registry) error {

	if err := validate.Required("crlim", "body", m.Crlim); err != nil {
		return err
	}

	if err := validate.Pattern("crlim", "body", string(*m.Crlim), `^(-)?[0-9]{1,17}$`); err != nil {
		return err
	}

	return nil
}

func (m *BalanceUpdateRequest) validateDualInd(formats strfmt.Registry) error {

	if swag.IsZero(m.DualInd) { // not required
		return nil
	}

	if err := validate.MinLength("dualInd", "body", string(*m.DualInd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("dualInd", "body", string(*m.DualInd), 1); err != nil {
		return err
	}

	return nil
}

func (m *BalanceUpdateRequest) validateResetHoldAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.ResetHoldAmt) { // not required
		return nil
	}

	if err := validate.MinLength("resetHoldAmt", "body", string(*m.ResetHoldAmt), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("resetHoldAmt", "body", string(*m.ResetHoldAmt), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BalanceUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BalanceUpdateRequest) UnmarshalBinary(b []byte) error {
	var res BalanceUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
