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

// CommSubyTable2ForLoanPlanAdd2 comm suby table2 for loan plan add2
//
// swagger:model CommSubyTable2ForLoanPlanAdd2
type CommSubyTable2ForLoanPlanAdd2 struct {

	//  Max length = 9, Commission Agent.
	// Max Length: 9
	// Min Length: 0
	// Pattern: ^[0-9]*$
	CommAgent *string `json:"commAgent,omitempty"`

	//  Max length = 17, Commision Amount:.
	// Pattern: ^(-)?[0-9]{1,17}$
	CommAmt string `json:"commAmt,omitempty"`
}

// Validate validates this comm suby table2 for loan plan add2
func (m *CommSubyTable2ForLoanPlanAdd2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommAmt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommSubyTable2ForLoanPlanAdd2) validateCommAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.CommAgent) { // not required
		return nil
	}

	if err := validate.MinLength("commAgent", "body", string(*m.CommAgent), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("commAgent", "body", string(*m.CommAgent), 9); err != nil {
		return err
	}

	if err := validate.Pattern("commAgent", "body", string(*m.CommAgent), `^[0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *CommSubyTable2ForLoanPlanAdd2) validateCommAmt(formats strfmt.Registry) error {

	if swag.IsZero(m.CommAmt) { // not required
		return nil
	}

	if err := validate.Pattern("commAmt", "body", string(m.CommAmt), `^(-)?[0-9]{1,17}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommSubyTable2ForLoanPlanAdd2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommSubyTable2ForLoanPlanAdd2) UnmarshalBinary(b []byte) error {
	var res CommSubyTable2ForLoanPlanAdd2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
