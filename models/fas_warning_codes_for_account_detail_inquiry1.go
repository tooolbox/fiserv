// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FasWarningCodesForAccountDetailInquiry1 fas warning codes for account detail inquiry1
//
// swagger:model FasWarningCodesForAccountDetailInquiry1
type FasWarningCodesForAccountDetailInquiry1 struct {

	//  Max length = 1, Warning Code: User defined field used by FirstVision authorisations to perform qualifying credit tests.  Values are:  '0' - Normal authorisations (Default)  '1' - Decline  '2' - Decline and pick up  '3' - Decline with fraud code  '4' - Decline with referral  '8' - Decline with a user exit  '9' - VIP account; always approve.
	WarningCode string `json:"warningCode,omitempty"`
}

// Validate validates this fas warning codes for account detail inquiry1
func (m *FasWarningCodesForAccountDetailInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FasWarningCodesForAccountDetailInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FasWarningCodesForAccountDetailInquiry1) UnmarshalBinary(b []byte) error {
	var res FasWarningCodesForAccountDetailInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
