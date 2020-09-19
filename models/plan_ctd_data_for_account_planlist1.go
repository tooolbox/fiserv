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

// PlanCtdDataForAccountPlanlist1 plan ctd data for account planlist1
//
// swagger:model PlanCtdDataForAccountPlanlist1
type PlanCtdDataForAccountPlanlist1 struct {

	//  Max length = 7, Interest rate on the plan.
	// Required: true
	CalcRate1 *string `json:"calcRate1"`

	//  Max length = 17, The Flexible Payment Amount for plan segments that have a value. The agreed payment amount for the Flexible Payment Plan.  This field is display only and cannot be manually updated for an existing plan segment.
	// Required: true
	FppAgrdPmtAmt *string `json:"fppAgrdPmtAmt"`

	//  Max length = 7, The AER rate of an FPP plan.  Populated for FPP plans only
	// Required: true
	FppCurrAer *string `json:"fppCurrAer"`

	//  Max length = 17, Current balance: Total current outstanding balance of the plan.
	// Required: true
	FppCurrBal *string `json:"fppCurrBal"`

	//  Max length = 17, The plan balance eligible for transfer. Calculated value.
	// Required: true
	FppEligBal *string `json:"fppEligBal"`

	//  Max length = 1, Indicates whether plan is FPP or not. Valid values are 0 and 1. 0 - Not an FPP. 1 - FPP.
	// Required: true
	FppInd *string `json:"fppInd"`

	//  Max length = 17, Flexi plan initial booking balance. the initial balance of the Flexible Payment Plan
	// Required: true
	FppInitBal *string `json:"fppInitBal"`

	//  Max length = 3, For FPP Plans, this field is the Original Term of the Flexible Payment Plan
	// Required: true
	FppOrigTerm *string `json:"fppOrigTerm"`

	//  Max length = 1, Cancel reason for FPP. Valid values are 0 to 5 and 9. 0 - No reason. 1 - Paid out balances. 2 - Due to Delinquency. 3 - Requested. 4 - Interest. 5 - DRP. 9 - Agreed payment amount not paid.
	// Required: true
	FppPlanCancelRsn *string `json:"fppPlanCancelRsn"`

	//  Max length = 1, Cancellation flag. Indicate that a plan has been cancelled and waiting to be processed in the overnight batch Valid values: 0 = Plan is not cancelled. 1 = Plan is cancelled.
	// Required: true
	FppPlanCancelSw *string `json:"fppPlanCancelSw"`

	//  Max length = 3, Remaining Term. The calculated Remaining Term of the Flexible Payment Plan
	// Required: true
	FppRemTerm *string `json:"fppRemTerm"`

	//  Max length = 17, Remaining unpaid agreed payment amount.
	// Required: true
	FppUnpaidPmtAmt *string `json:"fppUnpaidPmtAmt"`

	//  Max length = 40, Credit Plan Description: Description of the Credit Plan.
	// Required: true
	PlanDesc *string `json:"planDesc"`

	//  Max length = 5, Credit Plan Number: Identifies the plan number of the Credit Plan Master record associated with the Credit Plan Segment record.
	// Required: true
	PlanNbr *string `json:"planNbr"`

	// Format: YYYYMMDD. Effective/start date of the FPP plan. Date on which the Credit Plan Segment record was opened, or generated, for the account.
	// Required: true
	PlanOpenDte *string `json:"planOpenDte"`

	//  Max length = 2, Credit Plan Data Record Number: Sequence number of Credit Plan.
	// Required: true
	PlanRecNbr *string `json:"planRecNbr"`

	//  Max length = 17, Total due amount present in the plan.
	// Required: true
	PlanTotalAmtDue *string `json:"planTotalAmtDue"`

	//  Max length = 1, Credit Plan Type: indicates the type of Plan. Values are: A - Access Cash B - Balance Transfer Retail C - Cash D - Deferment K - Access Retail R - Retail T - Balance Transfer Cash L - Loan
	// Required: true
	PlanType *string `json:"planType"`

	//  Max length = 17, Payment posted cycle-to-date
	// Required: true
	PmtCtd *string `json:"pmtCtd"`

	//  Max length = 1, Indicates whether plan is promo or not. Valid values are 0 to 3. 0 - Not a promo plan. 1 - Promo plan. 2 - BNPL promo plan. 3 - Special plan.
	// Required: true
	PromoPlanInd *string `json:"promoPlanInd"`

	//  Max length = 17, Life to date payments applied to flexi plan
	// Required: true
	TotPlanPayments *string `json:"totPlanPayments"`
}

