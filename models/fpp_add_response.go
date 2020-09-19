// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FppAddResponse fpp add response
//
// swagger:model FppAddResponse
type FppAddResponse struct {

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 17, Unpaid Agreed Monthly minimum payment on Flexible Payment Plans. The value will calculated as the sum total of dues on Non-Flexible plans and unpaid instalment amounts on Flexible plans, and will be calculated daily to ensure that the effect of any interim payments, payment reversals or credits will be accounted for when a DD request is made
	FppUnpAgrdPmt string `json:"fppUnpAgrdPmt,omitempty"`

	//  Max length = 5, Plan Number: Number that identifies the Credit Plan Master record associated with the credit plan
	Plan string `json:"plan,omitempty"`

	//  Max length = 40, Credit Plan Description: Description of the FPP plan selected by the cardholder.
	PlanDesc string `json:"planDesc,omitempty"`

	//  Max length = 3, Credit Plan Data Record Number: Sequence number of Credit Plan.
	RecNbr string `json:"recNbr,omitempty"`
}

// Validate validates this fpp add response
func (m *FppAddResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FppAddResponse) validateCommon(formats strfmt.Registry) error {

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
func (m *FppAddResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FppAddResponse) UnmarshalBinary(b []byte) error {
	var res FppAddResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
