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

// WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1 waive txn fee grp for waive feesand interest update1
//
// swagger:model WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1
type WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1 struct {

	//  Max length = 1, Waive Transaction Fee Flag: Flag that indicates whether to waive transaction fees for prepaid accounts.  Values are:  0 - Do not waive transaction fee (Default)  1 - Waive transaction fee
	// Max Length: 1
	// Min Length: 0
	WaiveTxnFee *string `json:"waiveTxnFee,omitempty"`
}

// Validate validates this waive txn fee grp for waive feesand interest update1
func (m *WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWaiveTxnFee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1) validateWaiveTxnFee(formats strfmt.Registry) error {

	if swag.IsZero(m.WaiveTxnFee) { // not required
		return nil
	}

	if err := validate.MinLength("waiveTxnFee", "body", string(*m.WaiveTxnFee), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("waiveTxnFee", "body", string(*m.WaiveTxnFee), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1) UnmarshalBinary(b []byte) error {
	var res WaiveTxnFeeGrpForWaiveFeesandInterestUpdate1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
