// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssociatedPartiesInquiryResponse associated parties inquiry response
//
// swagger:model AssociatedPartiesInquiryResponse
type AssociatedPartiesInquiryResponse struct {

	//  Max length = 1,
	AssocPartiesGrp string `json:"assocPartiesGrp,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 3, Currency Code at Org Level: ISO currency code that identifies the unit of currency for this account.
	CurrencyCode string `json:"currencyCode,omitempty"`

	//  Max length = 1, Currency NOD: Number of decimal positions to the right of the decimal point (or other decimal character) in monetary amount fields.
	CurrencyNod string `json:"currencyNod,omitempty"`

	//  Max length = 1, Dual Account Indicator: Value received from Input.  Values are: L = Local account F = Foreign account
	DualInd string `json:"dualInd,omitempty"`

	//  Max length = 3, Identification number of the logo.
	Logo string `json:"logo,omitempty"`
}

// Validate validates this associated parties inquiry response
func (m *AssociatedPartiesInquiryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssociatedPartiesInquiryResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *AssociatedPartiesInquiryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssociatedPartiesInquiryResponse) UnmarshalBinary(b []byte) error {
	var res AssociatedPartiesInquiryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
