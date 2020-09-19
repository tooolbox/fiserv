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

// PlanLevelMonthlyInterestInquiryRequest plan level monthly interest inquiry request
//
// swagger:model PlanLevelMonthlyInterestInquiryRequest
type PlanLevelMonthlyInterestInquiryRequest struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero.
	// Max Length: 19
	// Min Length: 0
	AcctNbr *string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 1, Foreign use indicator: Value indicates whether the incoming action is applied to the local or foreign account. Values are: 0 = Local account (default) 1 = Foreign account SPACE = defaults to 0
	// Max Length: 1
	// Min Length: 0
	ForgnUse *string `json:"forgnUse,omitempty"`

	//  Max length = 2, Number of Statement Periods. Used in conjunction with the service type to pull the details.
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	NbrStatPrd *string `json:"nbrStatPrd,omitempty"`

	// Format: YYYYMMDD. Date of statement.
	StatDte string `json:"statDte,omitempty"`

	//  Max length = 1, Service request type to pull the data from a specific statement period or all data from current till the period value:  The values are: A = Number of statement periods, up to 12 months S = Specific statement date.
	// Max Length: 1
	// Min Length: 0
	SvcType *string `json:"svcType,omitempty"`
}

// Validate validates this plan level monthly interest inquiry request
func (m *PlanLevelMonthlyInterestInquiryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForgnUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNbrStatPrd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanLevelMonthlyInterestInquiryRequest) validateAcctNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctNbr) { // not required
		return nil
	}

	if err := validate.MinLength("acctNbr", "body", string(*m.AcctNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acctNbr", "body", string(*m.AcctNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *PlanLevelMonthlyInterestInquiryRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *PlanLevelMonthlyInterestInquiryRequest) validateForgnUse(formats strfmt.Registry) error {

	if swag.IsZero(m.ForgnUse) { // not required
		return nil
	}

	if err := validate.MinLength("forgnUse", "body", string(*m.ForgnUse), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("forgnUse", "body", string(*m.ForgnUse), 1); err != nil {
		return err
	}

	return nil
}

func (m *PlanLevelMonthlyInterestInquiryRequest) validateNbrStatPrd(formats strfmt.Registry) error {

	if swag.IsZero(m.NbrStatPrd) { // not required
		return nil
	}

	if err := validate.MinLength("nbrStatPrd", "body", string(*m.NbrStatPrd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("nbrStatPrd", "body", string(*m.NbrStatPrd), 2); err != nil {
		return err
	}

	if err := validate.Pattern("nbrStatPrd", "body", string(*m.NbrStatPrd), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PlanLevelMonthlyInterestInquiryRequest) validateSvcType(formats strfmt.Registry) error {

	if swag.IsZero(m.SvcType) { // not required
		return nil
	}

	if err := validate.MinLength("svcType", "body", string(*m.SvcType), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("svcType", "body", string(*m.SvcType), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanLevelMonthlyInterestInquiryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanLevelMonthlyInterestInquiryRequest) UnmarshalBinary(b []byte) error {
	var res PlanLevelMonthlyInterestInquiryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
