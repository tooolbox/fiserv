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

// StatementListRequest statement list request
//
// swagger:model StatementListRequest
type StatementListRequest struct {

	// FirstVision Account number of the card holder.
	// Required: true
	// Max Length: 19
	// Min Length: 13
	AccountNumber *string `json:"accountNumber"`

	// common
	Common *Header `json:"common,omitempty"`

	// Statement Type. 'A' for Annual Statement, 'M' for monthly statement
	// Pattern: [AM]
	StatementType string `json:"statementType,omitempty"`
}

// Validate validates this statement list request
func (m *StatementListRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatementListRequest) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("accountNumber", "body", m.AccountNumber); err != nil {
		return err
	}

	if err := validate.MinLength("accountNumber", "body", string(*m.AccountNumber), 13); err != nil {
		return err
	}

	if err := validate.MaxLength("accountNumber", "body", string(*m.AccountNumber), 19); err != nil {
		return err
	}

	return nil
}

func (m *StatementListRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *StatementListRequest) validateStatementType(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementType) { // not required
		return nil
	}

	if err := validate.Pattern("statementType", "body", string(m.StatementType), `[AM]`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatementListRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatementListRequest) UnmarshalBinary(b []byte) error {
	var res StatementListRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
