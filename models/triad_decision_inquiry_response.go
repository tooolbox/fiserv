// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TriadDecisionInquiryResponse triad decision inquiry response
//
// swagger:model TriadDecisionInquiryResponse
type TriadDecisionInquiryResponse struct {

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero.
	Acct string `json:"acct,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Account spending/credit limit: current credit limit assigned to the account.
	CrLim string `json:"crLim,omitempty"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`

	//  Max length = 5, TRIAD response code for the money transfer
	RespCd string `json:"respCd,omitempty"`

	//  Max length = 35, TRIAD response code decription for the money transfer
	RespMsg string `json:"respMsg,omitempty"`

	//  Max length = 20, Short name: short name of the customer
	ShortName string `json:"shortName,omitempty"`
}

// Validate validates this triad decision inquiry response
func (m *TriadDecisionInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriadDecisionInquiryResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *TriadDecisionInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriadDecisionInquiryResponse) UnmarshalBinary(b []byte) error {
	var res TriadDecisionInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
