// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PlanTableForClosingEndBalanceInquiry1 plan table for closing end balance inquiry1
//
// swagger:model PlanTableForClosingEndBalanceInquiry1
type PlanTableForClosingEndBalanceInquiry1 struct {

	//  Max length = 17, Accrued Interest: Amount of interest that has accrued for the Credit Plan Segment record in the current cycle-to-date.
	PlanAccruedInt string `json:"planAccruedInt,omitempty"`

	//  Max length = 17, Current balance: Total current outstanding balance of the plan.
	PlanCurrBal string `json:"planCurrBal,omitempty"`

	//  Max length = 17, Deferred Interest: Total amount of interest that has accrued during the interest deferment period to date.
	PlanDeferInt string `json:"planDeferInt,omitempty"`

	//  Max length = 17, Disputed balance of the plan: Total amount of transactions in dispute.
	PlanDispuBal string `json:"planDispuBal,omitempty"`

	//  Max length = 17, Per Diem Interest: Amount of interest that accrues per day at the current rate, also called daily interest accrual.
	PlanIntPerDiem string `json:"planIntPerDiem,omitempty"`

	//  Max length = 17, Outstanding Plan Balance: Sum of plan current balance, plan accrued interest, plan deferred interest and per diem interest.
	PlanOutstdBal string `json:"planOutstdBal,omitempty"`

	//  Max length = 2, Plan status: status of the Credit Plan Segment record.
	PlanStatus string `json:"planStatus,omitempty"`

	//  Max length = 1, Credit Plan Type: indicates the type of Plan. Values are: A - Access Cash B - Balance Transfer Retail C - Cash D - Deferment K - Access Retail R - Retail T - Balance Transfer Cash L - Loan
	PlanType string `json:"planType,omitempty"`
}

// Validate validates this plan table for closing end balance inquiry1
func (m *PlanTableForClosingEndBalanceInquiry1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanTableForClosingEndBalanceInquiry1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanTableForClosingEndBalanceInquiry1) UnmarshalBinary(b []byte) error {
	var res PlanTableForClosingEndBalanceInquiry1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
