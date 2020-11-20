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

// OptInCreditLimitUpdatesRequest opt in credit limit updates request
//
// swagger:model OptInCreditLimitUpdatesRequest
type OptInCreditLimitUpdatesRequest struct {

	//  Max length = 19, ACCOUNT
	// Max Length: 19
	// Min Length: 0
	Account *string `json:"account,omitempty"`

	//  Max length = 1, BYPS FLG
	// Max Length: 1
	// Min Length: 0
	// Pattern: ^[0-9]*$
	BypsFlg *string `json:"bypsFlg,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`
}

// Validate validates this opt in credit limit updates request
func (m *OptInCreditLimitUpdatesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBypsFlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OptInCreditLimitUpdatesRequest) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if err := validate.MinLength("account", "body", string(*m.Account), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("account", "body", string(*m.Account), 19); err != nil {
		return err
	}

	return nil
}

func (m *OptInCreditLimitUpdatesRequest) validateBypsFlg(formats strfmt.Registry) error {

	if swag.IsZero(m.BypsFlg) { // not required
		return nil
	}

	if err := validate.MinLength("bypsFlg", "body", string(*m.BypsFlg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("bypsFlg", "body", string(*m.BypsFlg), 1); err != nil {
		return err
	}

	if err := validate.Pattern("bypsFlg", "body", string(*m.BypsFlg), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *OptInCreditLimitUpdatesRequest) validateCommon(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *OptInCreditLimitUpdatesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OptInCreditLimitUpdatesRequest) UnmarshalBinary(b []byte) error {
	var res OptInCreditLimitUpdatesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}