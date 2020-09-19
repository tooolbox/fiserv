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

// CardholderDemographicListRequest2 cardholder demographic list request2
//
// swagger:model CardholderDemographicListRequest2
type CardholderDemographicListRequest2 struct {

	//  Max length = 1, Account Demographics switch: Indicates the demographics results option.  Valid values are: Space, Low Values and Zero - Return card demographic information only 1 - Return account and card demographic information
	// Max Length: 1
	// Min Length: 0
	AccountDemographicsSw *string `json:"accountDemographicsSw,omitempty"`

	//  Max length = 19, Account Number: Number of Customer's account. Must be numeric and greater than zero. This can be either Account number or Card number.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	Acct *string `json:"acct"`

	//  Max length = 1, Assoc. Demographics option: Indicates whether to include associated parties demographic information in the service response. The valid values: Space, Low Values and Zero - Exclude associated parties demogrphic data 1 - Include associated parties demographic data
	// Max Length: 1
	// Min Length: 0
	AssocDemographicsSw *string `json:"assocDemographicsSw,omitempty"`

	//  Max length = 1, Card Demographics options: Indicates the card results options. The valid values are: Space, Low Values and Zero - Includes all card in the response 1 - Include only cards with demographic data.
	// Max Length: 1
	// Min Length: 0
	CardDemographicsSw *string `json:"cardDemographicsSw,omitempty"`

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr,omitempty"`

	//  Max length = 3, Organization Number: Three digit Identification number of the organization.  Valid values are 001-998. Organization number must be on file. Required when Account Number not provided.
	// Max Length: 3
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardOrg *string `json:"cardOrg,omitempty"`

	//  Max length = 4, Card sequence number: record number assigned to the card (for card numbering schemes of 0, 1, and 3) and the sequence number assigned to the card (for card numbering schemes of 2). This number is part of the record key.
	// Max Length: 4
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CardSeq *string `json:"cardSeq,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 2, Number of Requested Records: Valid values are 01 to 20.
	// Max Length: 2
	// Min Length: 0
	// Pattern: ^[0-9]*$
	NbrOfRecsReq *string `json:"nbrOfRecsReq,omitempty"`

	//  Max length = 1, Prior Mail switch: Prior mailing information option.The values are: Space, Low Values and Zero - Do not include prior mailing information for the customer record. 1 - Include the prior mailing information for the customer record.
	// Max Length: 1
	// Min Length: 0
	PriorMailSw *string `json:"priorMailSw,omitempty"`

	//  Max length = 1, Service Function Code. Values are: Low Values, Spaces = Initial records N = Next P = Previous
	// Max Length: 1
	// Min Length: 0
	SvcFuncCd *string `json:"svcFuncCd,omitempty"`
}

// Validate validates this cardholder demographic list request2
func (m *CardholderDemographicListRequest2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountDemographicsSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssocDemographicsSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardDemographicsSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSeq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNbrOfRecsReq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriorMailSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcFuncCd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardholderDemographicListRequest2) validateAccountDemographicsSw(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountDemographicsSw) { // not required
		return nil
	}

	if err := validate.MinLength("accountDemographicsSw", "body", string(*m.AccountDemographicsSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("accountDemographicsSw", "body", string(*m.AccountDemographicsSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateAcct(formats strfmt.Registry) error {

	if err := validate.Required("acct", "body", m.Acct); err != nil {
		return err
	}

	if err := validate.MinLength("acct", "body", string(*m.Acct), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("acct", "body", string(*m.Acct), 19); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateAssocDemographicsSw(formats strfmt.Registry) error {

	if swag.IsZero(m.AssocDemographicsSw) { // not required
		return nil
	}

	if err := validate.MinLength("assocDemographicsSw", "body", string(*m.AssocDemographicsSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("assocDemographicsSw", "body", string(*m.AssocDemographicsSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateCardDemographicsSw(formats strfmt.Registry) error {

	if swag.IsZero(m.CardDemographicsSw) { // not required
		return nil
	}

	if err := validate.MinLength("cardDemographicsSw", "body", string(*m.CardDemographicsSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardDemographicsSw", "body", string(*m.CardDemographicsSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateCardNbr(formats strfmt.Registry) error {

	if swag.IsZero(m.CardNbr) { // not required
		return nil
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateCardOrg(formats strfmt.Registry) error {

	if swag.IsZero(m.CardOrg) { // not required
		return nil
	}

	if err := validate.MinLength("cardOrg", "body", string(*m.CardOrg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardOrg", "body", string(*m.CardOrg), 3); err != nil {
		return err
	}

	if err := validate.Pattern("cardOrg", "body", string(*m.CardOrg), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateCardSeq(formats strfmt.Registry) error {

	if swag.IsZero(m.CardSeq) { // not required
		return nil
	}

	if err := validate.MinLength("cardSeq", "body", string(*m.CardSeq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardSeq", "body", string(*m.CardSeq), 4); err != nil {
		return err
	}

	if err := validate.Pattern("cardSeq", "body", string(*m.CardSeq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateCommon(formats strfmt.Registry) error {

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

func (m *CardholderDemographicListRequest2) validateNbrOfRecsReq(formats strfmt.Registry) error {

	if swag.IsZero(m.NbrOfRecsReq) { // not required
		return nil
	}

	if err := validate.MinLength("nbrOfRecsReq", "body", string(*m.NbrOfRecsReq), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("nbrOfRecsReq", "body", string(*m.NbrOfRecsReq), 2); err != nil {
		return err
	}

	if err := validate.Pattern("nbrOfRecsReq", "body", string(*m.NbrOfRecsReq), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validatePriorMailSw(formats strfmt.Registry) error {

	if swag.IsZero(m.PriorMailSw) { // not required
		return nil
	}

	if err := validate.MinLength("priorMailSw", "body", string(*m.PriorMailSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("priorMailSw", "body", string(*m.PriorMailSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *CardholderDemographicListRequest2) validateSvcFuncCd(formats strfmt.Registry) error {

	if swag.IsZero(m.SvcFuncCd) { // not required
		return nil
	}

	if err := validate.MinLength("svcFuncCd", "body", string(*m.SvcFuncCd), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("svcFuncCd", "body", string(*m.SvcFuncCd), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardholderDemographicListRequest2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardholderDemographicListRequest2) UnmarshalBinary(b []byte) error {
	var res CardholderDemographicListRequest2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
