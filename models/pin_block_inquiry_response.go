// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PinBlockInquiryResponse pin block inquiry response
//
// swagger:model PinBlockInquiryResponse
type PinBlockInquiryResponse struct {

	//  Max length = 19, Card number for which the PIN is returned.
	CardNbr string `json:"cardNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 16, Encrypted Pin Block returned in case of successful retrieval.
	PinBlock string `json:"pinBlock,omitempty"`
}

// Validate validates this pin block inquiry response
func (m *PinBlockInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PinBlockInquiryResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *PinBlockInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PinBlockInquiryResponse) UnmarshalBinary(b []byte) error {
	var res PinBlockInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
