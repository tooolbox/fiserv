// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountScoresInquiryResponse account scores inquiry response
//
// swagger:model AccountScoresInquiryResponse
type AccountScoresInquiryResponse struct {

	//  Max length = 19, Account/Card Number
	AcctNbr string `json:"acctNbr,omitempty"`

	// file end
	AcctScoreGrp []*AcctScoreGrpForAccountScoresInquiry1 `json:"acctScoreGrp"`

	// common
	Common *Header `json:"common,omitempty"`

	// file end
	UserFieldGrp1 []*UserFieldGrp1ForAccountScoresInquiry1 `json:"userFieldGrp1"`

	// file end
	UserFieldGrp2 []*UserFieldGrp2ForAccountScoresInquiry1 `json:"userFieldGrp2"`

	// file end
	UserFieldGrp3 []*UserFieldGrp3ForAccountScoresInquiry1 `json:"userFieldGrp3"`

	// file end
	UserFieldGrp4 []*UserFieldGrp4ForAccountScoresInquiry1 `json:"userFieldGrp4"`

	//  Max length = 5, User Score 1. User score to be used as a decision key in TRIAD batch and online
	UserScore01 string `json:"userScore01,omitempty"`

	// file end
	UserScoreGrp1 []*UserScoreGrp1ForAccountScoresInquiry1 `json:"userScoreGrp1"`

	// file end
	UserScoreGrp2 []*UserScoreGrp2ForAccountScoresInquiry1 `json:"userScoreGrp2"`

	// file end
	UserScoreGrp3 []*UserScoreGrp3ForAccountScoresInquiry1 `json:"userScoreGrp3"`

	// file end
	UserScoreGrp4 []*UserScoreGrp4ForAccountScoresInquiry1 `json:"userScoreGrp4"`
}

// Validate validates this account scores inquiry response
func (m *AccountScoresInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcctScoreGrp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserFieldGrp1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserFieldGrp2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserFieldGrp3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserFieldGrp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserScoreGrp1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserScoreGrp2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserScoreGrp3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserScoreGrp4(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountScoresInquiryResponse) validateAcctScoreGrp(formats strfmt.Registry) error {

	if swag.IsZero(m.AcctScoreGrp) { // not required
		return nil
	}

	for i := 0; i < len(m.AcctScoreGrp); i++ {
		if swag.IsZero(m.AcctScoreGrp[i]) { // not required
			continue
		}

		if m.AcctScoreGrp[i] != nil {
			if err := m.AcctScoreGrp[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acctScoreGrp" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *AccountScoresInquiryResponse) validateUserFieldGrp1(formats strfmt.Registry) error {

	if swag.IsZero(m.UserFieldGrp1) { // not required
		return nil
	}

	for i := 0; i < len(m.UserFieldGrp1); i++ {
		if swag.IsZero(m.UserFieldGrp1[i]) { // not required
			continue
		}

		if m.UserFieldGrp1[i] != nil {
			if err := m.UserFieldGrp1[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userFieldGrp1" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserFieldGrp2(formats strfmt.Registry) error {

	if swag.IsZero(m.UserFieldGrp2) { // not required
		return nil
	}

	for i := 0; i < len(m.UserFieldGrp2); i++ {
		if swag.IsZero(m.UserFieldGrp2[i]) { // not required
			continue
		}

		if m.UserFieldGrp2[i] != nil {
			if err := m.UserFieldGrp2[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userFieldGrp2" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserFieldGrp3(formats strfmt.Registry) error {

	if swag.IsZero(m.UserFieldGrp3) { // not required
		return nil
	}

	for i := 0; i < len(m.UserFieldGrp3); i++ {
		if swag.IsZero(m.UserFieldGrp3[i]) { // not required
			continue
		}

		if m.UserFieldGrp3[i] != nil {
			if err := m.UserFieldGrp3[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userFieldGrp3" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserFieldGrp4(formats strfmt.Registry) error {

	if swag.IsZero(m.UserFieldGrp4) { // not required
		return nil
	}

	for i := 0; i < len(m.UserFieldGrp4); i++ {
		if swag.IsZero(m.UserFieldGrp4[i]) { // not required
			continue
		}

		if m.UserFieldGrp4[i] != nil {
			if err := m.UserFieldGrp4[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userFieldGrp4" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserScoreGrp1(formats strfmt.Registry) error {

	if swag.IsZero(m.UserScoreGrp1) { // not required
		return nil
	}

	for i := 0; i < len(m.UserScoreGrp1); i++ {
		if swag.IsZero(m.UserScoreGrp1[i]) { // not required
			continue
		}

		if m.UserScoreGrp1[i] != nil {
			if err := m.UserScoreGrp1[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userScoreGrp1" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserScoreGrp2(formats strfmt.Registry) error {

	if swag.IsZero(m.UserScoreGrp2) { // not required
		return nil
	}

	for i := 0; i < len(m.UserScoreGrp2); i++ {
		if swag.IsZero(m.UserScoreGrp2[i]) { // not required
			continue
		}

		if m.UserScoreGrp2[i] != nil {
			if err := m.UserScoreGrp2[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userScoreGrp2" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserScoreGrp3(formats strfmt.Registry) error {

	if swag.IsZero(m.UserScoreGrp3) { // not required
		return nil
	}

	for i := 0; i < len(m.UserScoreGrp3); i++ {
		if swag.IsZero(m.UserScoreGrp3[i]) { // not required
			continue
		}

		if m.UserScoreGrp3[i] != nil {
			if err := m.UserScoreGrp3[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userScoreGrp3" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountScoresInquiryResponse) validateUserScoreGrp4(formats strfmt.Registry) error {

	if swag.IsZero(m.UserScoreGrp4) { // not required
		return nil
	}

	for i := 0; i < len(m.UserScoreGrp4); i++ {
		if swag.IsZero(m.UserScoreGrp4[i]) { // not required
			continue
		}

		if m.UserScoreGrp4[i] != nil {
			if err := m.UserScoreGrp4[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userScoreGrp4" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountScoresInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountScoresInquiryResponse) UnmarshalBinary(b []byte) error {
	var res AccountScoresInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
