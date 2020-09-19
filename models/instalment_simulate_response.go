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

// InstalmentSimulateResponse instalment simulate response
//
// swagger:model InstalmentSimulateResponse
type InstalmentSimulateResponse struct {

	//  Max length = 19, ACCT NBR
	AcctNbr string `json:"acctNbr,omitempty"`

	// common
	Common *Header `json:"common,omitempty"`

	//  Max length = 3, CURR
	Curr string `json:"curr,omitempty"`

	//  Max length = 3, CURR NOD
	CurrNod string `json:"currNod,omitempty"`

	//  Max length = 3, INST OPT OCCURS
	InstOptOccurs string `json:"instOptOccurs,omitempty"`

	// inst option tbl
	InstOptionTbl []*InstOptionTblForInstalmentSimulate1 `json:"instOptionTbl"`

	//  Max length = 5, INST PLAN NBR
	InstPlanNbr string `json:"instPlanNbr,omitempty"`

	//  Max length = 3, INST SCHED OCCURS
	InstSchedOccurs string `json:"instSchedOccurs,omitempty"`

	// inst schedule tbl
	InstScheduleTbl []*InstScheduleTblForInstalmentSimulate1 `json:"instScheduleTbl"`

	//  Max length = 7, INTEREST APR
	InterestApr string `json:"interestApr,omitempty"`

	//  Max length = 3, LOGO
	Logo string `json:"logo,omitempty"`

	//  Max length = 7, PERCENT NOD
	PercentNod string `json:"percentNod,omitempty"`

	//  Max length = 2, REQ TERM
	ReqTerm string `json:"reqTerm,omitempty"`

	//  Max length = 13, SCHEDULE FINAL-PMT-AMT
	ScheduleFinalPmtAmt string `json:"scheduleFinalPmtAmt,omitempty"`

	//  Max length = 13, SCHEDULE FIRST-PMT-AMT
	ScheduleFirstPmtAmt string `json:"scheduleFirstPmtAmt,omitempty"`

	//  Max length = 13, SCHEDULE FIXED-PMT-AMT
	ScheduleFixedPmtAmt string `json:"scheduleFixedPmtAmt,omitempty"`

	//  Max length = 13, SCHEDULE TOT-FEE-AMT
	ScheduleTotFeeAmt string `json:"scheduleTotFeeAmt,omitempty"`

	//  Max length = 13, SCHEDULE TOT IFEE AMT
	ScheduleTotIfeeAmt string `json:"scheduleTotIfeeAmt,omitempty"`

	//  Max length = 13, SCHEDULE TOT INT AMT
	ScheduleTotIntAmt string `json:"scheduleTotIntAmt,omitempty"`

	//  Max length = 7, TFC RATE
	TfcRate string `json:"tfcRate,omitempty"`

	//  Max length = 13, TRAN PRIN
	TranPrin string `json:"tranPrin,omitempty"`
}

// Validate validates this instalment simulate response
func (m *InstalmentSimulateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstOptionTbl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstScheduleTbl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstalmentSimulateResponse) validateCommon(formats strfmt.Registry) error {

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

func (m *InstalmentSimulateResponse) validateInstOptionTbl(formats strfmt.Registry) error {

	if swag.IsZero(m.InstOptionTbl) { // not required
		return nil
	}

	for i := 0; i < len(m.InstOptionTbl); i++ {
		if swag.IsZero(m.InstOptionTbl[i]) { // not required
			continue
		}

		if m.InstOptionTbl[i] != nil {
			if err := m.InstOptionTbl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instOptionTbl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstalmentSimulateResponse) validateInstScheduleTbl(formats strfmt.Registry) error {

	if swag.IsZero(m.InstScheduleTbl) { // not required
		return nil
	}

	for i := 0; i < len(m.InstScheduleTbl); i++ {
		if swag.IsZero(m.InstScheduleTbl[i]) { // not required
			continue
		}

		if m.InstScheduleTbl[i] != nil {
			if err := m.InstScheduleTbl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instScheduleTbl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstalmentSimulateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstalmentSimulateResponse) UnmarshalBinary(b []byte) error {
	var res InstalmentSimulateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