// Validate validates this plan ctd data for account planlist1
func (m *PlanCtdDataForAccountPlanlist1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalcRate1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppAgrdPmtAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppCurrAer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppCurrBal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppEligBal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppInitBal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppOrigTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppPlanCancelRsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppPlanCancelSw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppRemTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFppUnpaidPmtAmt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanOpenDte(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanRecNbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanTotalAmtDue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmtCtd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromoPlanInd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotPlanPayments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateCalcRate1(formats strfmt.Registry) error {

	if err := validate.Required("calcRate1", "body", m.CalcRate1); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppAgrdPmtAmt(formats strfmt.Registry) error {

	if err := validate.Required("fppAgrdPmtAmt", "body", m.FppAgrdPmtAmt); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppCurrAer(formats strfmt.Registry) error {

	if err := validate.Required("fppCurrAer", "body", m.FppCurrAer); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppCurrBal(formats strfmt.Registry) error {

	if err := validate.Required("fppCurrBal", "body", m.FppCurrBal); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppEligBal(formats strfmt.Registry) error {

	if err := validate.Required("fppEligBal", "body", m.FppEligBal); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppInd(formats strfmt.Registry) error {

	if err := validate.Required("fppInd", "body", m.FppInd); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppInitBal(formats strfmt.Registry) error {

	if err := validate.Required("fppInitBal", "body", m.FppInitBal); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppOrigTerm(formats strfmt.Registry) error {

	if err := validate.Required("fppOrigTerm", "body", m.FppOrigTerm); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppPlanCancelRsn(formats strfmt.Registry) error {

	if err := validate.Required("fppPlanCancelRsn", "body", m.FppPlanCancelRsn); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppPlanCancelSw(formats strfmt.Registry) error {

	if err := validate.Required("fppPlanCancelSw", "body", m.FppPlanCancelSw); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppRemTerm(formats strfmt.Registry) error {

	if err := validate.Required("fppRemTerm", "body", m.FppRemTerm); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateFppUnpaidPmtAmt(formats strfmt.Registry) error {

	if err := validate.Required("fppUnpaidPmtAmt", "body", m.FppUnpaidPmtAmt); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePlanDesc(formats strfmt.Registry) error {

	if err := validate.Required("planDesc", "body", m.PlanDesc); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePlanNbr(formats strfmt.Registry) error {

	if err := validate.Required("planNbr", "body", m.PlanNbr); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePlanOpenDte(formats strfmt.Registry) error {

	if err := validate.Required("planOpenDte", "body", m.PlanOpenDte); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePlanRecNbr(formats strfmt.Registry) error {

	if err := validate.Required("planRecNbr", "body", m.PlanRecNbr); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePlanTotalAmtDue(formats strfmt.Registry) error {

	if err := validate.Required("planTotalAmtDue", "body", m.PlanTotalAmtDue); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePlanType(formats strfmt.Registry) error {

	if err := validate.Required("planType", "body", m.PlanType); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePmtCtd(formats strfmt.Registry) error {

	if err := validate.Required("pmtCtd", "body", m.PmtCtd); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validatePromoPlanInd(formats strfmt.Registry) error {

	if err := validate.Required("promoPlanInd", "body", m.PromoPlanInd); err != nil {
		return err
	}

	return nil
}

func (m *PlanCtdDataForAccountPlanlist1) validateTotPlanPayments(formats strfmt.Registry) error {

	if err := validate.Required("totPlanPayments", "body", m.TotPlanPayments); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanCtdDataForAccountPlanlist1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanCtdDataForAccountPlanlist1) UnmarshalBinary(b []byte) error {
	var res PlanCtdDataForAccountPlanlist1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
