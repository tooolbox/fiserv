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

// WaiveSvcChgGrpForWaiveFeesandInterestUpdate1 waive svc chg grp for waive feesand interest update1
//
// swagger:model WaiveSvcChgGrpForWaiveFeesandInterestUpdate1
type WaiveSvcChgGrpForWaiveFeesandInterestUpdate1 struct {

	//  Max length = 1, Waive Service Charge Flag: Flag that indicates whether to waive service charge for the account.  Values are:  0 - Do not waive svc charges if in A/R and waive in Profit/Loss (Default)  1 - Waive svc charges in A/R and P/L  2 - Do not waive svc charges in A/R and P/L  3 - Waive svc charges in A/R but do not waive in P/L.  4 - Waive svc charges in A/R & P/L for 1 cycle  5 - Waive svc charges in A/R & P/L for 2 cycles  6 - Waive svc charges in A/R & P/L for 3 cycles
	// Max Length: 1
	// Min Length: 0
	WaiveSvcChg *string `json:"waiveSvcChg,omitempty"`
}

// Validate validates this waive svc chg grp for waive feesand interest update1
func (m *WaiveSvcChgGrpForWaiveFeesandInterestUpdate1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWaiveSvcChg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WaiveSvcChgGrpForWaiveFeesandInterestUpdate1) validateWaiveSvcChg(formats strfmt.Registry) error {

	if swag.IsZero(m.WaiveSvcChg) { // not required
		return nil
	}

	if err := validate.MinLength("waiveSvcChg", "body", string(*m.WaiveSvcChg), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("waiveSvcChg", "body", string(*m.WaiveSvcChg), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WaiveSvcChgGrpForWaiveFeesandInterestUpdate1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WaiveSvcChgGrpForWaiveFeesandInterestUpdate1) UnmarshalBinary(b []byte) error {
	var res WaiveSvcChgGrpForWaiveFeesandInterestUpdate1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}