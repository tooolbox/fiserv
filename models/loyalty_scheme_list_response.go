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

// LoyaltySchemeListResponse loyalty scheme list response
//
// swagger:model LoyaltySchemeListResponse
type LoyaltySchemeListResponse struct {

	//  Max length = 19, LMS account number that identifies the Points Account record
	AcctNbr string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 40, LMS organization name
	LmsOrgName string `json:"lmsOrgName,omitempty"`

	//  Max length = 1, Code that indicates whether more records are on file N = No Y = Yes
	MoreData string `json:"moreData,omitempty"`

	//  Max length = 40, Name of the account holder
	Name1 string `json:"name1,omitempty"`

	//  Max length = 2, Number of occurrences for scheme related fields
	NbrItemRtn string `json:"nbrItemRtn,omitempty"`

	//  Max length = 3, Organization: Three digit Identification number of the organization.
	Org string `json:"org,omitempty"`

	// 20 Occurrences of scheme-id and scheme-name ( 45 bytes each group occurrence)
	Schm []*SchmForLoyaltySchemeList1 `json:"schm"`
}

// Validate validates this loyalty scheme list response
func (m *LoyaltySchemeListResponse) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *LoyaltySchemeListResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *LoyaltySchemeListResponse) validateSchm(formats strfmt.Registry) error {

	if swag.IsZero(m.Schm) { // not required
		return nil
	}

	for i := 0; i < len(m.Schm); i++ {
		if swag.IsZero(m.Schm[i]) { // not required
			continue
		}

		if m.Schm[i] != nil {
			if err := m.Schm[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schm" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoyaltySchemeListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoyaltySchemeListResponse) UnmarshalBinary(b []byte) error {
	var res LoyaltySchemeListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
