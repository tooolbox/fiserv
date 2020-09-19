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

// PinBlockUpdateRequest pin block update request
//
// swagger:model PinBlockUpdateRequest
type PinBlockUpdateRequest struct {

	//  Max length = 3, Association used to reference keys that are shared with an external association such as a card bureau, Visa, or MasterCard.
	// Max Length: 3
	// Min Length: 0
	Association *string `json:"association,omitempty"`

	//  Max length = 19, Card Number: Unique Card number embossed on the plastic card. 1. Must be numeric and greater than 0 2. Card number must be on file 3. Card number must be valid for Org provided
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CardNbr *string `json:"cardNbr"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 16, PIN Block encrypted under keys exchanged between First Data and Client. PIN Block cannot be spaces.
	// Required: true
	// Max Length: 16
	// Min Length: 0
	PinBlock *string `json:"pinBlock"`

	//  Max length = 1, PIN Set Reversals: Field to support PIN Set reversals. When the field value is set to 'Y' previous PIN offset will be used to populate the current offset at card level and the previous offset will be reset to zeroes. This action depends on the logo ATM PIN IND parameter value. Valid values : Spaces - Not a reversal request (Default) Y - PIN Set Reversal request
	// Max Length: 1
	// Min Length: 0
	PinSetReversalSw *string `json:"pinSetReversalSw,omitempty"`

	//  Max length = 1, Reset Invalid PIN Switch: Field to indicate the service to generate a log record (non-mon) for resetting invalid PIN tries count on embosser.  Valid Values:  Y - Reset invalid PIN tries count Low values, space - Do not reset invalid PIN tries count
	// Max Length: 1
	// Min Length: 0
	ResetInvPinSw *string `json:"resetInvPinSw,omitempty"`

	//  Max length = 6, Supervisor ID: 6 character id of the supervisor of the person requesting PIN assignment. Default value is 999999.
	// Required: true
	// Max Length: 6
	// Min Length: 0
	SupID *string `json:"supId"`

	//  Max length = 8, Terminal ID of the ATM from where the PIN change was requested. Default value is Zeroes.
	// Required: true
	// Max Length: 8
	// Min Length: 0
	TermID *string `json:"termId"`

	//  Max length = 6, User ID: 6 character ID number of the person requesting PIN assignment.  Default value is 999999.
	// Required: true
	// Max Length: 6
	// Min Length: 0
	UserID *string `json:"userId"`
}

// Validate validates this pin block update request
func (m *PinBlockUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePinBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePinSetReversalSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResetInvPinSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PinBlockUpdateRequest) validateAssociation(formats strfmt.Registry) error {

	if swag.IsZero(m.Association) { // not required
		return nil
	}

	if err := validate.MinLength("association", "body", string(*m.Association), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("association", "body", string(*m.Association), 3); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validateCardNbr(formats strfmt.Registry) error {

	if err := validate.Required("cardNbr", "body", m.CardNbr); err != nil {
		return err
	}

	if err := validate.MinLength("cardNbr", "body", string(*m.CardNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cardNbr", "body", string(*m.CardNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *PinBlockUpdateRequest) validatePinBlock(formats strfmt.Registry) error {

	if err := validate.Required("pinBlock", "body", m.PinBlock); err != nil {
		return err
	}

	if err := validate.MinLength("pinBlock", "body", string(*m.PinBlock), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("pinBlock", "body", string(*m.PinBlock), 16); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validatePinSetReversalSw(formats strfmt.Registry) error {

	if swag.IsZero(m.PinSetReversalSw) { // not required
		return nil
	}

	if err := validate.MinLength("pinSetReversalSw", "body", string(*m.PinSetReversalSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("pinSetReversalSw", "body", string(*m.PinSetReversalSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validateResetInvPinSw(formats strfmt.Registry) error {

	if swag.IsZero(m.ResetInvPinSw) { // not required
		return nil
	}

	if err := validate.MinLength("resetInvPinSw", "body", string(*m.ResetInvPinSw), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("resetInvPinSw", "body", string(*m.ResetInvPinSw), 1); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validateSupID(formats strfmt.Registry) error {

	if err := validate.Required("supId", "body", m.SupID); err != nil {
		return err
	}

	if err := validate.MinLength("supId", "body", string(*m.SupID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("supId", "body", string(*m.SupID), 6); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validateTermID(formats strfmt.Registry) error {

	if err := validate.Required("termId", "body", m.TermID); err != nil {
		return err
	}

	if err := validate.MinLength("termId", "body", string(*m.TermID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("termId", "body", string(*m.TermID), 8); err != nil {
		return err
	}

	return nil
}

func (m *PinBlockUpdateRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.MinLength("userId", "body", string(*m.UserID), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("userId", "body", string(*m.UserID), 6); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PinBlockUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PinBlockUpdateRequest) UnmarshalBinary(b []byte) error {
	var res PinBlockUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
