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

// InstantIssuanceStatusUpdateRequest instant issuance status update request
//
// swagger:model InstantIssuanceStatusUpdateRequest
type InstantIssuanceStatusUpdateRequest struct {

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 19, Customer number: Number that identifies the Customer Name/Address record to which this account is assigned.A valid customer name and address must be on file for this account.
	// Required: true
	// Max Length: 19
	// Min Length: 0
	CustNbr *string `json:"custNbr"`

	//  Max length = 1, Status of data delivery / card production. Must be one of valid values: 3  (Data prep failure) 4  (Sent to client) 5  (Failed delivery to client) 6  (Processing failure at client) 7  (Delivery failure to branch/store) 8  (Card not produced in-store) 9  (Completed successfully)
	// Required: true
	// Max Length: 1
	// Min Length: 0
	StatusFlag *string `json:"statusFlag"`
}

// Validate validates this instant issuance status update request
func (m *InstantIssuanceStatusUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstantIssuanceStatusUpdateRequest) validateCommon(formats strfmt.Registry) error {

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

func (m *InstantIssuanceStatusUpdateRequest) validateCustNbr(formats strfmt.Registry) error {

	if err := validate.Required("custNbr", "body", m.CustNbr); err != nil {
		return err
	}

	if err := validate.MinLength("custNbr", "body", string(*m.CustNbr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("custNbr", "body", string(*m.CustNbr), 19); err != nil {
		return err
	}

	return nil
}

func (m *InstantIssuanceStatusUpdateRequest) validateStatusFlag(formats strfmt.Registry) error {

	if err := validate.Required("statusFlag", "body", m.StatusFlag); err != nil {
		return err
	}

	if err := validate.MinLength("statusFlag", "body", string(*m.StatusFlag), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("statusFlag", "body", string(*m.StatusFlag), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstantIssuanceStatusUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstantIssuanceStatusUpdateRequest) UnmarshalBinary(b []byte) error {
	var res InstantIssuanceStatusUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
