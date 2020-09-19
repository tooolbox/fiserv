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

// ReoccuringPaymentListResponse reoccuring payment list response
//
// swagger:model reoccuringPaymentListResponse
type ReoccuringPaymentListResponse struct {

	// This holds the number of CPAs stored for the given card or account number.
	CpaCount string `json:"cpaCount,omitempty"`

	// cpa details
	CpaDetails []*CpaDetails `json:"cpaDetails"`

	// This is the result of the message request which siginifes the success or failure of the call and will be populated as 0000 or E001 as follows.  <BR/> 0000-Success, <BR/> E001-Unsuccessful.
	MessageResult string `json:"messageResult,omitempty"`
}

// Validate validates this reoccuring payment list response
func (m *ReoccuringPaymentListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpaDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReoccuringPaymentListResponse) validateCpaDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.CpaDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.CpaDetails); i++ {
		if swag.IsZero(m.CpaDetails[i]) { // not required
			continue
		}

		if m.CpaDetails[i] != nil {
			if err := m.CpaDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpaDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReoccuringPaymentListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReoccuringPaymentListResponse) UnmarshalBinary(b []byte) error {
	var res ReoccuringPaymentListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
