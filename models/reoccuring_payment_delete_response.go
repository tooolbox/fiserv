// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReoccuringPaymentDeleteResponse reoccuring payment delete response
//
// swagger:model reoccuringPaymentDeleteResponse
type ReoccuringPaymentDeleteResponse struct {

	// The message result siginifies the success or the  failure of the call and will be populated as 0000 or E001 as follows. <BR/> 0000-Sucesss, <BR/> E001-Unsuccessful.
	MessageResult string `json:"messageResult,omitempty"`
}

// Validate validates this reoccuring payment delete response
func (m *ReoccuringPaymentDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReoccuringPaymentDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReoccuringPaymentDeleteResponse) UnmarshalBinary(b []byte) error {
	var res ReoccuringPaymentDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
