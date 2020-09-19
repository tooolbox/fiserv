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

// MktingPrdtsDataForMarketingPreferencesUpdate1 mkting prdts data for marketing preferences update1
//
// swagger:model MktingPrdtsDataForMarketingPreferencesUpdate1
type MktingPrdtsDataForMarketingPreferencesUpdate1 struct {

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	CarInsurance *string `json:"carInsurance,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	CardProtection *string `json:"cardProtection,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	CoreBanking *string `json:"coreBanking,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	HomeInsurance *string `json:"homeInsurance,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	LifeInsurance *string `json:"lifeInsurance,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	Mortgage *string `json:"mortgage,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	Ppi *string `json:"ppi,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn1 *string `json:"prdtUserDefn1,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn10 *string `json:"prdtUserDefn10,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn11 *string `json:"prdtUserDefn11,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn2 *string `json:"prdtUserDefn2,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn3 *string `json:"prdtUserDefn3,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn4 *string `json:"prdtUserDefn4,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn5 *string `json:"prdtUserDefn5,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn6 *string `json:"prdtUserDefn6,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn7 *string `json:"prdtUserDefn7,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn8 *string `json:"prdtUserDefn8,omitempty"`

	//  Max length = 1, For future use.
	// Max Length: 1
	// Min Length: 0
	PrdtUserDefn9 *string `json:"prdtUserDefn9,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	Savings *string `json:"savings,omitempty"`

	//  Max length = 1, Confirms whether the product is opted or not. Values are: '0' -  Opt out '1' -  Opt in
	// Max Length: 1
	// Min Length: 0
	UnsecuredLoans *string `json:"unsecuredLoans,omitempty"`
}

// Validate validates this mkting prdts data for marketing preferences update1
func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarInsurance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoreBanking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomeInsurance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifeInsurance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMortgage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn10(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn11(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn7(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn8(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrdtUserDefn9(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSavings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnsecuredLoans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateCarInsurance(formats strfmt.Registry) error {

	if swag.IsZero(m.CarInsurance) { // not required
		return nil
	}

	if err := validate.MinLength("carInsurance", "body", string(*m.CarInsurance), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("carInsurance", "body", string(*m.CarInsurance), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateCardProtection(formats strfmt.Registry) error {

	if swag.IsZero(m.CardProtection) { // not required
		return nil
	}

	if err := validate.MinLength("cardProtection", "body", string(*m.CardProtection), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardProtection", "body", string(*m.CardProtection), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateCoreBanking(formats strfmt.Registry) error {

	if swag.IsZero(m.CoreBanking) { // not required
		return nil
	}

	if err := validate.MinLength("coreBanking", "body", string(*m.CoreBanking), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("coreBanking", "body", string(*m.CoreBanking), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateHomeInsurance(formats strfmt.Registry) error {

	if swag.IsZero(m.HomeInsurance) { // not required
		return nil
	}

	if err := validate.MinLength("homeInsurance", "body", string(*m.HomeInsurance), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("homeInsurance", "body", string(*m.HomeInsurance), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateLifeInsurance(formats strfmt.Registry) error {

	if swag.IsZero(m.LifeInsurance) { // not required
		return nil
	}

	if err := validate.MinLength("lifeInsurance", "body", string(*m.LifeInsurance), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lifeInsurance", "body", string(*m.LifeInsurance), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateMortgage(formats strfmt.Registry) error {

	if swag.IsZero(m.Mortgage) { // not required
		return nil
	}

	if err := validate.MinLength("mortgage", "body", string(*m.Mortgage), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("mortgage", "body", string(*m.Mortgage), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePpi(formats strfmt.Registry) error {

	if swag.IsZero(m.Ppi) { // not required
		return nil
	}

	if err := validate.MinLength("ppi", "body", string(*m.Ppi), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("ppi", "body", string(*m.Ppi), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn1(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn1) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn1", "body", string(*m.PrdtUserDefn1), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn1", "body", string(*m.PrdtUserDefn1), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn10(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn10) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn10", "body", string(*m.PrdtUserDefn10), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn10", "body", string(*m.PrdtUserDefn10), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn11(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn11) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn11", "body", string(*m.PrdtUserDefn11), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn11", "body", string(*m.PrdtUserDefn11), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn2(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn2) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn2", "body", string(*m.PrdtUserDefn2), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn2", "body", string(*m.PrdtUserDefn2), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn3(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn3) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn3", "body", string(*m.PrdtUserDefn3), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn3", "body", string(*m.PrdtUserDefn3), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn4(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn4) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn4", "body", string(*m.PrdtUserDefn4), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn4", "body", string(*m.PrdtUserDefn4), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn5(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn5) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn5", "body", string(*m.PrdtUserDefn5), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn5", "body", string(*m.PrdtUserDefn5), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn6(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn6) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn6", "body", string(*m.PrdtUserDefn6), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn6", "body", string(*m.PrdtUserDefn6), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn7(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn7) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn7", "body", string(*m.PrdtUserDefn7), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn7", "body", string(*m.PrdtUserDefn7), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn8(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn8) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn8", "body", string(*m.PrdtUserDefn8), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn8", "body", string(*m.PrdtUserDefn8), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validatePrdtUserDefn9(formats strfmt.Registry) error {

	if swag.IsZero(m.PrdtUserDefn9) { // not required
		return nil
	}

	if err := validate.MinLength("prdtUserDefn9", "body", string(*m.PrdtUserDefn9), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("prdtUserDefn9", "body", string(*m.PrdtUserDefn9), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateSavings(formats strfmt.Registry) error {

	if swag.IsZero(m.Savings) { // not required
		return nil
	}

	if err := validate.MinLength("savings", "body", string(*m.Savings), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("savings", "body", string(*m.Savings), 1); err != nil {
		return err
	}

	return nil
}

func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) validateUnsecuredLoans(formats strfmt.Registry) error {

	if swag.IsZero(m.UnsecuredLoans) { // not required
		return nil
	}

	if err := validate.MinLength("unsecuredLoans", "body", string(*m.UnsecuredLoans), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("unsecuredLoans", "body", string(*m.UnsecuredLoans), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MktingPrdtsDataForMarketingPreferencesUpdate1) UnmarshalBinary(b []byte) error {
	var res MktingPrdtsDataForMarketingPreferencesUpdate1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
